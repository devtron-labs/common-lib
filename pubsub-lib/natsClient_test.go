/*
 * Copyright (c) 2024. Devtron Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pubsub_lib

import (
	"fmt"
	"github.com/devtron-labs/common-lib/utils"
	"github.com/nats-io/nats.go"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"
)

func TestNewPubSubClient(t *testing.T) {

	t.SkipNow()
	const payload = "stop-msg"
	var globalVal = 0

	t.Run("command exec", func(t *testing.T) {
		err := os.Setenv("AWS_ACCESS_KEY_ID", "abcd")
		command := exec.Command("echo", "$AWS_ACCESS_KEY_ID", ">hello.world")
		err = command.Run()
		fmt.Println(err)
	})

	t.Run("subscriber", func(t *testing.T) {
		queueSubscriber(payload, true)
	})

	t.Run("subscriber1", func(t *testing.T) {
		queueSubscriber(payload, false)
	})

	t.Run("subscriber2", func(t *testing.T) {
		queueSubscriber(payload, false)
	})

	t.Run("pullSubscriber", func(t *testing.T) {
		sugaredLogger, _ := utils.NewSugardLogger()
		pubSubClient, _ := NewNatsClient(sugaredLogger)
		//streamConfig := &nats.StreamConfig{}
		//_ = AddStream(pubSubClient.JetStrCtxt, streamConfig, "New_Stream_2")
		subs, err := pubSubClient.JetStrCtxt.PullSubscribe("hello.world", WORKFLOW_STATUS_UPDATE_DURABLE, nats.BindStream("New_Stream_2"))
		if err != nil {
			fmt.Println("error occurred while subscribing pull reason: ", err)
			return
		}
		for subs.IsValid() {
			fmt.Println("fetching msgs")
			msgs, err := subs.Fetch(10)
			if err != nil && err == nats.ErrTimeout {
				fmt.Println(" timeout occurred but we have to try again")
				time.Sleep(5 * time.Second)
				continue
			} else if err != nil {
				fmt.Println("error occurred while extracting msg", err)
				return
			}
			for _, nxtMsg := range msgs {
				fmt.Println("Received a JetStream message: ", string(nxtMsg.Data))
				if string(nxtMsg.Data) == payload {
					return
				}
				defer nxtMsg.Ack()
			}
			time.Sleep(5 * time.Second)
		}
	})

	t.Run("publisher", func(t *testing.T) {
		publishMsg(globalVal)
	})

	t.Run("stopPublisher", func(t *testing.T) {
		sugaredLogger, _ := utils.NewSugardLogger()
		pubSubClient, _ := NewNatsClient(sugaredLogger)
		topic := "CD.TRIGGER" // for pull subs
		// topic := "CI-COMPLETE"
		streamName := ORCHESTRATOR_STREAM
		// streamName := util1.CI_RUNNER_STREAM

		WriteNatsEvent(pubSubClient, topic, payload, streamName)
	})
}

func publishMsg(globalVal int) {
	go handlePanic()
	sugaredLogger, _ := utils.NewSugardLogger()
	pubSubClient, _ := NewNatsClient(sugaredLogger)
	// topic := "CD.TRIGGER"
	// topic := "CI-COMPLETE"
	topic := "hello.world"
	// streamName := util1.ORCHESTRATOR_STREAM
	// streamName := util1.CI_RUNNER_STREAM
	streamName := "New_Stream_2"

	for true {
		globalVal++
		helloWorld := "Hello World " + strconv.Itoa(globalVal)
		WriteNatsEvent(pubSubClient, topic, helloWorld, streamName)
	}
}

func queueSubscriber(payload string, durable1 bool) {
	sugaredLogger, _ := utils.NewSugardLogger()
	pubSubClient, _ := NewNatsClient(sugaredLogger)
	globalVar := false
	//streamConfig := &nats.StreamConfig{}
	//_ = AddStream(pubSubClient.JetStrCtxt, streamConfig, "New_Stream_2")
	durable := "WORKFLOW_STATUS_UPDATE_DURABLE-1"
	if durable1 {
		durable = "WORKFLOW_STATUS_UPDATE_DURABLE-2"
	}
	subs, err := pubSubClient.JetStrCtxt.QueueSubscribe("hello.world", "CI-COMPLETE_GROUP-1", func(msg *nats.Msg) {
		println("msg received")
		defer msg.Ack()
		println(string(msg.Data))
		if string(msg.Data) == payload {
			globalVar = true
		}
	}, nats.Durable(durable), nats.DeliverAll(), nats.ManualAck(), nats.BindStream("New_Stream_2"))
	if err != nil {
		fmt.Println("error is ", err)
		return
	}
	for true {
		if globalVar {
			break
		}
		fmt.Println("looping & checking subs status: ", subs.IsValid())
		time.Sleep(5 * time.Second)
	}
}

func WriteNatsEvent(psc *NatsClient, topic string, payload string, streamName string) {
	//streamConfig := &nats.StreamConfig{}
	///_ = AddStream(psc.JetStrCtxt, streamConfig, streamName)
	// Generate random string for passing as Header Id in message
	randString := "MsgHeaderId-" + utils.Generate(10)
	_, err := psc.JetStrCtxt.Publish(topic, []byte(payload), nats.MsgId(randString))
	if err != nil {
		fmt.Println("error occurred while publishing event reason: ", err)
	} else {
		fmt.Println("msg published " + payload)
	}
}

func handlePanic() {
	if err := recover(); err != nil {
		fmt.Println("panic occurred:", err)
	}
}

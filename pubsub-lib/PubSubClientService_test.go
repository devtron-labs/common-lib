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
	"github.com/caarlos0/env"
	"github.com/devtron-labs/common-lib/pubsub-lib/model"
	"github.com/devtron-labs/common-lib/utils"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewPubSubClientServiceImpl(t *testing.T) {

	const payload = "stop-msg"

	// t.SkipNow()
	t.Run("PubAndSub", func(t *testing.T) {
		//	err := os.Setenv("STREAM_CONFIG_JSON", "{\"Devtron_Test_Stream\":{\"streamConfig\":{\"max_age\":900000000000,\"num_replicas\":4}}}")
		//	fmt.Println(err)
		//err := os.Setenv("CONSUMER_CONFIG_JSON", "{\"Test_Topic_Consumer\":{\"replicas\":2}}")
		//err := os.Setenv("CONSUMER_CONFIG_JSON", "{\"Test_Topic_Consumer\":{\"natsMsgProcessingBatchSize\":3,\"natsMsgBufferSize\":64,\"replicas\":2}}")

		//fmt.Println(err)
		sugaredLogger, _ := utils.NewSugardLogger()
		err := os.Setenv("STREAM_CONFIG_JSON", "{\"ORCHESTRATOR\":{\"streamConfig\":{\"num_replicas\":3}}}")
		fmt.Println(err)
		err = os.Setenv("CONSUMER_CONFIG_JSON", "{\"NOTIFICATION_EVENT_DURABLE\":{\"natsMsgBufferSize\":2,\"natsMsgProcessingBatchSize\":1}}")
		fmt.Println(err)
		fmt.Println(os.Getenv("CONSUMER_CONFIG_JSON"))
		fmt.Println(os.Getenv("STREAM_CONFIG_JSON"))
		err = ParseAndFillStreamWiseAndConsumerWiseConfigMaps()
		pubSubClient, err := NewPubSubClientServiceImpl(sugaredLogger)

		//for i := 0; i < 10; i++ {
		//	//time.Sleep(1 * time.Second)
		//	err = pubSubClient.Publish(NOTIFICATION_EVENT_TOPIC, fmt.Sprintf("published Msg %d", i))
		//}
		if err != nil {
			sugaredLogger.Fatalw("error occurred while publishing to topic")
		}

		err = pubSubClient.Subscribe(NOTIFICATION_EVENT_TOPIC,
			func(msg *model.PubSubMsg) {
				fmt.Println("Data received:", msg.Data)
			},
			func(msg model.PubSubMsg) (logMsg string, keysAndValues []interface{}) {
				return logMsg, keysAndValues
			},
			func(msg model.PubSubMsg) bool {
				return true
			})
		if err != nil {
			sugaredLogger.Fatalw("error occurred while subscribing to topic")
		}

		time.Sleep(time.Duration(1000) * time.Second)
	})

	//t.Run("SubOnly", func(t *testing.T) {
	//	sugaredLogger, _ := utils.NewSugardLogger()
	//	var pubSubClient = NewPubSubClientServiceImpl(sugaredLogger)
	//	err := pubSubClient.Subscribe(DEVTRON_TEST_TOPIC, func(msg *model.PubSubMsg) {
	//		fmt.Println("Data received:", msg.Data)
	//	},
	//		func(msg *model.PubSubMsg) {
	//
	//		})
	//	if err != nil {
	//		sugaredLogger.Fatalw("error occurred while subscribing to topic")
	//	}
	//	time.Sleep(time.Duration(500) * time.Second)
	//})

	//t.Run("SubOnly1", func(t *testing.T) {
	//	sugaredLogger, _ := utils.NewSugardLogger()
	//	var pubSubClient = NewPubSubClientServiceImpl(sugaredLogger)
	//	Consumed_Counter := 0
	//	lock := &sync.Mutex{}
	//	err := pubSubClient.Subscribe(DEVTRON_TEST_TOPIC, func(msg *model.PubSubMsg) {
	//		lock.Lock()
	//		Consumed_Counter++
	//		lock.Unlock()
	//		fmt.Println(time.Now(), "Data received:", msg.Data, " count", Consumed_Counter)
	//		time.Sleep(1 * time.Second)
	//	},
	//		func(msg *model.PubSubMsg) {
	//
	//		})
	//	if err != nil {
	//		sugaredLogger.Fatalw("error occurred while subscribing to topic")
	//	}
	//	time.Sleep(time.Duration(500) * time.Second)
	//})

	t.Run("PullSubs", func(t *testing.T) {
		sugaredLogger, _ := utils.NewSugardLogger()
		natsClient, err := NewNatsClient(sugaredLogger)
		subs, err := natsClient.JetStrCtxt.PullSubscribe(DEVTRON_TEST_TOPIC, DEVTRON_TEST_CONSUMER, nats.BindStream(DEVTRON_TEST_STREAM),
			nats.PullMaxWaiting(5))
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
			fmt.Println("msg recv count: ", len(msgs))
			for _, nxtMsg := range msgs {
				fmt.Println("Received a JetStream message: ", string(nxtMsg.Data))
				if string(nxtMsg.Data) == payload {
					return
				}
				nxtMsg.Ack()
			}
			time.Sleep(5 * time.Second)
		}

	})

	//t.Run("PubOnly", func(t *testing.T) {
	//	sugaredLogger, _ := utils.NewSugardLogger()
	//	var pubSubClient = NewPubSubClientServiceImpl(sugaredLogger)
	//	var ops uint64
	//	var msgId uint64
	//	channel := make(chan string, 64)
	//	wg := new(sync.WaitGroup)
	//	for index := 0; index < 3; index++ {
	//		wg.Add(1)
	//		go publishNatsMsg(pubSubClient, sugaredLogger, &ops, wg, channel)
	//	}
	//	for true {
	//		atomic.AddUint64(&msgId, 1)
	//		msg := "published Msg " + strconv.FormatUint(msgId, 10)
	//		channel <- msg
	//		// time.Sleep(1 * time.Second)
	//	}
	//	wg.Wait()
	//})

	t.Run("StreamWiseAndConsumerWiseConfig with default configs", func(t *testing.T) {
		ParseAndFillStreamWiseAndConsumerWiseConfigMaps()
		config := NatsClientConfig{}
		err := env.Parse(&config)
		if err != nil {
			log.Fatal("error occurred while parsing nats client config", "err", err)
		}
		var defaultStreamConfig = config.GetDefaultNatsStreamConfig()
		for _, streamWiseConfig := range NatsStreamWiseConfigMapping {
			assert.Equal(t, defaultStreamConfig.StreamConfig, streamWiseConfig.StreamConfig)
		}

		var defaultConsumerConfig = config.GetDefaultNatsConsumerConfig()

		for _, consumerWiseConfig := range NatsConsumerWiseConfigMapping {
			assert.Equal(t, defaultConsumerConfig, consumerWiseConfig)
		}
	})

	t.Run("StreamWiseAndConsumerWiseConfig with json configs", func(t *testing.T) {
		err := os.Setenv("STREAM_CONFIG_JSON", "{\"ORCHESTRATOR\":{\"streamConfig\":{\"replicas\":3}}}")
		fmt.Println(err)
		err = os.Setenv("CONSUMER_CONFIG_JSON", "{\"NOTIFICATION_EVENT_DURABLE\":{\"replicas\":3}}")
		fmt.Println(err)
		fmt.Println(os.Getenv("CONSUMER_CONFIG_JSON"))
		fmt.Println(os.Getenv("STREAM_CONFIG_JSON"))
		//err = ParseAndFillStreamWiseAndConsumerWiseConfigMaps()
		//config := NatsClientConfig{}
		//err = env.Parse(&config)
		//if err != nil {
		//	log.Fatal("error occurred while parsing nats client config", "err", err)
		//}
		//var defaultStreamConfig = config.GetDefaultNatsStreamConfig()
		//for streamName, streamWiseConfig := range NatsStreamWiseConfigMapping {
		//	if streamName == ORCHESTRATOR_STREAM {
		//		assert.NotEqual(t, defaultStreamConfig.StreamConfig, streamWiseConfig.StreamConfig)
		//	} else {
		//		assert.Equal(t, defaultStreamConfig.StreamConfig, streamWiseConfig.StreamConfig)
		//	}
		//}
		//
		//var defaultConsumerConfig = config.GetDefaultNatsConsumerConfig()
		//
		//defaultConsumerConfigForBulkCdTrigger := defaultConsumerConfig
		//
		//for consumerName, consumerWiseConfig := range NatsConsumerWiseConfigMapping {
		//	if consumerName == NOTIFICATION_EVENT_DURABLE {
		//		assert.NotEqual(t, defaultConsumerConfig, consumerWiseConfig)
		//	} else {
		//
		//		if consumerName == BULK_DEPLOY_DURABLE {
		//			assert.Equal(t, defaultConsumerConfigForBulkCdTrigger, consumerWiseConfig)
		//			continue
		//		}
		//		assert.Equal(t, defaultConsumerConfig, consumerWiseConfig)
		//	}
		//}
	})
}

func publishNatsMsg(pubSubClient *PubSubClientServiceImpl, sugaredLogger *zap.SugaredLogger, publishedCounter *uint64, wg *sync.WaitGroup, channel chan string) {
	defer wg.Done()
	for natsMsg := range channel {
		err := pubSubClient.Publish(DEVTRON_TEST_TOPIC, natsMsg)
		if err != nil {
			sugaredLogger.Fatalw("error occurred while publishing to topic")
		}
		atomic.AddUint64(publishedCounter, 1)
		fmt.Println("msg ", natsMsg, " count ", *publishedCounter)
	}
}

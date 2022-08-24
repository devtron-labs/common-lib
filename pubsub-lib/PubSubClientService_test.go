package pubsub_lib

import (
	"fmt"
	"github.com/devtron-labs/common-lib/utils"
	"github.com/nats-io/nats.go"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestNewPubSubClientServiceImpl(t *testing.T) {

	const payload = "stop-msg"

	t.SkipNow()
	t.Run("PubAndSub", func(t *testing.T) {
		sugaredLogger, _ := utils.NewSugardLogger()
		var pubSubClient = NewPubSubClientServiceImpl(sugaredLogger)
		err := pubSubClient.Subscribe(DEVTRON_TEST_TOPIC, func(msg *PubSubMsg) {
			fmt.Println("Data received:", msg.Data)
		})
		if err != nil {
			sugaredLogger.Fatalw("error occurred while subscribing to topic")
		}
		err = pubSubClient.Publish(DEVTRON_TEST_TOPIC, "published Msg "+strconv.Itoa(time.Now().Second()))
		if err != nil {
			sugaredLogger.Fatalw("error occurred while publishing to topic")
		}
		time.Sleep(time.Duration(5) * time.Second)
	})

	t.Run("SubOnly", func(t *testing.T) {
		sugaredLogger, _ := utils.NewSugardLogger()
		var pubSubClient = NewPubSubClientServiceImpl(sugaredLogger)
		err := pubSubClient.Subscribe(DEVTRON_TEST_TOPIC, func(msg *PubSubMsg) {
			fmt.Println("Data received:", msg.Data)
		})
		if err != nil {
			sugaredLogger.Fatalw("error occurred while subscribing to topic")
		}
		//err = pubSubClient.Publish(DEVTRON_TEST_TOPIC, "published Msg "+strconv.Itoa(time.Now().Second()))
		//if err != nil {
		//	sugaredLogger.Fatalw("error occurred while publishing to topic")
		//}
		time.Sleep(time.Duration(500) * time.Second)
	})

	t.Run("SubOnly1", func(t *testing.T) {
		sugaredLogger, _ := utils.NewSugardLogger()
		var pubSubClient = NewPubSubClientServiceImpl(sugaredLogger)
		Consumed_Counter := 0
		lock := &sync.Mutex{}
		err := pubSubClient.Subscribe(DEVTRON_TEST_TOPIC, func(msg *PubSubMsg) {
			lock.Lock()
			Consumed_Counter++
			lock.Unlock()
			fmt.Println(time.Now(), "Data received:", msg.Data, " count", Consumed_Counter)
			time.Sleep(1 * time.Second)
		})
		if err != nil {
			sugaredLogger.Fatalw("error occurred while subscribing to topic")
		}
		//err = pubSubClient.Publish(DEVTRON_TEST_TOPIC, "published Msg "+strconv.Itoa(time.Now().Second()))
		//if err != nil {
		//	sugaredLogger.Fatalw("error occurred while publishing to topic")
		//}
		time.Sleep(time.Duration(500) * time.Second)
	})

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

	t.Run("PubOnly", func(t *testing.T) {
		sugaredLogger, _ := utils.NewSugardLogger()
		var pubSubClient = NewPubSubClientServiceImpl(sugaredLogger)
		Published_Counter := 0
		//err := pubSubClient.Subscribe(DEVTRON_TEST_TOPIC, func(msg *PubSubMsg) {
		//	fmt.Println("Data received:", msg.Data)
		//})
		//if err != nil {
		//	sugaredLogger.Fatalw("error occurred while subscribing to topic")
		//}

		lock := &sync.Mutex{}
		for true {
			msg := "published Msg " + strconv.Itoa(time.Now().Second())
			err := pubSubClient.Publish(DEVTRON_TEST_TOPIC, msg)
			if err != nil {
				sugaredLogger.Fatalw("error occurred while publishing to topic")
			}
			lock.Lock()
			Published_Counter++
			lock.Unlock()
			fmt.Println("msg ", msg, " count ", Published_Counter)
			time.Sleep(1 * time.Second)
		}
		time.Sleep(time.Duration(5) * time.Second)
	})
}

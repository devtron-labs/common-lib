package pubsub_lib

import (
	"fmt"
	"github.com/devtron-labs/common-lib/utils"
	"strconv"
	"testing"
	"time"
)

func TestNewPubSubClientServiceImpl(t *testing.T) {

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
}

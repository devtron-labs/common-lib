package nats_lib

import (
	"github.com/devtron-labs/common-lib/utils"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type PubSubClientService interface {
	Publish(streamName string, subject string, msg string) error
	Subscribe(streamName string, subject string, callback func(msg PubSubMsg)) error
}

type PubSubMsg struct {
	msg string
}

type PubSubClientServiceImpl struct {
	logger     *zap.SugaredLogger
	natsClient *NatsClient
}

func NewPubSubClientServiceImpl(logger *zap.SugaredLogger) *PubSubClientServiceImpl {
	natsClient, err := NewNatsClient(logger)
	if err != nil {
		logger.Fatalw("error occurred while creating nats client stopping now!!")
	}
	pubSubClient := &PubSubClientServiceImpl{
		logger:     logger,
		natsClient: natsClient,
	}
	return pubSubClient
}

func (impl PubSubClientServiceImpl) Publish(streamName string, subject string, msg string) error {
	natsClient := impl.natsClient
	jetStrCtxt := natsClient.JetStrCtxt
	_ = AddStream(jetStrCtxt, natsClient.streamConfig, streamName)
	//Generate random string for passing as Header Id in message
	randString := "MsgHeaderId-" + utils.Generate(10)
	_, err := jetStrCtxt.Publish(CRON_EVENTS, []byte(msg), nats.MsgId(randString))
	if err != nil {
		impl.logger.Errorw("error while publishing message", "stream", streamName, "subject", subject, "error", err)
		return err
	}
	return nil
}

func (impl PubSubClientServiceImpl) Subscribe(streamName string, subject string, callback func(msg *PubSubMsg)) error {
	_ = AddStream(impl.natsClient.JetStrCtxt, impl.natsClient.streamConfig, streamName)
	_, err := impl.natsClient.JetStrCtxt.QueueSubscribe(subject, WORKFLOW_STATUS_UPDATE_GROUP, func(msg *nats.Msg) {
		defer msg.Ack()
		subMsg := &PubSubMsg{msg: string(msg.Data)}
		callback(subMsg)
	})
	if err != nil {
		impl.logger.Fatalw("error while subscribing", "stream", streamName, "subject", subject, "error", err)
		return err
	}
	
	return nil
}

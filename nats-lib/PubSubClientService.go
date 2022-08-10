package nats_lib

import (
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
	logger       *zap.SugaredLogger
	JetStrCtxt   nats.JetStreamContext
	streamConfig *nats.StreamConfig
	Conn         nats.Conn
}

func NewPubSubClientServiceImpl(logger *zap.SugaredLogger) {

}

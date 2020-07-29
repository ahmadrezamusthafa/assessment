package nsq

import (
	"fmt"
	"github.com/ahmadrezamusthafa/assessment/common/logger"
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/nsqio/go-nsq"
)

type ServiceHandler func(*nsq.Message, string) bool

type NSQ struct {
	Config   *config.Config `inject:"config"`
	producer *nsq.Producer
	consumer *nsq.Consumer
}

type AssessmentNSQ struct {
	Config config.Config `inject:"config"`
	NSQ
}

func (n *NSQ) Publish(topic string, jsonByte []byte) error {
	return n.producer.Publish(topic, jsonByte)
}

func (n *AssessmentNSQ) StartUp() {
	logger.Info("Initiating assessment nsq producer... ")
	producer, err := nsq.NewProducer(
		fmt.Sprintf("%s:%d", n.Config.NSQDHost, n.Config.NSQDPort),
		nsq.NewConfig(),
	)
	if err != nil {
		logger.Err("Error while init producer ", err)
	} else if err := producer.Ping(); err != nil {
		logger.Err("Error while connecting to nsq server")
	} else {
		logger.Info("Successfully connected to nsq server")
	}
	n.producer = producer
}

func (n *AssessmentNSQ) Shutdown() {
	if n.producer != nil {
		n.producer.Stop()
	}
}

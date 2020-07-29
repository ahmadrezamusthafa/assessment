package order

import (
	"fmt"
	"github.com/ahmadrezamusthafa/assessment/common/logger"
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/order"
	"github.com/ahmadrezamusthafa/assessment/domain/usecase/orderproduct"
	"github.com/ahmadrezamusthafa/assessment/domain/usecase/product"
	"github.com/ahmadrezamusthafa/assessment/pkg/cache"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
	assessmentnsq "github.com/ahmadrezamusthafa/assessment/pkg/nsq"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

const (
	TopicAddOrder = "Topic.AddOrder"
)

type OrderService struct {
	Config              config.Config                     `inject:"config"`
	DB                  *database.AssessmentDatabase      `inject:"database"`
	Cache               *cache.AssessmentCache            `inject:"cache"`
	NSQ                 *assessmentnsq.AssessmentNSQ      `inject:"nsq"`
	ProductService      *product.ProductService           `inject:"productService"`
	OrderProductService *orderproduct.OrderProductService `inject:"orderProductService"`
	OrderDomain         order.Domain
	nsqconfig           *nsq.Config
}

func (svc *OrderService) StartUp() {
	svc.nsqconfig = nsq.NewConfig()
	svc.nsqconfig.MaxInFlight = 5
	svc.nsqconfig.MaxAttempts = 10
	orderRepository := order.OrderRepository{
		DB: svc.DB,
	}
	svc.OrderDomain = order.NewDomainRepository(orderRepository)

	err := svc.addConsumer(TopicAddOrder, svc.addOrder)
	if err != nil {
		logger.Err("Error add consumer ", err)
	}
}

func (svc *OrderService) Shutdown() {}

func (svc *OrderService) addConsumer(topic string, serviceHandler assessmentnsq.ServiceHandler) error {
	logger.Info("Add nsq consumer for topic %s ", topic)
	consumer, err := nsq.NewConsumer(TopicAddOrder, ".ch", svc.nsqconfig)
	if err != nil {
		logger.Err("Error while init consumer ", err)
		return err
	}
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		logger.Info("Got a message: %s id:%s", message.Body, message.ID)
		isRequeue := serviceHandler(message, topic)
		if isRequeue {
			message.Requeue(2 * time.Second)
		} else {
			message.Finish()
		}
		return nil
	}))

	err = consumer.ConnectToNSQLookupd(svc.Config.NSQLookupHost + ":" + fmt.Sprint(svc.Config.NSQLookupPort))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

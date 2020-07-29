package order

import (
	"context"
	"github.com/ahmadrezamusthafa/assessment/common/logger"
	"github.com/ahmadrezamusthafa/assessment/shared"
	jsoniter "github.com/json-iterator/go"
	"github.com/nsqio/go-nsq"
)

func (svc *OrderService) addOrder(message *nsq.Message, topic string) bool {
	ctx := context.Background()
	var data shared.Order
	err := jsoniter.Unmarshal(message.Body, &data)
	if err != nil {
		logger.Err("Error unmarshal ", err)
		return false
	}

	err = svc.AddOrder(ctx, data, true)
	if err != nil {
		logger.Err("Error add order from nsq ", err)
		return true
	}
	return false
}

package nsq

import (
	"context"
	"fmt"
	"item-service/config"
	"item-service/internal/service"

	"github.com/nsqio/go-nsq"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	ConsumerModule = fx.Invoke(NewConsumers)
)

func NewConsumers(lifecycle fx.Lifecycle, config config.Config, groupService service.BusinessGroupService, logger *zap.Logger) (*nsq.Consumer, error) {

	consumer, err := buildConsumer(groupService, logger)

	if err != nil {
		return nil, err
	}

	address := fmt.Sprintf("%s:%d", config.Nsq.Host, config.Nsq.Port)

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Start Nsq Consumers....")
			go func() {
				err := consumer.ConnectToNSQD(address)
				if err != nil {
					logger.Error(err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping Nsq Consumers....")
			consumer.Stop()
			return nil
		},
	})

	return consumer, nil

}

type (
	AConsumerhandler struct {
		groupService service.BusinessGroupService
		logger       *zap.Logger
	}
)

func (c *AConsumerhandler) HandleMessage(message *nsq.Message) error {
	//todo 处理消息
	c.logger.Info(fmt.Sprintf("%s => %s", "NSQ_A", string(message.Body)))
	return nil
}

func buildConsumer(groupService service.BusinessGroupService, logger *zap.Logger) (*nsq.Consumer, error) {
	//todo 也可以 config 配置
	consumer, err := nsq.NewConsumer("test", "test-channel", nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	consumer.AddHandler(&AConsumerhandler{
		groupService: groupService,
		logger:       logger,
	})
	return consumer, err
}

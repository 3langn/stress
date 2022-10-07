package nsq

import (
	"context"
	"fmt"
	"item-service/config"

	"github.com/nsqio/go-nsq"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var ProducerModule = fx.Provide(NewProducer)

func NewProducer(lifecycle fx.Lifecycle, config config.Config, logger *zap.Logger) (*nsq.Producer, error) {
	address := fmt.Sprintf("%s:%d", config.Nsq.Host, config.Nsq.Port)
	producer, err := nsq.NewProducer(address, nsq.NewConfig())
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Start Nsq Producer....")
			go func() {
				err = producer.Ping()
				if err != nil {
					logger.Error(err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping Nsq Producer....")
			producer.Stop()
			return nil
		},
	})
	return producer, err
}

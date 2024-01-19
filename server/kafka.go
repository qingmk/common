package service

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

type KafkaService struct {
	Kafka sarama.SyncProducer
}

func NewKafkaService(Kafka sarama.SyncProducer) *KafkaService {
	return &KafkaService{
		Kafka: Kafka,
	}
}

func NewKafkaServiceV2(Brokers []string) *KafkaService {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	p, err := sarama.NewSyncProducer(Brokers, config)
	if err != nil {
		logx.Error("kafka connect ping failed, err:", zap.Error(err))
	}
	return NewKafkaService(p)

}

//kafka发送测试
/**initialize.Kafka()
msgx := &sarama.ProducerMessage{Topic: "test-ken-io", Value: sarama.StringEncoder("i am data2  2322")}
global.GVA_KAFKA.SendMessage(msgx)**/
//kafka发送的服务
func (exa *KafkaService) SendKafkaMessage(topic string, message string) (err error) {
	logx.Info("开始发送kafka消息", zap.String("topic", topic), zap.String("message", message))
	msgx := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(message)}
	exa.Kafka.SendMessage(msgx)
	return err
}

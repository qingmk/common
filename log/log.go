package log

import (
	"encoding/json"

	"github.com/qingmk/common/dto"
	service "github.com/qingmk/common/server"
)

type Logger struct {
	KafkaService *service.KafkaService

	// 定义一个chan
	MsgChan chan dto.LogMsg
	Topic   string
}

// 这里是不是还应该传kafka的集群地址
func NewLogger(Brokers []string, Topic string) *Logger {

	KafkaService := service.NewKafkaServiceV2(Brokers)

	return &Logger{
		KafkaService: KafkaService,
		MsgChan:      make(chan dto.LogMsg, 1000),
		Topic:        Topic,
	}
}

func (logger *Logger) ReceiveLog(message dto.LogMsg) (err error) {
	logger.MsgChan <- message
	return
}

func (logger *Logger) ReadLog() (err error) {

	/**select {
	case c := <-exa.MsgChan:
		fmt.Println(c)
		//logx.Error(c)
	default:
		fmt.Println("After one second!")
		time.Sleep(time.Second)

	}**/

	for c := range logger.MsgChan {
		//fmt.Println(c)
		bytes, _ := json.Marshal(c)
		logger.KafkaService.SendKafkaMessage(logger.Topic, string(bytes))
	}

	return
}

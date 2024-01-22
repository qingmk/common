package log

import (
	"encoding/json"
	"time"

	"github.com/qingmk/common/cont/log"
	"github.com/qingmk/common/dto"
	service "github.com/qingmk/common/server"
)

type Logger struct {
	KafkaService *service.KafkaService

	// 定义一个chan
	MsgChan chan dto.LogMsg
	Topic   string
	Node    string //节点名称
	Service string //那个服务发送的信息
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

func NewLoggerV2(Brokers []string, Topic string, Node string, Service string) *Logger {

	KafkaService := service.NewKafkaServiceV2(Brokers)

	logger := &Logger{
		KafkaService: KafkaService,
		MsgChan:      make(chan dto.LogMsg, 1000),
		Topic:        Topic,
		Node:         Node,
		Service:      Service,
	}

	go logger.ReadLog()
	return logger
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
		bytes, _ := json.Marshal(c)
		logger.KafkaService.SendKafkaMessage(logger.Topic, string(bytes))
	}

	return
}

func (logger *Logger) Error(mgs string) (err error) {
	message := dto.LogMsg{
		Service: logger.Service,
		Msg:     mgs,
		Level:   log.ERROR,
		Node:    logger.Node,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}
	logger.MsgChan <- message
	return
}

func (logger *Logger) Warn(mgs string) (err error) {
	message := dto.LogMsg{
		Service: logger.Service,
		Msg:     mgs,
		Level:   log.WARN,
		Node:    logger.Node,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}
	logger.MsgChan <- message
	return
}

func (logger *Logger) INFO(mgs string) (err error) {
	message := dto.LogMsg{
		Service: logger.Service,
		Msg:     mgs,
		Level:   log.INFO,
		Node:    logger.Node,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}
	logger.MsgChan <- message
	return
}

package dto

type LogMsg struct {
	Node    string //节点名称
	Time    string //日志时间
	Msg     string //日志信息
	Service string //那个服务发送的信息
	Level   string //日志级别
}

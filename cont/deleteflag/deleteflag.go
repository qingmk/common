package deleteFlag

type DeleteflagType int

const (
	//正常的状态
	Normal DeleteflagType = 1

	//删除的状态
	Deleted DeleteflagType = 2
)

package resource

const (

	//资源类型的菜单的id
	MENUID int64 = 2
	//资源类型的api的id
	APIID int64 = 3
	//资源类型的命名空间id
	NAMESPACEID int64 = 4
	//资源类型的资源类型id
	RESOURCETYPEID int64 = 5
	//规则类型中的角色
	TYPEROLE int64 = 1
	//规则类型中的用户
	TYPEUSER           int64 = 2
	DEFAULTNAMESPACEID int64 = 10
)

const (
	MENUOPERATE    string = "HAS"
	MENUNAME       string = "MENU"
	APINAME        string = "API"
	NAMESPACE      string = "NAMESPACE"
	DEFAULTOPERATE string = "HAS"
	RESOURCETYPE   string = "RESOURCETYPE"
)

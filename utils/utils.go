package utils

func ContainsGeneric[T comparable](slice []T, element T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}

	return false
}

func IsAdmin(roleIdList []int64, roleId int64) bool {
	return ContainsGeneric(roleIdList, roleId)
}

// DeleteSlice2 删除指定元素。
func DeleteSlice2[T comparable](a []T, elem T) []T {
	//tmp := make([]T, len(a))
	var tmp []T
	for _, v := range a {
		if v != elem {
			tmp = append(tmp, v)
		}
	}
	return tmp
}

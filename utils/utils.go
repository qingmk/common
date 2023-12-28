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

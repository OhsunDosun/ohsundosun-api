package util

func DeleteDuplicateItem(arr []string) []string {
	ret := []string{}
	m := make(map[string]struct{})

	for _, val := range arr {
		if _, ok := m[val]; !ok {
			m[val] = struct{}{}
			ret = append(ret, val)
		}
	}

	return ret
}

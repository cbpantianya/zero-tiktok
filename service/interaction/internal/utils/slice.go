package utils

// DeleteSlice 删除指定元素
func DeleteSlice(a []int64, elem int64) []int64 {
	j := 0
	for _, v := range a {
		if v != elem {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

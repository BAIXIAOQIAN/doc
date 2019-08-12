package golang

//字符串数组去重1
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

//字符串数组去重2
func RemoveReperateStringArr(arr []string) (newArr []string) {
	newMap := make(map[string]int)

	for i := 0; i < len(arr); i++ {
		newMap[arr[i]] = i
	}

	for k, _ := range newMap {
		newArr = append(newArr, k)
	}
	return
}

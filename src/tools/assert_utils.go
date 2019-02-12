package tools

import "fmt"

// SliceEqual check []int equals
func SliceEqual(actual interface{}, expected ...interface{}) string {
	actualArr := actual.([]int)
	expectedArr := expected[0].([]int)
	if len(actualArr) != len(expectedArr) {
		return fmt.Sprintf("length not equal !!! acutal:%v, expected:%v", actualArr, expectedArr)
	}
	for index := 0; index < len(actualArr); index++ {
		if actualArr[index] != expectedArr[index] {
			return fmt.Sprintf("elements not equal !!! acutal:%v, expected:%v", actualArr, expectedArr)
		}
	}
	return ""
}

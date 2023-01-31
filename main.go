package main

import (
	"fmt"
	"linq/common"
	vta "linq/valuetypearray"
)

func main() {
	array := &vta.LinqArray[int]{}
	array.Init(15)
	array.AddItem(1)
	array.AddRange(34, 21, 3, 4, 55)
	array.AddRange(2, 3, 4)
	array.AddRange(6, 7, 8, 99)
	fmt.Println(fmt.Sprintf("Array List %v", array.ToArray()))
	fmt.Println(fmt.Sprintf("Value 1 contains %v", array.Contains(1)))
	fmt.Println(fmt.Sprintf("Value 8 contains %v", array.Contains(8)))
	max := common.Max(array)
	fmt.Println(fmt.Sprintf("Int Array Max %v", max))
	fmt.Println(fmt.Sprintf("Array size Before Remove %v", array.Size()))
	_ = array.RemoveITem(max)
	fmt.Println(fmt.Sprintf("Array List After Remove %v", array.ToArray()))
	fmt.Println(fmt.Sprintf("Array size After Remove %v", array.Size()))
	fmt.Println(fmt.Sprintf("Even number array list %v", array.Where(IsEvenNumber).ToArray()))
	fmt.Println(fmt.Sprintf("Distinct array list %v", array.Distinct().ToArray()))
	fmt.Println(fmt.Sprintf("Order by asc list %v", common.OrderBy(array.Distinct()).ToArray()))
	fmt.Println(fmt.Sprintf("Order by Desc list %v", common.OrderByDesc(array.Distinct()).ToArray()))
	array.RemoveDuplicates()
	fmt.Println(fmt.Sprintf("Removed duplicates list %v", array.ToArray()))
	fmt.Println(fmt.Sprintf("Int array Sum : %v", common.Sum(array)))
	floatArray := &vta.LinqArray[float32]{}
	floatArray.Init(10)
	floatArray.AddRange(2.4, 3.3, 4.2, 5.8)
	fmt.Println(fmt.Sprintf("Float Array List %v", floatArray.ToArray()))
	fmt.Println(fmt.Sprintf("Float Array Max %v", common.Max(floatArray)))
	fmt.Println(fmt.Sprintf("Float Array Min %v", common.Min(floatArray)))
	stringArray := &vta.LinqArray[string]{}
	stringArray.Init(10)
	stringArray.AddRange("2", "z", "er", "ert")
	fmt.Println(fmt.Sprintf("Float Array List %v", stringArray.ToArray()))
	fmt.Println(fmt.Sprintf("Float Sum :  %v Int Sum : %v", common.Sum(floatArray), common.Sum(array)))
	fmt.Println(fmt.Sprintf("Float Avg :  %v Int Avg : %v", common.Avg(floatArray), common.Avg(array)))
}

// IsEvenNumber ...
func IsEvenNumber(val int) bool {
	if val%2 == 0 {
		return true
	}
	return false
}

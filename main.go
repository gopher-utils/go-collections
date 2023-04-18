package main

import (
	"fmt"
	"linq/collections"
)

func main() {
	list := collections.NewEmptyList[int](15)
	list.Add(1)
	list.Extend(collections.NewListFromArray([]int{21, 3, 4, 55}))
	list.Extend(collections.NewListFromArray([]int{2, 3, 4}))
	list.Extend(collections.NewListFromArray([]int{6, 7, 8, 99}))
	fmt.Println(fmt.Sprintf("Array List %v", list.ToArray()))
	fmt.Println(fmt.Sprintf("List %v", list))
	fmt.Println(fmt.Sprintf("Value 1 is present? : %v", list.Contains(1)))
	fmt.Println(fmt.Sprintf("Value 10 is present? : %v", list.Contains(10)))
	max := collections.Max[int](list)
	fmt.Println(fmt.Sprintf("Int Array Max %v", max))
	fmt.Println(fmt.Sprintf("List size Before Remove %v", list.Size()))
	_ = list.RemoveFirst(max)
	fmt.Println(fmt.Sprintf("List After Remove %v", list))
	fmt.Println(fmt.Sprintf("List size After Remove %v", list.Size()))
	fmt.Println(fmt.Sprintf("Even numbers in list %v", list.Where(func(i int) bool {
		return i%2 == 0
	})))
	fmt.Println(fmt.Sprintf("Distinct list %v", list.Distinct()))
	list.RemoveDuplicates()
	fmt.Println(fmt.Sprintf("Removed duplicates list %v", list))
	fmt.Println(fmt.Sprintf("List Sum : %v", collections.Sum[int](list)))

	// Sample Float list example
	floatList := collections.NewEmptyList[float64](10)
	floatList.Extend(collections.NewListFromArray([]float64{2.4, 3.3, 4.2, 5.8}))
	fmt.Println(fmt.Sprintf("Float64 Array List %v", floatList.ToArray()))
	fmt.Println(fmt.Sprintf("Float64 Max %v", collections.Max[float64](floatList)))
	fmt.Println(fmt.Sprintf("Float64 Min %v", collections.Min[float64](floatList)))

	// Sample Struct list example
	nesl := collections.NewEmptyList[ExampleStruct](5)
	nesl.Add(ExampleStruct{Field1: 1, Field2: false})
	nesl.Add(ExampleStruct{Field1: 1, Field2: false})
	nesl.Add(ExampleStruct{Field1: 3, Field2: true})
	nesl.Add(ExampleStruct{Field1: 3, Field2: false})
	nesl.Add(ExampleStruct{Field1: 2, Field2: true})
	fmt.Println(fmt.Sprintf("Struct Distinct list %v", nesl))
	nesl.RemoveDuplicates()
	fmt.Println(fmt.Sprintf("Struct Distinct list %v", nesl))
	fmt.Println(fmt.Sprintf("Struct List where Field2 is false %v", nesl.Where(func(es ExampleStruct) bool {
		return !es.Field2
	})))
}

// An example struct that implements CollectionElement
type ExampleStruct struct {
	Field1 int
	Field2 bool
}

func (obj ExampleStruct) String() string {
	return fmt.Sprintf("{ field1: %d, field2: %t }", obj.Field1, obj.Field2)
}

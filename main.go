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
	fmt.Printf("Array List %v\n", list.ToArray())
	fmt.Printf("List %v\n", list)
	fmt.Printf("Value 1 is present? : %v\n", list.Contains(1))
	fmt.Printf("Value 10 is present? : %v\n", list.Contains(10))
	max := collections.Max[int](list)
	fmt.Printf("Int Array Max %v\n", max)
	fmt.Printf("List size Before Remove %v\n", list.Size())
	_ = list.RemoveFirst(max)
	fmt.Printf("List After Remove %v\n", list)
	fmt.Printf("List size After Remove %v\n", list.Size())
	fmt.Printf("Even numbers in list %v\n", list.Where(func(i int) bool {
		return i%2 == 0
	}))
	fmt.Printf("Distinct list %v\n", list.Distinct())
	list.RemoveDuplicates()
	fmt.Printf("Removed duplicates list %v\n", list)
	fmt.Printf("List Sum : %v\n", collections.Sum[int](list))

	// Sample Float list example
	floatList := collections.NewEmptyList[float64](10)
	floatList.Extend(collections.NewListFromArray([]float64{2.4, 3.3, 4.2, 5.8}))
	fmt.Printf("Float64 Array List %v\n", floatList.ToArray())
	fmt.Printf("Float64 Max %v\n", collections.Max[float64](floatList))
	fmt.Printf("Float64 Min %v\n", collections.Min[float64](floatList))

	// Sample Struct list example
	nesl := collections.NewEmptyList[ExampleStruct](5)
	nesl.Add(ExampleStruct{Field1: 1, Field2: false})
	nesl.Add(ExampleStruct{Field1: 1, Field2: false})
	nesl.Add(ExampleStruct{Field1: 3, Field2: true})
	nesl.Add(ExampleStruct{Field1: 3, Field2: false})
	nesl.Add(ExampleStruct{Field1: 2, Field2: true})
	fmt.Printf("Struct Distinct list %v\n", nesl)
	nesl.RemoveDuplicates()
	fmt.Printf("Struct Distinct list %v\n", nesl)
	fmt.Printf("Struct List where Field2 is false %v\n", nesl.Where(func(es ExampleStruct) bool {
		return !es.Field2
	}))
}

// An example struct that implements CollectionElement
type ExampleStruct struct {
	Field1 int
	Field2 bool
}

func (obj ExampleStruct) String() string {
	return fmt.Sprintf("{ field1: %d, field2: %t }", obj.Field1, obj.Field2)
}

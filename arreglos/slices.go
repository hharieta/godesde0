package arreglos

import "fmt"

func Slices() {
	var slice1 []int = []int{2, 3, 4, 5}
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice2 := array[3:]
	slice3 := array[:5]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func CapacidadSlices() {
	slice4 := make([]int, 5, 20)

	fmt.Printf("Slice: %d \nLongitud actual: %d Capacidad total: %d \n", slice4, len(slice4), cap(slice4))
}

func RellenarSlice() {
	slice5 := make([]int, 0, 0)

	for i := 0; i < 160; i++ {
		slice5 = append(slice5, i*2)
	}
	fmt.Printf("Slice: %d \nLongitud actual: %d Capacidad total: %d \n", slice5, len(slice5), cap(slice5))
}

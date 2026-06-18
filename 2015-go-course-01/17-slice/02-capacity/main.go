package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 10)

	fmt.Println("---------------------")
	fmt.Println(mySlice, "length:", len(mySlice), "capacity:", cap(mySlice))
	fmt.Println("---------------------")

	// neat - each time append exceeds the capacity of mySlice - the underlying array is doubled in size
	// so in this example; 10, 20, 40 ...
	// at somepoint (~512?) it switches to binary sensible increments - 1536, 2048, ...

	capacity := cap(mySlice)
	for i := 0; i < 1000000; i++ {
		mySlice = append(mySlice, i)
		if cap(mySlice) != capacity {
			fmt.Println(i, ":", "\tlength:", len(mySlice), "\tcapacity:", cap(mySlice))
			capacity = cap(mySlice)
		}
	}
	fmt.Println("---------------------")

	mySlice2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	mySlice2 = append(mySlice2[:4], mySlice2[5:]...)
	fmt.Println(mySlice2)

}

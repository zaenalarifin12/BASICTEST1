package main

import (
	"fmt"
)

func miniMaxSum(array [5]int) {

	var dataSum []int
	var sumMin, sumMax int;

	// for i := 0; i <= (len(array) - 2); i++ {
    //     sumMin += array[i]
    // }
	
	inc := 0
	sumTemp := 0;

	for i := 0; i <= (len(array) - 1); i++ {
	
		for j := 0; j <= (len(array) - 1); j++ {
		
			if (inc != j){
				sumTemp += array[j]
			}
		}


		// for initial
		sumMin = sumTemp
		sumMax = sumTemp

		dataSum = append(dataSum, sumTemp)
		inc = i+1
		sumTemp = 0

		// summ
	}

	for _, value := range dataSum {
        if value < sumMin {
            sumMin = value
        }
        if value > sumMax {
            sumMax = value
        }
    }

	fmt.Printf("all data sum = %d\n", dataSum)
	fmt.Printf("Minimum value: %d\n", sumMin)
    fmt.Printf("Maximum value: %d\n", sumMax)
}

func main()  {
	// integerArray := [5]int{1, 3, 5, 7, 9}
	integerArray2 := [5]int{1, 2, 3, 4, 5}

	// minimum sum and maximum sum
	
	miniMaxSum(integerArray2)


}
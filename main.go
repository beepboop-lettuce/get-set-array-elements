package main

import (
	"fmt"
)

func main() {

	names := [3]string{
		"Rachel",
		"Shavonda",
		"Jen",
	}

	print("names")

	ordinaryLoopString(names[:])

	distances := [5]int{
		14,
		6,
		12,
		87,
		9,
	}

	print("distances")

	ordinaryLoopInt(distances[:])

	data := [5]uint8{
		23,
		84,
		12,
		1,
		97,
	}

	print("data")

	ordinaryLoopUint(data[:])

	ratios := [1]float64{
		3.14,
	}

	print("ratios")

	ordinaryLoopFloat(ratios[:])

	alives := [4]bool{
		true,
		false,
		false,
		false,
	}

	print("alives")

	ordinaryLoopBool(alives[:])

	zero := [0]uint8{}

	print("zero")

	ordinaryLoopUint(zero[:])

	fmt.Printf(`
	
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
FOR RANGES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

`)

	print("names")

	for index, element := range names {
		fmt.Printf("names[%v]: %s\n", index, element)
	}

	print("distances")

	for index, element := range distances {
		fmt.Printf("distances[%v]: %d\n", index, element)
	}

	print("data")

	for index, element := range data {
		fmt.Printf("data[%v]: %d\n", index, element)
	}

	print("ratios")

	for index, element := range ratios {
		fmt.Printf("ratios[%v]: %.2f\n", index, element)
	}

	print("alives")

	for index, element := range alives {
		fmt.Printf("alives[%v]: %t\n", index, element)
	}

	print("zero")

	for index, element := range zero {
		fmt.Printf("zero[%v]: %d\n", index, element)
	}

}

func print(s string) {
	fmt.Printf(`
%s

=====================
`, s)

}

func ordinaryLoopString(a []string) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("ratios[%v]: %s\n", i, a[i])
	}
}

func ordinaryLoopInt(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("ratios[%v]: %d\n", i, a[i])
	}
}

func ordinaryLoopFloat(a []float64) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("ratios[%v]: %.2f\n", i, a[i])
	}
}

func ordinaryLoopBool(a []bool) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("ratios[%v]: %t\n", i, a[i])
	}
}

func ordinaryLoopUint(a []uint8) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("ratios[%v]: %d\n", i, a[i])
	}
}

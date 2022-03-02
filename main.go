package main

import (
	"fmt"
)

func main() {

	names := [3]string{
		"Rachel", "Shavonda", "Jen",
	}

	distances := [5]int{
		14, 6, 12, 87, 9,
	}

	data := [5]uint8{
		23, 84, 12, 1, 97,
	}

	ratios := [1]float64{
		3.14,
	}

	alives := [4]bool{
		true, false, false, false,
	}

	zero := [0]uint8{}

	print("names")
	ordinaryLoopString(names[:])

	print("distances")
	ordinaryLoopInt(distances[:])

	print("data")
	ordinaryLoopUint(data[:])

	print("ratios")
	ordinaryLoopFloat(ratios[:])

	print("alives")
	ordinaryLoopBool(alives[:])

	print("zero")
	ordinaryLoopUint(zero[:])

	fmt.Printf(`
	
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
FOR RANGES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

`)

	print("names")
	rangeString(names[:])

	print("distances")
	rangeInt(distances[:])

	print("data")
	rangeUint(data[:])

	print("ratios")
	rangeFloat(ratios[:])

	print("alives")
	rangeBool(alives[:])

	print("zero")
	rangeUint(zero[:])

}

func print(s string) {
	fmt.Printf(`
%s

=====================
`, s)

}

func ordinaryLoopString(a []string) {
	for i := 0; i < len(a); i++ {
		fmt.Print("names")
		fmt.Printf("[%v]: %s\n", i, a[i])
	}
}

func ordinaryLoopInt(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Print("distances")
		fmt.Printf("[%v]: %d\n", i, a[i])
	}
}

func ordinaryLoopFloat(a []float64) {
	for i := 0; i < len(a); i++ {
		fmt.Print("ratios")
		fmt.Printf("[%v]: %.2f\n", i, a[i])
	}
}

func ordinaryLoopBool(a []bool) {
	for i := 0; i < len(a); i++ {
		fmt.Print("alives")
		fmt.Printf("[%v]: %t\n", i, a[i])
	}
}

func ordinaryLoopUint(a []uint8) {
	for i := 0; i < len(a); i++ {
		fmt.Print("data")
		fmt.Printf("[%v]: %d\n", i, a[i])
	}
}

func rangeString(a []string) {
	for index, element := range a {
		fmt.Printf("[%v]: %s\n", index, element)
	}
}

func rangeInt(a []int) {
	for index, element := range a {
		fmt.Printf("[%v]: %d\n", index, element)
	}
}

func rangeFloat(a []float64) {
	for index, element := range a {
		fmt.Printf("[%v]: %.2f\n", index, element)
	}
}

func rangeBool(a []bool) {
	for index, element := range a {
		fmt.Printf("[%v]: %t\n", index, element)
	}
}

func rangeUint(a []uint8) {
	for index, element := range a {
		fmt.Printf("[%v]: %d\n", index, element)
	}
}

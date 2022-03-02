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

	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%v]: %s\n", i, names[i])
	}

	distances := [5]int{
		14,
		6,
		12,
		87,
		9,
	}

	print("distances")

	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%v]: %d\n", i, distances[i])
	}

	data := [5]uint8{
		23,
		84,
		12,
		1,
		97,
	}

	print("data")

	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%v]: %d\n", i, data[i])
	}

	ratios := [1]float64{
		3.14,
	}

	print("ratios")

	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%v]: %.2f\n", i, ratios[i])
	}

	alives := [4]bool{
		true,
		false,
		false,
		false,
	}

	print("alives")

	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%v]: %t\n", i, alives[i])
	}

	zero := [0]uint8{}

	print("zero")

	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%v]: %d\n", i, zero[i])
	}

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

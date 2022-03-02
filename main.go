package main

import "fmt"

func main() {

	names := [3]string{
		"Rachel",
		"Shavonda",
		"Jen",
	}

	fmt.Printf(`
names

=====================
`)

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

	fmt.Printf(`
distances

=====================
`)

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

	fmt.Printf(`
data

=====================
`)

	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%v]: %d\n", i, data[i])
	}

	ratios := [1]float64{
		3.14,
	}

	fmt.Printf(`
ratios

=====================
`)

	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%v]: %.2f\n", i, ratios[i])
	}

	alives := [4]bool{
		true,
		false,
		false,
		false,
	}

	fmt.Printf(`
alives

=====================
`)

	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%v]: %t\n", i, alives[i])
	}

	zero := [0]uint8{}

	fmt.Printf(`
zero

=====================
`)

	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%v]: %d\n", i, zero[i])
	}

	fmt.Printf(`
	
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
FOR RANGES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

`)

	fmt.Printf(`
names

=====================
`)

	for index, element := range names {
		fmt.Printf("names[%v]: %s\n", index, element)
	}

	fmt.Printf(`
distances

=====================
`)

	for index, element := range distances {
		fmt.Printf("distances[%v]: %d\n", index, element)
	}

	fmt.Printf(`
data

=====================
`)

	for index, element := range data {
		fmt.Printf("data[%v]: %d\n", index, element)
	}

	fmt.Printf(`
ratios

=====================
`)

	for index, element := range ratios {
		fmt.Printf("ratios[%v]: %.2f\n", index, element)
	}

	fmt.Printf(`
alives

=====================
`)

	for index, element := range alives {
		fmt.Printf("alives[%v]: %t\n", index, element)
	}

	fmt.Printf(`
zero

=====================
`)

	for index, element := range zero {
		fmt.Printf("zero[%v]: %d\n", index, element)
	}

}

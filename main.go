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

	/*	distances := [5]int{
			14,
			6,
			12,
			87,
			9,
		}

		data := [5]uint8{
			23,
			84,
			12,
			1,
			97,
		}

		ratios := [1]float64{
			3.14,
		}

		alives := [4]bool{
			true,
			false,
			false,
			false,
		}

		zero := [0]uint8{}
	*/
}

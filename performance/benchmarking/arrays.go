package main

func simpleArray() []int {

	a := []int{}

	for i := 0; i < 1250; i++ {
		a = append(a, i)
	}

	return a
}

func capacityArray() []int {

	l := 1250
	a := make([]int, 0, l)

	for i := 0; i < l; i++ {
		a = append(a, i)
	}

	return a
}

func lengthArray() []int {

	l := 1250
	a := make([]int, l)

	for i := 0; i < l; i++ {
		a[i] = i
	}

	return a
}

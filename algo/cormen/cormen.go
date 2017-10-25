package main

// Chapter 2 part 1 exercice 2: insertion-sort in decreasing order
func ex212(a []int) []int {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		// Insert a[j] into the sorted sequence
		for i>-1 && a[i]<key {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
	return a
}

// Chapter 2 part 1 exercise 3: linear search
func ex213(a []int, v int) int {
	for i, key := range a {
		if key == v {
			return i
		}
	}
	return -1
}

// Chapter 2 part 1 exercise 4: adding n-bit integers
func ex214(a []int, b []int) []int {
	if len(a) != len(b) {
		return []int{-1}
	}
	c := make([]int, len(a)+1)
	for i, _ := range a {
		tmp := a[i] + b[i] + c[i]
		switch tmp {
		case 0: fallthrough
		case 2:
			c[i] = 0
		case 1: fallthrough
		case 3:
			c[i] = 1
		}
		if tmp > 1 {
			c[i+1] = 1
		}
	}
	return c
}

func main() {
}

package main

func main() {

	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range x {
		if v%2 == 0 {
			println(v, " is even")
		} else {
			println(v, " is odd")
		}
	}
}

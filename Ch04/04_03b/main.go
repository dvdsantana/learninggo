package main

func main() {
	println("Loops")

	// colors := []string{"Red", "Green", "Blue"}

	// for i := 0; i < len(colors); i++ {
	// 	println(colors[i])
	// }

	// for i, color := range colors {
	// 	println(i, color)
	// }

	// for i := range colors {
	// 	println(i)
	// }

	// for _, color := range colors {
	// 	println(color)
	// }

	value := 0
	sum := 0
	for value < 10 {
		sum += value
		println(value)
		println(sum)
		value++
	}
	println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
		if sum > 100 {
			goto theEnd
		}
		println(sum)
	}

theEnd:
	println("The end")
}
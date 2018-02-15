package oppgaver

import "fmt"

func Oppgave2a() {
	list := []int{6, 14, 4, 18, 17, 1, 20, 7, 3, 15, 5, 2, 16, 10, 19, 8, 11, 13, 9, 12}
	fmt.Println("Usortert liste:")
	fmt.Println(list)
	fmt.Println("Sortert liste:")
	Bubble(list)
	fmt.Println(list)
}
func Bubble(list []int) {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(list)-1; i++ {
			if list[i+1] < list[i] {
				DoSwap(list, i, i+1)
				swap = true
			}
		}
	}
}

// External swap function for re-use
func DoSwap(list []int, i1, i2 int) {
	tmp := list[i1]
	list[i1] = list[i2]
	list[i2] = tmp
}

// Forklare bigO
func Oppgave2Counts() {
	list := []int{6, 14, 4, 18, 17, 1, 20, 7, 3, 15, 5, 2, 16, 10, 19, 8, 11, 13, 9, 12}
	BubbleBigO(list)
}
func BubbleBigO(list []int) {
	swapcount := 0
	counter := 0
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(list)-1; i++ {
			counter++
			if list[i+1] < list[i] {
				DoSwap(list, i, i+1)
				swapcount++
				swap = true
			}
		}
	}
	fmt.Println("Total checks:")
	fmt.Println(counter)
	fmt.Println("Total swaps:")
	fmt.Println(swapcount)
}

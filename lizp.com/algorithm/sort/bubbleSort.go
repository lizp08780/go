package sort

func BubbleSort(a []int) {
	length := len(a)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func BubbleSort2(a []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(a)-1; i++ {
			if a[i+1] < a[i] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
	}
}

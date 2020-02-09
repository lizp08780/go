package sort

func sortMax(a []int, length int) {
	if length <= 1 {
		return
	}
	depth := length/2 - 1
	for i := depth; i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		maxIndex := i
		if left < length && a[left] > a[maxIndex] {
			maxIndex = left
		}
		if right < length && a[right] > a[maxIndex] {
			maxIndex = right
		}
		if i != maxIndex {
			a[i], a[maxIndex] = a[maxIndex], a[i]
		}
	}

}

func HeapSort(a []int) {
	length := len(a)
	for i := length; i > 0; i-- {
		sortMax(a, i)
		a[0], a[i-1] = a[i-1], a[0]
	}
}

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

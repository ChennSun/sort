package sortplus

import (
	"sort"
)

// 插入排序
// 时间复杂度：O(n^2)
// 空间复杂度：
func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		// cursor为要和前面有序数组的中的元素对比的元素
		// a[i]的值可能会被向后移动的值覆盖，用变量暂存起来
		cursor := a[i]
		// cursor对比的起始位置
		j := i - 1
		// 如果元素大于cursor，则将这个元素向后移一位
		// j不断自减直到碰到一个元素小于cursor, 将cursor放到该元素后面即可。
		for ; j >= 0 && a[j] > cursor; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = cursor
	}
	return a
}

// 冒泡排序
// 时间复杂度：O(n^2) 计算方式 = (n-1) + (n-2) + (n-3) + .... + 0 = (n - 1) * n / 2
// 空间复杂度：O(1)
func BubbleSort(a []int) []int {
	// 控制循环次数
	for i := 0; i < len(a)-1; i++ {
		// 循环结束后，最大的元素会移动到右边
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				// 将大的元素交换到右边
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

// 归并排序 O(nlgn) 递归合并？
// 时间复杂度 O(nlgn)
// 空间复杂度 O(n)
/*
MERGE-SORT(A, p, r)
	1 ifp<r
		z q = L<P+r)/2
		3 MERGE-SORT(A, p, q)
		4 MERGE-SORT(A, q+l, r)
		5 MERGE(A, p, q, r)
*/
func MergeSort(a []int) []int {
	// 1. 拆分
	n := len(a)
	// 递归终止
	if n <= 1 {
		return a
	}
	// 开始递归！！！
	left := MergeSort(a[:n/2])
	right := MergeSort(a[n/2:])
	// 2. 排序
	// 3. 归并
	return merge(left, right)
}

// 排序 && 归并
func merge(a []int, b []int) (c []int) {
	for len(a) > 0 && len(b) > 0 {
		if a[0] > b[0] {
			c = append(c, b[0])
			b = b[1:]
		} else {
			c = append(c, a[0])
			a = a[1:]
		}
	}
	if len(a) > 0 {
		c = append(c, a...)
	} else {
		c = append(c, b...)
	}
	return
}

// 快速排序
// 时间复杂度 O(nlgn)
// 空间复杂度 O(1)
/*
QUICKSORT(A, p, r)
	1 if p<r
		2 q = PARTITION(A, p, r)
		3 QUICKSORT(A, p, q-1)
		4 QUICKSORT(A, q+1, r)
*/
func QuickSort(a []int, p, r int) []int {
	if p >= r {
		return a
	}
	q := partition(a, p, r)
	QuickSort(a, p, q-1)
	QuickSort(a, q+1, r)
	return a
}

func partition(a []int, p, r int) int {
	// a[r]的值可能会被覆盖，暂存起来
	x := a[r]
	// 下标i
	i := p - 1
	// 下标j
	// 以a[r]作为主元进行对比，将大于a[r]的都交换到左边
	// p <= 下标 <= i 的元素均小于 a[r]
	// 下标 = i+1的元素 等于 a[r]
	// i + 1 <= 下标 <= r 均大于 a[r]
	for j := p; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

// 计数排序
// 时间复杂度 O(n+k)
// 空间复杂度 O(n+k)
func CountingSort(a []int) []int {
	var k int
	for _, v := range a {
		if v > k {
			k = v
		}
	}
	// 数组下标从0开始，长度为k，下标最大值其实为 k - 1，所以数组长度需要为 k + 1
	b := make([]int, k+1)
	for _, v := range a {
		b[v]++
	}
	// 开始计数
	for i := 1; i < k+1; i++ {
		b[i] = b[i] + b[i-1]
	}
	c := make([]int, len(a))
	for j := len(a) - 1; j >= 0; j-- {
		// b中的计数从1开始，因此还原到数组下标需要减一
		c[b[a[j]]-1] = a[j]
		// 可能存在重复，所以每次取值后自减
		b[a[j]]--
	}
	return c
}

// 基数排序
// 时间复杂度 O(k*n)
// 空间复杂度 O(n+k)
func RadixSort(a []int) []int {
	var max int
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	// 获取最大数的位数
	var k int = 1
	for ; max >= 10; max /= 10 {
		k++
	}
	var radix = 1
	for i := 0; i < k; i++ {
		var b [10]int
		for _, v := range a {
			b[v/radix%10]++
		}
		for j := 1; j < 10; j++ {
			b[j] = b[j] + b[j-1]
		}
		c := make([]int, len(a))
		for j := len(a) - 1; j >= 0; j-- {
			c[b[a[j]/radix%10]-1] = a[j]
			b[a[j]/radix%10]--
		}
		a = c
		radix *= 10
	}
	return a
}

// 桶排序
// 时间复杂度 O()
// 空间复杂度 O()
func BucketSort(a []int) []int {
	var b [100][]int
	for _, v := range a {
		b[v/100] = append(b[v/100], v)
	}
	for i := 0; i < 100; i++ {
		sort.Ints(b[i])
	}
	var c = make([]int, 0, len(a))
	for _, v := range b {
		c = append(c, v...)
	}
	return c
}

// 希尔排序
func ShellSort(a []int) []int {
	return a
}

// 选择排序

// 堆排序

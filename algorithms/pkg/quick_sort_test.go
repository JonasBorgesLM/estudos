package pkg_test

import (
	"github.com/JonasBorgesLM/estudos/algorithms/pkg"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

// TestQuickSort sets up the test environment for the QuickSort function.
var _ = ginkgo.Describe("QuickSort", func() {

	ginkgo.Context("when sorting an array", func() {

		ginkgo.It("should return an empty array when input is empty", func() {
			arr := []int{}

			sortedArr := pkg.QuickSort(arr)

			gomega.Expect(sortedArr).To(gomega.Equal([]int{}))
		})

		ginkgo.It("should return the same array when it contains one element", func() {
			arr := []int{5}

			sortedArr := pkg.QuickSort(arr)

			gomega.Expect(sortedArr).To(gomega.Equal([]int{5}))
		})

		ginkgo.It("should return the same array when it is already sorted", func() {
			arr := []int{1, 2, 3, 4, 5}

			sortedArr := pkg.QuickSort(arr)

			gomega.Expect(sortedArr).To(gomega.Equal([]int{1, 2, 3, 4, 5}))
		})

		ginkgo.It("should sort an unsorted array", func() {
			arr := []int{5, 3, 4, 1, 2}

			sortedArr := pkg.QuickSort(arr)

			gomega.Expect(sortedArr).To(gomega.Equal([]int{1, 2, 3, 4, 5}))
		})

		ginkgo.It("should handle arrays with duplicates", func() {
			arr := []int{3, 1, 2, 1, 3}

			sortedArr := pkg.QuickSort(arr)

			gomega.Expect(sortedArr).To(gomega.Equal([]int{1, 1, 2, 3, 3}))
		})
	})
})

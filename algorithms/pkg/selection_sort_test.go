package pkg_test

import (
	"github.com/JonasBorgesLM/estudos/algorithms/pkg"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

// TestSelectionSort sets up the test environment for the SelectionSort function.
var _ = ginkgo.Describe("SelectionSort", func() {

	ginkgo.Context("when sorting an array", func() {

		ginkgo.It("should return an empty array when input is empty", func() {
			gomega.Expect(pkg.SelectionSort([]int{})).To(gomega.Equal([]int{}))
		})

		ginkgo.It("should return the same array when it contains one element", func() {
			gomega.Expect(pkg.SelectionSort([]int{5})).To(gomega.Equal([]int{5}))
		})

		ginkgo.It("should return the same array when it is already sorted", func() {
			gomega.Expect(pkg.SelectionSort([]int{1, 2, 3, 4, 5})).To(gomega.Equal([]int{1, 2, 3, 4, 5}))
		})

		ginkgo.It("should sort an unsorted array", func() {
			gomega.Expect(pkg.SelectionSort([]int{5, 3, 4, 1, 2})).To(gomega.Equal([]int{1, 2, 3, 4, 5}))
		})

		ginkgo.It("should handle arrays with duplicates", func() {
			gomega.Expect(pkg.SelectionSort([]int{3, 1, 2, 1, 3})).To(gomega.Equal([]int{1, 1, 2, 3, 3}))
		})

		ginkgo.It("should return an empty array when input is nil", func() {
			gomega.Expect(pkg.SelectionSort(nil)).To(gomega.Equal([]int{}))
		})
	})
})

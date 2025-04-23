package pkg_test

import (
	"github.com/JonasBorgesLM/estudos/algorithms/pkg"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

// TestBinarySearch sets up the test environment for the BinarySearch function.
var _ = ginkgo.Describe("Binary Search", func() {

	ginkgo.Context("when the array is sorted", func() {

		ginkgo.It("should return the index of the found element", func() {
			arr := []int{1, 2, 3, 4, 5}
			target := 3

			index, err := pkg.BinarySearch(arr, target)

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(index).To(gomega.Equal(2))
		})

		ginkgo.It("should return an error if the element is not found", func() {
			arr := []int{1, 2, 3, 4, 5}
			target := 6

			index, err := pkg.BinarySearch(arr, target)

			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal("target not found"))
			gomega.Expect(index).To(gomega.Equal(-1))
		})

		ginkgo.It("should return the index of the first occurrence of a repeated element", func() {
			arr := []int{1, 2, 2, 2, 3}
			target := 2

			index, err := pkg.BinarySearch(arr, target)

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(index).To(gomega.Equal(1))
		})

		ginkgo.It("should handle negative numbers", func() {
			arr := []int{-5, -3, -1, 0, 2, 4}
			target := -3

			index, err := pkg.BinarySearch(arr, target)

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(index).To(gomega.Equal(1))
		})
	})

	ginkgo.Context("when the array is empty", func() {

		ginkgo.It("should return an error", func() {
			arr := []int{}
			target := 1

			index, err := pkg.BinarySearch(arr, target)

			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal("array is empty"))
			gomega.Expect(index).To(gomega.Equal(-1))
		})
	})
})

package pkg_test

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

// The TestPKG function registers a fail handler and runs the test suite for the package named "PKG".
func TestPKG(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "PKG Suite")
}

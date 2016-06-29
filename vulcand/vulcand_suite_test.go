package vulcand_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVulcand(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vulcand Suite")
}

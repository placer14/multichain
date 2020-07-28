package acala_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAcala(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acala Suite")
}

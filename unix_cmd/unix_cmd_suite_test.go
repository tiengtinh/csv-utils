package unix_cmd_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUnixCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UnixCmd Suite")
}

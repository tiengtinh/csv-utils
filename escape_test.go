package main_test

import (
	. "github.com/tiengtinh/csv-escape"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Escape", func() {
	It("should be able to escape", func() {
		Expect(Escape("12,5")).To(Equal("12/&#44/5"))
		Expect(Escape("this is a test, really?")).To(Equal("this is a test/&#44/ really?"))
		Expect(Escape("this is a test\n just to be clear")).To(Equal("this is a test/&#013/ just to be clear"))
		Expect(Escape(";;;;")).To(Equal("/&#009//&#009//&#009//&#009/"))
		Expect(Escape(`first line
second line`)).To(Equal("first line/&#013/second line"))
	})
})

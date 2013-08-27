package matcher_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/matchers"
)

var _ = Describe("ContainElement", func() {
	Context("when passed a supported type", func() {
		Context("and expecting a non-matcher", func() {
			It("should do the right thing", func() {
				Ω([2]int{1, 2}).Should(ContainElement(2))
				Ω([2]int{1, 2}).ShouldNot(ContainElement(3))

				Ω([]int{1, 2}).Should(ContainElement(2))
				Ω([]int{1, 2}).ShouldNot(ContainElement(3))

				arr := make([]myCustomType, 2)
				arr[0] = myCustomType{s: "foo", n: 3, f: 2.0, arr: []string{"a", "b"}}
				arr[1] = myCustomType{s: "foo", n: 3, f: 2.0, arr: []string{"a", "c"}}
				Ω(arr).Should(ContainElement(myCustomType{s: "foo", n: 3, f: 2.0, arr: []string{"a", "b"}}))
				Ω(arr).ShouldNot(ContainElement(myCustomType{s: "foo", n: 3, f: 2.0, arr: []string{"b", "c"}}))
			})
		})

		Context("and expecting a matcher", func() {
			It("should pass each element through the matcher", func() {
				Ω([]int{1, 2, 3}).Should(ContainElement(BeNumerically(">=", 3)))
				Ω([]int{1, 2, 3}).ShouldNot(ContainElement(BeNumerically(">", 3)))
			})

			It("should fail if the matcher ever fails", func() {
				actual := []interface{}{1, 2, "3", 4}
				success, _, err := (&ContainElementMatcher{Element: BeNumerically(">=", 3)}).Match(actual)
				Ω(success).Should(BeFalse())
				Ω(err).Should(HaveOccured())
			})
		})
	})

	Context("when passed an unsupported type", func() {
		It("should error", func() {
			success, _, err := (&ContainElementMatcher{Element: 0}).Match(0)
			Ω(success).Should(BeFalse())
			Ω(err).Should(HaveOccured())

			success, _, err = (&ContainElementMatcher{Element: 0}).Match("abc")
			Ω(success).Should(BeFalse())
			Ω(err).Should(HaveOccured())

			success, _, err = (&ContainElementMatcher{Element: 0}).Match(map[string]int{"a": 1, "b": 2})
			Ω(success).Should(BeFalse())
			Ω(err).Should(HaveOccured())

			success, _, err = (&ContainElementMatcher{Element: 0}).Match(nil)
			Ω(success).Should(BeFalse())
			Ω(err).Should(HaveOccured())
		})
	})
})

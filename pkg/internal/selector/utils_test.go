package selector_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/fields"

	. "sigs.k8s.io/controller-runtime/pkg/internal/selector"
)

var _ = Describe("RequiresExactMatch function", func() {

	It("Returns false when the selector matches everything", func() {
		_, _, requiresExactMatch := RequiresExactMatch(fields.Everything())
		Expect(requiresExactMatch).To(BeFalse())
	})

	It("Returns false when the selector matches nothing", func() {
		_, _, requiresExactMatch := RequiresExactMatch(fields.Nothing())
		Expect(requiresExactMatch).To(BeFalse())
	})

	It("Returns false when the selector has the form key1!=val1", func() {
		_, _, requiresExactMatch := RequiresExactMatch(fields.ParseSelectorOrDie("key1!=val1"))
		Expect(requiresExactMatch).To(BeFalse())
	})

	It("Returns false when the selector has the form key1==val1,key2==val2", func() {
		_, _, requiresExactMatch := RequiresExactMatch(fields.ParseSelectorOrDie("key1==val1,key2==val2"))
		Expect(requiresExactMatch).To(BeFalse())
	})

	It("Returns true when the selector has the form key1==val1", func() {
		_, _, requiresExactMatch := RequiresExactMatch(fields.ParseSelectorOrDie("key1==val1"))
		Expect(requiresExactMatch).To(BeTrue())
	})

	It("Returns true when the selector has the form key1=val1", func() {
		_, _, requiresExactMatch := RequiresExactMatch(fields.ParseSelectorOrDie("key1=val1"))
		Expect(requiresExactMatch).To(BeTrue())
	})

	It("Returns true when the selector has the form key1==val1", func() {
		key, val, _ := RequiresExactMatch(fields.ParseSelectorOrDie("key1==val1"))
		Expect(requiresExactMatch).To(BeTrue())
	})
})

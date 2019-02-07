package manifest_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/happytobi/autopilot/manifest"
)

func TestManifestPartser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Manifest Testsuite")
}

var _ = Describe("Parse Manifest", func() {
	It("parses complete manifest", func() {
		document, err := ParseManifest("./fixtures/manifest.yml")
		Expect(err).ToNot(HaveOccurred())
	})
})

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
		manifest, err := ParseManifest("../fixtures/manifest.yml")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(manifest.ApplicationManifests[0].Name).Should(Equal("myApp"))
	})
})

var _ = Describe("Parse multi Application Manifest", func() {
	It("parses complete manifest", func() {
		manifest, err := ParseManifest("../fixtures/multiManifest.yml")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(manifest.ApplicationManifests[0].Name).Should(Equal("myApp"))
		Expect(manifest.ApplicationManifests[1].Name).Should(Equal("myApp2"))
	})
})

var _ = Describe("Parse invalid Application Manifest", func() {
	It("parses invalid manifest", func() {
		manifest, err := ParseManifest("../fixtures/invalidManifest.yml")
		Expect(err).ShouldNot(BeNil())
		Expect(manifest.ApplicationManifests).Should(BeNil())
	})
})

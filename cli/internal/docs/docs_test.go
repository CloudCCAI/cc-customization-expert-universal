package docs

import (
	"strings"
	"testing"
)

func TestReadEmbeddedDoc(t *testing.T) {
	content, err := Read("platform/object", "introduction")
	if err != nil {
		t.Fatal(err)
	}
	if content == "" {
		t.Fatal("object introduction doc is empty")
	}
}

func TestConfigDefaultsToDevGuide(t *testing.T) {
	content, err := Read("platform/config", "")
	if err != nil {
		t.Fatal(err)
	}
	if content == "" {
		t.Fatal("config devguide doc is empty")
	}
}

func TestPageComponentDocAliasesPluginDoc(t *testing.T) {
	content, err := Read("platform/pagecomponent", "devguide")
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(content, "pagecomponent", "cloudcc create pagecomponent") {
		t.Fatalf("pagecomponent doc did not expose canonical naming:\n%s", content[:min(len(content), 300)])
	}
}

func TestPluginDocIsNotPublic(t *testing.T) {
	if _, err := Read("plugin", "devguide"); err == nil {
		t.Fatal("plugin doc should not be exposed as a public module")
	}
	if _, err := Read("platform/plugin", "devguide"); err == nil {
		t.Fatal("platform/plugin doc should not be exposed as a public module")
	}
}

func TestFlatDocModulesAreNotPublic(t *testing.T) {
	if _, err := Read("object", "introduction"); err == nil {
		t.Fatal("flat object doc should not be exposed after two-level docs layout")
	}
}

func TestPlatformKnowledgeDocsAreEmbedded(t *testing.T) {
	for _, module := range []string{
		"platform/overview",
		"platform/capabilityMap",
		"platform/security",
		"platform/automation",
		"platform/dataModeling",
		"platform/integrationArchitecture",
		"platform/integrationPatterns",
		"platform/lowcodeHighcode",
		"platform/mobileCapabilities",
		"platform/almRelease",
	} {
		content, err := Read(module, "introduction")
		if err != nil {
			t.Fatalf("%s introduction should be embedded: %v", module, err)
		}
		if !strings.Contains(content, "CloudCC") {
			t.Fatalf("%s introduction should mention CloudCC", module)
		}
		devguide, err := Read(module, "devguide")
		if err != nil {
			t.Fatalf("%s devguide should be embedded: %v", module, err)
		}
		if devguide == "" {
			t.Fatalf("%s devguide is empty", module)
		}
	}
}

func TestMethodologyAndPlaybookDocsAreEmbedded(t *testing.T) {
	for _, module := range []string{
		"methodology/blueprint",
		"methodology/moduleDesign",
		"methodology/integrationDesign",
		"methodology/deliveryPlan",
		"playbooks/manufacturingCrm",
		"playbooks/serviceWorkOrder",
		"playbooks/miniProgramPortal",
	} {
		content, err := Read(module, "introduction")
		if err != nil {
			t.Fatalf("%s introduction should be embedded: %v", module, err)
		}
		if !strings.Contains(content, "CloudCC") {
			t.Fatalf("%s introduction should mention CloudCC", module)
		}
		devguide, err := Read(module, "devguide")
		if err != nil {
			t.Fatalf("%s devguide should be embedded: %v", module, err)
		}
		if devguide == "" {
			t.Fatalf("%s devguide is empty", module)
		}
	}
}

func TestModulesListsTwoLevelDocs(t *testing.T) {
	modules, err := Modules()
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(modules, "\n")
	for _, module := range []string{
		"platform/object",
		"methodology/blueprint",
		"playbooks/serviceWorkOrder",
	} {
		if !strings.Contains(joined, module) {
			t.Fatalf("Modules() should include %s; got:\n%s", module, joined)
		}
	}
}

func containsAll(content string, values ...string) bool {
	for _, value := range values {
		if !strings.Contains(content, value) {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

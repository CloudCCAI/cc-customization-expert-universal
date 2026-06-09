package templates

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteProjectUsesCanonicalLayout(t *testing.T) {
	tmp := t.TempDir()
	target := filepath.Join(tmp, "demo")
	if err := WriteProject(target, "Demo Project"); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{
		"cloudcc-cli.config.json",
		".gitignore",
		"frontend/package.json",
		"frontend/vue.config.js",
		"frontend/babel.config.js",
		"frontend/public/index.html",
		"frontend/src/main.js",
		"frontend/src/App.vue",
		"frontend/pagecomponents",
		"backend/classes",
		"backend/triggers",
		"backend/schedule",
		"backend/lib/ccopenapi-0.1.3.jar",
		"backend/lib/fastjson-1.2.83.jar",
		"backend/lib/reflections-0.9.12.jar",
		"sidecar",
	} {
		if _, err := os.Stat(filepath.Join(target, filepath.FromSlash(path))); err != nil {
			t.Fatalf("expected %s: %v", path, err)
		}
	}
	for _, path := range []string{
		".cloudcc",
		".claude",
		"doc",
		"package.json",
		"src",
		"public",
		"lib",
		"plugins",
	} {
		if _, err := os.Stat(filepath.Join(target, filepath.FromSlash(path))); !os.IsNotExist(err) {
			t.Fatalf("unexpected legacy path %s, statErr=%v", path, err)
		}
	}
	packageJSON, err := os.ReadFile(filepath.Join(target, "frontend", "package.json"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(packageJSON), "cloudcc-plugin") {
		t.Fatalf("package.json still contains legacy plugin naming: %s", packageJSON)
	}
	if !strings.Contains(string(packageJSON), "pagecomponent") {
		t.Fatalf("package.json should use pagecomponent naming: %s", packageJSON)
	}
	indexHTML, err := os.ReadFile(filepath.Join(target, "frontend", "public", "index.html"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(indexHTML), "Plugin Dev") || !strings.Contains(string(indexHTML), "PageComponent Dev") {
		t.Fatalf("index.html naming mismatch: %s", indexHTML)
	}
}

func TestWriteProjectRejectsNonEmptyTarget(t *testing.T) {
	tmp := t.TempDir()
	if err := os.WriteFile(filepath.Join(tmp, "existing.txt"), []byte("x"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := WriteProject(tmp, "demo"); err == nil {
		t.Fatal("expected non-empty target rejection")
	}
}

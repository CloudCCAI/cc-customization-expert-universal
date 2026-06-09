package modules

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"cloudcc-customization-expert-go/internal/jsonx"
)

func TestPageComponentCreateWritesCanonicalFiles(t *testing.T) {
	tmp := t.TempDir()
	var stderr bytes.Buffer
	if err := pageComponentCreate([]string{"cc-demo"}, &stderr, tmp); err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(tmp, "frontend", "pagecomponents", "cc-demo")
	for _, file := range []string{
		filepath.Join(dir, "cc-demo.vue"),
		filepath.Join(dir, "components", "HelloWorld.vue"),
		filepath.Join(dir, "config.json"),
	} {
		if _, err := os.Stat(file); err != nil {
			t.Fatalf("expected generated file %s: %v", file, err)
		}
	}
	cfg, err := jsonx.ReadObjectFile(filepath.Join(dir, "config.json"))
	if err != nil {
		t.Fatal(err)
	}
	if cfg["component"] != "component-cc-demo" {
		t.Fatalf("unexpected component: %#v", cfg["component"])
	}
}

func TestParsePageComponentVueData(t *testing.T) {
	content := `<template><div /></template>
<script>
export default {
  data() {
    return {
      componentInfo: {
        component: 'component-cc-demo',
        compName: "Demo",
        loadModel: "start"
      },
      propObj: { title: "hello" },
      events: {}
    };
  },
};
</script>`
	data, err := parsePageComponentVueData(content)
	if err != nil {
		t.Fatal(err)
	}
	info, _ := data["componentInfo"].(map[string]any)
	if info["component"] != "component-cc-demo" || info["loadModel"] != "start" {
		t.Fatalf("unexpected componentInfo: %#v", info)
	}
}

func TestCollectPageComponentDependencies(t *testing.T) {
	tmp := t.TempDir()
	root := filepath.Join(tmp, "frontend", "pagecomponents", "cc-demo")
	if err := os.MkdirAll(filepath.Join(root, "components"), 0755); err != nil {
		t.Fatal(err)
	}
	entry := filepath.Join(root, "cc-demo.vue")
	if err := os.WriteFile(entry, []byte(`<template><Child /></template>
<script>
import Child from "./components/Child.vue";
import helper from "./helper";
export default {};
</script>
<style>@import "./style.scss";</style>`), 0644); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		filepath.Join(root, "components", "Child.vue"): "<template><div /></template>",
		filepath.Join(root, "helper.js"):               "export default {};",
		filepath.Join(root, "style.scss"):              ".demo { color: red; }",
	}
	for file, content := range files {
		if err := os.WriteFile(file, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
	deps := collectPageComponentDependencies(entry, tmp)
	for _, key := range []string{
		"frontend/pagecomponents/cc-demo/cc-demo.vue",
		"frontend/pagecomponents/cc-demo/components/Child.vue",
		"frontend/pagecomponents/cc-demo/helper.js",
		"frontend/pagecomponents/cc-demo/style.scss",
	} {
		if _, ok := deps[key]; !ok {
			t.Fatalf("missing dependency %s in %#v", key, deps)
		}
	}
}

func TestPluginResourceIsNotAlias(t *testing.T) {
	tmp := t.TempDir()
	var stdout, stderr bytes.Buffer
	err := Handle("create", "plugin", []string{"cc-demo"}, &stdout, &stderr, tmp)
	if err == nil {
		t.Fatal("plugin resource should not be accepted as a pagecomponent alias")
	}
	if _, statErr := os.Stat(filepath.Join(tmp, "plugins")); !os.IsNotExist(statErr) {
		t.Fatalf("plugin alias should not create files, statErr=%v", statErr)
	}
	if _, statErr := os.Stat(filepath.Join(tmp, "frontend", "pagecomponents")); !os.IsNotExist(statErr) {
		t.Fatalf("plugin alias should not create pagecomponent files, statErr=%v", statErr)
	}
}

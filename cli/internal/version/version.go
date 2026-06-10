package version

import (
	"fmt"
	"io"
	"strings"
)

const Version = "2.0.4"
const CompatVersion = "2.5.3"

func IsVersionAction(action string) bool {
	switch action {
	case "version", "-version", "--version", "-v", "--v", "help", "-help", "--help", "-h", "--h", "changelog", "doctor", "docs", "stats":
		return true
	default:
		return false
	}
}

func Handle(action string, args []string, stdout io.Writer, stderr io.Writer) error {
	switch normalize(action) {
	case "get":
		fmt.Fprintf(stdout, "\ncloudcc-cli-go version: %s\ncompat cloudcc-cli version: %s\n\n", Version, CompatVersion)
	case "help":
		Help(stdout, stderr)
		return nil
	case "changelog":
		fmt.Fprintln(stderr)
		fmt.Fprintln(stderr, "CloudCC Go skill CLI")
		fmt.Fprintln(stderr, "- P0-P3 Go rewrite: doc/config/openapi/metadata/file publish foundations.")
		fmt.Fprintln(stderr, "- P4 partial: pagecomponent create/get/detail/pull/delete/doc plus publish via local vue-cli-service.")
		fmt.Fprintln(stderr, "- P4 pure-Go Vue build replacement, JSP migration, and full MCP tool registration are intentionally deferred.")
		fmt.Fprintln(stderr)
	case "doctor":
		fmt.Fprintln(stdout, "cloudcc doctor")
		fmt.Fprintln(stdout, "- runtime: go binary")
		fmt.Fprintf(stdout, "- version: %s\n", Version)
		fmt.Fprintln(stdout, "- node/npm: not required for P0-P3 commands")
		fmt.Fprintln(stdout, "- pagecomponent publish: requires project-local Vue CLI build toolchain")
	case "docs":
		fmt.Fprintln(stdout, "Use: cloudcc doc <layer>/<module> introduction|devguide")
	case "stats":
		fmt.Fprintln(stdout, "Command stats are not collected in the Go offline skill.")
	default:
		return fmt.Errorf("version: unsupported command %s", action)
	}
	return nil
}

func Help(stdout io.Writer, stderr io.Writer) int {
	fmt.Fprintln(stdout, "CloudCC CLI Go")
	fmt.Fprintln(stdout)
	fmt.Fprintln(stdout, "Usage:")
	fmt.Fprintln(stdout, "  cloudcc --version")
	fmt.Fprintln(stdout, "  cloudcc doc <layer>/<module> introduction|devguide")
	fmt.Fprintln(stdout, "  cloudcc get config [projectPath]")
	fmt.Fprintln(stdout, "  cloudcc use config <env> [projectPath]")
	fmt.Fprintln(stdout, "  cloudcc create project <name|.>")
	fmt.Fprintln(stdout, "  cloudcc <query|pageQuery|create|update|delete|upsert> openapi <projectPath> <encodedBodyJson> [isMcp]")
	fmt.Fprintln(stdout, "  cloudcc create pagecomponent <name>")
	fmt.Fprintln(stdout, "  cloudcc publish pagecomponent <name> [projectPath]")
	fmt.Fprintln(stdout, "  cloudcc get pagecomponent [projectPath] [compName]")
	fmt.Fprintln(stdout, "  cloudcc detail pagecomponent <name> [id] [projectPath]")
	fmt.Fprintln(stdout, "  cloudcc pull pagecomponent <nameOrId> [projectPath]")
	fmt.Fprintln(stdout, "  cloudcc delete pagecomponent <nameOrId> [projectPath]")
	fmt.Fprintln(stdout)
	fmt.Fprintln(stdout, "Naming: pagecomponent is the only supported resource name for CloudCC page custom components.")
	fmt.Fprintln(stdout, "Deferred: pure-Go Vue build replacement, JSP migration, full MCP tool registration.")
	return 0
}

func normalize(action string) string {
	switch strings.TrimSpace(action) {
	case "version", "-version", "--version", "-v", "--v":
		return "get"
	case "help", "-help", "--help", "-h", "--h":
		return "help"
	default:
		return action
	}
}

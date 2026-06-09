package command

import (
	"fmt"
	"io"
	"os"

	"cloudcc-customization-expert-go/internal/docs"
	"cloudcc-customization-expert-go/internal/modules"
	"cloudcc-customization-expert-go/internal/openapi"
	"cloudcc-customization-expert-go/internal/version"
)

func Run(args []string, stdout io.Writer, stderr io.Writer, cwd string) int {
	if cwd == "" {
		var err error
		cwd, err = os.Getwd()
		if err != nil {
			fmt.Fprintln(stderr, "cloudcc: cannot resolve current directory:", err)
			return 1
		}
	}
	if len(args) == 0 {
		return version.Help(stdout, stderr)
	}
	if len(args) == 1 {
		if err := version.Handle(args[0], nil, stdout, stderr); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		return 0
	}

	action, resource := args[0], args[1]
	rest := args[2:]

	if version.IsVersionAction(action) && resource != "version" {
		if err := version.Handle(action, append([]string{resource}, rest...), stdout, stderr); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		return 0
	}

	var err error
	switch {
	case action == "doc":
		err = docs.Print(resource, firstOr(rest, ""), stdout)
	case resource == "openapi":
		err = openapi.Handle(action, rest, stdout, stderr, cwd)
	default:
		err = modules.Handle(action, resource, rest, stdout, stderr, cwd)
	}
	if err != nil {
		fmt.Fprintln(stderr)
		fmt.Fprintln(stderr, "请查看帮助：cloudcc --help", err)
		fmt.Fprintln(stderr)
		return 1
	}
	return 0
}

func firstOr(values []string, fallback string) string {
	if len(values) == 0 {
		return fallback
	}
	return values[0]
}

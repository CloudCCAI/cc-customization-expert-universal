package modules

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"cloudcc-customization-expert-go/internal/config"
	"cloudcc-customization-expert-go/internal/httpclient"
	"cloudcc-customization-expert-go/internal/jsonx"
	projecttemplates "cloudcc-customization-expert-go/internal/templates"
)

type endpoint struct {
	base string
	path string
}

var genericEndpoints = map[string]map[string]endpoint{
	"application": {
		"get":    {"setup", "/api/application/list"},
		"create": {"setup", "/api/application/save"},
		"update": {"setup", "/api/application/save"},
		"delete": {"setup", "/api/application/delete"},
	},
	"button": {
		"get":    {"setup", "/api/button/list"},
		"create": {"setup", "/api/button/save"},
		"update": {"setup", "/api/button/save"},
		"delete": {"setup", "/api/button/delete"},
	},
	"customSetting": {
		"get":    {"setup", "/api/customsetting/list"},
		"detail": {"setup", "/api/customsetting/detail"},
		"create": {"setup", "/api/customsetting/save"},
		"delete": {"setup", "/api/customsetting/delete"},
		"modify": {"setup", "/api/customsetting/modify"},
	},
	"dupeCatcher": {
		"get":    {"setup", "/api/duplication/list"},
		"detail": {"setup", "/api/duplication/detail"},
		"create": {"setup", "/api/duplication/save"},
		"delete": {"setup", "/api/duplication/delete"},
	},
	"fields": {
		"detail": {"setup", "/api/fieldSetup/detail"},
		"create": {"setup", "/api/fieldSetup/save"},
		"update": {"setup", "/api/fieldSetup/save"},
		"delete": {"setup", "/api/fieldSetup/delete"},
	},
	"globalSelectList": {
		"get":    {"setup", "/api/globalSelectList/list"},
		"detail": {"setup", "/api/globalSelectList/detail"},
		"create": {"setup", "/api/globalSelectList/save"},
		"delete": {"setup", "/api/globalSelectList/delete"},
	},
	"identityProvider": {
		"get":      {"setup", "/api/identityProvider/list"},
		"create":   {"setup", "/api/identityProvider/save"},
		"delete":   {"setup", "/api/identityProvider/delete"},
		"download": {"setup", "/api/identityProvider/download"},
	},
	"menu": {
		"get":    {"setup", "/api/menu/list"},
		"create": {"setup", "/api/menu/save"},
		"update": {"setup", "/api/menu/save"},
		"delete": {"setup", "/api/menu/delete"},
	},
	"pagelayout": {
		"get":    {"setup", "/api/layout/list"},
		"create": {"setup", "/api/layout/save"},
		"delete": {"setup", "/api/layout/delete"},
	},
	"permission": {
		"get":    {"setup", "/api/permission/list"},
		"assign": {"setup", "/api/permission/assign"},
		"add":    {"setup", "/api/permission/add"},
		"remove": {"setup", "/api/permission/remove"},
	},
	"profile": {
		"get":    {"setup", "/api/profile/list"},
		"create": {"setup", "/api/profile/save"},
		"delete": {"setup", "/api/profile/delete"},
	},
	"recordType": {
		"get":         {"api", "/api/batch/getRecordType"},
		"getList":     {"setup", "/api/recordtype/list"},
		"newInfo":     {"setup", "/api/recordtype/newInfo"},
		"create":      {"setup", "/api/recordtype/save"},
		"editInfo":    {"setup", "/api/recordtype/editInfo"},
		"editSave":    {"setup", "/api/recordtype/save"},
		"validDelete": {"setup", "/api/recordtype/validDelete"},
		"delete":      {"setup", "/api/recordtype/delete"},
	},
	"role": {
		"get":    {"setup", "/api/role/list"},
		"create": {"setup", "/api/role/save"},
		"delete": {"setup", "/api/role/delete"},
	},
	"scheduleJob": {
		"get":     {"setup", "/api/scheduleJob/list"},
		"getList": {"setup", "/api/scheduleJob/list"},
		"detail":  {"setup", "/api/scheduleJob/detail"},
		"create":  {"setup", "/api/scheduleJob/save"},
		"delete":  {"setup", "/api/scheduleJob/delete"},
	},
	"singleSignOn": {
		"get":    {"setup", "/api/singleSignOn/list"},
		"delete": {"setup", "/api/singleSignOn/delete"},
	},
	"user": {
		"get":    {"setup", "/api/user/list"},
		"view":   {"setup", "/api/user/view"},
		"create": {"setup", "/api/user/save"},
		"update": {"setup", "/api/user/save"},
		"delete": {"setup", "/api/user/delete"},
	},
	"validationRule": {
		"get":    {"setup", "/api/validationRule/list"},
		"create": {"setup", "/api/validationRule/save"},
		"delete": {"setup", "/api/validationRule/delete"},
	},
	"view": {
		"get":    {"setup", "/api/view/list"},
		"update": {"setup", "/api/view/save"},
	},
	"sharingRule": {
		"get": {"setup", "/api/sharingRule/list"},
	},
}

func Handle(action string, resource string, args []string, stdout io.Writer, stderr io.Writer, cwd string) error {
	switch resource {
	case "config":
		return handleConfig(action, args, stdout, cwd)
	case "object":
		return handleObject(action, args, stdout, stderr, cwd)
	case "brief":
		return handleBrief(action, args, stdout, cwd)
	case "fields":
		if action == "get" {
			return handleFieldsGet(args, stdout, cwd)
		}
	case "classes":
		return handleCodeResource(action, resource, "classes", "ccfag", args, stdout, stderr, cwd)
	case "triggers":
		return handleTrigger(action, args, stdout, stderr, cwd)
	case "timer", "schedule":
		return handleCodeResource(action, "timer", "schedule", "ccPeak", args, stdout, stderr, cwd)
	case "script":
		return handleScript(action, args, stdout, stderr, cwd)
	case "html":
		return handleHTML(action, args, stdout, stderr, cwd)
	case "staticResource":
		return handleStaticResource(action, args, stdout, stderr, cwd)
	case "pagecomponent":
		return handlePageComponent(action, resource, args, stdout, stderr, cwd)
	case "project":
		return handleProject(action, args, stderr, cwd)
	case "jsp", "site", "skill":
		return fmt.Errorf("%s %s is deferred to P4 or a later task in the Go rewrite", action, resource)
	}
	if byAction, ok := genericEndpoints[resource]; ok {
		if ep, ok := byAction[action]; ok {
			return callGeneric(ep, action, resource, args, stdout, cwd)
		}
	}
	return fmt.Errorf("unsupported command: cloudcc %s %s", action, resource)
}

func handleProject(action string, args []string, stderr io.Writer, cwd string) error {
	if action != "create" {
		return fmt.Errorf("unsupported project action: %s", action)
	}
	if len(args) < 1 || strings.TrimSpace(args[0]) == "" {
		return fmt.Errorf("cloudcc create project <name|.>")
	}
	target := strings.TrimSpace(args[0])
	if target == "." {
		target = cwd
	} else if !filepath.IsAbs(target) {
		target = filepath.Join(cwd, target)
	}
	if err := projecttemplates.WriteProject(target, filepath.Base(target)); err != nil {
		return err
	}
	fmt.Fprintf(stderr, "Created CloudCC project: %s\n", target)
	return nil
}

func handleConfig(action string, args []string, stdout io.Writer, cwd string) error {
	projectPath := firstArg(args, cwd)
	switch action {
	case "get":
		cfg, err := config.Load(projectPath)
		if err != nil {
			return err
		}
		return printJSON(stdout, cfg)
	case "use":
		if len(args) == 0 {
			return fmt.Errorf("cloudcc use config <env> [projectPath]")
		}
		projectPath = cwd
		if len(args) > 1 {
			projectPath = args[1]
		}
		return config.Use(projectPath, args[0])
	default:
		return fmt.Errorf("unsupported config action: %s", action)
	}
}

func callGeneric(ep endpoint, action string, resource string, args []string, stdout io.Writer, cwd string) error {
	projectPath := firstArg(args, cwd)
	body := map[string]any{}
	if len(args) > 1 && args[1] != "" {
		parsed, err := jsonx.ParseEncodedObject(args[1], action+" "+resource)
		if err != nil {
			if action == "delete" {
				body = map[string]any{"id": args[1], resource + "id": args[1]}
			} else {
				return err
			}
		} else {
			body = parsed
		}
	}
	cfg, err := config.Load(projectPath)
	if err != nil {
		return err
	}
	return postClass(stdout, cfg, ep.base, ep.path, body)
}

func handleBrief(action string, args []string, stdout io.Writer, cwd string) error {
	if action != "get" {
		return fmt.Errorf("unsupported brief action: %s", action)
	}
	cfg, err := config.Load(firstArg(args, cwd))
	if err != nil {
		return err
	}
	var res map[string]any
	if err := httpclient.New().PostClass(config.String(cfg, "setupSvc")+"/api/customObject/newPage", map[string]any{"id": ""}, config.String(cfg, "accessToken"), &res); err != nil {
		return err
	}
	if data, _ := res["data"].(map[string]any); data != nil {
		if list, ok := data["objList"]; ok {
			return printJSON(stdout, list)
		}
	}
	return printJSON(stdout, res)
}

func handleObject(action string, args []string, stdout io.Writer, stderr io.Writer, cwd string) error {
	switch action {
	case "get":
		return objectGet(args, stdout, cwd)
	case "create":
		return objectCreate(args, stdout, stderr, cwd)
	case "delete":
		if len(args) < 2 {
			return fmt.Errorf("cloudcc delete object <projectPath> <objid>")
		}
		cfg, err := config.Load(args[0])
		if err != nil {
			return err
		}
		return postClass(stdout, cfg, "setup", "/api/customObject/deleteLogic", map[string]any{"objid": args[1]})
	default:
		return fmt.Errorf("unsupported object action: %s", action)
	}
}

func objectGet(args []string, stdout io.Writer, cwd string) error {
	projectPath := firstArg(args, cwd)
	kind := ""
	if len(args) > 1 {
		kind = args[1]
	}
	cfg, err := config.Load(projectPath)
	if err != nil {
		return err
	}
	client := httpclient.New()
	var out []map[string]any
	if kind == "" || kind == "chat" || kind == "standard" {
		var res map[string]any
		if err := client.PostClass(config.String(cfg, "setupSvc")+"/api/customObject/standardObjList", map[string]any{}, config.String(cfg, "accessToken"), &res); err != nil {
			return err
		}
		if list, ok := res["data"].([]any); ok {
			for _, item := range list {
				if m, _ := item.(map[string]any); m != nil {
					out = append(out, map[string]any{
						"id":              m["id"],
						"label":           firstAny(m["objname"], m["label"]),
						"objname":         m["label"],
						"objprefix":       m["objprefix"],
						"schemetableName": m["label"],
						"apiname":         m["label"],
					})
				}
			}
		}
	}
	if kind == "" || kind == "chat" || kind == "custom" {
		var res map[string]any
		if err := client.PostClass(config.String(cfg, "setupSvc")+"/api/customObject/list", map[string]any{}, config.String(cfg, "accessToken"), &res); err != nil {
			return err
		}
		if data, _ := res["data"].(map[string]any); data != nil {
			if list, ok := data["objList"].([]any); ok {
				for _, item := range list {
					if m, _ := item.(map[string]any); m != nil {
						out = append(out, map[string]any{
							"id":              m["id"],
							"label":           firstAny(m["objLabel"], m["label"]),
							"objname":         firstAny(m["schemetable_name"], m["objname"]),
							"objprefix":       m["prefix"],
							"schemetableName": m["schemetableName"],
							"apiname":         m["schemetableName"],
						})
					}
				}
			}
		}
	}
	if kind == "trigger" {
		var res map[string]any
		if err := client.PostClass(config.String(cfg, "setupSvc")+"/api/trigger/newobjtrigger", map[string]any{}, config.String(cfg, "accessToken"), &res); err != nil {
			return err
		}
		if data, _ := res["data"].(map[string]any); data != nil {
			if list, ok := data["targetObjList"].([]any); ok {
				for _, item := range list {
					if m, _ := item.(map[string]any); m != nil {
						out = append(out, m)
					}
				}
			}
		}
	}
	return printJSON(stdout, out)
}

func objectCreate(args []string, stdout io.Writer, stderr io.Writer, cwd string) error {
	if len(args) < 3 {
		return fmt.Errorf("cloudcc create object <projectPath> <label> [nameLabel] <businessDescription>")
	}
	projectPath, label := args[0], args[1]
	nameLabel, remarkSuffix := "", ""
	if len(args) == 3 {
		nameLabel = labelToSlug(label) + "_custom_object"
		remarkSuffix = args[2]
	} else {
		nameLabel = args[2]
		remarkSuffix = args[3]
	}
	cfg, err := config.Load(projectPath)
	if err != nil {
		return err
	}
	profiles, err := loadProfileIDs(cfg)
	if err != nil {
		return err
	}
	perms := make([]map[string]any, 0, len(profiles))
	for i, id := range profiles {
		perm := "1,1,1,1,0,0"
		if i == 0 {
			perm = "1,1,1,1,1,1"
		}
		perms = append(perms, map[string]any{"profileid": id, "permission": perm})
	}
	body := map[string]any{
		"iscreatperm":     "true",
		"profileFieldArr": perms,
		"nameLabel":       nameLabel,
		"obj": map[string]any{
			"label":           label,
			"schemetableName": nameLabel,
			"dataType":        "V",
			"showFormat":      "{yyyy}{mm}{dd}{000}",
			"beginIndex":      "0",
			"isquickcreated":  "false",
			"islbs":           "false",
			"isreportcreated": "false",
			"remark":          fmt.Sprintf("用于管理与「%s」相关的业务数据 %s", label, strings.TrimSpace(remarkSuffix)),
		},
	}
	fmt.Fprintln(stderr, "Creating, please wait...")
	return postClass(stdout, cfg, "setup", "/api/customObject/saveButton", body)
}

func loadProfileIDs(cfg config.Config) ([]string, error) {
	var res map[string]any
	if err := httpclient.New().PostClass(config.String(cfg, "setupSvc")+"/api/customObject/newPage", map[string]any{"id": ""}, config.String(cfg, "accessToken"), &res); err != nil {
		return nil, err
	}
	var ids []string
	if data, _ := res["data"].(map[string]any); data != nil {
		if list, ok := data["objList"].([]any); ok {
			for _, item := range list {
				if m, _ := item.(map[string]any); m != nil {
					if id := fmt.Sprint(m["id"]); id != "" && id != "<nil>" {
						ids = append(ids, id)
					}
				}
			}
		}
	}
	return ids, nil
}

func handleFieldsGet(args []string, stdout io.Writer, cwd string) error {
	if len(args) < 2 {
		return fmt.Errorf("cloudcc get fields <projectPath> <prefix>")
	}
	cfg, err := config.Load(args[0])
	if err != nil {
		return err
	}
	return postClass(stdout, cfg, "setup", "/api/fieldSetup/queryField", map[string]any{"prefix": args[1]})
}

func handleCodeResource(action string, resource string, dir string, apiName string, args []string, stdout io.Writer, stderr io.Writer, cwd string) error {
	switch action {
	case "create":
		return createJavaResource(dir, resource, args, stderr, cwd)
	case "publish":
		return publishJavaResource(dir, apiName, args, stdout, stderr, cwd)
	case "get", "pullList":
		projectPath := firstArg(args, cwd)
		cfg, err := config.Load(projectPath)
		if err != nil {
			return err
		}
		listPath := "/api/" + apiName + "/list"
		if action == "pullList" {
			listPath = "/api/" + apiName + "/detail"
		}
		body := map[string]any{}
		if len(args) > 1 {
			body["name"] = args[1]
		}
		return postClass(stdout, cfg, "setup", listPath, body)
	case "detail", "pull":
		if len(args) < 2 {
			return fmt.Errorf("cloudcc %s %s <projectPath> <name>", action, resource)
		}
		cfg, err := config.Load(args[0])
		if err != nil {
			return err
		}
		return postClass(stdout, cfg, "setup", "/api/"+apiName+"/detail", map[string]any{"name": args[1]})
	case "delete":
		if len(args) < 2 {
			return fmt.Errorf("cloudcc delete %s <projectPath> <id>", resource)
		}
		cfg, err := config.Load(args[0])
		if err != nil {
			return err
		}
		return postClass(stdout, cfg, "setup", "/api/"+apiName+"/delete", map[string]any{"id": args[1]})
	default:
		return fmt.Errorf("unsupported %s action: %s", resource, action)
	}
}

func backendResourcePath(projectPath string, dir string, parts ...string) string {
	all := append([]string{projectPath, "backend", dir}, parts...)
	return filepath.Join(all...)
}

func handleTrigger(action string, args []string, stdout io.Writer, stderr io.Writer, cwd string) error {
	switch action {
	case "create":
		return createJavaResource("triggers", "triggers", args, stderr, cwd)
	case "publish":
		if len(args) < 1 {
			return fmt.Errorf("cloudcc publish triggers <object/name>")
		}
		namePath := args[0]
		name := filepath.Base(namePath)
		srcDir := backendResourcePath(cwd, "triggers", namePath)
		source, err := readMarkedSource(filepath.Join(srcDir, name+".java"))
		if err != nil {
			return err
		}
		cfgContent, _ := jsonx.ReadObjectFile(filepath.Join(srcDir, "config.json"))
		cfg, err := config.Load(cwd)
		if err != nil {
			return err
		}
		body := map[string]any{
			"id":             configID(cfgContent),
			"apiname":        cfgContent["apiname"],
			"isactive":       cfgContent["isactive"],
			"targetObjectId": cfgContent["targetObjectId"],
			"triggerTime":    cfgContent["triggerTime"],
			"name":           name,
			"version":        firstAny(cfgContent["version"], "2"),
			"triggerSource":  url.PathEscape(strings.TrimSpace(source)),
			"folderId":       "wgd",
		}
		return postClass(stdout, cfg, "setup", "/api/triggerSetup/saveTrigger", body)
	default:
		return handleCodeResource(action, "triggers", "triggers", "trigger", args, stdout, stderr, cwd)
	}
}

func createJavaResource(dir string, resource string, args []string, stderr io.Writer, cwd string) error {
	if len(args) < 1 {
		return fmt.Errorf("cloudcc create %s <name>", resource)
	}
	name := filepath.Base(args[0])
	target := backendResourcePath(cwd, dir, args[0])
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}
	pkg := dir
	if dir == "schedule" {
		pkg = "schedule"
	}
	source := fmt.Sprintf("package %s.%s;\n\nimport com.cloudcc.core.*;\n// @SOURCE_CONTENT_START\npublic class %s {\n    public void execute() {\n        // TODO: implement business logic\n    }\n}\n// @SOURCE_CONTENT_END\n", pkg, name, name)
	if resource == "timer" {
		source = fmt.Sprintf("package schedule.%s;\n\nimport com.cloudcc.core.*;\n\npublic class %s extends CCSchedule {\n    public %s() {\n        // @SOURCE_CONTENT_START\n        // TODO: implement schedule logic\n        // @SOURCE_CONTENT_END\n    }\n}\n", name, name, name)
	}
	if err := os.WriteFile(filepath.Join(target, name+".java"), []byte(source), 0644); err != nil {
		return err
	}
	test := fmt.Sprintf("package %s.%s;\n\npublic class %sTest {\n}\n", pkg, name, name)
	_ = os.WriteFile(filepath.Join(target, name+"Test.java"), []byte(test), 0644)
	_ = jsonx.WriteObjectFile(filepath.Join(target, "config.json"), map[string]any{"name": name, "version": "2"})
	fmt.Fprintf(stderr, "Created %s resource: %s\n", resource, target)
	return nil
}

func publishJavaResource(dir string, apiName string, args []string, stdout io.Writer, stderr io.Writer, cwd string) error {
	if len(args) < 1 {
		return fmt.Errorf("cloudcc publish <resource> <name>")
	}
	name := filepath.Base(args[0])
	srcDir := backendResourcePath(cwd, dir, args[0])
	source, err := readMarkedSource(filepath.Join(srcDir, name+".java"))
	if err != nil {
		return err
	}
	cfgContent, _ := jsonx.ReadObjectFile(filepath.Join(srcDir, "config.json"))
	cfg, err := config.Load(cwd)
	if err != nil {
		return err
	}
	endpoint := "/api/" + apiName + "/save"
	body := map[string]any{
		"id":       configID(cfgContent),
		"name":     name,
		"source":   url.PathEscape(strings.TrimSpace(source)),
		"version":  firstAny(cfgContent["version"], "2"),
		"folderId": "wgd",
	}
	fmt.Fprintln(stderr, "Posting, please wait...")
	return postClass(stdout, cfg, "setup", endpoint, body)
}

func handleScript(action string, args []string, stdout io.Writer, stderr io.Writer, cwd string) error {
	switch action {
	case "create":
		if len(args) < 1 {
			return fmt.Errorf("cloudcc create script <object/name>")
		}
		name := filepath.Base(args[0])
		target := filepath.Join(cwd, "script", args[0])
		if err := os.MkdirAll(target, 0755); err != nil {
			return err
		}
		body := "function main($CCDK, obj) {\n    // TODO: implement client script\n}\n"
		if err := os.WriteFile(filepath.Join(target, name+".js"), []byte(body), 0644); err != nil {
			return err
		}
		return jsonx.WriteObjectFile(filepath.Join(target, "config.json"), map[string]any{"name": name})
	case "publish":
		if len(args) < 1 {
			return fmt.Errorf("cloudcc publish script <object/name>")
		}
		name := filepath.Base(args[0])
		srcDir := filepath.Join(cwd, "script", args[0])
		content, err := os.ReadFile(filepath.Join(srcDir, name+".js"))
		if err != nil {
			return err
		}
		bodyText, err := extractMainBody(string(content))
		if err != nil {
			return err
		}
		local, _ := jsonx.ReadObjectFile(filepath.Join(srcDir, "config.json"))
		local["scriptContent"] = strings.TrimSpace(bodyText)
		cfg, err := config.Load(cwd)
		if err != nil {
			return err
		}
		var res map[string]any
		if err := httpclient.New().PostEnvelope(baseURL(cfg)+"/devconsole/script/saveClientScript", local, map[string]any(cfg), &res); err != nil {
			return err
		}
		return printJSON(stdout, res)
	case "get", "pull", "pullList", "detail", "delete":
		return callDevConsoleScript(action, args, stdout, cwd)
	default:
		return fmt.Errorf("unsupported script action: %s", action)
	}
}

func callDevConsoleScript(action string, args []string, stdout io.Writer, cwd string) error {
	projectPath := firstArg(args, cwd)
	body := map[string]any{}
	if len(args) > 1 {
		parsed, err := jsonx.ParseEncodedObject(args[1], action+" script")
		if err == nil {
			body = parsed
		}
	}
	cfg, err := config.Load(projectPath)
	if err != nil {
		return err
	}
	var res map[string]any
	if err := httpclient.New().PostEnvelope(baseURL(cfg)+"/devconsole/script/pageClientScript", body, map[string]any(cfg), &res); err != nil {
		return err
	}
	return printJSON(stdout, res)
}

func handleHTML(action string, args []string, stdout io.Writer, stderr io.Writer, cwd string) error {
	switch action {
	case "create":
		if len(args) < 1 {
			return fmt.Errorf("cloudcc create html <apiName>")
		}
		apiName := args[0]
		target := filepath.Join(cwd, "html", apiName)
		if err := os.MkdirAll(target, 0755); err != nil {
			return err
		}
		_ = os.WriteFile(filepath.Join(target, "index.html"), []byte("<!doctype html>\n<html><body></body></html>\n"), 0644)
		return jsonx.WriteObjectFile(filepath.Join(target, "config.json"), map[string]any{"apiName": apiName, "htmlLabel": apiName})
	case "publish":
		if len(args) < 1 {
			return fmt.Errorf("cloudcc publish html <apiName> [projectPath]")
		}
		projectPath := cwd
		if len(args) > 1 {
			projectPath = args[1]
		}
		apiName := args[0]
		dir := filepath.Join(projectPath, "html", apiName)
		local, err := jsonx.ReadObjectFile(filepath.Join(dir, "config.json"))
		if err != nil {
			return err
		}
		html, err := os.ReadFile(filepath.Join(dir, "index.html"))
		if err != nil {
			return err
		}
		local["htmlContent"] = string(html)
		cfg, err := config.Load(projectPath)
		if err != nil {
			return err
		}
		var res map[string]any
		if err := httpclient.New().PostRaw(baseURL(cfg)+"/devconsole/htmlComponent/saveHtmlComponent", local, map[string]string{"accessToken": config.String(cfg, "pluginToken")}, &res); err != nil {
			return err
		}
		return printJSON(stdout, res)
	default:
		return fmt.Errorf("unsupported html action: %s", action)
	}
}

func handleStaticResource(action string, args []string, stdout io.Writer, stderr io.Writer, cwd string) error {
	switch action {
	case "create":
		if len(args) < 2 {
			return fmt.Errorf("cloudcc create staticResource <name> <filePath>")
		}
		cfg, err := config.Load(cwd)
		if err != nil {
			return err
		}
		body := map[string]any{"name": args[0], "filePath": args[1]}
		return postClass(stdout, cfg, "setup", "/api/staticResource/save", body)
	case "get":
		return callGeneric(endpoint{"setup", "/api/staticResource/list"}, action, "staticResource", args, stdout, cwd)
	case "detail":
		return callGeneric(endpoint{"setup", "/api/staticResource/detail"}, action, "staticResource", args, stdout, cwd)
	case "delete":
		return callGeneric(endpoint{"setup", "/api/staticResource/delete"}, action, "staticResource", args, stdout, cwd)
	case "count":
		return callGeneric(endpoint{"setup", "/api/staticResource/count"}, action, "staticResource", args, stdout, cwd)
	default:
		return fmt.Errorf("unsupported staticResource action: %s", action)
	}
}

func postClass(stdout io.Writer, cfg config.Config, base string, apiPath string, body map[string]any) error {
	var svc string
	if base == "api" {
		svc = strings.TrimRight(config.String(cfg, "apiSvc"), "/")
	} else {
		svc = strings.TrimRight(config.String(cfg, "setupSvc"), "/")
	}
	var res map[string]any
	if err := httpclient.New().PostClass(svc+apiPath, body, config.String(cfg, "accessToken"), &res); err != nil {
		return err
	}
	return printJSON(stdout, res)
}

func readMarkedSource(file string) (string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	text := string(b)
	re := regexp.MustCompile(`(?s)//\s*@SOURCE_CONTENT_START\r?\n(.*?)\r?\n\s*//\s*@SOURCE_CONTENT_END`)
	if m := re.FindStringSubmatch(text); len(m) > 1 {
		return m[1], nil
	}
	return text, nil
}

func extractMainBody(content string) (string, error) {
	idx := strings.Index(content, "function main")
	if idx < 0 {
		return "", fmt.Errorf("no function main($CCDK, obj) found")
	}
	open := strings.Index(content[idx:], "{")
	if open < 0 {
		return "", fmt.Errorf("main function has no body")
	}
	start := idx + open + 1
	depth := 1
	inSingle, inDouble, inTemplate, escaped := false, false, false, false
	for i := start; i < len(content); i++ {
		ch := content[i]
		if escaped {
			escaped = false
			continue
		}
		if ch == '\\' {
			escaped = true
			continue
		}
		if inSingle {
			if ch == '\'' {
				inSingle = false
			}
			continue
		}
		if inDouble {
			if ch == '"' {
				inDouble = false
			}
			continue
		}
		if inTemplate {
			if ch == '`' {
				inTemplate = false
			}
			continue
		}
		switch ch {
		case '\'':
			inSingle = true
		case '"':
			inDouble = true
		case '`':
			inTemplate = true
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return content[start:i], nil
			}
		}
	}
	return "", fmt.Errorf("main function braces are not balanced")
}

func configID(data map[string]any) any {
	if v := data["id"]; v != nil {
		return v
	}
	return ""
}

func baseURL(cfg config.Config) string {
	if v := config.String(cfg, "baseUrl"); v != "" {
		return strings.TrimRight(v, "/")
	}
	return "https://developer.apis.cloudcc.cn"
}

func printJSON(w io.Writer, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(b))
	return err
}

func firstArg(args []string, fallback string) string {
	if len(args) == 0 || args[0] == "" {
		return fallback
	}
	return args[0]
}

func firstAny(values ...any) any {
	for _, v := range values {
		if v != nil && fmt.Sprint(v) != "" && fmt.Sprint(v) != "<nil>" {
			return v
		}
	}
	return nil
}

func labelToSlug(label string) string {
	s := strings.ToLower(strings.TrimSpace(label))
	re := regexp.MustCompile(`[^a-z0-9]+`)
	s = re.ReplaceAllString(s, "_")
	s = strings.Trim(s, "_")
	if s == "" {
		return "custom_object"
	}
	return s
}

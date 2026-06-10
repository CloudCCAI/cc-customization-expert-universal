# cc-customization-expert-universal v2.0.4

CloudCC CRM 二开技能的 Universal Go 版实现。目标是替代 Node 技能中的 `node_modules` 打包方式，使用平台二进制 + embedded docs/templates 完成离线分发。

- Package name: `cc-customization-expert-universal`
- Version: `2.0.4`
- Release date: `2026-06-10`

## Scope

- Included: P0-P3 from `docs/specs/FEAT-001-cloudcc-expert-go-rewrite.md`
- Included P4 partial: canonical `pagecomponent` commands, with `publish` delegating the Vue build to local `vue-cli-service`
- Included project template: `cloudcc create project <name>` creates `frontend/`, `backend/`, `sidecar/`, and root config files
- Included docs are organized as three two-level groups: `platform/*`, `methodology/*`, and `playbooks/*`
- Included platform knowledge: `platform/overview`, `platform/capabilityMap`, `platform/security`, `platform/automation`, `platform/dataModeling`, `platform/integrationArchitecture`, `platform/integrationPatterns`, `platform/lowcodeHighcode`, `platform/mobileCapabilities`, `platform/almRelease`, and concrete modules such as `platform/object` and `platform/pagecomponent`
- Included methodology docs: `methodology/blueprint`, `methodology/moduleDesign`, `methodology/integrationDesign`, and `methodology/deliveryPlan`
- Included playbooks: `playbooks/manufacturingCrm`, `playbooks/serviceWorkOrder`, and `playbooks/miniProgramPortal`
- Deferred: pure-Go Vue build replacement, JSP migration, and full MCP tool registration

## Runtime

The skill uses `tools/bin/cloudcc` on macOS/Linux and `tools/bin/cloudcc.cmd` on Windows as stable wrappers.

```bash
tools/bin/cloudcc --version
tools/bin/cloudcc doc platform/overview introduction
tools/bin/cloudcc doc platform/capabilityMap introduction
tools/bin/cloudcc doc platform/object introduction
tools/bin/cloudcc doc platform/pagecomponent devguide
tools/bin/cloudcc doc methodology/blueprint devguide
tools/bin/cloudcc doc playbooks/serviceWorkOrder introduction
tools/bin/cloudcc create project demo-cloudcc
```

## Knowledge Docs

Use platform docs before jumping into single-module implementation:

```bash
tools/bin/cloudcc doc platform/overview introduction
tools/bin/cloudcc doc platform/capabilityMap introduction
tools/bin/cloudcc doc platform/security introduction
tools/bin/cloudcc doc platform/automation introduction
tools/bin/cloudcc doc platform/dataModeling introduction
tools/bin/cloudcc doc platform/integrationArchitecture introduction
tools/bin/cloudcc doc platform/integrationPatterns introduction
tools/bin/cloudcc doc platform/lowcodeHighcode introduction
tools/bin/cloudcc doc platform/mobileCapabilities introduction
tools/bin/cloudcc doc platform/almRelease introduction
```

Use methodology and playbooks for project design:

```bash
tools/bin/cloudcc doc methodology/blueprint devguide
tools/bin/cloudcc doc methodology/moduleDesign devguide
tools/bin/cloudcc doc methodology/integrationDesign devguide
tools/bin/cloudcc doc methodology/deliveryPlan devguide
tools/bin/cloudcc doc playbooks/manufacturingCrm introduction
tools/bin/cloudcc doc playbooks/serviceWorkOrder devguide
tools/bin/cloudcc doc playbooks/miniProgramPortal devguide
```

On Windows:

```bat
tools\bin\cloudcc.cmd --version
tools\bin\cloudcc.cmd doc platform/object introduction
```

The wrapper looks for a platform binary under:

```text
tools/bin-<os>-<arch>/cloudcc
tools/bin-windows-<arch>/cloudcc.exe
```

If the binary is missing and Go is installed, it can run from source with:

```bash
cd cli
go run ./cmd/cloudcc --version
```

## Build

```bash
./tools/build.sh
```

Build outputs are written to `tools/bin-<os>-<arch>/cloudcc`.
Windows builds are written to `tools/bin-windows-<arch>/cloudcc.exe`.

## PageComponent

`pagecomponent` is the only supported resource name for CloudCC page custom components.
Local pagecomponent sources live under `frontend/pagecomponents/<name>/`.

```bash
tools/bin/cloudcc create project demo-cloudcc
tools/bin/cloudcc create pagecomponent cc-demo
tools/bin/cloudcc publish pagecomponent cc-demo /path/to/project
tools/bin/cloudcc get pagecomponent /path/to/project
tools/bin/cloudcc detail pagecomponent cc-demo "" /path/to/project
tools/bin/cloudcc pull pagecomponent <nameOrId> /path/to/project
tools/bin/cloudcc delete pagecomponent <nameOrId> /path/to/project
```

`publish pagecomponent` requires the target project to have its normal Vue 2 build toolchain available. The Go CLI creates the temporary build entry, runs local `npx vue-cli-service build`, reads the generated UMD bundle, and uploads it with the collected source dependency map.

## Generated Project Layout

```text
demo-cloudcc/
├── cloudcc-cli.config.json
├── .gitignore
├── frontend/
│   ├── package.json
│   ├── vue.config.js
│   ├── babel.config.js
│   ├── public/index.html
│   ├── src/
│   └── pagecomponents/
├── backend/
│   ├── classes/
│   ├── triggers/
│   ├── schedule/
│   └── lib/
└── sidecar/
```

- Frontend pagecomponents and related Vue build files are under `frontend/`.
- Backend classes, triggers, and schedule code are under `backend/`.
- Sidecar middleware programs belong under `sidecar/`.
- Global config files stay at the project root.

## Release

The canonical release folder name is:

```text
cc-customization-expert-universal
```

The version is embedded in `SKILL.md`, `README.md`, `config.json`, and the CLI `--version` output. Archive filenames may also include the version, for example:

```text
cc-customization-expert-universal-2.0.4.tar.gz
cc-customization-expert-universal-2.0.4.zip
```

## Layout

```text
.
├── SKILL.md
├── README.md
├── config.json
├── cli/
│   ├── cmd/cloudcc/
│   └── internal/
└── tools/
    ├── bin/cloudcc
    ├── build.sh
    └── bin-<os>-<arch>/
```

---
name: cc-customization-expert-universal
version: 2.0.4
description: CloudCC CRM/PaaS 实施设计与开发的 Universal Go 离线技能。优先使用技能目录内置 `tools/bin/cloudcc` 获取 `platform/*` 平台能力、`methodology/*` 实施方法论、`playbooks/*` 业务方案库与模块文档，再执行 config/openapi/metadata/pagecomponent/file publish 命令。自定义页面组件统一使用 pagecomponent 命令；JSP 迁移和完整 MCP 工具注册暂不包含。
---

# CloudCC CRM 实施专家技能 Universal v2.0.4

当前技能版本：`2.0.4`。
技能分发名：`cc-customization-expert-universal`。

## 必须保留

- 先确认有可用的 config 配置信息；如果没有，优先建议运行 `cloudcc create project <name|.>` 生成 `cloudcc-cli.config.json`，或让用户按默认模板提供 `use/dev/safetyMark/CloudCCDev`。
- 不要把 `username/baseUrl/orgId/clientId/openSecretKey` 作为新项目最小必需配置展示；这些是兼容旧明文配置或 `CloudCCDev` 解析后的字段。
- 本技能使用 Universal Go 版内置 CLI，优先调用技能目录内的 `tools/bin/cloudcc`。
- 不要在线安装 `cloudcc-cli`，不要依赖全局 Node/npm。
- 若 `tools/bin/cloudcc --version` 不可用，说明当前平台二进制缺失；先构建或补齐 `tools/bin-<os>-<arch>/cloudcc`。

## 快速规则

- 后文命令中的 `cloudcc` 均指技能内置 wrapper：`tools/bin/cloudcc`。
- 先确认内置 CLI 可用：`tools/bin/cloudcc --version`。
- 方案设计前先读取平台级文档：`tools/bin/cloudcc doc platform/overview introduction`、`tools/bin/cloudcc doc platform/capabilityMap introduction`。
- 蓝图、模块拆解、接口设计、阶段计划先读取方法论文档：`methodology/blueprint`、`methodology/moduleDesign`、`methodology/integrationDesign`、`methodology/deliveryPlan`。
- 制造业营销服一体化、售后工单、小程序门户类方案先读取 playbook：`playbooks/manufacturingCrm`、`playbooks/serviceWorkOrder`、`playbooks/miniProgramPortal`。
- 最后读取具体平台模块 `introduction`：`tools/bin/cloudcc doc platform/<module> introduction`。
- 方案开发前读取模块 `devguide`：`tools/bin/cloudcc doc platform/<module> devguide`。
- 特例：`platform/config` 仅支持 `devguide`。

## 当前实现范围

- P0：Go CLI skeleton、版本命令、help、文档 embed。
- P1：config、HTTP client、OpenAPI common executor、version/doctor/changelog 基础命令。
- P2：metadata 模块的高频读取与 encoded JSON body 写入入口。
- P3：classes、triggers、timer、script、html、staticResource 的本地文件/发布基础能力；classes/triggers/timer 本地源码位于 `backend/`。
- P4 部分：project 初始化、pagecomponent create/get/detail/pull/delete/doc，以及调用项目 `frontend/` 本地 `vue-cli-service` 的 publish。
- 文档知识库调整为三层二级目录：`platform/*`、`methodology/*`、`playbooks/*`。
- 平台能力层：`platform/overview`、`platform/capabilityMap`、`platform/security`、`platform/automation`、`platform/dataModeling`、`platform/integrationArchitecture`、`platform/integrationPatterns`、`platform/lowcodeHighcode`、`platform/mobileCapabilities`、`platform/almRelease`，以及 `platform/object`、`platform/fields`、`platform/pagecomponent` 等具体模块。
- 实施方法论层：`methodology/blueprint`、`methodology/moduleDesign`、`methodology/integrationDesign`、`methodology/deliveryPlan`。
- 业务方案库层：`playbooks/manufacturingCrm`、`playbooks/serviceWorkOrder`、`playbooks/miniProgramPortal`。

## 暂缓范围

- P4 暂不做：纯 Go Vue 构建替代、JSP analyze/split 迁移、完整 MCP 工具注册。
- 自定义页面组件只使用 `pagecomponent` resource，不保留旧命名入口。
- 遇到暂缓命令时，Go CLI 会返回明确的 deferred 提示。

## 常用命令

```bash
tools/bin/cloudcc --version
tools/bin/cloudcc doc platform/overview introduction
tools/bin/cloudcc doc platform/capabilityMap introduction
tools/bin/cloudcc doc platform/security introduction
tools/bin/cloudcc doc platform/automation introduction
tools/bin/cloudcc doc platform/integrationPatterns introduction
tools/bin/cloudcc doc methodology/blueprint devguide
tools/bin/cloudcc doc methodology/moduleDesign devguide
tools/bin/cloudcc doc playbooks/serviceWorkOrder introduction
tools/bin/cloudcc doc platform/object introduction
tools/bin/cloudcc doc platform/object devguide
tools/bin/cloudcc create project demo-cloudcc
tools/bin/cloudcc get config /path/to/project
tools/bin/cloudcc use config test /path/to/project
tools/bin/cloudcc query openapi /path/to/project '<encodeURI(JSON.stringify(body))>'
tools/bin/cloudcc create pagecomponent cc-demo
tools/bin/cloudcc publish pagecomponent cc-demo /path/to/project
tools/bin/cloudcc get pagecomponent /path/to/project
```

## 初始化目录约定

- 前端页面组件及相关构建文件：`frontend/`，组件源码位于 `frontend/pagecomponents/<name>/`。
- 后端 classes、triggers、timer/schedule 代码：`backend/classes/`、`backend/triggers/`、`backend/schedule/`。
- 外挂中间程序：`sidecar/`。
- 全局配置文件：项目根目录，例如 `cloudcc-cli.config.json`。

默认配置模板：

```json
{
  "use": "dev",
  "dev": {
    "safetyMark": "请设置一个安全标识",
    "CloudCCDev": "请设置开发者密钥"
  }
}
```

## 响应要求

- 给方案或代码前，先明确引用了哪些 `tools/bin/cloudcc doc` 平台级与模块级文档。
- 复杂需求必须先输出能力路径，例如 `object -> fields -> pagelayout -> validationRule -> profile`。
- 事实依据来自手册、技能文档或源码扫描；行业通用方法论必须标注为建议实践，不要冒充平台事实。
- 不确定的接口、后台配置项或执行顺序必须标注“待确认”，并建议查看后台页面或源码。
- 输出命令时保持可直接复制执行。
- 涉及 JSP、完整 MCP 或纯 Go Vue 构建替代时，说明当前 Go 版暂不包含，并建议继续使用 Node 版技能或创建单独迁移任务。

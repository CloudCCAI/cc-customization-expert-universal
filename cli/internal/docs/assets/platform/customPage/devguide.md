# CloudCC 自定义页面开发指南

## 1. 模块定位

`customPage` 模块用于通过 CLI 管理自定义页面，当前提供：

- 创建页面：`cloudcc create customPage ...`
- 更新页面：`cloudcc update customPage ...`
- 查询页面：`cloudcc get customPage ...`
- 删除页面：`cloudcc delete customPage ...`
- 文档查看：`cloudcc doc platform/customPage introduction|devguide`

---

## 2. 开发前准备

执行命令前请确认：

- 已完成 `cloudcc doc platform/project devguide` 的环境准备
- 项目路径下存在可用配置，且包含 `accessToken`
- 当前组织已存在至少一个自定义组件（创建页面时会依赖组件列表）

若配置缺失，命令会报错：

```text
Error: Configuration not found or accessToken is missing
```

---

## 3. 命令总览（以代码实现为准）

```bash
cloudcc create customPage [<pageLabel> <pageApi> <pluginId|compLabel>] [projectPath]
cloudcc update customPage <id> <pageLabel> <pageApi> <pluginId|compLabel> [projectPath]
cloudcc get customPage [pageNo] [pageSize] [projectPath]
cloudcc delete customPage <id> [projectPath]
cloudcc doc platform/customPage <introduction|devguide>
```

---

## 4. 创建自定义页面

## 4.1 标准创建模式

```bash
cloudcc create customPage <pageLabel> <pageApi> <pluginId|compLabel> [projectPath]
```

参数说明：

- `pageLabel`：页面名称
- `pageApi`：页面 API 名称
- `pluginId|compLabel`：组件标识，支持以下任一匹配：
  - 组件 `id`
  - 组件 `compLabel`
  - 组件 `compUniName`
- `projectPath`：项目路径，默认当前目录

示例：

```bash
cloudcc create customPage "合同助手页面" contract_assistant_page 2f9d0d6d2a ./
cloudcc create customPage "合同助手页面" contract_assistant_page 合同助手组件
```

## 4.2 自动创建模式（无参数）

```bash
cloudcc create customPage
```

当不传 `pageLabel/pageApi/plugin` 三个参数时，CLI 会自动：

1. 拉取自定义组件列表
2. 使用列表第一个组件
3. 生成页面名：`CLI自动页面_<timestamp>`
4. 生成页面 API：`cc_cli_page_<timestamp>`

这是快速验证环境是否可用的便捷方式。

## 4.3 参数校验规则

- 三个核心参数要么**同时传入**，要么**全部省略**
- 若只传部分参数，会报错并提示正确用法：

```text
Error: pageLabel、pageApi、pluginId/compLabel 需同时提供，或全部省略以使用第一个自定义组件
Usage: cloudcc create customPage [<pageLabel> <pageApi> <pluginId|compLabel>] [projectPath]
```

## 4.4 创建过程说明

创建命令内部会执行以下步骤：

1. 读取项目配置（token、org 等）
2. 拉取自定义组件列表
3. 按 `id/compLabel/compUniName` 匹配组件
4. 读取组件 `vueData`，并拼装页面内容
5. 执行页面保存并返回创建结果

常见失败场景：

- 组件列表为空：`no custom components found`
- 指定组件未找到：`component "<x>" not found in custom component list`
- `vueData` 非合法 JSON：`plugin vueData is not valid JSON`

---

## 5. 查询自定义页面

```bash
cloudcc get customPage [pageNo] [pageSize] [projectPath]
```

参数说明：

- `pageNo`：页码，默认 `1`
- `pageSize`：每页数量，默认 `20`
- `projectPath`：项目路径，默认当前目录

示例：

```bash
cloudcc get customPage
cloudcc get customPage 1 50 ./
```

执行成功后会输出总数，并逐条打印页面：

- `ID`
- `Label`（`pageLabel`）
- `API`（`pageApi`）

---

## 6. 更新自定义页面

```bash
cloudcc update customPage <id> <pageLabel> <pageApi> <pluginId|compLabel> [projectPath]
```

参数说明：

- `id`：页面 ID（必填，更新时必须有值）
- `pageLabel`：页面名称（必填）
- `pageApi`：页面 API 名称（必填）
- `pluginId|compLabel`：组件标识（必填，支持 `id/compLabel/compUniName`）
- `projectPath`：项目路径，默认当前目录

说明：

- 更新与新建使用同一套保存逻辑
- 更新通过传入 `id` 实现；新建时 `id` 为空字符串
- 更新时会按目标组件重新生成 `pageContent`，从而支持“更新自定义组件”

示例：

```bash
cloudcc update customPage 2f9d0d6d2a "合同助手页面V2" contract_assistant_page_v2 3a8b7c6d5e ./
cloudcc update customPage 2f9d0d6d2a "合同助手页面V2" contract_assistant_page_v2 合同助手组件
```

---

## 7. 删除自定义页面

```bash
cloudcc delete customPage <id> [projectPath]
```

参数说明：

- `id`：页面 ID（必填）
- `projectPath`：项目路径，默认当前目录

示例：

```bash
cloudcc delete customPage 2f9d0d6d2a
cloudcc delete customPage 2f9d0d6d2a ./
```

若缺少 `id`，会报错：

```text
Error: Custom page ID is required
Usage: cloudcc delete customPage <id> [projectPath]
```

---

## 8. 文档命令

```bash
cloudcc doc platform/customPage introduction
cloudcc doc platform/customPage devguide
```

说明：

- 仅支持 `introduction` 与 `devguide`
- 传入其他子命令会抛错

---

## 9. 推荐操作流程

```bash
# 1) 先查列表，确认现有页面
cloudcc get customPage

# 2) 创建页面（标准模式）
cloudcc create customPage "合同助手页面" contract_assistant_page 2f9d0d6d2a

# 3) 更新页面并切换到新组件
cloudcc update customPage <id> "合同助手页面V2" contract_assistant_page_v2 <pluginId|compLabel>

# 4) 再查列表，确认创建/更新成功
cloudcc get customPage

# 5) 如需回滚，按 ID 删除
cloudcc delete customPage <id>
```

---

## 10. 注意事项

- `create` 强依赖自定义组件，先确保组件已存在
- `update` 同样依赖自定义组件匹配，组件不存在会报错
- 建议优先使用可读性高的 `pageApi` 命名（如业务域+功能）
- 删除前先确认页面未被菜单或其他入口依赖
- 生产环境操作前，先在测试环境验证参数与组件匹配

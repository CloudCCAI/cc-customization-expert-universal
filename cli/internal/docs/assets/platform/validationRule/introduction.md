# CloudCC 验证规则使用总结

验证规则（Validation Rule）用于在数据保存前进行校验，确保数据的完整性和正确性。当数据不符合规则时，系统会显示错误信息阻止保存。

---

## 快速开始（CLI 命令）

### 支持的验证规则操作

| 操作 | 说明 |
|------|------|
| `create` | 创建新验证规则 |
| `get` | 查询验证规则列表 |
| `delete` | 删除验证规则 |

---

## CLI 命令详解

### 创建验证规则

创建一个新的 CloudCC 验证规则。

```bash
cloudcc create validationRule <path> <objectPrefix> [ruleName] [ruleContent] [errorMessage]
```

**参数说明：**

| 参数 | 必填 | 说明 |
|------|------|------|
| `path` | 是 | 项目路径，`.` 表示当前目录 |
| `objectPrefix` | 是 | 对象前缀，如 `b00` |
| `ruleName` | 否 | 规则名称（不传则交互式输入） |
| `ruleContent` | 否 | 规则内容，如 `Batch_Size__c__f==5`（不传则交互式输入） |
| `errorMessage` | 否 | 错误提示信息（不传则交互式输入） |

**示例：**

```bash
# 交互式输入所有信息
cloudcc create validationRule . "b00"

# 非交互式创建
cloudcc create validationRule . "b00" "规则1" "Batch_Size__c__f==5" "数量必须是5"
```

---

### 查询验证规则列表

获取指定对象的所有验证规则列表。

```bash
cloudcc get validationRule <projectPath> <objectPrefix>
```

**参数说明：**

| 参数 | 必填 | 说明 |
|------|------|------|
| `projectPath` | 否 | 项目路径，默认当前目录 |
| `objectPrefix` | 是 | 对象前缀，如 `b00` |

**示例：**

```bash
# 获取对象 b00 的所有验证规则
cloudcc get validationRule . "b00"
```

---

### 删除验证规则

删除指定的验证规则。

```bash
cloudcc delete validationRule <projectPath> <ruleId>
```

**参数说明：**

| 参数 | 必填 | 说明 |
|------|------|------|
| `projectPath` | 否 | 项目路径，默认当前目录 |
| `ruleId` | 是 | 规则 ID |

**示例：**

```bash
# 删除指定规则
cloudcc delete validationRule . 202689E55795D38oAQlN
```

---

## 完整工作流示例

### 场景：为订单对象创建验证规则

```bash
# 1. 确认项目已初始化（有 cloudcc-cli.config.js）
cat cloudcc-cli.config.js

# 2. 查询对象现有的验证规则
cloudcc get validationRule . "b00"

# 3. 创建新验证规则
cloudcc create validationRule . "b00"

# 4. 验证规则创建成功
cloudcc get validationRule . "b00"

# 5. 如需删除
# cloudcc delete validationRule . <ruleId>
```

---

*文档版本：1.0 | 最后更新：2026-03-27*

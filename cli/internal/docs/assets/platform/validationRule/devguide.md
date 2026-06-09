# CloudCC 验证规则 CLI 命令说明

## 支持的命令

| 操作 | 说明 |
|------|------|
| `create` | 创建新验证规则 |
| `get` | 查询验证规则列表 |
| `delete` | 删除验证规则 |

## CLI 命令详解

### 创建验证规则

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

### 查询验证规则列表

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

### 删除验证规则

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
cloudcc delete validationRule . 202689E55795D38oAQlN
```

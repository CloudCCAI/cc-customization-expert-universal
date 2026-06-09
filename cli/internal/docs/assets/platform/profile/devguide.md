# CloudCC 简档 CLI 命令说明

## 支持的命令

| 操作 | 说明 |
|------|------|
| `create` | 创建自定义简档 |
| `get` | 查询简档列表 |
| `delete` | 删除自定义简档 |

## CLI 命令详解

### 创建简档

```bash
cloudcc create profile <path> <profileName> [description]
```

**参数说明：**

| 参数 | 必填 | 说明 |
|------|------|------|
| `path` | 是 | 项目路径，`.` 表示当前目录 |
| `profileName` | 是 | 简档名称 |
| `description` | 否 | 简档描述 |

**示例：**

```bash
# 创建简档
cloudcc create profile . "销售经理简档"

# 创建带描述的简档
cloudcc create profile . "销售代表简档" "适用于销售团队的权限配置"
```

### 查询简档列表

```bash
cloudcc get profile <projectPath> [encodedCondJson]
```

**参数说明：**

| 参数 | 必填 | 说明 |
|------|------|------|
| `projectPath` | 否 | 项目路径，默认当前目录 |
| `encodedCondJson` | 否 | URI 编码后的查询条件 JSON |

**示例：**

```bash
# 获取所有简档
cloudcc get profile .

# 带查询条件
cloudcc get profile . '%7B%22type%22%3A%22custom%22%7D'
```

### 删除简档

```bash
cloudcc delete profile <projectPath> <profileId>
```

**参数说明：**

| 参数 | 必填 | 说明 |
|------|------|------|
| `projectPath` | 否 | 项目路径，默认当前目录 |
| `profileId` | 是 | 简档 ID |

**示例：**

```bash
cloudcc delete profile . aaa202672F656B7VfEjL
```

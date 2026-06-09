# CloudCC 简档使用总结

简档（Profile）用于定义 CloudCC 系统中用户的权限集合，包括对象访问权限、字段权限、功能权限等。

---

## 快速开始（CLI 命令）

### 支持的简档操作

| 操作 | 说明 |
|------|------|
| `create` | 创建自定义简档 |
| `get` | 查询简档列表 |
| `delete` | 删除自定义简档 |

---

## CLI 命令详解

### 创建简档

创建一个新的自定义简档，基于现有简档进行克隆。

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

---

### 查询简档列表

获取当前环境中的所有简档列表。

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

---

### 删除简档

删除指定的自定义简档。

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
# 删除指定简档
cloudcc delete profile . a0I9D000000XXXXUAI
```

---

## 完整工作流示例

### 场景：为销售团队创建新的权限简档

```bash
# 1. 确认项目已初始化（有 cloudcc-cli.config.js）
cat cloudcc-cli.config.js

# 2. 查询现有简档
cloudcc get profile .

# 3. 创建新的销售经理简档
cloudcc create profile . "销售经理简档" "销售团队经理级别的权限配置"

# 4. 验证简档创建成功
cloudcc get profile .

# 5. 如需删除
# cloudcc delete profile . <profileId>
```

---

*文档版本：1.0 | 最后更新：2026-03-26*

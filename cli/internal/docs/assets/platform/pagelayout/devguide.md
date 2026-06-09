# CloudCC 页面布局 CLI 命令说明

## 支持的命令

| 操作 | 说明 |
|------|------|
| `get` | 查询页面布局列表 |
| `create` | 创建/复制页面布局 |
| `delete` | 删除页面布局 |

## CLI 命令详解

### 查询页面布局列表

```bash
cloudcc get pagelayout <projectPath> <prefix>
```

**参数说明：**

| 参数 | 必填 | 说明 |
|------|------|------|
| `projectPath` | 否 | 项目路径，`.` 表示当前目录 |
| `prefix` | 是 | 对象前缀（如 001, b25） |

**示例：**

```bash
# 查询对象 b25 的页面布局列表
cloudcc get pagelayout . b25

# 查询客户对象（001）的页面布局列表
cloudcc get pagelayout . 001
```

### 创建页面布局

```bash
cloudcc create pagelayout <projectPath> <objId> <layoutName> [sourceLayoutId] [isCloneDynamic]
```

**参数说明：**

| 参数 | 必填 | 说明 |
|------|------|------|
| `projectPath` | 否 | 项目路径，`.` 表示当前目录 |
| `objId` | 是 | 对象 ID |
| `layoutName` | 是 | 新页面布局名称 |
| `sourceLayoutId` | 否 | 要复制的源布局 ID，不传则使用默认第一个 |
| `isCloneDynamic` | 否 | 是否复制动态布局规则，默认 `true` |

**示例：**

```bash
# 创建页面布局（自动使用默认布局作为模板）
cloudcc create pagelayout . 20267D1465464C5OB6m5 "课程表2"

# 指定源布局 ID 进行复制
cloudcc create pagelayout . 20267D1465464C5OB6m5 "课程表2" add20261DA7347CZPAUz

# 不复制动态布局规则
cloudcc create pagelayout . 20267D1465464C5OB6m5 "课程表2" add20261DA7347CZPAUz false
```

### 删除页面布局

```bash
cloudcc delete pagelayout <projectPath> <layoutId>
```

**参数说明：**

| 参数 | 必填 | 说明 |
|------|------|------|
| `projectPath` | 否 | 项目路径，`.` 表示当前目录 |
| `layoutId` | 是 | 要删除的页面布局 ID |

**示例：**

```bash
# 删除指定页面布局
cloudcc delete pagelayout . add202610BD89F09XyGT
```

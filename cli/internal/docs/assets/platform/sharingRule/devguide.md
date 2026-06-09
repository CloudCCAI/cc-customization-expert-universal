# CloudCC 共享规则 CLI 命令说明

> 对应 one-setup-web 接口：`{setupSvc}/api/sharingSettings/queryRule`

## 1. 支持命令

```bash
# 按对象查询共享规则
cloudcc get sharingRule <objid> [projectPath]
```

## 2. 参数说明

### 2.1 查询：`cloudcc get sharingRule <objid> [projectPath]`

| 参数 | 必填 | 说明 |
| :--- | :--- | :--- |
| `objid` | 是 | 对象 id |
| `projectPath` | 否 | 项目根目录，默认当前目录（可传 `.`） |

## 3. 示例

### 3.1 按对象查询规则

```bash
cloudcc get sharingRule account .
```


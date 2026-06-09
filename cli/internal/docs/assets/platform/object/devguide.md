# CloudCC 自定义对象开发指南

## 1. 模块定位

`object` 模块用于管理 CloudCC 自定义对象，当前提供查询、创建、删除三类能力。

可通过以下命令读取文档：

```bash
cloudcc doc platform/object introduction
cloudcc doc platform/object devguide
```

---

## 2. 当前支持的命令

```bash
cloudcc get object <projectPath> [type]
cloudcc create object <projectPath> <label> <业务功能描述>
cloudcc create object <projectPath> <label> <nameLabel> <业务功能描述>
cloudcc delete object <projectPath> <objid>
```

说明：

- `projectPath`：本地项目路径，用于读取 `cloudcc-cli.config.js`
- `label`：对象中文名称或展示名称
- `nameLabel`：对象 API 名称；仅在上一种四参数形式中传入；不传（三参数形式）时基于 `label` 自动生成
- `业务功能描述`（**必填**，放在最后）：拼接到默认 remark「用于管理与「label」相关的业务数据」之后，整体写入 `obj.remark`
- `objid`：对象 ID

---

## 3. 查询对象

### 3.1 查询全部对象

```bash
cloudcc get object <projectPath>
```

返回标准对象和自定义对象的合并列表。

### 3.2 按类型查询

```bash
cloudcc get object <projectPath> standard
cloudcc get object <projectPath> custom
```

常见含义：

- `standard`：标准对象
- `custom`：自定义对象

---

## 4. 创建对象

### 4.1 基本命令

```bash
cloudcc create object <projectPath> <label> "记录客户拜访与跟进计划"
cloudcc create object <projectPath> <label> <nameLabel> "主数据：客户主档"
```


### 4.2 创建过程

创建对象时，模块会：

1. 读取项目配置
2. 查询角色列表
3. 自动生成对象权限配置
4. 调用 `/api/customObject/saveButton` 创建对象

---

## 5. 删除对象

### 5.1 基本命令

```bash
cloudcc delete object <projectPath> <objid>
```

删除前建议：

- 先确认对象未被页面、菜单、字段或自动化逻辑依赖
- 先保留对象 ID 与对象名称，避免误删

---

## 6. 开发前检查

- 已完成 `cloudcc doc platform/project devguide` 中的环境准备
- 项目根目录存在可用的 `cloudcc-cli.config.js`
- 当前环境密钥配置正确
- 已确认对象名称、API 名称、角色权限需求

---

## 7. 占位说明

这份文档当前为初版开发指南，后续可继续补充：

- 对象建模规范
- 命名建议
- 角色权限说明
- 与字段、菜单、应用的联动关系

# CloudCC ALM 与发布指南

## 本地项目结构

当前 Go 技能生成结构：

```text
project/
├── cloudcc-cli.config.json
├── frontend/
│   └── pagecomponents/
├── backend/
│   ├── classes/
│   ├── triggers/
│   └── schedule/
└── sidecar/
```

## 发布前检查

### 1. 环境配置

确认：

- `cloudcc-cli.config.json` 存在
- `use` 指向正确环境
- `safetyMark` 和 `CloudCCDev` 已配置
- 不要把生产环境误设为当前环境

### 2. 元数据清单

列出本次变更涉及：

- 对象
- 字段
- 布局
- 菜单
- 应用
- 权限
- 自动化
- 代码和组件

### 3. 权限验收

至少用目标简档用户验证：

- 应用可见
- 菜单可见
- 对象操作权限
- 字段可见/只读
- 记录共享范围
- 移动端入口

### 4. 自动化验收

验证：

- 验证规则
- 查重过滤器
- 触发器
- 工作流/批准
- 定时作业

### 5. 高代码发布

按能力选择命令：

```bash
cloudcc publish classes <name>
cloudcc publish triggers <object/name>
cloudcc publish timer <name>
cloudcc publish pagecomponent <name> [projectPath]
cloudcc publish staticResource ...
```

具体参数以对应模块 `devguide` 为准。

## 软件包使用边界

手册确认软件包可迁移类、自定义页面、自定义组件、应用程序等。具体哪些组件类型可添加、依赖关系如何解析、是否覆盖已有配置，需要以目标租户的软件包页面为准。

## 行业通用发布建议

- 每次发布准备变更清单和回滚方案。
- 配置变更和代码变更分开验收。
- 先发布后端类和触发器，再发布依赖它们的页面或组件。
- 生产发布前冻结目标环境手工配置。
- 记录已安装软件包 id、版本和安装 URL。
- sidecar 独立管理部署、环境变量、日志和回滚。

## 待确认

当前技能未完整实现软件包 CLI，也未覆盖所有后台配置迁移能力。涉及软件包安装、上载、组件依赖解析时，应以后台页面或后续补充文档为准。

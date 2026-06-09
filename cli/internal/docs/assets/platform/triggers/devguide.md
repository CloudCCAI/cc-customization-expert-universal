# CloudCC 触发器开发规范

## 1. 核心原则

触发器只做事件入口，复杂业务逻辑统一写入自定义类。

推荐职责划分：

- 触发器：获取参数、判断条件、调用自定义类、必要时报错
- 自定义类：承载业务规则、聚合其他类、负责编排逻辑

## 2. 开发规范

- 一个对象的一个触发时机，只创建一个触发器
- 创建触发器时，推荐使用自动创建自定义类模式：`cloudcc create triggers <encodedCreateJson> true`
- 尽量不要在触发器中直接开发业务逻辑
- 需要复用、编排、扩展的逻辑，统一下沉到自定义类

## 3. 触发时间

- `beforeInsert`：当新建、复制记录行满足触发器条件，执行触发器操作
- `beforeUpdate`：编辑记录行满足触发器条件，执行触发器操作
- `beforeUpsert`：新建和编辑、复制记录行满足触发器条件，执行触发器操作
- `beforeDelete`：删除记录行满足触发器条件，执行触发器操作
- `afterInsert`：新建、复制记录行满足触发器条件，执行触发器操作
- `afterUpdate`：修改记录行满足触发器条件，执行触发器操作
- `afterUpsert`：新建或编辑、复制记录行满足触发器条件，执行触发器操作
- `afterDelete`：删除记录行满足触发器条件，执行触发器操作
- `afterInsertCommit`：插入提交后，报错不会回滚
- `afterUpdateCommit`：更新提交后，报错不会回滚
- `afterUpsertCommit`：保存提交后，报错不会回滚
- `afterDeleteCommit`：删除提交后，报错不会回滚
- `approval`：审批

## 4. 可直接引用的参数

- `(UserInfo)userInfo`
- `(Map<String, Object> )record_old`
- `(Map<String, Object> )record_new`

取记录值时，直接使用字段 API 名：

```java
record_old.get("name");
record_new.get("name");
```

## 5. 保存前提示

在记录保存前可使用：

```java
trigger.addErrorMessage("提示内容");
```

用于向用户提示并阻断保存。

## 8. 推荐写法

触发器中只保留入口逻辑：

```java
public class MyTriggerService extends CCTrigger {
    public MyTriggerService() {
        super(userInfo);
        // @SOURCE_CONTENT_START
        try {
            MyTriggerService service = new MyTriggerService(userInfo,(CCService)this);
            service.execute(record_old, record_new);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
        // @SOURCE_CONTENT_END
    }
}
```

自定义类中承载业务实现：

```java
import java.util.Map;
// @SOURCE_CONTENT_START
public class MyTriggerService {
    private UserInfo userInfo;
    private CCService cs;

    public MyTriggerService(UserInfo userInfo) {
        this(userInfo, new CCService(userInfo));
    }

    public MyTriggerService(UserInfo userInfo, CCService cs) {
        this.userInfo = userInfo;
        this.cs = cs;
    }

    public void execute(Map<String, Object> record_old, Map<String, Object> record_new) {
            // 在这里编写业务逻辑
    }
}
// @SOURCE_CONTENT_END
```

## 8. triggers 模块支持的 CLI 命令总览（重点：入参）

说明：

- 下文命令中的资源名使用 `triggers`。
- `projectPath` 未传时，默认使用当前工作目录。

#### 1) 创建触发器

命令：

```bash
cloudcc create triggers <encodedCreateJson> [autoCreateClass]
```

`encodedCreateJson`（先 JSON.stringify，再 encodeURI）推荐字段：

| 字段              | 必填 | 类型     | 说明                                                   |
| ----------------- | ---- | -------- | ------------------------------------------------------ |
| `schemetableName` | 是   | `string` | 目标对象 API 名（目录中会转小写）                      |
| `targetObjectId`  | 是   | `string` | 目标对象 ID                                            |
| `triggerTime`     | 是   | `string` | 触发时机，如 `beforeInsert`、`afterUpdate`、`approval` |
| `name`            | 是   | `string` | 触发器名称，必须使用英文                                             |
| `apiname`         | 建议 | `string` | 触发器 API 名                                          |

`autoCreateClass`（可选）：

| 参数              | 必填 | 类型      | 说明                                                                 |
| ----------------- | ---- | --------- | -------------------------------------------------------------------- |
| `autoCreateClass` | 否   | `boolean` | 是否自动创建配套自定义类；默认 `false`，传 `true` 开启推荐自动创建模式 |

#### 2) 发布触发器

命令：

```bash
cloudcc publish triggers <namePath>
```

参数：

| 参数       | 必填 | 类型     | 说明                                    |
| ---------- | ---- | -------- | --------------------------------------- |
| `namePath` | 是   | `string` | 触发器路径，格式：`对象小写名/触发器名` |

#### 3) 拉取触发器（按本地路径）

命令：

```bash
cloudcc pull triggers <namePath>
```

参数：

| 参数       | 必填 | 类型     | 说明                                                                             |
| ---------- | ---- | -------- | -------------------------------------------------------------------------------- |
| `namePath` | 是   | `string` | 触发器路径，格式：`对象小写名/触发器名`；会读取该目录 `config.json` 的 `id` 拉取 |

#### 4) 查询触发器列表（支持条件查询）

命令：

```bash
cloudcc get triggers <listQueryJson> [projectPath]
```

`listQueryJson` 推荐结构：

```json
{
    "shownum": 2000,
    "showpage": 1,
    "sname": "",
    "objId": ""
}
```

字段语义：

| 字段       | 含义              | 类型     | 是否推荐 | 说明                           |
| ---------- | ----------------- | -------- | -------- | ------------------------------ |
| `shownum`  | 每页条数          | `number  | string`  | 推荐                           |
| `showpage` | 页码              | `number  | string`  | 推荐                           |
| `sname`    | 触发器名字        | `string` | 可选     | 按名称模糊筛选，模糊查询       |
| `objId`    | 触发器作用对象 ID | `string` | 可选     | 对象id（对象筛选优先用该字段） |

#### 5) 查看触发器详情

命令：

```bash
cloudcc detail triggers <namePath> <id>
```

参数规则（实现口径）：

| 参数       | 必填     | 类型     | 说明                                             |
| ---------- | -------- | -------- | ------------------------------------------------ |
| `namePath` | 条件必填 | `string` | 传 `namePath` 时优先查本地；本地不完整时再走线上 |
| `id`       | 条件必填 | `string` | 当 `namePath` 为空时，必须传 `id` 走线上查询     |

等价理解：`namePath` 与 `id` 至少传一个，优先使用 `namePath` 路径。

#### 6) 按 ID 拉取并落地到本地目录

命令：

```bash
cloudcc pullList triggers <id> <projectPath>
```

参数：

| 参数          | 必填 | 类型     | 说明                                                        |
| ------------- | ---- | -------- | ----------------------------------------------------------- |
| `id`          | 是   | `string` | 线上触发器 ID                                               |
| `projectPath` | 是   | `string` | 项目根目录；会写入到 `<projectPath>/triggers/<obj>/<name>/` |

#### 7) 删除触发器

命令：

```bash
cloudcc delete triggers <namePathOrId> [projectPath]
```

参数规则：

| 参数           | 必填 | 类型     | 说明                                                                                  |
| -------------- | ---- | -------- | ------------------------------------------------------------------------------------- |
| `namePathOrId` | 是   | `string` | 可传触发器路径或线上 ID；若本地路径存在且 `config.json` 含 `id`，优先使用该 `id` 删除 |
| `projectPath`  | 否   | `string` | 项目根目录，默认当前目录                                                              |

#### 8) 文档命令

命令：

```bash
cloudcc doc platform/triggers <introduction|devguide>
```

参数：

| 参数          | 必填      | 类型 | 说明     |
| ------------- | --------- | ---- | -------- |
| `introduction | devguide` | 是   | `string` |

## 7. 当前项目中触发器的真实约束

这是 AI 最容易忽视、但必须先接受的现实约束。

### 7.1 触发器不是普通 Java 类容器

根据当前项目的实际文件结构，触发器目录位于：

- `triggers/<对象 API 名小写>/<触发器名>/`

主类通常形如：

- `package triggers.<对象小写>.<触发器名>;`
- `public class Xxx extends CCTrigger`
- SOURCE 区域位于构造函数中

这意味着：

- 触发器天然是“薄入口”
- 不适合承载超长业务逻辑
- 更不适合在触发器里构建一整套复杂服务

## 8. AI 必须遵守的硬规则

### 8.1 只能通过 cloudcc-cli 管理触发器目录

不得手工创建或复制 `triggers/...` 目录，不得私自构造 `config.json`。

### 8.2 只能在 SOURCE 区域内写业务逻辑

AI 修改已有触发器时，只能修改：

```java
// @SOURCE_CONTENT_START
// @SOURCE_CONTENT_END
```

之间的内容。

不得破坏：

- 包路径
- 类名
- 目录结构
- `config.json`
- SOURCE 标记外的框架代码

### 8.3 不得私改 `config.json` 的身份字段

尤其不要私自改动：

- `id`
- `apiname`
- `targetObjectId`
- `schemetableName`
- `triggerTime`

这些字段应通过创建、发布、拉取、云端同步保持一致。

### 8.4 必须保留 `userInfo` 上下文

所有 SDK 能力都应围绕当前上下文用户执行。

不得：

- 硬编码用户
- 绕过 `userInfo` 伪造“当前用户”
- 用固定时区或固定组织代替上下文

### 8.5 触发器里必须显式考虑递归风险

凡是触发器里再去更新当前对象、父对象、子对象、共享对象，都要先判断是否会再次触发相关逻辑。

AI 不能默认“更新一下没事”。

### 8.6 触发器里必须显式考虑幂等性

尤其是以下动作：

- 自动生成记录
- 自动发待办
- 自动发邮件
- 自动推送外部系统
- 自动创建共享

必须先判断是否已执行过，否则很容易重复生成、重复发送、重复推送。

### 8.7 涉及时间必须优先使用 `TimeUtil`

只要出现以下需求，AI 默认当成“有时区风险”处理：

- 当前时间
- 到期日比较
- 提醒时间
- 审批时间
- 统计周期
- 格式化时间字符串

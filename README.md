
## GIN-DDD - 领域设计版本

___
###  1.项目信息

#### coding项目路径: https://dencit.coding.net/public/golang-framework/gin-ddd/git/files
#### github项目路径: https://github.com/Dencit/GIN-DDD

#### 作者: Dencit - 陈鸿扬
#### 邮箱: 632112883@qq.com

___
###  2.Demo模块-目录结构
~~~


http.go [执行http应用]
http [端口/应用层]
  |--- demo [服务模块/示例]
        |--- controler [控制器][AO层]
        |--- request [输入拦截][DTO层]
        |--- query [查询参数][DTO层]
        |--- trans [输出转化][DTO层]
        |--- logic [应用逻辑][BO层]

grpc.go [执行grpc应用]
grpc [端口/应用层]
  |--- demo [服务模块/示例]
        |--- controler [控制器][AO层]
        |--- request [输入拦截][DTO层]
        |--- query [查询参数][DTO层]
        |--- trans [输出转化][DTO层]
        |--- logic [应用逻辑][BO层]
        
cron.go [任务调度]
schedule[任务/应用层]
  |--- demo [任务模块/示例]

console.go [指令控制台]
command[指令/应用层]
  |--- demo [指令模块/示例]

domain [领域层]
  |--- demo [服务模块/示例]
        |--- srv [本地业务逻辑][BO层]
        |--- manager [跨域业务管理][BO层]
        |--- repo [数据仓储][DAO层]
        |--- entity [数据实体][DO层]/[数据持久化][PO层]
        |--- err [异常描述]
base [基础层/基类]
  |--- config [配置基础]
  |--- repo [数据仓储基础]
  |--- err [异常描述基础]
  |--- exception [异常基础]
  |--- midware [中间件基础]
  |--- respond [输出基础]

extend [扩展库]
database [SQL表结构]
postman [Postman导出接口]


~~~

---
### 3.Demo模块-代码调用顺序

~~~

业务简单时,业务逻辑写在logic类; 业务复杂 或 跨用户端时, logic类中共性业务,必须抽象为srv服务函数,再由logi类进行组合,
以促进 "应用层低耦合,领域层高内聚"的条件, 利于后期迭代扩展.

|[端口层]                       |[应用层]                        | |[领域层]                       |[基础]                      |
|                              |                               | |                              |                            |
|==========MVC_架构=============|*应用层,logic类代码 高冗余       | |*公共代码,业务耦合,没按模块划分   |*底层对象统一控制             |
|                              |*散弹式业务代码,1次迭代n个地方修改 | |*迭代修改,容易产生关联错误        |                           |
|                              |                               | |                              |                            |
|                              |                               | |                              |                            |
| http --> route -->           | controler --> logic -->       | | -->                          | entity                     |
|                              |                |              | | --> helper                   | edoc/es_orm                |
|                              |                |              | | --> common                   | entity(curl/RPC/OpenAPI)   |
|                              |                |              | |                              |                            |
|                              |                |              | |                              |                            |
|                              |                |              | |                              |                            |
|                              | job -->        |              | |                              |                            |
|                              |                |              | |                              |                            |
|                              | console -->    |              | |                              |                            |
|                              |                               | |                              |                            |
|                              |                               | |                              |                            |
|==========DDD_架构=============|*应用层只组合领域层对象,低代码冗余 | |*高复用代码,solid原则,按模块划分  |*底层对象统一控制             |
|                              |*高复用方法,不怕应用层拷贝修改     | |*可将MVC业务对象抽出,渐进式重构   |                            |
|                              |                               | |                              |                            |
| http --> route --> demo -->  | controler --> request         | |                              |                            |
|                              |   ^               |           | |                              |                            |
|                              |   |             logic -->     | | -> srv |--> repo -->         | entity                     |
|                              |   |               |           | |    |   |                     |                            |
|                              |   |____________ trans         | |    |   |--> manager |-->     | edoc/es_orm                |
|                              |                               | |    |   |            |-->     | entity(curl/RPC/OpenAPI)   |
|                              |                               | |    |   |                     |                            |
|                              |                               | |    |   |--> enum             |                            |
|                              | job -->                       | | ___|   |--> error            |                            |
|                              |                               | |    |                         |                            |
|                              | console -->                   | | ___|                         |                            |
|                              |                               | |                              |                            |
|                              |                               | |                              |                            |

以上示例模块代码,能够根据项目规范定制,这样就可以把最优代码整合起来,作为自动生成代码的模板.

~~~

---
### 4.接口文档&注释-命名规范
[./RE_DOC.md](./RE_DOC.md)

---
### 5.Demo模块-接口调用规范 (query查询表达式-参数获取工具)
[./extend/match-query/MatchQuery.md](./extend/match-query/MatchQuery.md)

---
### 6.api缓存工具
[./extend/api-cache/api-cache.md](./extend/api-cache/api-cache.md)

---
### 7.控制台指令
[./console.md](./console.md)

---
### 8.任务调度
[./cron.md](./cron.md)

___
### 9.grpc-流服务
[./grpc.md](./grpc.md)


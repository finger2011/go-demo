## 基于beego的项目

由于没有写过go项目，所以以一个beego的新项目作为毕业项目

### Go 项目工程化

+ conf: 配置文件
+ controllers: 控制器 - BFF
+ models: 持久化对象 - PO
+ routers: 路由
    + router_v1,v2,v3..:路由版本
+ services: 服务层
+ daos: 领域层 - DO
+ parameters: 输入和输出DTO对象 
+ untils:工具类以及中间件
    + log:日志
    + error:错误
+ jobs: 定时任务类
+ logs:日志
+ error.md:错误码
+ grpc: grpc 处理

### 微服务架构
+ BFF:controllers
    + 通过parameters对应的参数，以及调用beego的ParseJson方法生成DTO对象
    + 生成&调用services方法，返回各个service的结果
    + 组装service的结果，返回给前端
+ Service: services，
    + 创建&调用PO(简单对象)或者DO(复杂对象，需要多个PO组成)对象
    + 具体业务逻辑
    + 返回结果
+ Job:jobs

### API 设计

参考restful api设计
+ 接口:/version(如v1,v2)/resource(如user,order)/action(如create/update/...)/parameters(如果需要)
+ 参数:通过json传递参数，beego可以直接使用parseJson处理，包括对参数的校验如必填，不能为空等
+ 返回: 若有一些如接口不存在，直接返回404 - 使用responseWrite.WriteHeader(404) 
    + code: 错误码，可定义好0 或者 200为正确返回，其他为错误返回
    + msg:错误提示，当code不为正确返回时，为string类型的错误提示
    + data:正确返回时，返回结构化数据

### gRPC 的使用

+ protoFile: proto文件 protoc --go_out=. ./test.proto   protoc --go-grpc_out=. ./test.proto
+ service:生成的proto文件
+ 

### 并发的使用

### 微服务中间件的使用
在untils中创建对应的中间件对象，在services中调用对应的until工具

### 缓存的使用优化

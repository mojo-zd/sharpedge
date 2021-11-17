## 通过Request对象定义WEB接口

### `pkg/util/server`说明
该包下`server`文件主要通过注册对象来动态识别需要注册的`api`，注册的`instance`对象的所有方法需要包含输入参数(2个,一个context.Context对象用于传递链路信息， 另一个为xxReq对象
此对象定义了Meta信息用于restful信息解析。xxxReq对象动态注册路由到restful.WebService对象中),同时需要包含两个输出参数(1:用于response到客户端, 2:error对象,该方法是否出错)

### `RouteFunction`生成
通过反射在构建的`RouteFunction`中调用`instance`指定的方法并获取返回值，对返回值做统一处理

## 运行
```
go run cmd/main.go server
```
浏览器中打开 `http://localhost:8888/swagger`可查看效果
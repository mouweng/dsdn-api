# DSDN-API
> 这是DSDN项目后端API部分，使用Go语言实现，一个基于gin框架的web server服务。

## 目录说明
 - api：本项目所有提供的api接口
 - config：本项目配置文件
 - middleware：中间件
 - model：数据库操作持久层
 - modelgenerator：根据建表语句自动生成model层代码
 - utils：工具包
 - build.sh：编译脚本
 - main.go：本项目入口
 - README.md：项目说明文档

## 启动流程
```
sh build.sh
./main -p 8080
```

## 新建接口方式
1. 新建文件夹，创建handler.go和task.go
- handler.go
```go
package test

import "ginTemplate/api"

func init() {
	api.SetRouterRegister(func(group *api.RouterGroup) {
		rootRouter := group.Group("/test")
		// ping测试链接
		rootRouter.StdGET("ping", Pong)
	})
}
```
- task.go
```go
package test

import (
	"ginTemplate/api"
)

// ping测试链接
func Pong(c *api.Context) (int, string, interface{}) {
	return 0, "success", api.H{"ping": "pong"}
}
```
2. 在main.go中添加引用
```go
import(
	_ "ginTemplate/api/test"
)
```
3. 添加model层
- 使用modelgenerator根据建表语句生成代码
- 在api层里引用model层并调用相关函数
```go
import center "ginTemplate/model/center"
func DoAddCenter(c *api.Context) (int, string, interface{}) {
    ...
    center, err := center.GetCenter("1 = 1 AND id = ?", id)
    ...
	return 0, "success", api.H{"id": id}
}
```

## 支持功能
- 自动post传入的相关参数，生成数据库查询的条件【详见接口/test/testParam】
- 支持事务【详见接口/message/delCenter】
- 支持自动生成数据库增删改查代码【详见modelgenerator】
- 支持登录验证扩展【详见middleware】
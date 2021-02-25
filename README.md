# Modeltools
#### GO语言连接Mysql生成对应的model，包括对应字段类型、注释等。生成基础的结构体，不局限于某一个ORM。
  
  **备注:<br/>**
     如果字段类型设置允许为null，则为该字段赋予相应的类型null ,建议数据库中设置为不允许为null
  
  **源码码地址---------**
  ##### github：[https://github.com/eline001/modeltools](https://github.com/eline001/modeltools)
 
 

 **生成示例---------**

```go 
  package models
  
// 用户表
type User struct {
	Id        int          `json:"id"`         // 用户ID
	UserPhone string       `json:"user_phone"` // 用手机号
	UserName  string       `json:"user_name"`  // 用户名称
	Gender    int          `json:"gender"`     // 性别: 0:保密,1:女,2:男
	IsDeleted int          `json:"is_deleted"` // 是否删除;0:未删除,1:已删除
	DeletedAt sql.NullTime `json:"deleted_at"` // 删除日期
	CreatedAt time.Time    `json:"created_at"` // 创建时间
	UpdatedAt time.Time    `json:"updated_at"` // 最后更新时间
}

```

**参数配置--------conf.go**

```go 
  package conf
  
  // model保存路径
  const ModelPath = "./model/"
  // 是否覆盖已存在model
  const ModelReplace = true
  // 数据库驱动
  const DriverName = "mysql"
  
  type DbConf struct {
  	Host   string
  	Port   string
  	User   string
  	Pwd    string
  	DbName string
  }
  // 数据库链接配置
  var MasterDbConfig DbConf = DbConf{
  	Host:   "127.0.0.1",
  	Port:   "3306",
  	User:   "eline",
  	Pwd:    "test123456",
  	DbName: "goa_order",
  }
```

**生成model--------**
```go
package main

import (
	"goa-org/internal/modeltools/dbtools"
	"goa-org/internal/modeltools/generate"
)

func main() {
	//初始化数据库
	dbtools.Init()
	generate.Genertate() //生成所有表信息
	//generate.Genertate("order") //生成指定表信息，可变参数可传入多个表名
}


```
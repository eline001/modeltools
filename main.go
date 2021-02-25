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

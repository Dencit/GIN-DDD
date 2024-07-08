package main

import (
	"app/command/demo"
	"os"
)

// 指令列表-实例

var CmdList *CmdListStruct

// 指令列表-结构

type CmdListStruct struct {
	SampleCmd *demo.SampleCmd //指令例子
}

// 注册指令
func init() {
	CmdList = &CmdListStruct{
		//demo 模块指令
		SampleCmd: &demo.SampleCmd{},
	}
}

func main() {
	args := os.Args
	if len(args) > 1 {
		arg := args[1]

		//示例指令 - .\console.exe sample --name foobar --age 18
		if arg == "sample" {
			CmdList.SampleCmd.Exec()
		}

	}
}

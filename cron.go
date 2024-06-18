package main

import (
	"app/extend/crontab"
	"app/schedule/demo"
)

/**
desc: 任务调度注册-包
doc: https://pkg.go.dev/github.com/robfig/cron/v3#section-readme
*/

// 任务列表-实例

var CmdList *CmdListStruct

// 任务列表-结构

type CmdListStruct struct {
	SampleCmd *demo.SampleCmd //任务例子
}

// 注册任务
func init() {
	CmdList = &CmdListStruct{
		//demo 模块任务
		SampleCmd: &demo.SampleCmd{},
	}
}

func main() {
	cronTab := &crontab.CronTab{}
	//开始
	cronTab.Start()

	//例子
	cronTab.Command("sample action-1", "* * * * *", CmdList.SampleCmd.Exec, map[string]any{"action": 1})

	//结束
	cronTab.Stop()
}

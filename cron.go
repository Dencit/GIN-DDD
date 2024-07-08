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

var TaskList *TaskListStruct

// 任务列表-结构

type TaskListStruct struct {
	SampleTask *demo.SampleTask //任务例子
}

// 注册任务
func init() {
	TaskList = &TaskListStruct{
		//demo 模块任务
		SampleTask: &demo.SampleTask{},
	}
}

func main() {
	cronTab := &crontab.CronTab{}
	//开始
	cronTab.Start()

	//例子
	cronTab.Command("sample action-1", "* * * * *", TaskList.SampleTask.Exec, map[string]any{"action": 1})

	//结束
	cronTab.Stop()
}

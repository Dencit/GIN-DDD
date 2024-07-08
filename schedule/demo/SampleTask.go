package demo

import (
	"log"
	"os"
)

//任务结构

type SampleTask struct {
	sysLog *log.Logger //系统日志
}

//任务执行

func (receiver *SampleTask) Exec(param map[string]any) {

	// 系统日志
	sysLog := log.New(os.Stdout, "cron: ", log.LstdFlags)
	receiver.sysLog = sysLog

	//test
	receiver.sysLog.Println("SampleTask.Execute:", param["action"]) //

	return
}

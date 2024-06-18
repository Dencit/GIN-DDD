package demo

import (
	"log"
	"os"
)

//任务结构

type SampleCmd struct {
	sysLog *log.Logger //系统日志
}

//任务执行

func (receiver *SampleCmd) Exec(param map[string]any) {

	// 系统日志
	sysLog := log.New(os.Stdout, "cron: ", log.LstdFlags)
	receiver.sysLog = sysLog

	//test
	receiver.sysLog.Println("sampleCmd.Execute:", param["action"]) //

	return
}

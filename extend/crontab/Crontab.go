package crontab

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
desc: 任务调度功能-包
doc: https://pkg.go.dev/github.com/robfig/cron/v3#section-readme
*/

// 定时器-结构

type CronTab struct {
	sysLog         *log.Logger             //系统日志
	cronLog        cron.Logger             //调度日志
	cronTab        *cron.Cron              //定时器
	taskEntryIdMap map[string]cron.EntryID //记录启动的任务EntryId
	end            uint                    //结束条件
}

// 定时器-启动

func (instance *CronTab) Start() {
	// 系统日志
	sysLog := log.New(os.Stdout, "cron: ", log.LstdFlags)
	instance.sysLog = sysLog
	// 调度日志
	cronLog := cron.VerbosePrintfLogger(sysLog)
	cronLogSet := cron.WithLogger(cronLog)
	instance.cronLog = cronLog
	//调度实例化
	instance.cronTab = cron.New(cronLogSet)
	//结束条件 默认0-不结束
	instance.end = 0
	//初始化 任务列表
	instance.taskEntryIdMap = make(map[string]cron.EntryID, 0)
	//启动
	instance.cronTab.Start()
}

// 定时器-添加任务

func (instance *CronTab) Command(taskId string, cron string, closure func(map[string]any), param map[string]any) bool {

	entryId, err := instance.cronTab.AddFunc(cron, func() {
		closure(param)
		instance.cronLog.Info("task run success: " + taskId)
	})
	if err != nil {
		instance.cronLog.Error(err, "add task error: "+taskId, entryId)
		return false
	}
	instance.addEntryId(taskId, entryId)
	return true
}

// 添加任务ID
func (instance *CronTab) addEntryId(taskId string, entryId cron.EntryID) {
	instance.taskEntryIdMap[taskId] = entryId
}

// 删除任务ID
func (instance *CronTab) deleteEntryId(taskId string) {
	delete(instance.taskEntryIdMap, taskId)
}

// 定时器-结束

func (instance *CronTab) Stop() {
	//截取退出信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	sig := <-ch
	sigStr := fmt.Sprintf("%v", sig)
	if sigStr == "interrupt" {
		//退出
		instance.end = 1
		instance.cronTab.Stop()
	}
	//守护进程
	for {
		time.Sleep(1 * time.Second)
		if instance.end == 1 {
			break
		}
	}
}

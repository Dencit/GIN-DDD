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

func (receiver *CronTab) Start() {
	// 系统日志
	sysLog := log.New(os.Stdout, "cron: ", log.LstdFlags)
	receiver.sysLog = sysLog
	// 调度日志
	cronLog := cron.VerbosePrintfLogger(sysLog)
	cronLogSet := cron.WithLogger(cronLog)
	receiver.cronLog = cronLog
	//调度实例化
	receiver.cronTab = cron.New(cronLogSet)
	//结束条件 默认0-不结束
	receiver.end = 0
	//初始化 任务列表
	receiver.taskEntryIdMap = make(map[string]cron.EntryID, 0)
	//启动
	receiver.cronTab.Start()
}

// 定时器-添加任务

func (receiver *CronTab) Command(taskId string, cron string, closure func(map[string]any), param map[string]any) bool {

	entryId, err := receiver.cronTab.AddFunc(cron, func() {
		closure(param)
		receiver.cronLog.Info("task run success: " + taskId)
	})
	if err != nil {
		receiver.cronLog.Error(err, "add task error: "+taskId, entryId)
		return false
	}
	receiver.addEntryId(taskId, entryId)
	return true
}

// 添加任务ID
func (receiver *CronTab) addEntryId(taskId string, entryId cron.EntryID) {
	receiver.taskEntryIdMap[taskId] = entryId
}

// 删除任务ID
func (receiver *CronTab) deleteEntryId(taskId string) {
	delete(receiver.taskEntryIdMap, taskId)
}

// 定时器-结束

func (receiver *CronTab) Stop() {
	//截取退出信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	sig := <-ch
	sigStr := fmt.Sprintf("%v", sig)
	if sigStr == "interrupt" {
		//退出
		receiver.end = 1
		receiver.cronTab.Stop()
	}
	//守护进程
	for {
		time.Sleep(1 * time.Second)
		if receiver.end == 1 {
			break
		}
	}
}

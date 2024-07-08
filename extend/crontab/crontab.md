# 任务调度


## 说明
~~~
运行 go run cron.go
~~~

### 示例
~~~

func main() {
	cronTab := &crontab.CronTab{}
	//开始
	cronTab.Start()

	//例子
	cronTab.Command("sample action-1", "* * * * *", TaskList.SampleTask.Exec, map[string]any{"action": 1})

	//结束
	cronTab.Stop()
}

~~~
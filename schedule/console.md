
# 控制台指令

## 说明
~~~

编译: go run console.go

运行 .\console.exe sample --name foobar --age 18

~~~

### 示例
~~~

> ./console.go->main()

func main() {
	args := os.Args
	if len(args) > 1 {
		arg := args[1]
		if arg == "sample" {
			CmdList.SampleCmd.Exec()
		}
	}
}


> schedule/demo/SampleTask.go->Exec()

func (receiver *SampleTask) Exec(param map[string]any) {

	// 系统日志
	sysLog := log.New(os.Stdout, "cron: ", log.LstdFlags)
	receiver.sysLog = sysLog
	
	//test
	receiver.sysLog.Println("SampleTask.Execute:", param["action"])
	
	return
}

~~~
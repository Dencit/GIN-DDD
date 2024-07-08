package demo

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

//指令结构

type SampleCmd struct {
	sysLog *log.Logger //系统日志
	name   string      //姓名
	age    int         //年龄
}

//指令执行

func (receiver *SampleCmd) Exec() {

	// 系统日志
	sysLog := log.New(os.Stdout, "cmd: ", log.LstdFlags)
	receiver.sysLog = sysLog

	var cmd = &cobra.Command{
		Use:   "sample",
		Short: "sample cmd",
		Long:  `demo-sample desc`,
	}
	cmd.Run = func(cobraCmd *cobra.Command, args []string) {
		if args[0] == "sample" {
			receiver.Command(cobraCmd, args)
		}
	}
	cmd.Flags().StringVarP(&receiver.name, "name", "n", "", "person's name")
	cmd.Flags().IntVarP(&receiver.age, "age", "a", 0, "person's age")

	if err := cmd.Execute(); err != nil {
		fmt.Println("cmd error::", err)
		os.Exit(-1)
	}
	return
}

func (receiver *SampleCmd) Command(cobraCmd *cobra.Command, args []string) {

	//test
	receiver.sysLog.Printf("My Name is %s, My age is %d\n", receiver.name, receiver.age) //

	return
}

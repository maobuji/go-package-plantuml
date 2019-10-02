package codeanalysis

import (
	"os/exec"
	"time"
	"context"
)


// 任务执行结果
type JobExecuteResult struct {
	Cmd string
	Output []byte // 脚本输出
	Err error // 脚本错误原因
	StartTime time.Time // 启动时间
	EndTime time.Time // 结束时间
}

func Exec(c string) *JobExecuteResult {
	var (
		cmd *exec.Cmd
		err error
		output []byte
		result *JobExecuteResult
	)

	result = &JobExecuteResult{}
	// 上锁成功后，重置任务启动时间
	result.StartTime = time.Now()

	// 执行shell命令
	cancelCtx, cancelFunc := context.WithCancel(context.TODO())
	_ = cancelFunc

	cmd = exec.CommandContext(cancelCtx, "bash", "-c", c)

	// 执行并捕获输出
	output, err = cmd.CombinedOutput()

	// 记录任务结束时间
	result.Cmd = c
	result.EndTime = time.Now()
	result.Output = output
	result.Err = err

	return result
}
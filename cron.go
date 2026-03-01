package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type CronJob struct {
	Name     string // 작업 이름
	Schedule string // 크론 주기
	Task     func() // 실제 실행할 함수
}

// TODO: nateon 알림으로 특정 시간동안 어떤 것이 연동되지 않았는지 톡

func StartCronJobs() {
	c := cron.New(cron.WithSeconds(), cron.WithChain(
		cron.SkipIfStillRunning(cron.DefaultLogger),
		cron.Recover(cron.DefaultLogger),
	))

	jobList := []CronJob{
		{
			Name:     "SetInvoiceNumber",
			Schedule: "*/5 * * * * *",
			Task:     CronSetInvoiceNumber,
		},
	}

	for _, j := range jobList {
		job := j
		_, err := c.AddFunc(job.Schedule, func() {
			logger.Info("Cron Task Started", "jobName", job.Name)
			job.Task()
		})
		if err != nil {
			logger.Error("Cron Registration Failed", "jobName", job.Name)
		}
	}

	c.Start()
	logger.Info("Cron Start")
}

func CronSetInvoiceNumber() {
	orderNum := "20260224-0000497"
	token := "eyJ0eXAiOiJKV1QiLCJyZWdEYXRlIjoxNzcyMzc4MzI4NjA4LCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE3NzIzODkxMjgsImFjY2Vzc1VzZXIiOnsidXNlcklkIjoiOTgyNzE4IiwidXNlck5hbWUiOiLso7zsi53tmozsgqwg7ZWc7LyA7Ja0IiwiZGVwdElkIjpudWxsLCJkZXB0TmFtZSI6bnVsbCwicm9sZXMiOm51bGwsInBlcm1pc3Npb25zIjpudWxsLCJhdXRoZW50aWNhdGlvblRpbWUiOjE3NzIzNzY1ODAzNDMsImFjY2Vzc1RpbWUiOjE3NzIzNzgzMjg2MDgsInN5c3RtSWQiOiIzIiwibWFjQWRkcmVzcyI6Im5vcm1hbC1icm93c2VyIiwibGdpbklwIjoiMjIxLjE1My4xMzQuMTUwIn19.OFb4gBJxLH5eN9UqJo1inMsItRBEpOtFXnsPkKZsl74"
	data, err := GetAlpsOrders(orderNum, token)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get data: %v (orderNum: %s)", err, orderNum))
		return
	}
	fmt.Println(data)
	return
}

//package main

package tasks

import (
	"github.com/jasonlvhit/gocron"
	"graph-service/db"
	"time"
)

func Start() {
	db.InitMysql()

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	_ = s.Every(5).Minutes().From(gocron.NextTick()).Do(JobClaimRecord)
	<-s.Start()

}

package main

import (
	"InfoBot/app/handler"
	"InfoBot/apptype"
	"InfoBot/fmtogram"
	"time"
)

func main() {
	apptype.DB = apptype.ConnectToDatabase(true)
	//client.Start()
	//admin.Start()
	//channel.Start()
	//notification.Start()
	go fmtogram.Start()
	handler.Notification()
	for {
		time.Sleep(time.Second * 10)
		handler.TimeGuard()
		time.Sleep(time.Hour * 3)
	}
}

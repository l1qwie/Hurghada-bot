package main

import (
	"InfoBot/apptype"
	"InfoBot/tests/client"
)

func main() {
	apptype.DB = apptype.ConnectToDatabase(true)
	client.Start()
	//admin.Start()
	//go fmtogram.Start()
	//handler.Notification()
	//channel.Start()
	//notification.Start()
}

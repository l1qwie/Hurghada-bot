package main

import (
	"InfoBot/apptype"
	"InfoBot/tests/admin"
)

func main() {
	apptype.DB = apptype.ConnectToDatabase(true)
	//client.Start()
	admin.Start()
	//fmtogram.Start()
	//handler.Notification()
	//channel.Start()
	//notification.Start()
}

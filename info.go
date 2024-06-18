package main

import (
	"InfoBot/apptype"
	"InfoBot/tests/channel"
	"InfoBot/tests/notification"
)

func main() {
	apptype.DB = apptype.ConnectToDatabase(true)
	//client.Start()
	//admin.Start()
	//fmtogram.Start()
	channel.Start()
	notification.Start()
}

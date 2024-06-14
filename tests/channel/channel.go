package channel

import (
	"InfoBot/api"
	"InfoBot/app"
	"InfoBot/apptype"
	"time"
)

func Start() {
	defer rewriteToSheet()
	defer deleteActivity()
	defer deleteClientAcivity()
	apptype.DB = apptype.ConnectToDatabase(true)
	fm := api.StartTest()
	time.Sleep(time.Second * 2)
	checkProgress(fm)
	req := prepareReq()
	app.Receiving(req)
	checkResponse()
}

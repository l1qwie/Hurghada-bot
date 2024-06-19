package channel

import (
	"InfoBot/api"
	"InfoBot/app"
	"InfoBot/tests/settest"
	"log"
)

func channelMessage() {
	fm := api.StartTest()
	checkProgress(fm)
	log.Print("Test channelMessage() completed")
}

func regClient() {
	req := prepareReq()
	fm := app.Receiving(req)
	checkResponse(fm)
	log.Print("Test regClient() completed")
}

func Start() {
	defer rewriteToSheet()
	defer deleteActivity()
	defer deleteClientAcivity()
	defer settest.DeleteUser()
	channelMessage()
	regClient()
	log.Print("All channel tests completed")
}

package notification

import (
	"InfoBot/app"
	"InfoBot/app/handler"
)

func notif1() {
	checkProgress1(handler.Notificationtest1())
	req := createReq1()
	fm := app.Receiving(req)
	checkAnswer1(fm)
}

func notif2() {
	checkProgress2(handler.Notificationtest2())
	req := createReq2()
	fm := app.Receiving(req)
	checkAnswer2(fm)
}

func Start() {
	defer deleteDvij()
	defer deleteClientDvij()
	defer reSetSeqDvij()
	defer deleteClient()
	createDvij()
	createClientDvij()
	createClient()
	notif1()
	notif2()
}

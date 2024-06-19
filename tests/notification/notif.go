package notification

import (
	"InfoBot/app"
	"InfoBot/app/handler"
	"InfoBot/tests/settest"
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
	defer settest.DeleteDvij()
	defer settest.DeleteClientDvij()
	defer settest.RestartDvijSeq()
	defer settest.DeleteUser()
	settest.CreateDvij()
	settest.CreateClientDvij()
	settest.CreateUser()
	notif1()
	notif2()
}

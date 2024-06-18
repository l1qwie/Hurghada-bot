package notification

import "InfoBot/app/handler"

func notif1() {
	checkProgress1(handler.Notificationtest1())
}

func notif2() {
	checkProgress2(handler.Notificationtest2())
}

func Start() {
	defer deleteDvij()
	defer deleteClientDvij()
	defer reSetSeqDvij()
	createDvij()
	createClientDvij()
	notif1()
	notif2()
}

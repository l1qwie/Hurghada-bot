package client

import "InfoBot/fmtogram/types"

func sayHello() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "/start"
	return tr
}

func chClient() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "client"
	return tr
}

func chWorship() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "1"
	return tr
}

func chYouth() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "2"
	return tr
}

func chHomeGroups() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "3"
	return tr
}

func chPrayers() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "4"
	return tr
}

func chMandW() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "5"
	return tr
}

func chFilm() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "6"
	return tr
}

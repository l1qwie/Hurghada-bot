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
	tr.Result[0].Query.Data = "worship"
	return tr
}

func chYouth() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "youthmeeting"
	return tr
}

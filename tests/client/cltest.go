package client

import (
	"InfoBot/apptype"
	"InfoBot/fmtogram/formatter"
	"InfoBot/fmtogram/types"
	"InfoBot/tests/settest"
)

func testShowSchedule() {
	defer settest.DeleteUser()
	settest.CreateUser()
	ts := new(settest.TestStuct)
	ts.Round = 2
	ts.Name = "clientTest"
	ts.FuncReq = []func() *types.TelegramResponse{sayHello, chClient}
	ts.FuncRes = []func(*formatter.Formatter){greeteings, showSchedule}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3}
	ts.UpdtLevel = []int{0, 1}
	ts.DoTest()
}

func Start() {
	apptype.DB = apptype.ConnectToDatabase(true)
	testShowSchedule()
}

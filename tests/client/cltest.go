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
	ts.Name = "ShowSchedule"
	ts.FuncReq = []func() *types.TelegramResponse{sayHello, chClient}
	ts.FuncRes = []func(*formatter.Formatter){greeteings, showSchedule}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3}
	ts.UpdtLevel = []int{0, 1}
	ts.DoTest()
}

func testShowWorship() {
	defer settest.DeleteUser()
	settest.CreateUser()
	ts := new(settest.TestStuct)
	ts.Round = 3
	ts.Name = "ShowWorship"
	ts.FuncReq = []func() *types.TelegramResponse{sayHello, chClient, chWorship}
	ts.FuncRes = []func(*formatter.Formatter){greeteings, showSchedule, showWorship}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3, trfunc4, trfunc5}
	ts.UpdtLevel = []int{0, 1, 2}
	ts.DoTest()
}

func testShowYouthMeeting() {
	defer settest.DeleteUser()
	settest.CreateUser()
	ts := new(settest.TestStuct)
	ts.Round = 3
	ts.Name = "ShowYouthMeeting"
	ts.FuncReq = []func() *types.TelegramResponse{sayHello, chClient, chYouth}
	ts.FuncRes = []func(*formatter.Formatter){greeteings, showSchedule, showYouth}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3, trfunc4, trfunc5}
	ts.UpdtLevel = []int{0, 1, 2}
	ts.DoTest()
}

func testShowHomeGroups() {
	defer settest.DeleteUser()
	settest.CreateUser()
	ts := new(settest.TestStuct)
	ts.Round = 3
	ts.Name = "ShowHomeGroups"
	ts.FuncReq = []func() *types.TelegramResponse{sayHello, chClient, chHomeGroups}
	ts.FuncRes = []func(*formatter.Formatter){greeteings, showSchedule, showHomeGroups}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3, trfunc4, trfunc5}
	ts.UpdtLevel = []int{0, 1, 2}
	ts.DoTest()
}

func Start() {
	apptype.DB = apptype.ConnectToDatabase(true)
	//testShowSchedule()
	//testShowWorship()
	//testShowYouthMeeting()
	testShowHomeGroups()
}

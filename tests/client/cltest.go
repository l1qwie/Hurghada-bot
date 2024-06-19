package client

import (
	"InfoBot/fmtogram/formatter"
	"InfoBot/fmtogram/types"
	"InfoBot/tests/settest"
)

func testShowSchedule() {
	defer settest.DeleteUser()
	settest.CreateUser(0, "new")
	ts := new(settest.TestStuct)
	ts.Round = 2
	ts.Name = "ShowSchedule"
	ts.FuncReq = []func() *types.TelegramResponse{sayHello, chClient}
	ts.FuncRes = []func(*formatter.Formatter){greeteings, showSchedule}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3}
	ts.UpdtLevel = []int{0, 1}
	ts.DoTest(false, false)
}

func testShowWorship() {
	defer settest.DeleteUser()
	settest.CreateUser(0, "divarication")
	ts := new(settest.TestStuct)
	ts.Round = 2
	ts.Name = "ShowWorship"
	ts.FuncReq = []func() *types.TelegramResponse{chClient, chWorship}
	ts.FuncRes = []func(*formatter.Formatter){showSchedule, showWorship}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3}
	ts.UpdtLevel = []int{1, 2}
	ts.DoTest(false, false)
}

func testShowYouthMeeting() {
	defer settest.DeleteUser()
	settest.CreateUser(0, "divarication")
	ts := new(settest.TestStuct)
	ts.Round = 2
	ts.Name = "ShowYouthMeeting"
	ts.FuncReq = []func() *types.TelegramResponse{chClient, chYouth}
	ts.FuncRes = []func(*formatter.Formatter){showSchedule, showYouth}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3}
	ts.UpdtLevel = []int{1, 2}
	ts.DoTest(false, false)
}

func testShowHomeGroups() {
	defer settest.DeleteUser()
	settest.CreateUser(0, "divarication")
	ts := new(settest.TestStuct)
	ts.Round = 2
	ts.Name = "ShowHomeGroups"
	ts.FuncReq = []func() *types.TelegramResponse{chClient, chHomeGroups}
	ts.FuncRes = []func(*formatter.Formatter){showSchedule, showHomeGroups}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3}
	ts.UpdtLevel = []int{1, 2}
	ts.DoTest(false, false)
}

func testShowPrayers() {
	defer settest.DeleteUser()
	settest.CreateUser(0, "divarication")
	ts := new(settest.TestStuct)
	ts.Round = 2
	ts.Name = "ShowPrayers"
	ts.FuncReq = []func() *types.TelegramResponse{chClient, chPrayers}
	ts.FuncRes = []func(*formatter.Formatter){showSchedule, showPrayers}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3}
	ts.UpdtLevel = []int{1, 2}
	ts.DoTest(false, false)
}

func testShowMenAndWomen() {
	defer settest.DeleteUser()
	settest.CreateUser(0, "divarication")
	ts := new(settest.TestStuct)
	ts.Round = 2
	ts.Name = "ShowMenAndWomen"
	ts.FuncReq = []func() *types.TelegramResponse{chClient, chMandW}
	ts.FuncRes = []func(*formatter.Formatter){showSchedule, showMandW}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3}
	ts.UpdtLevel = []int{1, 2}
	ts.DoTest(false, false)
}

func testShowFilm() {
	defer settest.DeleteUser()
	settest.CreateUser(0, "divarication")
	ts := new(settest.TestStuct)
	ts.Round = 2
	ts.Name = "ShowFilm"
	ts.FuncReq = []func() *types.TelegramResponse{chClient, chFilm}
	ts.FuncRes = []func(*formatter.Formatter){showSchedule, showFilm}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3}
	ts.UpdtLevel = []int{1, 2}
	ts.DoTest(false, false)
}

func Start() {
	testShowSchedule()
	testShowWorship()
	testShowYouthMeeting()
	testShowHomeGroups()
	testShowPrayers()
	testShowMenAndWomen()
	testShowFilm()
}

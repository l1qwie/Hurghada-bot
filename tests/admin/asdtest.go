package admin

import (
	"InfoBot/apptype"
	"InfoBot/fmtogram/formatter"
	"InfoBot/fmtogram/types"
	"InfoBot/tests/settest"
)

func testCreateActivity() {
	defer settest.RestartSeq()
	defer settest.DeleteActivity()
	defer settest.DeleteUser()
	settest.CreateUser()
	ts := new(settest.TestStuct)
	ts.Round = 8
	ts.Name = "CreateActivity"
	ts.FuncReq = []func() *types.TelegramResponse{sayHello, chAdmin, chCreate, chNameRu, chNameEn, chTextRu, chTextEn, chSave}
	ts.FuncRes = []func(*formatter.Formatter){greeteings, showOptions, chCreateOpt, sentNameRu, sentNameEn, sentTextRu, sentTextEn, sentSave}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3, trfunc4, trfunc5, trfunc4, trfunc5}
	ts.UpdtLevel = []int{0, 1, 3, 1, 2, 3, 4, 5}
	ts.DoTest(true, false)
}

func testDeleteActivity() {
	defer settest.RestartSeq()
	defer settest.DeleteActivity()
	defer settest.DeleteUser()
	settest.CreateUser()
	settest.CreateActivity()
	ts := new(settest.TestStuct)
	ts.Round = 4
	ts.Name = "DeleteActivity"
	ts.FuncReq = []func() *types.TelegramResponse{sayHello, chAdmin, chDelete, chActivity}
	ts.FuncRes = []func(*formatter.Formatter){greeteings, showOptions, chDeleteOpt, sentActivity}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3, trfunc4, trfunc5, trfunc4, trfunc5}
	ts.UpdtLevel = []int{0, 1, 3, 1}
	ts.DoTest(false, false)
}

func testChangeActivity() {
	defer settest.RestartSeq()
	defer settest.DeleteActivity()
	defer settest.DeleteUser()
	settest.CreateUser()
	settest.CreateActivity()
	ts := new(settest.TestStuct)
	ts.Round = 8
	ts.Name = "ChangeActivity"
	ts.FuncReq = []func() *types.TelegramResponse{sayHello, chAdmin, chChange, chActivity, chChangeable, chNewTitleRu, chNewTitleEn, chSave}
	ts.FuncRes = []func(*formatter.Formatter){greeteings, showOptions, chChangeOpt, sentActivityCh, sentChangeable, sentNewTitleRu, sentNewTitleEn, sentSaveCh}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3, trfunc4, trfunc5, trfunc4, trfunc5, trfunc, trfunc1, trfunc2, trfunc3}
	ts.UpdtLevel = []int{0, 1, 3, 1, 2, 3, 4, 5}
	ts.DoTest(true, true)
}

func Start() {
	apptype.DB = apptype.ConnectToDatabase(true)
	testCreateActivity()
	testDeleteActivity()
	testChangeActivity()
}

package admin

import (
	"InfoBot/apptype"
	"InfoBot/fmtogram/formatter"
	"InfoBot/fmtogram/types"
	"InfoBot/tests/settest"
)

func testCreateActivity() {
	defer settest.DeleteUser()
	settest.CreateUser()
	ts := new(settest.TestStuct)
	ts.Round = 8
	ts.Name = "CreateActivity"
	ts.FuncReq = []func() *types.TelegramResponse{sayHello, chAdmin, chCreate, chNameRu, chNameEn, chTextRu, chTextEn, chSave}
	ts.FuncRes = []func(*formatter.Formatter){greeteings, showOptions, chCreateOpt, sentNameRu, sentNameEn, sentTextRu, sentTextEn, sentSave}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3, trfunc4, trfunc5, trfunc, trfunc1, trfunc2, trfunc3, trfunc4, trfunc5, trfunc, trfunc1, trfunc2, trfunc3}
	ts.UpdtLevel = []int{0, 1, 3, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func Start() {
	apptype.DB = apptype.ConnectToDatabase(true)
	testCreateActivity()
	//testDeleteActivity()
	//testChangeActivity()
}

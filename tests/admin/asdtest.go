package admin

import (
	"InfoBot/fmtogram/formatter"
	"InfoBot/fmtogram/types"
	"InfoBot/tests/settest"
)

func testCreateActivity() {
	defer settest.RestartSeq()
	defer settest.DeleteActivity()
	defer settest.DeleteUser()
	settest.CreateUser(1, "create")
	ts := new(settest.TestStuct)
	ts.Round = 5
	ts.Name = "CreateActivity"
	ts.FuncReq = []func() *types.TelegramResponse{chNameRu, chNameEn, chTextRu, chTextEn, chSave}
	ts.FuncRes = []func(*formatter.Formatter){sentNameRu, sentNameEn, sentTextRu, sentTextEn, sentSave}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc4, trfunc5}
	ts.UpdtLevel = []int{1, 2, 3, 4, 5}
	ts.DoTest(true, false)
}

func testDeleteActivity() {
	defer settest.RestartSeq()
	defer settest.DeleteActivity()
	defer settest.DeleteUser()
	settest.CreateUser(1, "delete")
	settest.CreateActivity()
	ts := new(settest.TestStuct)
	ts.Round = 1
	ts.Name = "DeleteActivity"
	ts.FuncReq = []func() *types.TelegramResponse{chActivity}
	ts.FuncRes = []func(*formatter.Formatter){sentActivity}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1}
	ts.UpdtLevel = []int{1}
	ts.DoTest(true, false)
}

func testChangeActivity() {
	defer settest.RestartSeq()
	defer settest.DeleteActivity()
	defer settest.DeleteUser()
	settest.CreateUser(1, "change")
	settest.CreateActivity()
	ts := new(settest.TestStuct)
	ts.Round = 5
	ts.Name = "ChangeActivity"
	ts.FuncReq = []func() *types.TelegramResponse{chActivity, chChangeable, chNewTitleRu, chNewTitleEn, chSave}
	ts.FuncRes = []func(*formatter.Formatter){sentActivityCh, sentChangeable, sentNewTitleRu, sentNewTitleEn, sentSaveCh}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3, trfunc4, trfunc5}
	ts.UpdtLevel = []int{1, 2, 3, 4, 5}
	ts.DoTest(true, true)
}

func testActivityList() {
	defer settest.DeleteUser()
	defer settest.DeleteClientDvij()
	defer settest.DeleteDvij()
	defer settest.RestartDvijSeq()
	settest.CreateUser(1, "list")
	settest.CreateDvij()
	settest.CreateClientDvij()
	ts := new(settest.TestStuct)
	ts.Round = 3
	ts.Name = "ActivityList"
	ts.FuncReq = []func() *types.TelegramResponse{chList, chActivity, chPerson}
	ts.FuncRes = []func(*formatter.Formatter){showActivity, showClients, showPersonInf}
	ts.FuncTrsh = []func() *types.TelegramResponse{trfunc, trfunc1, trfunc2, trfunc3, trfunc4, trfunc5, trfunc4, trfunc5, trfunc, trfunc1}
	ts.UpdtLevel = []int{0, 1, 2}
	ts.DoTest(false, false)
}

func Start() {
	testCreateActivity()
	testDeleteActivity()
	testChangeActivity()
	testActivityList()
}

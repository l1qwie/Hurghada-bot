package app

import (
	"InfoBot/app/handler"
	"InfoBot/apptype"
	"InfoBot/fmtogram/formatter"
	"InfoBot/fmtogram/helper"
	"InfoBot/fmtogram/types"
)

func Receiving(tr *types.TelegramResponse) *formatter.Formatter {
	comm := new(apptype.Common)
	fm := new(formatter.Formatter)
	apptype.DB = apptype.ConnectToDatabase(true)
	comm.Request = helper.ReturnText(tr)
	comm.Id = helper.ReturnChatId(tr)
	comm.Language = helper.ReturnLanguage(tr)
	handler.Redirect(comm, fm)
	return fm
}

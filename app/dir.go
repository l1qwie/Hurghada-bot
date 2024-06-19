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
	comm.Request = helper.ReturnText(tr)
	comm.Id = helper.ReturnChatId(tr)
	comm.Name = helper.ReturnName(tr)
	comm.Lastname = helper.ReturnLastName(tr)
	comm.Phone = helper.ReturnPhone(tr)
	comm.Nickname = helper.ReturnUsername(tr)
	comm.Language = helper.ReturnLanguage(tr)
	typeChat := helper.ReturnTypeOfChat(tr)
	if typeChat == "private" {
		handler.Redirect(comm, fm)
	} else if typeChat == "channel" {
		handler.RegToActivity(comm, fm)
	}
	return fm
}

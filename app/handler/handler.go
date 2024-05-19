package handler

import (
	"InfoBot/app/admin"
	"InfoBot/app/client"
	"InfoBot/apptype"
	"InfoBot/fmtogram/formatter"
)

func retrieveUser(req *apptype.Common, fm *formatter.Formatter) {
	if find(req.Id, fm.Error) {
		dbRetrieveUser(req, fm.Error)
	} else {
		createUser(req, fm.Error)
	}
}

func Redirect(req *apptype.Common, fm *formatter.Formatter) {
	retrieveUser(req, fm)
	isadmin, err := wichWay(req.Id, fm.Error)
	if err == nil {
		if isadmin {
			admin.Dispatcher(req, fm)
		} else {
			client.Dispatcher(req, fm)
		}
	}
	retainUser(req, fm)
}

func retainUser(req *apptype.Common, fm *formatter.Formatter) {
	dbRetainUser(req, fm.Error)
}

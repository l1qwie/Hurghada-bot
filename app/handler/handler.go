package handler

import (
	"InfoBot/app/admin"
	"InfoBot/app/client"
	"InfoBot/apptype"
	"InfoBot/fmtogram/formatter"
	"log"
)

func retrieveUser(req *apptype.Common, fm *formatter.Formatter) {
	if find(req.Id, fm.Error) {
		dbRetrieveUser(req, fm.Error)
		if req.Request == "admin" {
			req.Action = "divarication"
			changeStatus(req.Id, 1, fm.Error)
		} else if req.Request == "client" {
			req.Action = "divarication"
			changeStatus(req.Id, 0, fm.Error)
			log.Print("How it could be possiple?")
		}
	} else {
		createUser(req, fm.Error)
	}
}

func logs(req *apptype.Common) {
	log.Printf("req.Request = %s", req.Request)
	log.Printf("req.Id = %d", req.Id)
	log.Printf("req.Language = %s", req.Language)
	log.Printf("req.ExMessageId = %d", req.ExMessageId)
	log.Printf("req.Action = %s", req.Action)
	log.Printf("req.Level = %d", req.Level)
	log.Printf("req.TitleRu = %s", req.TitleRu)
	log.Printf("req.TitleEn = %s", req.TitleEn)
	log.Printf("req.DiscrpRu = %s", req.DiscrpRu)
	log.Printf("req.DiscrpEn = %s", req.DiscrpEn)
	log.Printf("req.DelActivity = %d", req.DelActivity)
	log.Printf("req.Changeable = %s", req.Changeable)
	log.Printf("req.ChangesRu = %s", req.ChangesRu)
	log.Printf("req.ChangesEn = %s", req.ChangesEn)
}

func Redirect(req *apptype.Common, fm *formatter.Formatter) {
	retrieveUser(req, fm)
	logs(req)
	isadmin, err := wichWay(req.Id, fm.Error)
	if err == nil {
		if isadmin {
			admin.Dispatcher(req, fm)
		} else {
			client.Dispatcher(req, fm)
		}
	}
	retainUser(req, fm.Error)
}

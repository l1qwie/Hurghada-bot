package handler

import (
	"InfoBot/api"
	"InfoBot/app/admin"
	"InfoBot/app/client"
	"InfoBot/app/dict"
	"InfoBot/apptype"
	"InfoBot/fmtogram/formatter"
	"InfoBot/fmtogram/types"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type userAct struct {
	userid int
	actid  int
}

func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

func retrieveUser(req *apptype.Common, fm *formatter.Formatter) {
	if find(req.Id, fm.Error) {
		dbRetrieveUser(req, fm.Error)
		if req.Request == "/admin" {
			req.Action = "divarication"
			changeStatus(req.Id, 1, fm.Error)
		} else if req.Request == "/client" {
			req.Action = "divarication"
			changeStatus(req.Id, 0, fm.Error)
		} else if req.Request == "/refresh" {
			api.Start()
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
	if fm.Err != nil {
		log.Print(fmt.Errorf("YOU HAVE AN ERROR! %d", fm.Err))
	}
	retainUser(req, fm.Error)
}

func registerTheClient(req *apptype.Common, fm *formatter.Formatter, id int) {
	registerTheClientDB(req.Id, id)
	fm.WriteChatId(req.Id)
	fm.WriteString(dict.Dictionary[req.Language]["Registred"])
}

func RegToActivity(req *apptype.Common, fm *formatter.Formatter) {
	ok, id := client.IntCheck(req.Request)
	log.Print(ok, id)
	if ok {
		if findActivity(id) {
			if findClientWithActivity(req.Id, id) {
				registerTheClient(req, fm, id)
			}
		}
	}
}

func prepareKbJson(row *userAct, secondtime bool) (string, string) {
	var jsonbut1, jsonbut2 []byte
	but1 := &apptype.Answer{
		Userid: row.userid,
		Actid:  row.actid,
		Answer: "yes",
		First:  true,
	}
	but2 := &apptype.Answer{
		Userid: row.userid,
		Actid:  row.actid,
		Answer: "no",
		First:  true,
	}
	if secondtime {
		but1.Second = true
		but2.Second = true
	}
	jsonbut1, err := json.Marshal(but1)
	if err == nil {
		jsonbut2, err = json.Marshal(but2)
		if err != nil {
			log.Printf("!ERROR! after jsonbut2, err = json.Marshal(but2) in prepareKbJson(): %v", err)
		}
	} else {
		log.Printf("!ERROR! after jsonbut1, err := json.Marshal(but1) in prepareKbJson(): %v", err)
	}
	return string(jsonbut1), string(jsonbut2)
}

func notifbody(row *userAct, interval, but1, but2 string) *formatter.Formatter {
	fm := new(formatter.Formatter)
	fm.WriteChatId(row.userid)
	caption, discp, datetime, link := selectActInf(row.actid, fm.Error)
	parsed, err := time.Parse(time.RFC3339Nano, datetime)
	if err == nil {
		date := parsed.Format("2006-01-02")
		time := parsed.Format("15:04:05")
		fm.WriteString(fmt.Sprint(dict.Dictionary["ru"][interval], "\n\n\n", fmt.Sprintf(dict.Dictionary["ru"]["SampleChannel"], caption, discp, fmt.Sprintf("%s %s", date, time), link)))
		setKb(fm, []int{1, 1}, []string{dict.Dictionary["ru"]["I will"], dict.Dictionary["ru"]["I won't"]}, []string{but1, but2})
		fm.WriteParseMode(types.HTML)
	} else {
		fm.Error(err)
	}
	return fm
}

func send(fm *formatter.Formatter) {
	if err := fm.Complete(); err == nil {
		err = fm.SendToChannel()
		if err != nil {
			log.Printf("YOU HAVE AN ERROR DURING THE PROCCESS OF SENDING A MESSAGE TO A USER (notification): %v", err)
		}
	} else {
		log.Printf("YOU HAVE AN ERROR DURING THE PROCCESS OF PREPARE A MESSAGE (notification): %v", err)
	}
}

func notifDay() {
	useract, err := selectUserId("1 day")
	if err == nil {
		for _, row := range useract {
			but1, but2 := prepareKbJson(row, false)
			fm := notifbody(row, "1 day", but1, but2)
			changeNotifStatus("notifeone", row.userid, row.actid, fm.Error)
			send(fm)
		}
	} else {
		log.Printf("!ERROR! in notifDay(): %v", err)
	}
}

func notifHours() {
	useract, err := selectUserId("2 hours")
	if err == nil {
		for _, row := range useract {
			but1, but2 := prepareKbJson(row, true)
			fm := notifbody(row, "1 day", but1, but2)
			changeNotifStatus("notiftwo", row.userid, row.actid, fm.Error)
			send(fm)
		}
	} else {
		log.Printf("!ERROR! in notifHours(): %v", err)
	}
}

func Notification() {
	for {
		notifDay()
		notifHours()
		time.Sleep(time.Second * 120)
	}
}

func Notificationtest1() *formatter.Formatter {
	useract, err := selectUserId("12 seconds")
	if err != nil {
		panic(err)
	}
	but1, but2 := prepareKbJson(useract[0], false)
	fm := notifbody(useract[0], "1 day", but1, but2)
	changeNotifStatus("notifeone", useract[0].userid, useract[0].actid, fm.Error)
	return fm
}

func Notificationtest2() *formatter.Formatter {
	time.Sleep(time.Second * 2)
	useract, err := selectUserId("10 second")
	if err != nil {
		panic(err)
	}
	but1, but2 := prepareKbJson(useract[0], true)
	fm := notifbody(useract[0], "2 hours", but1, but2)
	changeNotifStatus("notiftwo", useract[0].userid, useract[0].actid, fm.Error)
	return fm
}

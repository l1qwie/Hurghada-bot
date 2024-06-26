package admin

import (
	"InfoBot/app/client"
	"InfoBot/app/dict"
	"InfoBot/apptype"
	"InfoBot/fmtogram/formatter"
	"InfoBot/fmtogram/types"
	"fmt"
	"log"
)

const (
	START  int = 0
	LEVEL1 int = 1
	LEVEL2 int = 2
	LEVEL3 int = 3
	LEVEL4 int = 4
	LEVEL5 int = 5
	OPT    int = 3
)

// Sets any keyboard
func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

func mainMenu(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Action = "divarication"
	req.Level = OPT
	fm.WriteString(dict["chooseOpt"])
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["create"], dict["change"], dict["delete"], dict["list"]}, []string{"create", "change", "delete", "list"})
}

func checkLanguage(req *apptype.Common) {
	if req.Language != "ru" && req.Language != "en" {
		req.Language = "en"
	}
}

func titleRu(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 1
	fm.WriteString(dict["nameRu"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func titleEn(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 2
	req.TitleRu = req.Request
	fm.WriteString(dict["nameEn"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func discrpRu(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 3
	req.TitleEn = req.Request
	fm.WriteString(dict["textRu"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func discrpEn(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 4
	req.DiscrpRu = req.Request
	fm.WriteString(dict["textEn"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func sendAlmoseDone(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 5
	fm.WriteString(fmt.Sprintf(dict["almostDone"], req.TitleRu, req.TitleEn, req.DiscrpRu, req.DiscrpEn))
	setKb(fm, []int{1, 1}, []string{dict["save"], dict["MainMenu"]}, []string{"save", "MainMenu"})
}

func almostDone(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.DiscrpEn = req.Request
	sendAlmoseDone(req, fm, dict)
}

func done(fm *formatter.Formatter, dict map[string]string) {
	fm.WriteString(dict["done"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func Createdone(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	saveInDB(req, fm.Error)
	done(fm, dict)
}

func doneOrChange(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Request == "save" {
		Createdone(req, fm, dict)
	} else if req.Request == "change" {
		req.Action = req.Request
		//change()
	} else {
		sendAlmoseDone(req, fm, dict)
	}
}

func create(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Level == START {
		titleRu(req, fm, dict)
	} else if req.Level == LEVEL1 {
		titleEn(req, fm, dict)
	} else if req.Level == LEVEL2 {
		discrpRu(req, fm, dict)
	} else if req.Level == LEVEL3 {
		discrpEn(req, fm, dict)
	} else if req.Level == LEVEL4 {
		almostDone(req, fm, dict)
	} else if req.Level == LEVEL5 {
		doneOrChange(req, fm, dict)
	}
}

func chooseActivity(req *apptype.Common, fm *formatter.Formatter, str string, dict map[string]string) {
	req.Level = 1
	crd := client.LengthOfNames(fm.Error)
	names := client.SelectNames(req.Language, len(crd), fm.Error)
	data := client.SelectValues(names, req.Language, len(crd), fm.Error)
	crd = append(crd, 1)
	names = append(names, dict["MainMenu"])
	data = append(data, "MainMenu")
	setKb(fm, crd, names, data)
	fm.WriteString(str)
}

func whichToDel(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Request == "delete" {
		chooseActivity(req, fm, dict["chooseAnActiv"], dict)
	} else {
		mainMenu(req, fm, dict)
	}
}

func goToDelete(req *apptype.Common, fm *formatter.Formatter, num int, dict map[string]string) {
	req.DelActivity = num
	deleteFromDB(req.DelActivity, fm.Error)
	done(fm, dict)
}

func deleted(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	det, num := client.IntCheck(req.Request)
	if det {
		if findActivity(num, fm.Error) {
			goToDelete(req, fm, num, dict)
		} else {
			chooseActivity(req, fm, dict["chooseAnActiv"], dict)
		}
	} else {
		chooseActivity(req, fm, dict["chooseAnActiv"], dict)
	}
}

func delete(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Level == START {
		whichToDel(req, fm, dict)
	} else if req.Level == LEVEL1 {
		deleted(req, fm, dict)
	}
}

func showDvij(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 1
	ids, names := selectAllDvij(fm.Error)
	crd := make([]int, len(ids)+1)
	for i := range crd {
		crd[i] = 1
	}
	crd[len(ids)] = 1
	ids = append(ids, "MainMenu")
	names = append(names, dict["MainMenu"])
	setKb(fm, crd, names, ids)
	fm.WriteString(dict["ChAct"])
}

func showClients(req *apptype.Common, fm *formatter.Formatter, dict map[string]string, actid int) {
	req.Level = 2
	req.DvijId = actid
	userids, nameslastname := selectAllClients(actid, fm.Error)
	crd := make([]int, len(userids)+1)
	for i := range crd {
		crd[i] = 1
	}
	crd[len(userids)] = 1
	userids = append(userids, "MainMenu")
	nameslastname = append(nameslastname, dict["MainMenu"])
	setKb(fm, crd, nameslastname, userids)
	fm.WriteString(dict["ChClient"])
}

func chooseClient(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	ok, actid := client.IntCheck(req.Request)
	if ok {
		if findDvij(actid, fm.Error) {
			showClients(req, fm, dict, actid)
		} else {
			showDvij(req, fm, dict)
		}
	} else {
		showDvij(req, fm, dict)
	}
}

func showClientInf(req *apptype.Common, fm *formatter.Formatter, dict map[string]string, userid int) {
	name, lastname, phonem, nickname, confirmed := selectAllClientInf(userid, req.DvijId, fm.Error)
	fm.WriteString(fmt.Sprintf(dict["ClientInf"], name, lastname, phonem, confirmed))
	fm.SetIkbdDim([]int{1, 1})
	fm.WriteInlineButtonUrl(dict["Chat"], fmt.Sprintf("t.me/%s", nickname))
	fm.WriteInlineButtonCmd(dict["MainMenu"], "MainMenu")
}

func clientInf(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	ok, userid := client.IntCheck(req.Request)
	if ok {
		if findDvijClient(req.DvijId, userid, fm.Error) {
			showClientInf(req, fm, dict, userid)
		} else {
			showClients(req, fm, dict, req.DvijId)
		}
	} else {
		showClients(req, fm, dict, req.DvijId)
	}
}

func list(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Level == START {
		showDvij(req, fm, dict)
	} else if req.Level == LEVEL1 {
		chooseClient(req, fm, dict)
	} else if req.Level == LEVEL2 {
		clientInf(req, fm, dict)
	}
}

func startChanges(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Request == "change" {
		chooseActivity(req, fm, dict["cooseAnActivToChange"], dict)
	} else {
		mainMenu(req, fm, dict)
	}
}

func chooseChangAble(req *apptype.Common, num int, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 2
	req.DelActivity = num
	fm.WriteString(dict["WhatyouwantChange"])
	setKb(fm, []int{1, 1, 1}, []string{dict["title"], dict["discrp"], dict["MainMenu"]}, []string{"title", "discrp", "MainMenu"})
}

func changeable(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	det, num := client.IntCheck(req.Request)
	if det {
		if findActivity(num, fm.Error) {
			chooseChangAble(req, num, fm, dict)
		} else {
			chooseActivity(req, fm, dict["cooseAnActivToChange"], dict)
		}
	} else {
		chooseActivity(req, fm, dict["cooseAnActivToChange"], dict)
	}
}

func bodyChanges(req *apptype.Common, lang string, fm *formatter.Formatter, dict map[string]string) {
	dataRu, dataEn := selectData(req.Changeable, req.DelActivity, fm.Error)
	fm.WriteString(fmt.Sprintf(dict["thisisdata"], dict[req.Changeable], dataRu, dataEn, dict[lang]))
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func changingRu(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 3
	req.Changeable = req.Request
	bodyChanges(req, "Rus", fm, dict)
}

func sendChangesRu(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Request == "title" || req.Request == "discrp" {
		changingRu(req, fm, dict)
	} else {
		chooseChangAble(req, req.DelActivity, fm, dict)
	}
}

func sendChangeEn(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 4
	req.ChangesRu = req.Request
	bodyChanges(req, "Engl", fm, dict)
}

func bodysendAlomostDone(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	ttru, tten, disru, disen := selectAllData(req.DelActivity, fm.Error)
	if req.Changeable == "title" {
		ttru = req.ChangesRu
		tten = req.ChangesEn
	} else {
		disru = req.ChangesRu
		disen = req.ChangesEn
	}
	fm.WriteString(fmt.Sprintf(dict["almoseChdone"], ttru, tten, disru, disen))
	setKb(fm, []int{1, 1}, []string{dict["save"], dict["MainMenu"]}, []string{"save", "MainMenu"})
}

func sendAlmostDoneCh(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 5
	req.ChangesEn = req.Request
	bodysendAlomostDone(req, fm, dict)
}

func saveChanges(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	saveChangesDB(req, fm.Error)
	done(fm, dict)
}

func changesDone(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Request == "save" {
		saveChanges(req, fm, dict)
	} else {
		bodysendAlomostDone(req, fm, dict)
	}
}

func change(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Level == START {
		startChanges(req, fm, dict)
	} else if req.Level == LEVEL1 {
		changeable(req, fm, dict)
	} else if req.Level == LEVEL2 {
		sendChangesRu(req, fm, dict)
	} else if req.Level == LEVEL3 {
		sendChangeEn(req, fm, dict)
	} else if req.Level == LEVEL4 {
		sendAlmostDoneCh(req, fm, dict)
	} else if req.Level == LEVEL5 {
		changesDone(req, fm, dict)
	}
}

func changeLvlAndAct(req *apptype.Common) {
	req.Level = START
	req.Action = req.Request
}

func divarication(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	changeLvlAndAct(req)
	if req.Request == "create" {
		create(req, fm, dict)
	} else if req.Request == "change" {
		change(req, fm, dict)
	} else if req.Request == "delete" {
		delete(req, fm, dict)
	} else if req.Request == "list" {
		list(req, fm, dict)
	} else {
		mainMenu(req, fm, dict)
	}
}

func Dispatcher(req *apptype.Common, fm *formatter.Formatter) {
	checkLanguage(req)
	if req.Request == "MainMenu" {
		mainMenu(req, fm, dict.Dictionary[req.Language])
	} else if req.Action == "divarication" {
		divarication(req, fm, dict.Dictionary[req.Language])
	} else if req.Action == "create" {
		create(req, fm, dict.Dictionary[req.Language])
	} else if req.Action == "change" {
		change(req, fm, dict.Dictionary[req.Language])
	} else if req.Action == "delete" {
		delete(req, fm, dict.Dictionary[req.Language])
	} else if req.Action == "list" {
		list(req, fm, dict.Dictionary[req.Language])
	} else {
		mainMenu(req, fm, dict.Dictionary[req.Language])
	}
	fm.WriteChatId(req.Id)
}

func makeTheDataCooler(row []interface{}) []string {
	res := make([]string, len(row))
	for i := 0; i < len(row); i++ {
		res[i] = fmt.Sprintf("%v", row[i])
	}
	return res
}

func prepareMessage(id int, inf []string) *formatter.Formatter {
	fm := new(formatter.Formatter)
	fm.WriteChatName("testdvijhurghada")
	fm.WriteString(fmt.Sprintf(dict.Dictionary["ru"]["SampleChannel"], inf[1], inf[2], fmt.Sprintf("%s %s", inf[3], inf[4]), inf[5]))
	fm.WriteParseMode(types.HTML)
	setKb(fm, []int{1}, []string{"Я приду!"}, []string{fmt.Sprint(id)})
	if err := fm.Complete(); err == nil {
		err = fm.SendToChannel()
		if err != nil {
			log.Printf("YOU HAVE AN ERROR DURING THE PROCCESS OF SENDING A MESSAGE TO A CHANNEL: %v", err)
		}
	} else {
		log.Printf("YOU HAVE AN ERROR DURING THE PROCCESS OF PREPARE A MESSAGE: %v", err)
	}
	return fm
}

func FromGoogle(row []interface{}) (fm *formatter.Formatter) {
	inf := makeTheDataCooler(row)
	id, err := saveToDatabase(inf)
	if err == nil {
		fm = prepareMessage(id, inf)
	}
	if err != nil {
		log.Printf("You made a mistake: %v", err)
	}
	return fm
}

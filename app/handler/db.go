package handler

import (
	"InfoBot/apptype"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func wichWay(userId int, f func(error)) (bool, error) {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Users WHERE isadmin = 1 AND userId = $1", userId).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0, err
}

func find(userId int, f func(error)) bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Users WHERE userId = $1", userId).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

func dbRetrieveUser(req *apptype.Common, f func(error)) {
	dbreq := "SELECT nickname, dvijid, action, level, titleru, titleen, discrpru, discrpen, delactivity, changeable, changesru, changesen FROM Users WHERE userId = $1"
	err := apptype.DB.QueryRow(dbreq, req.Id).Scan(&req.Nickname, &req.DvijId, &req.Action, &req.Level, &req.TitleRu, &req.TitleEn, &req.DiscrpRu, &req.DiscrpEn, &req.DelActivity, &req.Changeable, &req.ChangesRu, &req.ChangesEn)
	if err != nil {
		f(err)
	}
}

func createUser(req *apptype.Common, f func(error)) {
	_, err := apptype.DB.Exec("INSERT INTO Users (userId, action, isadmin) VALUES ($1, 'new', 0)", req.Id)
	if err != nil {
		f(err)
	}
	if req.Name != "" {
		_, err = apptype.DB.Exec("UPDATE Users SET name = $1 WHERE userid = $2", req.Name, req.Id)
		if err != nil {
			f(err)
		}
	}
	if req.Lastname != "" {
		_, err = apptype.DB.Exec("UPDATE Users SET lastname = $1 WHERE userid = $2", req.Lastname, req.Id)
		if err != nil {
			f(err)
		}
	}
	if req.Phone != "" {
		_, err = apptype.DB.Exec("UPDATE Users SET phone = $1 WHERE userid = $2", req.Phone, req.Id)
		if err != nil {
			f(err)
		}
	}
	req.Action = "new"
	req.Level = 0
}

func retainUser(req *apptype.Common, f func(error)) {
	_, err := apptype.DB.Exec(`UPDATE Users SET action = $1, level = $2, titleru = $4, titleen = $5, discrpru = $6, 
	discrpen = $7, delactivity = $8, changeable = $9, changesru = $10, changesen = $11, nickname = $12, dvijid = $13 WHERE userId = $3`,
		req.Action, req.Level, req.Id, req.TitleRu, req.TitleEn, req.DiscrpRu, req.DiscrpEn, req.DelActivity, req.Changeable, req.ChangesRu, req.ChangesEn, req.Nickname, req.DvijId)
	if err != nil {
		f(err)
	}
}

func changeStatus(userId, status int, f func(error)) {
	_, err := apptype.DB.Exec("UPDATE Users SET isadmin = $1 WHERE userId = $2", status, userId)
	if err != nil {
		f(err)
	}
}

func findActivity(actid int) bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Dvij WHERE id = $1 AND status != -1", actid).Scan(&count)
	if err != nil {
		log.Printf("YOU HAVE AN ERROR IN findActivity(): %v", err)
	}
	return count > 0
}

func findClientWithActivity(userid, actid int) bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM DvijClients WHERE id = $1 AND userid = $2 AND status != -1", actid, userid).Scan(&count)
	if err != nil {
		log.Printf("YOU HAVE AND ERROR IN findClientWithActivity(): %v", err)
	}
	return count == 0
}

func registerTheClientDB(userid, actid int) {
	_, err := apptype.DB.Exec("INSERT INTO DvijClients (id, userid, notiftime) VALUES ($1, $2, CURRENT_TIMESTAMP)", actid, userid)
	if err != nil {
		log.Printf("YOU HAVE AND ERROR IN registerTheClient(): %v", err)
	}
}

func selectUserId(interval, cond string) ([]*userAct, error) {
	var (
		useract    []*userAct
		userid, id int
	)
	query := fmt.Sprintf(`
        SELECT dcl.userid, dcl.id
        FROM DvijClients dcl
        JOIN Dvij d ON dcl.id = d.id
        WHERE d.datetime <= CURRENT_TIMESTAMP + INTERVAL '%s' %s AND dcl.status != -1`, interval, cond)
	rows, err := apptype.DB.Query(query)
	if err == nil {
		log.Print(rows)
		defer rows.Close()
		for rows.Next() && err == nil {
			row2 := new(userAct)
			err = rows.Scan(&userid, &id)
			if err == nil {
				row2.userid = userid
				row2.actid = id
				useract = append(useract, row2)
			}
		}
	}
	return useract, err
}

func selectActInf(actid int, f func(error)) (string, string, string, string) {
	var capt, dis, datetime, link string
	err := apptype.DB.QueryRow("SELECT caption, description, datetime, link FROM Dvij WHERE id = $1", actid).Scan(&capt, &dis, &datetime, &link)
	if err != nil {
		f(err)
	}
	return capt, dis, datetime, link
}

func changeNotifStatus(column string, userid, actid int, f func(error)) {
	query := fmt.Sprintf("UPDATE DvijClients SET %s = true WHERE userid = $1 AND id = $2 AND status != -1", column)
	_, err := apptype.DB.Exec(query, userid, actid)
	if err != nil {
		f(err)
	}
}

func changeAnswerStat(column string, userid, actid int, f func(error)) {
	_, err := apptype.DB.Exec(fmt.Sprintf("UPDATE DvijClients SET %s = 1 WHERE userid = $1 AND id = $2 AND status != -1", column), userid, actid)
	if err != nil {
		f(err)
	}
}

func deleteClientAct(userid, actid int, f func(error)) {
	_, err := apptype.DB.Exec("UPDATE DvijClients status = -1 WHERE userid = $1 AND id = $2", userid, actid)
	if err != nil {
		f(err)
	}
}

func TimeGuard() {
	rows, err := apptype.DB.Query("SELECT id FROM Dvij WHERE datetime < CURRENT_TIMESTAMP AND status != -1")
	if err == nil {
		log.Print("????", rows)
		defer rows.Close()
		id := 0
		for rows.Next() && err == nil {
			err = rows.Scan(&id)
			if err == nil {
				_, err = apptype.DB.Exec("UPDATE Dvij SET status = -1 WHERE id = $1", id)
				if err == nil {
					_, err = apptype.DB.Exec("UPDATE DvijClients SET status = -1 WHERE id = $1", id)
					if err != nil {
						log.Printf("Error after the app made a database request (3 time) in TimeGuard(): %s", err)
					}
				} else {
					log.Printf("Error after the app made a database request (2 time) in TimeGuard(): %s", err)
				}
			} else {
				log.Printf("Error while the app was scaning the rows from database in TimeGuard(): %s", err)
			}
		}
	} else {
		log.Printf("Error after the app made a database request (1 time) in TimeGuard(): %s", err)
	}
}

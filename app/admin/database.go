package admin

import (
	"InfoBot/apptype"
	"fmt"

	_ "github.com/lib/pq"
)

func saveInDB(req *apptype.Common, f func(error)) {
	_, err := apptype.DB.Exec("INSERT INTO Phrases (titleru, titleen, discrpru, discrpen, status) VALUES ($1, $2, $3, $4, 1)",
		req.TitleRu, req.TitleEn, req.DiscrpRu, req.DiscrpEn)
	if err != nil {
		f(err)
	}
}

func findActivity(name int, f func(error)) bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Phrases WHERE name = $1 AND status = 1", name).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

func deleteFromDB(name int, f func(error)) {
	_, err := apptype.DB.Exec("UPDATE Phrases SET status = 0 WHERE name = $1", name)
	if err != nil {
		f(err)
	}
}

func selectData(which string, nameAct int, f func(error)) (string, string) {
	var ru, en string
	err := apptype.DB.QueryRow(fmt.Sprintf("SELECT %sru, %sen FROM Phrases WHERE name = $1 AND status = 1", which, which), nameAct).Scan(&ru, &en)
	if err != nil {
		f(err)
	}
	return ru, en
}

func selectAllData(nameAct int, f func(error)) (string, string, string, string) {
	var ttru, tten, disru, disen string
	err := apptype.DB.QueryRow("SELECT titleru, titleen, discrpru, discrpen FROM Phrases WHERE name = $1 AND status = 1", nameAct).Scan(&ttru, &tten, &disru, &disen)
	if err != nil {
		f(err)
	}
	return ttru, tten, disru, disen
}

func saveChangesDB(req *apptype.Common, f func(error)) {
	_, err := apptype.DB.Exec(fmt.Sprintf("UPDATE Phrases SET %sru = $1, %sen = $2 WHERE name = $3 AND status = 1", req.Changeable, req.Changeable),
		req.ChangesRu, req.ChangesEn, req.DelActivity)
	if err != nil {
		f(err)
	}
}

func saveToDatabase(row []string) (int, error) {
	var (
		id       int
		datetime string
	)
	err := apptype.DB.QueryRow("SELECT nextval('dvij_id_seq')").Scan(&id)
	if err == nil {
		err = apptype.DB.QueryRow("SELECT TO_TIMESTAMP($1, 'DD.MM.YYYY HH24:MI:SS')::TIMESTAMPTZ", fmt.Sprintf("%s %s", row[3], row[4])).Scan(&datetime)
		if err == nil {
			_, err = apptype.DB.Exec("INSERT INTO Dvij (id, caption, description, datetime, link) VALUES ($1, $2, $3, $4, $5)",
				id, row[1], row[2], datetime, row[5])
		}
	}
	return id, err
}

func selectAllDvij(f func(error)) ([]string, []string) {
	var (
		ids        []string
		names      []string
		counter, i int
	)
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Dvij WHERE status != -1").Scan(&counter)
	if err != nil {
		f(err)
	} else {
		rows, err := apptype.DB.Query("SELECT id, caption FROM Dvij WHERE status != -1")
		if err != nil {
			f(err)
		} else {
			defer rows.Close()
			ids = make([]string, counter)
			names = make([]string, counter)
			for rows.Next() && err == nil {
				err = rows.Scan(&ids[i], &names[i])
				if err != nil {
					f(err)
				}
				i++
			}
		}
	}
	return ids, names
}

func findDvij(actid int, f func(error)) bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Dvij WHERE id = $1 AND status != -1", actid).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

func selectAllClients(actid int, f func(error)) ([]string, []string) {
	var (
		userids        []string
		nameslastnames []string
		name, lastname string
		counter, i     int
	)
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM DvijClients WHERE id = $1 AND status != -1", actid).Scan(&counter)
	if err != nil {
		f(err)
	} else {
		rows, err := apptype.DB.Query(`SELECT DISTINCT(u.userid), name, lastname 
				FROM Users u
				JOIN DvijClients dcl ON dcl.userid = u.userid `)
		if err != nil {
			f(err)
		} else {
			defer rows.Close()
			userids = make([]string, counter)
			nameslastnames = make([]string, counter)
			for rows.Next() && err == nil {
				err = rows.Scan(&userids[i], &name, &lastname)
				if err != nil {
					f(err)
				}
				nameslastnames[i] = fmt.Sprintf("%s %s", name, lastname)
				i++
			}
		}
	}
	return userids, nameslastnames
}

func findDvijClient(actid, userid int, f func(error)) bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM DvijClients WHERE userid = $1 AND id = $2 AND status != -1", userid, actid).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

func selectAllClientInf(userid, actid int, f func(error)) (string, string, string, string, int) {
	var (
		name, lastname, phone, nickname string
		confirmed, i, j                 int
	)
	err := apptype.DB.QueryRow("SELECT name, lastname, phone, nickname FROM Users WHERE userid = $1", userid).Scan(&name, &lastname, &phone, &nickname)
	if err != nil {
		f(err)
	}
	err = apptype.DB.QueryRow("SELECT answerone, answertwo FROM DvijClients WHERE userid = $1 and id = $2", userid, actid).Scan(&i, &j)
	if err != nil {
		f(err)
	}
	confirmed = i + j
	return name, lastname, phone, nickname, confirmed
}

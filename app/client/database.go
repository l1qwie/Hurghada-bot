package client

import (
	"InfoBot/apptype"
	"fmt"
	"log"
)

func selectDetails(req, lang string, f func(error)) string {
	var value string
	err := apptype.DB.QueryRow(fmt.Sprintf("SELECT discrp%s FROM Phrases WHERE name = $1 AND status = 1", lang), req).Scan(&value)
	if err != nil {
		f(err)
	}
	return value
}

func findName(req string, f func(error)) bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Phrases WHERE name = $1 AND status = 1", req).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

func lengthOfNames(f func(error)) []int {
	var length int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Phrases WHERE status = 1").Scan(&length)
	if err != nil {
		f(err)
	}
	crd := make([]int, length)
	for i := 0; i < length; i++ {
		crd[i] = 1
	}
	return crd
}

func selectNames(lang string, length int, f func(error)) []string {
	names := make([]string, length)
	rows, err := apptype.DB.Query(fmt.Sprintf("SELECT title%s FROM Phrases WHERE status = 1", lang))
	if err != nil {
		f(err)
	}
	i := 0
	for rows.Next() && err == nil {
		err := rows.Scan(&names[i])
		if err != nil {
			f(err)
		}
		i++
	}
	log.Print("NAMES = ", names)
	return names
}

func selectValues(names []string, lang string, length int, f func(error)) []string {
	var err error
	data := make([]string, length)
	for i := 0; i < length && err == nil; i++ {
		err = apptype.DB.QueryRow(fmt.Sprintf("SELECT name FROM Phrases WHERE status = 1 AND title%s = $1", lang), names[i]).Scan(&data[i])
		if err != nil {
			f(err)
		}
	}
	return data
}

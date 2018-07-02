package model

import (
	_ "github.com/mattn/go-sqlite3"
	"../../config"
	"fmt"
)

// Lender data structure
type Lend struct {
	LendID int
	KundenID  int
	ArtikelID int
	Beginn    string
	Rueckgabe string
}

type Lende map[int]*Lend

func GetAllLende() (verleihe [] Lend) {
	rows, err := config.Db.Query("select * from Lend")

	if err != nil {
		return
	}
	for rows.Next() {
		verleih := Lend{}

		err = rows.Scan(&verleih.LendID, &verleih.KundenID, &verleih.ArtikelID, &verleih.Beginn, &verleih.Rueckgabe)

		if err != nil {
			return
		}
		verleihe = append(verleihe, verleih)
	}
	rows.Close()
	return
}

// Delete User by id
func DeleteKundeByLend(id int) (err error) {
	//Get Items IDs und Update die Items Tabelle (Anzahl)
	_, err = config.Db.Exec("delete from Lend where KundenID = $1", id)
	return
}

func UpdateLend(date string, id int) (err error) {
	_, err = config.Db.Exec("update Items set RueckgabeAm = $1 where ItemID = $2", date, id)
	return
}

func LendItems(itemID int, username string, userID int, today string, rueckgabe string, anz int) (err error) {

	statement := "insert into Lend (UserID, ItemID) values (?,?)"
	stmt, err := config.Db.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = config.Db.Exec("update Items set AusgeliehenAm = $1 where ItemID = $2", today, itemID)
	_, err = config.Db.Exec("update Items set RueckgabeAm = $1 where ItemID = $2", rueckgabe, itemID)
	_, err = config.Db.Exec("update Items set entliehen = 1 where ItemID = $1", itemID)
	_, err = config.Db.Exec("update Items set Anzahl = Anzahl -1 where ItemID = $1", itemID)
	_, err = config.Db.Exec("update Items set Status = $1 where ItemID = $2", username, itemID)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID,itemID)

	return
}

func GetReturnDate(id int) (strings []Date) {
	rows, err := config.Db.Query("select rueckgabeAm from Items where ItemID=$1", id)

	if err != nil {
		return
	}
	for rows.Next() {
		string := Date{}
		err = rows.Scan(&string.RueckgabeAm)

		if err != nil {
			return
		}
		strings = append(strings, string)
	}
	rows.Close()
	return
}

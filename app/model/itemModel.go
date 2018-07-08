package model

import (
	"../../config"
	"fmt"
)

func CreateItem(bez string, kat string, invNum int, lgo string, anz int, in string, hin string, url string) {

	statement := "insert into Items (Bezeichnung, Kategorie, InventarNummer, Lagerort, Anzahl, Inhalt, Hinweis, BildURL, Status, AusgeliehenAm, RueckgabeAm, Entliehen, Noticed) values (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := config.Db.Prepare(statement)

	if err != nil {
		return
	}
	defer stmt.Close()

	if url == "" {
		_, err = stmt.Exec(bez, kat, invNum, lgo, anz, in, hin, "item_placeholder.png", "Verfügbar", "-", "-", 0,0)
	} else {
		_, err = stmt.Exec(bez, kat, invNum, lgo, anz, in, hin, url, "Verfügbar", "-", "-", 0,0)
	}

	return
}

// Update Item
func UpdateItem(id int, bez string, kat string, invNum int, lago string, in string, anz int, hin string, url string) (err error) {
	_, err = config.Db.Exec("update Items set Bezeichnung = $1 where ItemID = $2", bez, id)
	_, err = config.Db.Exec("update Items set Kategorie = $1 where ItemID = $2", kat, id)
	_, err = config.Db.Exec("update Items set InventarNummer = $1 where ItemID = $2", invNum, id)
	_, err = config.Db.Exec("update Items set Lagerort = $1 where ItemID = $2", lago, id)
	_, err = config.Db.Exec("update Items set Inhalt = $1 where ItemID = $2", in, id)
	_, err = config.Db.Exec("update Items set Anzahl = $1 where ItemID = $2", anz, id)
	_, err = config.Db.Exec("update Items set Hinweis = $1 where ItemID = $2", hin, id)

	if url != "" {
		_, err = config.Db.Exec("update Items set BildURL = $1 where ItemID = $2", url, id)
	}
	return
}

// Delete Items by id
func DeleteItem(id int) (err error) {
	_, err = config.Db.Exec("delete from Items where ItemID = $1", id)
	return
}

func GetEquipment() (Equipments []Equipment) {
	rows, err := config.Db.Query("select ItemID, BildURL, Bezeichnung, Anzahl, Hinweis FROM Items")

	if err != nil {
		return
	}

	equipment := Equipment{}

	for rows.Next() {
		err = rows.Scan(&equipment.ItemID, &equipment.BildURL, &equipment.Bezeichnung, &equipment.Anzahl, &equipment.Hinweis)

		Equipments = append(Equipments, equipment)

		if err != nil {
			return
		}
	}
	rows.Close()

	return
}

func GetAllBezeichnungenFromKundenArtikel(kunde_id int) (Bezeichnungen []string) {

	rows, err := config.Db.Query("select Items.Bezeichnung from Items,Lend where Lend.KundenID=$1 and Items.ItemID = Lend.ArtikelID", kunde_id)

	if err != nil {
		return
	}

	var temp = ""

	for rows.Next() {
		err = rows.Scan(&temp)

		Bezeichnungen = append(Bezeichnungen, temp)

		if err != nil {
			return
		}
	}
	rows.Close()

	return
}

func GetUserEquipment(kunde_id int) (equipments []MyItem) {
	rows, err := config.Db.Query("select Items.ItemID, Items.BildURL, Items.Bezeichnung, Items.InventarNummer, Items.Hinweis, Items.AusgeliehenAm, Items.RueckgabeAm, Items.Noticed from Items,Lend WHERE (Items.ItemID = Lend.ItemID AND Lend.UserID=$1) OR Items.Noticed = 1", kunde_id)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		equipment := MyItem{}

		err = rows.Scan(&equipment.ItemID, &equipment.BildURL, &equipment.Bezeichnung, &equipment.InventarNummer, &equipment.Hinweis, &equipment.Beginn, &equipment.Rueckgabe)

		if err != nil {
			return
		}

		equipments = append(equipments, equipment)
	}
	rows.Close()
	defer config.Db.Close()
	config.InitSQLiteDB()

	return
}

func GetNoticedEquipment() (equipments []MyItem) {
	rows, err := config.Db.Query("select ItemID, BildURL, Bezeichnung, InventarNummer, Hinweis, AusgeliehenAm, RueckgabeAm, Noticed from Items WHERE Noticed = 1")

	if err != nil {
		return
	}

	for rows.Next() {
		equipment := MyItem{}

		err = rows.Scan(&equipment.ItemID, &equipment.BildURL, &equipment.Bezeichnung, &equipment.InventarNummer, &equipment.Hinweis, &equipment.Beginn, &equipment.Rueckgabe, &equipment.Noticed)

		if err != nil {
			return
		}

		equipments = append(equipments, equipment)
	}
	rows.Close()
	return
}

func GetAdminEquipment() (adminEquipments []AdminItem) {
	rows, err := config.Db.Query("select ItemID, BildURL, Bezeichnung, InventarNummer, Lagerort, Hinweis, Status, RueckgabeAm, Entliehen FROM Items")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		adminEquipment := AdminItem{}

		err = rows.Scan(&adminEquipment.ItemID,&adminEquipment.BildURL,&adminEquipment.Bezeichnung, &adminEquipment.InventarNummer, &adminEquipment.Lagerort, &adminEquipment.Hinweis, &adminEquipment.Status, &adminEquipment.Rueckgabe, &adminEquipment.Entliehen)

		if err != nil {
			return
		}
		adminEquipments = append(adminEquipments, adminEquipment)
	}
	rows.Close()
	return
}

func GetChangeItem(id int) (changeItems []ChangeItem) {
	rows, err := config.Db.Query("select ItemID, BildURL, Bezeichnung, Kategorie, InventarNummer, Lagerort, Inhalt, Anzahl, Hinweis FROM Items WHERE ItemID = $1",id)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		changeItem := ChangeItem{}

		err = rows.Scan(&changeItem.ItemID,&changeItem.BildURL,&changeItem.Bezeichnung, &changeItem.Kategorie, &changeItem.InventarNummer, &changeItem.Lagerort, &changeItem.Inhalt, &changeItem.Anzahl, &changeItem.Hinweis)

		if err != nil {
			fmt.Println(err)
			return
		}
		changeItems = append(changeItems, changeItem)
	}

	rows.Close()
	return
}

func GetCartItems(id int) (cartItems []Cart) {
	rows, err := config.Db.Query("select ItemID, BildURL, Bezeichnung, InventarNummer, Hinweis FROM Items WHERE ItemID = $1",id)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		cartItem := Cart{Anz:1}

		err = rows.Scan(&cartItem.ItemID,&cartItem.BildURL,&cartItem.Bez, &cartItem.InvNum, &cartItem.Hinweis)

		if err != nil {
			fmt.Println(err)
			return
		}
		cartItems = append(cartItems, cartItem)
	}
	rows.Close()
	return
}

func UpdateNoticed(id int)(err error) {
	_, err = config.Db.Exec("update Items set Noticed = 1 where ItemID = $1", id)
	fmt.Println(err)
	defer config.Db.Close()
	config.InitSQLiteDB()
	return
}

func UpdateNoticedOff(id int)(err error) {
	_, err = config.Db.Exec("update Items set Noticed = 0 where ItemID = $1", id)
	defer config.Db.Close()
	config.InitSQLiteDB()
	fmt.Println(err)
	return
}


package model

import (
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"encoding/base64"
	"../../config"
	"fmt"
)

//type Kunden map[int]*User

func RegisterKunden(name string, mail string, psw string) {

	statement := "insert into User (Name, Passwort, Email, BildURL, Typ, Status, PasswortKlar) values (?,?,?,?,?,?,?)"
	stmt, err := config.Db.Prepare(statement)

	if err != nil {
		return
	}
	defer stmt.Close()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(psw), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)

	_, err = stmt.Exec(name,b64HashedPwd, mail, "user_placeholder.png","Benutzer","aktiv", psw)

	return
}

// GetAll Kunden
func GetAllUser() (users [] User){
	rows, err := config.Db.Query("select UserID, Name, Status, BildURL, Typ from User where Typ = 'Benutzer'")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.UserID,&user.Name,&user.Status,&user.BildUrl, &user.Typ)

		if err != nil {
			fmt.Println(err)
			return
		}

		users = append(users, user)
	}
	rows.Close()
	return
}

func GetProfile(kunde_id int) (profiles []Profile) {
	rows, err := config.Db.Query("select User.UserID, User.Name,User.BildUrl,User.Email,User.Status from User WHERE User.UserID = $1", kunde_id)

	if err != nil {
		return
	}
	for rows.Next() {
		profile := Profile{}

		err = rows.Scan(&profile.UserID,&profile.Name, &profile.BildURL, &profile.Mail, &profile.Status)

		if err != nil {
			return
		}

		profiles = append(profiles, profile)
	}
	rows.Close()
	return
}

// Delete User by id
func DeleteProfile(id int) (err error) {
	_, err = config.Db.Exec("delete from User where UserID = $1", id)

	DeleteKundeByLend(id)
	return
}

func GetKundenTyp(kunde_id int) (strings []Typ) {
	rows, err := config.Db.Query("select typ from User where UserID=$1", kunde_id)

	if err != nil {
		return
	}
	for rows.Next() {
		string := Typ{}
		err = rows.Scan(&string.Typ)

		if err != nil {
			return
		}
		strings = append(strings, string)
	}
	rows.Close()
	return
}

// Update the Profile Data
func UpdateProfile(id int, user string, mail string, psw string, url string) (err error) {

	if(psw != "") {
		hashedPwd, err := bcrypt.GenerateFromPassword([]byte(psw), 14)
		b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)
		if err != nil {
			fmt.Println(err)
		}
		_, err = config.Db.Exec("update User set Passwort = $1 where UserID = $2", b64HashedPwd, id)
		_, err = config.Db.Exec("update User set PasswortKlar = $1 where UserID = $2", psw, id)
	}

	_, err = config.Db.Exec("update User set Name = $1 where UserID = $2", user, id)
	_, err = config.Db.Exec("update User set Email = $1 where UserID = $2", mail, id)

	if(url != "") {
		_, err = config.Db.Exec("update User set BildURL = $1 where UserID = $2", url, id)
	}
	return
}

// Lock the Profile
func LockProfile(id int) (err error) {
	_, err = config.Db.Exec("update User set Status = gesperrt where UserID = $1", id)
	return
}

// GetUserByUsername retrieve Session by username
func GetUserByUsername(username string) (user SessionData, err error) {
	user = SessionData{}
	err = config.Db.QueryRow("select UserID, Name, Typ, Passwort from User where Name = $1", username).Scan(&user.ID, &user.Username, &user.Typ, &user.Password)

	return
}


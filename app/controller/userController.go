package controller

import (
	"html/template"
	"net/http"
	"../model"
	"strconv"
	"os"
	"io"
	"time"
)

type MyEquipment struct {
	Items []model.MyItem
}

type Equipment struct {
	Items      []model.Equipment
}

type Profiles struct {
	Items []model.Profile
}

func EquipmentPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		p := menu{
			Title:     "borgdir.media, index",
			Item1:     "Equipment,equipment",
			Item2:     "Login,login",
			Item3:     "",
			Basket:    false,
			Name:      "",
			Type:      "",
			ID:        nr,
			EmptySide: false,
			Profile:   false}

		EquipmentArr := model.GetEquipment()

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/equipment.html", "template/header.html", "template/layout.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "layout", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "equipment", Equipment{Items: EquipmentArr})
	} else {
		p := menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment",
			Item2:     "Meine Geräte,myequipment",
			Item3:     "Logout,logout",
			Basket:    true,
			Name:      session.Values["username"].(string),
			Type:      session.Values["typ"].(string),
			ID:        session.Values["id"].(int),
			EmptySide: false,
			Profile:   true}

		EquipmentArr := model.GetEquipment()

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/equipment.html", "template/header.html", "template/layout.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "layout", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "equipment", Equipment{Items: EquipmentArr})
	}
}

func Myequipment(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {
		p := menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment",
			Item2:     "Meine Geräte,myequipment",
			Item3:     "Logout,logout",
			Basket:    true,
			Name:      session.Values["username"].(string),
			Type:      session.Values["typ"].(string),
			ID:        session.Values["id"].(int),
			EmptySide: false,
			Profile:   true}

		ArtikelArr := model.GetUserEquipment(session.Values["id"].(int))

		if len(ArtikelArr) == 0 {
			ArtikelArr = model.GetNoticedEquipment()
		}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/myequipment.html", "template/header.html", "template/layout.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "layout", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "myequipment", MyEquipment{Items: ArtikelArr})
	}
}

func ExtendLend(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {
		i := r.URL.Path[len("/extend-lend/"):]
		id, err := strconv.Atoi(i)
		if err != nil {
			return
		}
		//Get old date
		date := model.GetReturnDate(id)[0].RueckgabeAm
		//Parse old date to Date format
		temp, err := time.Parse("02.01.2006", date)
		if err != nil {
			return
		}
		//Extend lend about 14 days
		temp = temp.AddDate(0,0,14)
		newTime := temp.Format("02.01.2006")

		model.UpdateLend(newTime, id)
		http.Redirect(w, r, "/myequipment", 301)
	}
}

func Profile(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {
		if r.Method == "POST" {
			i := r.URL.Path[len("/profile/"):]

			id, err := strconv.Atoi(i)
			if err != nil {
				return
			}

			user := r.FormValue("user")
			mail := r.FormValue("mail")
			psw := r.FormValue("psw")

			//Handle Uploading Image
			file, handler, err := r.FormFile("img")
			if (handler.Filename != "") {
				if err != nil {
					return
				}
				defer file.Close()

				src, err := handler.Open()
				if err != nil {
					return
				}
				defer src.Close()

				dst, err := os.Create("./static/images/"+handler.Filename)
				if err != nil {
					return
				}
				defer dst.Close()

				//Save Image in destination (static/images/"name")
				io.Copy(dst,src)
			}

			model.UpdateProfile(id, user, mail, psw,handler.Filename)
			session.Values["username"] = user
			session.Save(r,w)
			http.Redirect(w, r, "/profile/"+i, 301)
		} else {
			p := menu{
				Title:     "borgdir.media,index",
				Item1:     "Equipment,equipment",
				Item2:     "Meine Geräte,myequipment",
				Item3:     "Logout,logout",
				Basket:    true,
				Name:      session.Values["username"].(string),
				Type:      session.Values["typ"].(string),
				ID:        session.Values["id"].(int),
				EmptySide: false,
				Profile:   true}

			ProfilesArr := model.GetProfile(session.Values["id"].(int))

			tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/profile.html", "template/header.html", "template/layout.html"))

			tmpl.ExecuteTemplate(w, "main", p)
			tmpl.ExecuteTemplate(w, "layout", p)
			tmpl.ExecuteTemplate(w, "header", p)

			tmpl.ExecuteTemplate(w, "profile", Profiles{Items: ProfilesArr})
		}
	}
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {

		i := r.URL.Path[len("/delete/profile/"):]

		id, err := strconv.Atoi(i)
		if err != nil {
			return
		}
		model.DeleteProfile(id)
		http.Redirect(w, r, "/logout", 301)
	}
}

func NoticeItem (w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {
		i := r.URL.Path[len("/notice/"):]

		id, err := strconv.Atoi(i)
		if err != nil {
			return
		}
		model.UpdateNoticed(id)
		http.Redirect(w, r, "/equipment", 301)
	}
}

func NoticeOff (w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {
		i := r.URL.Path[len("/notice/off/"):]

		id, err := strconv.Atoi(i)
		if err != nil {
			return
		}
		model.UpdateNoticedOff(id)
		http.Redirect(w, r, "/myequipment", 301)
	}
}

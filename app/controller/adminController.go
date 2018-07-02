package controller

import (
	"html/template"
	"net/http"
	"../model"
	"strconv"
	"os"
	"io"
)

// /admin/clients Seiten Struct
type Clients struct {
	Items []Client
}
type Client struct {
	BildUrl       string
	Name          string
	UserID        int
	Typ           string
	Bezeichnungen []Bez
	Status        string
}
type Bez struct {
	Bezeichnung string
}

type AdminEquipments struct {
	Items []model.AdminItem
}

type ChangeItems struct {
	Item []model.ChangeItem
}



func Admin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {
		if session.Values["typ"].(string) == "Verleiher" {
			p := menu{
				Title:     "borgdir.media,index",
				Item1:     "Equipment,equipment",
				Item2:     "Kunden,admin/clients",
				Item3:     "Logout,logout",
				Basket:    false,
				Name:      session.Values["username"].(string),
				Type:      session.Values["typ"].(string),
				ID:        session.Values["id"].(int),
				EmptySide: false,
				Profile:   true}

			tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/admin.html", "template/header.html", "template/layout.html"))

			tmpl.ExecuteTemplate(w, "main", p)
			tmpl.ExecuteTemplate(w, "layout", p)
			tmpl.ExecuteTemplate(w, "header", p)
			tmpl.ExecuteTemplate(w, "admin", p)
		} else {
			http.Redirect(w, r, "/index", 301)
		}
	}
}

func AdminItems(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {
		if session.Values["typ"].(string) == "Verleiher" {
			p := menu{
				Title:     "borgdir.media,index",
				Item1:     "Equipment,equipment",
				Item2:     "Kunden,admin/clients",
				Item3:     "Logout,logout",
				Basket:    false,
				Name:      session.Values["username"].(string),
				Type:      session.Values["typ"].(string),
				ID:        session.Values["id"].(int),
				EmptySide: false,
				Profile:   true}

			EquipmentArr := model.GetAdminEquipment()

			tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminEquipment.html", "template/header.html", "template/layout.html"))

			tmpl.ExecuteTemplate(w, "main", p)
			tmpl.ExecuteTemplate(w, "layout", p)
			tmpl.ExecuteTemplate(w, "header", p)
			tmpl.ExecuteTemplate(w, "adminEquipment", AdminEquipments{Items: EquipmentArr})
		} else {
			http.Redirect(w, r, "/login", 301)
		}
	}
}

func AdminAddItem(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {
		if session.Values["typ"].(string) == "Verleiher" {
			if r.Method == "POST" {
				bez := r.FormValue("bz")
				kat := r.FormValue("kat")
				invNumStr:= r.FormValue("invNum")
				lgo := r.FormValue("lgo")
				in:= r.FormValue("inhalt")
				hin:= r.FormValue("hinweis")
				anzStr := r.FormValue("anz")

				invNum, err := strconv.Atoi(invNumStr)
				if err != nil {
					return
				}
				anz, err := strconv.Atoi(anzStr)
				if err != nil {
					return
				}

				model.CreateItem(bez, kat, invNum, lgo, anz, in, hin,)
				http.Redirect(w, r, "/admin/equipment", 301)
			} else {
				p := menu{
					Title:     "borgdir.media,index",
					Item1:     "Equipment,equipment",
					Item2:     "Kunden,admin/clients",
					Item3:     "Logout,logout",
					Basket:    false,
					Name:      session.Values["username"].(string),
					Type:      session.Values["typ"].(string),
					ID:        session.Values["id"].(int),
					EmptySide: false,
					Profile:   true}

				tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminAddEquipment.html", "template/header.html", "template/layout.html"))

				tmpl.ExecuteTemplate(w, "main", p)
				tmpl.ExecuteTemplate(w, "layout", p)
				tmpl.ExecuteTemplate(w, "header", p)
				tmpl.ExecuteTemplate(w, "adminAddEquipment", p)
			}
		} else {
			http.Redirect(w, r, "/login", 301)
		}
	}
}

func AdminUser(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {
		if session.Values["typ"].(string) == "Verleiher" {
			p := menu{
				Title:     "borgdir.media,index",
				Item1:     "Equipment,equipment",
				Item2:     "Kunden,admin/clients",
				Item3:     "Logout,logout",
				Basket:    false,
				Name:      session.Values["username"].(string),
				Type:      session.Values["typ"].(string),
				ID:        session.Values["id"].(int),
				EmptySide: false,
				Profile:   true}

			//Alle Kunden auslesen
			KundenArr := model.GetAllUser()
			ClientsArr := []Client{}

			for _, element := range KundenArr {
				EquipmentString := []Bez{}

				artikelFromUser := model.GetAllBezeichnungenFromKundenArtikel(element.UserID)

				for _, element := range artikelFromUser {

					EquipmentString = append(EquipmentString, Bez{element})
				}
				ClientsArr = append(ClientsArr, Client{element.BildUrl, element.Name, element.UserID, element.Typ, EquipmentString, element.Status})
			}
			data := Clients{
				Items: ClientsArr,
			}

			tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/clients.html", "template/header.html", "template/layout.html"))

			tmpl.ExecuteTemplate(w, "main", nil)
			tmpl.ExecuteTemplate(w, "layout", p)
			tmpl.ExecuteTemplate(w, "header", p)
			tmpl.ExecuteTemplate(w, "clients", data)
		} else {
			http.Redirect(w, r, "/index", 301)
		}
	}
}

func AdminChangeItem(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {
		if session.Values["typ"].(string) == "Verleiher" {
			if r.Method == "POST" {
				i := r.URL.Path[len("/admin/change/item/"):]
				id, err := strconv.Atoi(i)

				if err != nil {
					return
				}

				bez := r.FormValue("bz")
				kat := r.FormValue("kat")
				invNumStr:= r.FormValue("invNum")
				lgo := r.FormValue("lgo")
				in:= r.FormValue("inhalt")
				hin:= r.FormValue("hinweis")
				anzStr := r.FormValue("anz")

				invNum, err := strconv.Atoi(invNumStr)
				if err != nil {
					return
				}
				anz, err := strconv.Atoi(anzStr)
				if err != nil {
					return
				}

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
				model.UpdateItem(id,bez, kat, invNum, lgo, in, anz, hin, handler.Filename)
				http.Redirect(w, r, "/admin/change/item/"+i, 301)
			} else {
				i := r.URL.Path[len("/admin/change/item/"):]
				id, err := strconv.Atoi(i)

				if err != nil {
					return
				}

				ItemArr := model.GetChangeItem(id)

				p := menu{
					Title:     "borgdir.media,index",
					Item1:     "Equipment,equipment",
					Item2:     "Kunden,admin/clients",
					Item3:     "Logout,logout",
					Basket:    false,
					Name:      session.Values["username"].(string),
					Type:      session.Values["typ"].(string),
					ID:        session.Values["id"].(int),
					EmptySide: false,
					Profile:   true}

				tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminChangeItem.html", "template/header.html", "template/layout.html"))

				tmpl.ExecuteTemplate(w, "main", p)
				tmpl.ExecuteTemplate(w, "layout", p)
				tmpl.ExecuteTemplate(w, "header", p)
				tmpl.ExecuteTemplate(w, "adminChangeItem", ChangeItems{Item: ItemArr})
			}
		} else {
			http.Redirect(w, r, "/login", 301)
		}
	}
}

func AdminEditUser(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {
		if session.Values["typ"].(string) == "Verleiher" {
			if r.Method == "POST" {
				i := r.URL.Path[len("/admin/edit-client/"):]

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
				http.Redirect(w, r, "/admin/edit-client/"+i, 301)
			} else {
				i := r.URL.Path[len("/admin/edit-client/"):]
				id, err := strconv.Atoi(i)

				if err != nil {
					return
				}

				p := menu{
					Title:     "borgdir.media,index",
					Item1:     "Equipment,equipment",
					Item2:     "Kunden,admin/clients",
					Item3:     "Logout,logout",
					Basket:    false,
					Name:      session.Values["username"].(string),
					Type:      session.Values["typ"].(string),
					ID:        session.Values["id"].(int),
					EmptySide: false,
					Profile:   true}

				ClientArr := model.GetProfile(id)

				tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/adminEditClients.html", "template/header.html", "template/layout.html"))

				tmpl.ExecuteTemplate(w, "main", p)
				tmpl.ExecuteTemplate(w, "layout", p)
				tmpl.ExecuteTemplate(w, "header", p)

				tmpl.ExecuteTemplate(w, "adminEditClients", Profiles{Items: ClientArr})
			}
		} else {
			http.Redirect(w, r, "/index", 301)
		}
	}
}

func AdminDeleteItem(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Redirect(w, r, "/login", 301)
	} else {

		i := r.URL.Path[len("/admin/delete/item/"):]

		id, err := strconv.Atoi(i)
		if err != nil {
			return
		}
		model.DeleteItem(id)
		http.Redirect(w, r, "/admin/equipment", 301)
	}
}
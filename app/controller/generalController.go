package controller

import (
	"html/template"
	"net/http"
	"strings"
	"github.com/gorilla/sessions"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"crypto/rand"
	"../model"
	"encoding/gob"
	"strconv"
	"time"
)

//Session ID
var nr = 0
var store *sessions.CookieStore
//var tmpl *template.Template
//Session Data from Session

type menu struct {
	Title     string
	Item1     string
	Item2     string
	Item3     string
	Basket    bool
	Name      string
	Type      string
	ID        int
	Profil    bool
	EmptySide bool
	Profile   bool
}

type cartItems struct {
	Items []model.Cart
}

func init() {
	//tmpl = template.Must(template.ParseGlob("template/*.tmpl"))

	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key := make([]byte, 32)
	rand.Read(key)
	store = sessions.NewCookieStore(key)
	store.MaxAge(3600)

	gob.Register(model.Cart{})
	gob.Register(cartItems{})
}

func (cItems *cartItems) AddItem(item model.Cart) []model.Cart {
	cItems.Items = append(cItems.Items, item)
	return cItems.Items
}

var funcMap = template.FuncMap{
	"split": func(s string, index int) string {
		arr := strings.Split(s, ",")

		if s == "" {
			return ""
		} else {
			return arr[index]
		}

	},
}

func Index(w http.ResponseWriter, r *http.Request) {
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
			ID:        0,
			EmptySide: false,
			Profile:   false}

		var tmpl = template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/header.html", "template/layout.html", "template/index.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "layout", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "index", p)
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

		var tmpl = template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/header.html", "template/layout.html", "template/index.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "layout", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "index", p)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if (r.Method == "POST") {

		username := r.FormValue("user")
		password := r.FormValue("psw")

		// Authentication
		user, _ := model.GetUserByUsername(username)

		// decode base64 String to []byte
		passwordDB, _ := base64.StdEncoding.DecodeString(user.Password)
		err := bcrypt.CompareHashAndPassword(passwordDB, []byte(password))

		if user.Status == "gesperrt" {
			http.Redirect(w, r, "/login", 301)
		} else {
			if err == nil {
				// Set user as authenticated
				session.Values["authenticated"] = true
				session.Values["username"] = username
				session.Values["id"] = user.ID
				session.Values["typ"] = user.Typ
				session.Save(r, w)

				if (model.GetKundenTyp(session.Values["id"].(int))[0].Typ == "Benutzer") {
					http.Redirect(w, r, "/index", 301)
				} else {
					http.Redirect(w, r, "/admin", 301)
				}
			} else {
				http.Redirect(w, r, "/login", 301)
			}
		}

	} else {
		//Header Variablen setzen
		p := menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment",
			Item2:     "Registrieren,register",
			Item3:     "",
			Basket:    false,
			Name:      "",
			Type:      "",
			ID:        0,
			EmptySide: true,
			Profile:   false}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/header.html", "template/layout.html", "template/login.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "layout", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "login", p)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	session.Values["authenticated"] = false
	session.Values["username"] = ""
	session.Values["id"] = 0
	session.Values["typ"] = ""
	session.Values["save"] = false
	session.Values["cartLength"] = nil
	session.Values["cart"] = cartItems{}
	session.Save(r, w)

	p := menu{
		Title:     "borgdir.media, index",
		Item1:     "Equipment,equipment",
		Item2:     "Login,login",
		Item3:     "",
		Basket:    false,
		Name:      "",
		Type:      "",
		ID:        0,
		EmptySide: false,
		Profile:   false}

	var tmpl = template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/header.html", "template/layout.html", "template/index.html"))

	tmpl.ExecuteTemplate(w, "main", p)
	tmpl.ExecuteTemplate(w, "layout", p)
	tmpl.ExecuteTemplate(w, "header", p)
	tmpl.ExecuteTemplate(w, "index", p)
}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("user")
		mail := r.FormValue("mail")
		psw := r.FormValue("psw")

		model.RegisterKunden(name, mail, psw)
		http.Redirect(w, r, "/index", 301)
	} else {

		// REGISTER
		p := menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment",
			Item2:     "Login,login",
			Item3:     "",
			Basket:    false,
			Name:      "",
			Type:      "",
			ID:        0,
			EmptySide: true,
			Profile:   false}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/register.html", "template/layout.html", "template/header.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "layout", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "register", p)
	}
}

func Cart(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		p := menu{
			Title:     "borgdir.media,index",
			Item1:     "Equipment,equipment",
			Item2:     "Login,login",
			Item3:     "",
			Basket:    false,
			Name:      "",
			Type:      "",
			ID:        0,
			EmptySide: false,
			Profile:   false}

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/cartGuest.html", "template/header.html", "template/layout.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "layout", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "cartGuest", session.Values["cart"])
	} else {
		if r.Method == "POST" {
			username := session.Values["username"].(string)
			userID := session.Values["id"].(int)

			if session.Values["cartLength"] != nil {
				for i := 0; i < session.Values["cartLength"].(int); i++ {
					anz := session.Values["cart"].(cartItems).Items[i].Anz
					itemID := session.Values["cart"].(cartItems).Items[i].ItemID

					timeToday := time.Now()
					today := timeToday.String()
					today = timeToday.Format("02.01.2006")

					timeToday = timeToday.AddDate(0, 0, 14)
					rueckgabe := timeToday.Format("02.01.2006")

					model.LendItems(itemID, username, userID, today, rueckgabe, anz)
				}
			}

			session.Values["save"] = false
			session.Values["cartLength"] = nil
			session.Values["cart"] = cartItems{}
			session.Save(r, w)
		}
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

		tmpl := template.Must(template.New("main").Funcs(funcMap).ParseFiles("template/cartUser.html", "template/header.html", "template/layout.html"))

		tmpl.ExecuteTemplate(w, "main", p)
		tmpl.ExecuteTemplate(w, "layout", p)
		tmpl.ExecuteTemplate(w, "header", p)
		tmpl.ExecuteTemplate(w, "cartUser", session.Values["cart"])
	}
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	i := r.URL.Path[len("/add-to-cart/"):]
	id, err := strconv.Atoi(i)

	if err != nil {
		return
	}

	cItems := cartItems{}

	if session.Values["save"] == true {
		cItems = session.Values["cart"].(cartItems)
	}

	//logic
	cartArr := model.GetCartItems(id)
	item := model.Cart{
		cartArr[0].ItemID,
		cartArr[0].BildURL,
		cartArr[0].Bez,
		cartArr[0].InvNum,
		cartArr[0].Hinweis,
		cartArr[0].Anz,
		cartArr[0].Rueckgabe}

	cItems.AddItem(item)

	session.Values["save"] = true
	session.Values["cart"] = cItems
	if (session.Values["cartLength"] == nil) {
		session.Values["cartLength"] = 1
	} else {
		session.Values["cartLength"] = session.Values["cartLength"].(int) + 1
	}
	session.Save(r, w)

	http.Redirect(w, r, "/equipment", 301)
}

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	i := r.URL.Path[len("/remove-from-cart/"):]
	id, err := strconv.Atoi(i)

	if err != nil {
		return
	}
	temp := cartItems{}

	for i := 0; i < session.Values["cartLength"].(int); i++ {
		if (id != session.Values["cart"].(cartItems).Items[i].ItemID) {
			temp.AddItem(session.Values["cart"].(cartItems).Items[i])
		}
	}

	session.Values["cartLength"] = session.Values["cartLength"].(int) - 1
	session.Values["cart"] = temp

	if session.Values["cartLength"] == 0 {
		session.Values["save"] = false
	}

	session.Save(r, w)
	http.Redirect(w, r, "/cart", 301)
}

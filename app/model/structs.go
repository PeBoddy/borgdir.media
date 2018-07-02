package model

//Session Data from Session
type SessionData struct {
	ID       int
	Username string
	Typ      string
	Password string
}

// User data structure
type User struct {
	UserID  int
	Name    string
	BildUrl string
	Typ     string
	Status  string
}

// Typ
type Typ struct {
	Typ string
}

//Date
type Date struct {
	RueckgabeAm string
}

type Cart struct {
	ItemID    int
	BildURL   string
	Bez       string
	InvNum    int
	Hinweis   string
	Anz       int
	Rueckgabe string
}

// /myequipment Seiten Struct
type MyItem struct {
	ItemID         int
	BildURL        string
	Bezeichnung    string
	InventarNummer int
	Hinweis        string
	Beginn         string
	Rueckgabe      string
}

// /admin/equipment Seiten Struct
type AdminItem struct {
	ItemID         int
	BildURL        string
	Bezeichnung    string
	InventarNummer int
	Lagerort       string
	Hinweis        string
	Status         string
	Rueckgabe      string
	Entliehen      int
}

// /equipment Seiten Struct
type Equipment struct {
	ItemID      int
	BildURL     string
	Bezeichnung string
	Anzahl      int
	Hinweis     string
}

type ChangeItem struct {
	ItemID         int
	BildURL        string
	Bezeichnung    string
	Kategorie      string
	InventarNummer int
	Lagerort       string
	Inhalt         string
	Anzahl         int
	Hinweis        string
}

// /admin/edit-clients Seiten Struct
type Profile struct {
	UserID  int
	Name    string
	BildURL string
	Mail    string
	Status  string
}

package main

import (
	"net/http"
	"./app/controller"
	"./config"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", controller.Index)
	r.HandleFunc("/index", controller.Index)
	r.HandleFunc("/admin", controller.Admin)
	r.HandleFunc("/admin/equipment", controller.AdminItems)
	r.HandleFunc("/admin/add", controller.AdminAddItem)
	r.HandleFunc("/admin/change/item/{id}", controller.AdminChangeItem)
	r.HandleFunc("/admin/delete/item/{id}", controller.AdminDeleteItem)
	r.HandleFunc("/admin/clients", controller.AdminUser)
	r.HandleFunc("/admin/edit-client/{id}", controller.AdminEditUser)
	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/logout", controller.Logout)
	r.HandleFunc("/equipment", controller.EquipmentPage)
	r.HandleFunc("/myequipment", controller.Myequipment)
	r.HandleFunc("/extend-lend/{id}", controller.ExtendLend)
	r.HandleFunc("/register", controller.Register)
	r.HandleFunc("/cart", controller.Cart)
	r.HandleFunc("/add-to-cart/{id}", controller.AddToCart)
	r.HandleFunc("/remove-from-cart/{id}", controller.RemoveFromCart)
	r.HandleFunc("/profile/{id}", controller.Profile)
	r.HandleFunc("/delete/profile/{id}", controller.DeleteProfile)
	r.HandleFunc("/lock/profile/{id}", controller.LockProfile)

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/static/", fs)

	config.InitSQLiteDB()
	//config.InitPostgresDB()

	http.Handle("/", r)
	http.ListenAndServe(":80",nil)
}


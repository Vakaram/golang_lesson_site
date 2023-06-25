package main

//
//import (
//	"fmt"
//	"html/template"
//	"log"
//	"net/http"
//)
//
//type User struct {
//	Name       string
//	Age        uint16
//	Money      int16
//	Avg_grades float64
//	Happines   float64
//	Hobbies    []string
//}
//
//func (u User) getAllInfo() string {
//	return fmt.Sprintf("User %s,Age %d, money %d ", u.Name, u.Age, u.Money)
//}
//
//func (u *User) setNewName(newName string) {
//	u.Name = newName
//}
//
//func home_page(w http.ResponseWriter, r *http.Request) {
//	bob := User{"Boby", 24, -50, 4.5, 0.8, []string{"Money",
//		"Paty Barge", "Sport"}}
//	//fmt.Fprintf(w, bob.getAllInfo())
//	tmpl, _ := template.ParseFiles("templates/home_page.html")
//	tmpl.Execute(w, bob)
//}
//
//func contact_page(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Contact page")
//}
//func handeleRequest() {
//	http.HandleFunc("/", home_page)             // each request calls handler
//	http.HandleFunc("/contacts/", contact_page) // each request calls handler
//	log.Fatal(http.ListenAndServe("localhost:8000", nil))
//
//}
//
//func main() {
//	handeleRequest()
//}

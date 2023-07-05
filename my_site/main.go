package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Article struct {
	Id                      uint16
	Title, Anons, Full_text string
}

var posts = []Article{}
var showPost = Article{}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("my_site/templates/index.html", "my_site/templates/header.html", "my_site/templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "db_user:db_user_pass@tcp(193.17.92.177:3306)/app_db") // подключаемся к БД рот рот это лог пар
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// выборка данных
	res, err := db.Query("SELECT * FROM `articalce`")
	if err != nil {
		panic(err)
	}
	posts = []Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
		//fmt.Println(fmt.Sprintf("Post: %s Id: %d ", post.Title, post.Id))
	}
	t.ExecuteTemplate(w, "index", posts)
}
func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("my_site/templates/create.html", "my_site/templates/header.html", "my_site/templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "create", nil) // подключение блока

}
func save_arcticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	//проверка
	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Одно из полей не заполненно нужно заполнить все поля")
		//time.Sleep(time.Second * 1)
		//http.Redirect(w, r, "/create/", http.StatusSeeOther)
	} else {

		db, err := sql.Open("mysql", "db_user:db_user_pass@tcp(193.17.92.177:3306)/app_db") // подключаемся к БД рот рот это лог пар
		if err != nil {
			panic(err)
		}
		defer db.Close()
		//добавляем данные в бд
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articalce` (`title`, `anons`,`full_text`) VALUES('%s', '%s','%s')", title, anons, full_text))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("перевожу в нужную страницу")
	}
}

func show_post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "ID: %v\n", vars["id"]) // сюда пишем такой же id  кка и написаали в rtr. post/{id...}
	db, err := sql.Open("mysql", "db_user:db_user_pass@tcp(193.17.92.177:3306)/app_db") // подключаемся к БД рот рот это лог пар
	if err != nil {
		panic(err)
	}
	defer db.Close()

	t, err := template.ParseFiles("my_site/templates/show.html", "my_site/templates/header.html", "my_site/templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	//получаем данные
	res, err := db.Query(fmt.Sprintf("SELECT * FROM `articalce` WHERE `id`= '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}
	showPost = Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text)
		if err != nil {
			panic(err)
		}
		showPost = post
		//fmt.Println(fmt.Sprintf("Post: %s Id: %d ", post.Title, post.Id))
	}
	t.ExecuteTemplate(w, "show", showPost)

}

func hendleFunc() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create/", create).Methods("GET")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")
	rtr.HandleFunc("/save_article/", save_arcticle).Methods("POST")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe("193.17.92.177:1111", nil)

}

func main() {
	hendleFunc()
}

//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//)
//
//type User struct {
//	Name string `json:"name"`
//	Age  uint16 `json:"age"`
//}
//
//func main() {
//
//	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang_users") // подключаемся к БД рот рот это лог пар
//	if err != nil {
//		panic(err)
//	}
//
//	defer db.Close()
//
//	//insert, err := db.Query("INSERT INTO `user` (`name`, `age`) VALUES('Jon', 33)")
//	//if err != nil {
//	//	panic(err)
//	//}
//	//defer insert.Close()
//
//	res, err := db.Query("SELECT `name`,`age` FROM `user` ")
//
//	if err != nil {
//		panic(err)
//	}
//	for res.Next() {
//		var user User
//		err = res.Scan(&user.Name, &user.Age)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(fmt.Sprintf("Name: %s wint age %d ", user.Name, user.Age))
//
//	}
//
//	fmt.Println("Connect in MySQL")
//
//}

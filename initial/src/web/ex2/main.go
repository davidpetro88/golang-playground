package main

import (
	"fmt"
	"net/http"
	"html/template"
	_"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"database/sql"
)

type Post struct {
	Id int
	Title string
	Body string
}

var db, err = sql.Open("mysql","root:123456@/go_course?charset=utf8")

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{id}/view", ViewHandler)
	fmt.Println(http.ListenAndServe(":8082",r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("templates/layout.html","templates/list.html"))
	if err := t.ExecuteTemplate(w,"layout.html",ListsPosts()); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request)  {
	id := mux.Vars(r)["id"]
	t := template.Must(template.ParseFiles("templates/layout.html","templates/view.html"))
	t.ExecuteTemplate(w,"layout.html",GetPostById(id))
}

func ListsPosts() []Post  {
	rows, err := db.Query("Select * from posts")
	checkErr(err)
	items := []Post{}
	for rows.Next(){
		post := Post{}
		rows.Scan(&post.Id,&post.Title,&post.Body)
		items = append(items, post)
	}
	return items
}

func GetPostById(id string) Post  {
	row := db.QueryRow("select * from posts where id = ?", id)
	post := Post{}
	row.Scan(&post.Id,&post.Title,&post.Body)
	return post
}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}
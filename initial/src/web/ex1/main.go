package main

import (
	"fmt"
	"net/http"
	"html/template"
	_"github.com/go-sql-driver/mysql"
	//"database/sql"
	//_"github.com/go-sql-driver/mysql"
	"database/sql"
)

type Post struct {
	Id int
	Title string
	Body string
}

var db, err = sql.Open("mysql","root:123456@/go_course?charset=utf8")

func main() {

	//stmt, err := db.Prepare("Insert into posts(title,body) values(?,?) ")
	//checkErr(err)
	//_, err = stmt.Exec("My First Post","My first content" )
	//checkErr(err)
	//db.Close()

	rows, err := db.Query("Select * from posts")
	checkErr(err)
	items := []Post{}
	for rows.Next(){
		post := Post{}
		rows.Scan(&post.Id,&post.Title,&post.Body)
		items = append(items, post)
	}
	db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w,"index.html",items); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	http.HandleFunc("/ex", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	http.HandleFunc("/ex2", func(w http.ResponseWriter, r *http.Request) {
		// http://localhost:8082/ex2
		// http://localhost:8082/ex2?title=My First Post
		post := Post{Id:1, Title: "Unamed Post", Body:"No Content"}
		if title := r.FormValue("title"); title != "" {
			post.Title = title
		}

		if title := r.FormValue("title"); title != "" {
			post.Title = title
		}

		t := template.Must(template.ParseFiles("templates/ex2.html"))

		if err := t.ExecuteTemplate(w,"ex2.html",post); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8082",nil))
}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}
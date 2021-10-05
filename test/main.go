package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/ddynamic/godatatables"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	env, _ := godotenv.Read(".env")

	if env["DATABASE_URL"] == "" {
		env["DATABASE_URL"] = "root:@tcp(127.0.0.1:3306)/datatables?parseTime=true&charset=utf8mb4,utf8"
	}

	godotenv.Write(env, ".env")

	godotenv.Load(".env")

	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err)
	}

	tmpl := template.Must(template.ParseFiles("test.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		godatatables.DataTables(w, r, db, "person" /*"nama = 'udean' && tempat = 'jember'"*/, "", "",
			godatatables.Column{Name: "id", Display: "id"},
			godatatables.Column{Name: "nama", Display: "nama"},
			godatatables.Column{Name: "nik", Display: "nik"},
			godatatables.Column{Name: "telp", Display: "telp"},
			godatatables.Column{Name: "tgl", Display: "tgl"},
			godatatables.Column{Name: "tempat", Display: "tempat"})
	})

	http.ListenAndServe(":8080", nil)
}

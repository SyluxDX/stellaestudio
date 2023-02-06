package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"stellaeestudio/cmd/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

// func debug() {
// 	// command line flags
// 	var databaseFilepath string
// 	flag.StringVar(&databaseFilepath, "db", "./stellae.db", "Path to sqlite3 database.")
// 	flag.Parse()

// 	log.Printf("Connect to db %s\n", databaseFilepath)
// 	conn, err := sqlite.Open(databaseFilepath)
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	defer conn.Close()

// 	profs, err := sqlite.GetProfessors(conn)
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	log.Println("professors results:")
// 	fmt.Printf("%+v\n", profs)

// 	classes, err := sqlite.GetClasses(conn)
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	log.Println("Classes results:")
// 	fmt.Printf("%+v\n", classes)

// 	studs, err := sqlite.GetStudents(conn)
// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	log.Println("Students results:")
// 	fmt.Printf("%+v\n", studs)

// 	ocpu, err := sqlite.GetClassOcupancy(conn)

// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	log.Println("ClassOcupancy results:")
// 	fmt.Printf("%+v\n", ocpu)
// }

var (
	DBConn *sql.DB
)

func searchHandler(students []sqlite.StudentName) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templatePage, err := template.ParseFiles("html/index.html")
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := templatePage.Execute(w, students); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func getStudentInfo(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	id := req.Form["studentid"][0]
	// for key, value := range req.Form {
	// 	fmt.Printf("%s = %s\n", key, value)
	// }
	student, _ := sqlite.GetStudentById(DBConn, id)
	fmt.Println(student)

	// write output
	w.Header().Add("Content-Type", "application/json")
	log.Println("Endpoint Hit: students")
	json.NewEncoder(w).Encode(student)
}

func main() {
	// command line flags
	var databaseFilepath string
	flag.StringVar(&databaseFilepath, "db", "./stellae.db", "Path to sqlite3 database.")
	flag.Parse()

	log.Printf("Connect to db %s\n", databaseFilepath)
	// conn, err := sqlite.Open(databaseFilepath)
	Conn, err := sqlite.Open(databaseFilepath)
	if err != nil {
		log.Panicln(err)
	}
	DBConn = Conn
	defer DBConn.Close()

	studs, err := sqlite.GetStudentsNames(Conn)
	if err != nil {
		log.Panicln(err)
	}

	id := "5"
	student, _ := sqlite.GetStudentById(Conn, id)
	fmt.Println(student)
	fmt.Println()
	fmt.Println("Conn", Conn)

	// Server Configurations
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "html/index.html")
	// })
	// http.HandleFunc("/confirm", uploadFiles)
	http.HandleFunc("/", searchHandler(studs))
	http.HandleFunc("/studentinfo", getStudentInfo)

	server := "localhost"
	port := 8080

	log.Printf("Serving on %[1]s port %[2]d (http://%[1]s:%[2]d/)\n", server, port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", server, port), nil)
}

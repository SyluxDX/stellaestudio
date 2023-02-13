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

// type formData struct {
// }

type StudentDetail struct {
	Name string
	NIF  string
}

type formInfo struct {
	StudentsNames []sqlite.StudentName
	Subscriptions []sqlite.Subscription
	Promotions    []sqlite.Promotion
	Detail        StudentDetail
}

var (
	DBConn   *sql.DB
	FormInfo formInfo
)

func searchHandler(students []sqlite.StudentName) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Search handler hit")
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

func searchFormHandler(data formInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Form Search handler hit")
		templatePage, err := template.ParseFiles("html/form.html")
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := templatePage.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
func parseForm(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	for key, value := range req.Form {
		fmt.Printf("%s = %s\n", key, value)
	}
}

// func searchFormHandler(w http.ResponseWriter, req *http.Request) {
// 	log.Println("Form Search handler hit")

// 	// get student info
// 	req.ParseForm()
// 	id := req.Form["studentid"][0]
// 	student, _ := sqlite.GetStudentById(DBConn, id)

// 	detail := StudentDetail{Name: student.Name, NIF: fmt.Sprint(student.Nif)}

// 	FormInfo.Detail = detail
// 	templatePage, err := template.ParseFiles("html/index.html")
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	if err := templatePage.Execute(w, FormInfo); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

func getStudentInfo(w http.ResponseWriter, req *http.Request) {
	log.Println("Endpoint Hit: students")
	// var bodyBytes []byte

	// bodyBytes, _ = io.ReadAll(req.Body)
	// // defer req.Body.Close()
	// fmt.Println("body", string(bodyBytes))

	// req.ParseForm()
	// for key, value := range req.Form {
	// 	fmt.Printf("%s = %s\n", key, value)
	// }

	req.ParseForm()
	id := req.Form["studentid"][0]
	student, _ := sqlite.GetStudentById(DBConn, id)
	fmt.Println(student)

	// write output
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func serveForm(w http.ResponseWriter, req *http.Request) {
	log.Println("Endpoint Hit: Serve Form")
	log.Println(req.Method)
	templatePage, err := template.ParseFiles("html/form.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templatePage.Execute(w, FormInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	// command line flags
	var databaseFilepath string
	flag.StringVar(&databaseFilepath, "db", "./stellae.db", "Path to sqlite3 database.")
	flag.Parse()

	log.Printf("Connect to db %s\n", databaseFilepath)
	var err error
	// conn, err := sqlite.Open(databaseFilepath)
	DBConn, err = sqlite.Open(databaseFilepath)
	if err != nil {
		log.Panicln(err)
	}

	defer DBConn.Close()

	FormInfo.StudentsNames, err = sqlite.GetStudentsNames(DBConn)
	if err != nil {
		log.Panicln(err)
	}

	FormInfo.Subscriptions, err = sqlite.GetSubscriptions(DBConn)
	if err != nil {
		log.Panicln(err)
	}

	FormInfo.Promotions, err = sqlite.GetPromos(DBConn)
	if err != nil {
		log.Panicln(err)
	}

	data, _ := json.MarshalIndent(FormInfo, "", " ")
	fmt.Println(string(data))

	// use handler to send variable to http function
	http.HandleFunc("/", searchHandler(FormInfo.StudentsNames))

	// use global variables to send variable to http function
	http.HandleFunc("/form", searchFormHandler(FormInfo))

	// set favivon.ico
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/studentinfo", getStudentInfo)
	http.HandleFunc("/parseForm", parseForm)

	// Server Configurations
	server := "localhost"
	port := 8080
	// start server
	log.Printf("Serving on %[1]s port %[2]d (http://%[1]s:%[2]d/)\n", server, port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", server, port), nil)
}

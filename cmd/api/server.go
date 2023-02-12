package api

import (
	"fmt"
	"log"
	"net/http"
)

// func searchHandler(students []sqlite.StudentName) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		templatePage, err := template.ParseFiles("html/index.html")
// 		if err != nil {
// 			fmt.Println(err)
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		if err := templatePage.Execute(w, students); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 	}
// }

// func menuFormHandler(students []sqlite.StudentName) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		templatePage, err := template.ParseFiles("html/index.html")
// 		if err != nil {
// 			fmt.Println(err)
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		if err := templatePage.Execute(w, students); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 	}
// }

// func getStudentInfo(w http.ResponseWriter, req *http.Request) {
// 	req.ParseForm()
// 	id := req.Form["studentid"][0]
// 	// for key, value := range req.Form {
// 	// 	fmt.Printf("%s = %s\n", key, value)
// 	// }
// 	student, _ := sqlite.GetStudentById(DBConn, id)
// 	fmt.Println(student)

// 	// write output
// 	w.Header().Add("Content-Type", "application/json")
// 	log.Println("Endpoint Hit: students")
// 	json.NewEncoder(w).Encode(student)
// }

func parseForm(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	for key, value := range req.Form {
		fmt.Printf("%s = %s\n", key, value)
	}
}

func SetupServer(address string, port int) {

	// setup routes

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "html/index.html")
	// })
	// http.HandleFunc("/confirm", uploadFiles)
	// http.HandleFunc("/", searchHandler(studs))
	// http.HandleFunc("/studentinfo", getStudentInfo)
	// http.HandleFunc("/static", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "html/static.html")
	// })
	// http.HandleFunc("/parseform", parseForm)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/static.html")
	})
	http.HandleFunc("/parseform", parseForm)

	// start server
	log.Printf("Serving on %[1]s port %[2]d (http://%[1]s:%[2]d/)\n", address, port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), nil)
}

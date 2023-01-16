package main

import (
	"fmt"
	"log"

	// To import a package solely for its side-effects (initialization), use the
	// blank identifier as explicit package name
	// _ "github.com/mattn/go-sqlite3"
	"stellaeestudio/cmd/sqlite"
)

func main() {
	log.Println("Connect to db")
	conn, err := sqlite.Open("./stellae.db")
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	profs, err := sqlite.GetProfessors(conn)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("professors results:")
	fmt.Printf("%+v\n", profs)

	classes, err := sqlite.GetClasses(conn)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Classes results:")
	fmt.Printf("%+v\n", classes)

	studs, err := sqlite.GetStudents(conn)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Students results:")
	fmt.Printf("%+v\n", studs)

	ocpu, err := sqlite.GetClassOcupancy(conn)

	if err != nil {
		log.Panicln(err)
	}

	log.Println("ClassOcupancy results:")
	fmt.Printf("%+v\n", ocpu)
}

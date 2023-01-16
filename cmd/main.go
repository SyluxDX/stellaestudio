package main

import (
	"flag"
	"fmt"
	"log"

	"stellaeestudio/cmd/sqlite"
)

func main() {
	// command line flags
	var databaseFilepath string
	flag.StringVar(&databaseFilepath, "db", "./stellae.db", "Path to sqlite3 database.")
	flag.Parse()

	log.Printf("Connect to db %s\n", databaseFilepath)
	conn, err := sqlite.Open(databaseFilepath)
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

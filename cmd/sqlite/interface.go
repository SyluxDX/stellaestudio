package sqlite

import (
	"database/sql"
	"fmt"

	// To import a package solely for its side-effects (initialization), use the
	// blank identifier as explicit package name
	_ "github.com/mattn/go-sqlite3"
)

type Professor struct {
	Id     int
	Name   string
	Active bool
}

type Class struct {
	Id           int
	Name         string
	Weekday      int
	Time         string
	Professor_id int
	Active       bool
}

type Student struct {
	Id          int
	Name        string
	PhoneNumber sql.NullString
	Email       sql.NullString
	Nif         int
	Active      bool
}

type StudentName struct {
	Id   int
	Name string
}

type ClassOcupancy struct {
	ClassId   int
	StudentId int
}

type Subscription struct {
	Id    int
	Name  string
	Value int
}

type Promotion struct {
	Id    int
	Name  string
	Value int
}

// Open opens the sqlite database, requires a string with
// the path to the sqlite db file and retuns:
//
//	connection object: *sql.DB
//	error : error
func Open(dbPath string) (*sql.DB, error) {
	dbConn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

func GetProfessors(dbConn *sql.DB) ([]Professor, error) {
	query := "SELECT id, name, active FROM professors;"
	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []Professor
	for rows.Next() {
		var row Professor
		err = rows.Scan(&row.Id, &row.Name, &row.Active)
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}

func GetClasses(dbConn *sql.DB) ([]Class, error) {
	query := "SELECT id, name, weekday, time, professor_id, active FROM classes;"
	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []Class
	for rows.Next() {
		var row Class
		err = rows.Scan(
			&row.Id,
			&row.Name,
			&row.Weekday,
			&row.Time,
			&row.Professor_id,
			&row.Active,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}

func GetStudents(dbConn *sql.DB) ([]Student, error) {
	query := "SELECT id, name, phone_number, email, nif, active FROM students;"
	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []Student
	for rows.Next() {
		var row Student
		err = rows.Scan(
			&row.Id,
			&row.Name,
			&row.PhoneNumber,
			&row.Email,
			&row.Nif,
			&row.Active,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}

func GetStudentById(dbConn *sql.DB, id string) (*Student, error) {
	query := fmt.Sprintf("SELECT id, name, phone_number, email, nif, active FROM students where id=%s;", id)
	rows, err := dbConn.Query(query)
	if err != nil {
		fmt.Println("Error connection query")
		return nil, err
	}
	defer rows.Close()

	var data Student
	for rows.Next() {
		var row Student
		err = rows.Scan(
			&row.Id,
			&row.Name,
			&row.PhoneNumber,
			&row.Email,
			&row.Nif,
			&row.Active,
		)
		if err != nil {
			return nil, err
		}
		data = row
	}
	return &data, nil
}

func GetStudentsNames(dbConn *sql.DB) ([]StudentName, error) {
	query := "SELECT id, name FROM students;"
	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []StudentName
	for rows.Next() {
		var row StudentName
		err = rows.Scan(
			&row.Id,
			&row.Name,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}

func GetClassOcupancy(dbConn *sql.DB) ([]ClassOcupancy, error) {
	query := "SELECT class_id, student_id FROM class_ocupancy;"
	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []ClassOcupancy
	for rows.Next() {
		var row ClassOcupancy
		err = rows.Scan(
			&row.ClassId,
			&row.StudentId,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}

func GetSubscriptions(dbConn *sql.DB) ([]Subscription, error) {
	query := "SELECT id, name, value FROM subscription;"
	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []Subscription
	for rows.Next() {
		var row Subscription
		err = rows.Scan(
			&row.Id,
			&row.Name,
			&row.Value,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}
func GetPromos(dbConn *sql.DB) ([]Promotion, error) {
	query := "SELECT id, name, value FROM promo;"
	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []Promotion
	for rows.Next() {
		var row Promotion
		err = rows.Scan(
			&row.Id,
			&row.Name,
			&row.Value,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}

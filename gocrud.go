package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connect()
	var option int

	fmt.Println("Insert the row into the database")
	fmt.Println("Update the row into the database")
	fmt.Println("Delete the row from the database")
	fmt.Println("View All the records from the database")
	fmt.Println("Enter the choice")
	fmt.Scanln(&option)
	switch option {
	case 1:
		insert()
		break
	case 2:
		update()
		break

	case 3:
		delete()
		break

	case 4:
		findAll()
		break

	default:
		fmt.Println("Invalid Choice")
	}

}

func connect() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root123"
	dbName := "newdb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// error Handling for database conecting issues
	if err != nil {
		log.Println("Database Not connected ?")
		panic(err.Error())
	}
	log.Println("Database connect Success")
	return db
}

func insert() {

	db := connect()
	var id int
	var name string
	fmt.Println("Enter the Id")
	fmt.Scanln(&id)
	fmt.Println("Enter the name")
	fmt.Scanln(&name)
	insrt, err := db.Query("INSERT INTO employee(Id, Name) VALUES (?,?)", id, name)
	if err != nil {

		panic(err.Error())
	}
	fmt.Println("Successfully Inserted")
	defer insrt.Close()
	defer db.Close()

}

func update() {

	db := connect()
	var name string
	var id int
	fmt.Println("Enter the name")
	fmt.Scanln(&name)
	fmt.Println("Enter the Id by which u can modify the row")
	fmt.Scanln(&id)

	updt, err := db.Query("UPDATE employee SET name=? WHERE id=?", name, id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Row is successfully updated")

	defer updt.Close()
	defer db.Close()
}

func delete() {

	db := connect()
	var id int
	fmt.Println("Enter the id which u want to delete")
	fmt.Scanln(&id)

	del, err := db.Query("DELETE FROM employee WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Row is successfully deleted")
	defer del.Close()
	defer db.Close()
}

func findAll() {

	db := connect()
	var id int
	var name string
	find, err := db.Query("Select * from employee")
	if err != nil {
		panic(err.Error())
	}
	for find.Next() {
		find.Scan(&id, &name)
		fmt.Printf("Id: %v : Name:%v ", id, name)
	}
	defer find.Close()
	defer db.Close()

}

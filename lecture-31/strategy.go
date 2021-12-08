package main

import (
	"fmt"
)

type IDBconnection interface {
	Connect()
}
type SqlConnection struct {
	connectionString string
}

func (con SqlConnection) Connect() {
	fmt.Println(("Sql " + con.connectionString))
}

type OracleConnection struct {
	connectionString string
}

func (con OracleConnection) Connect() {
	fmt.Println("Oracle " + con.connectionString)
}

type DBConnection struct {
	someContextValue int
	db               IDBconnection
}

func (con DBConnection) DBConnect() {
	con.db.Connect()
}

func main() {
	sqlConnection := SqlConnection{connectionString: "Connection is connected"}
	con := DBConnection{db: sqlConnection, someContextValue: 1}
	con.DBConnect()

	orcConnection := OracleConnection{connectionString: "Connection is connected"}
	con2 := DBConnection{db: orcConnection, someContextValue: 2}
	con2.DBConnect()
}

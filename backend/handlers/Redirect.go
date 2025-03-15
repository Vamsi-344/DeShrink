package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"database/sql"

	_ "github.com/lib/pq"
)

type ReqBody struct {
	After string `json:"after"`
}

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "example"
// 	dbname   = "postgres"
// )

func Redirect(c *gin.Context) {

	afterurl, _ := c.Params.Get("after")
	var ReqBody ReqBody

	if c.BindJSON(&ReqBody) != nil {
		fmt.Println("error")
	}
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	//retrieve
	// dynamic
	// fmt.Println(ReqBody.After)
	retrieveDynStmt := `select * from urls where after = $1`
	rows := db.QueryRow(retrieveDynStmt, afterurl)
	// CheckError(e)
	// rows, err := db.Query(`SELECT "Name", "Roll_Number" FROM "Students"`)
	// CheckError(err)
	// fmt.Println(rows)
	// defer rows.Close()
	var before string
	var after string
	err = rows.Scan(&before, &after)
	CheckError(err)

	fmt.Println(before, after)

	// defer rows.Close()
	// for rows.Next() {
	// 	fmt.Println(rows)
	// 	var before string
	// 	var after string
	// 	// var name string
	// 	// var roll_number int

	// 	err = rows.Scan(&before, &after)
	// 	CheckError(err)

	// 	fmt.Println(before, after)
	// }

	c.Redirect(http.StatusMovedPermanently, before)
	CheckError(err)
	c.JSON(http.StatusOK, "not yet ready")
}

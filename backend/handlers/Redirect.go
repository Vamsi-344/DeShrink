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
	var before string
	var after string
	retrieveDynStmt := `select * from urls where after = $1`
	err = db.QueryRow(retrieveDynStmt, afterurl).Scan(&before, &after)
	CheckError(err)

	fmt.Println(before, after)

	c.Redirect(http.StatusMovedPermanently, before)
	CheckError(err)
	c.JSON(http.StatusOK, "not yet ready")
}

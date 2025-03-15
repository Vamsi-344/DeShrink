package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

type ShortenReqBody struct {
	Url string `json:"url"`
}

// GenerateRandomString generates a random string of a given length
func GenerateRandomString(length int) (string, error) {
	// Create a byte slice to store the random string
	bytes := make([]byte, length)

	// Generate random bytes
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Convert random bytes to a hex string and return it
	return hex.EncodeToString(bytes)[:length], nil
}

func ShortURLGenerator(c *gin.Context) {

	var ReqBody ShortenReqBody
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

	//insert
	// dynamic
	insertDynStmt := `insert into urls (before, after) values($1, $2)`
	shortPath, err := GenerateRandomString(6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated short path:", shortPath)
	_, e := db.Exec(insertDynStmt, ReqBody.Url, shortPath)
	CheckError(e)
	// fmt.Println("url")
	c.JSON(http.StatusOK, "Record inserted successfully")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

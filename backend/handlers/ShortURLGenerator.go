package handlers

import (
	"crypto/rand"
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

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// GenerateRandomString generates a random string of a given length
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		randomByte := make([]byte, 1)
		_, err := rand.Read(randomByte)
		if err != nil {
			panic(err)
		}
		b[i] = charset[randomByte[0]%byte(len(charset))]
	}

	return string(b), nil
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

	var shortPath string
	exists := true

	// Retry generating shortPath until it is unique
	for exists {
		shortPath, err = GenerateRandomString(6)
		if err != nil {
			log.Fatal(err)
		}

		// Check if the generated shortPath already exists in DB
		var count int
		checkQuery := `SELECT COUNT(*) FROM urls WHERE after = $1`
		err = db.QueryRow(checkQuery, shortPath).Scan(&count)
		if err != nil {
			log.Fatal(err) // Handle DB error properly
		}

		exists = (count > 0) // If count > 0, shortPath already exists, so retry
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated short path:", shortPath)

	insertDynStmt := `insert into urls (before, after) values($1, $2)`

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

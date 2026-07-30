package main

import (
	"database/sql"
	"go-sqlc/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db    *sql.DB
	query *db.Queries
}

func main() {
	// initializing the db
	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	query := db.New(conn)

	// initializing the server
	s := Server{
		db:    conn,
		query: query,
	}

	router := gin.Default()
	router.GET("/get-users", func(c *gin.Context) {
		users, err := s.query.Listusers(c)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
			"data":    users,
		})
	})

	router.GET("/get-user/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}
		users, err := s.query.Detailusers(c, int32(id))
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
			"data":    users,
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}

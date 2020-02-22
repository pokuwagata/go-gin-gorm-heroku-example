package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Ping struct {
	ID       string
	Occurred time.Time
}

const tableName = "ping_timestamp"

func registerPing(db *gorm.DB) {
	// _, err := db.Exec("INSERT INTO ping_timestamp (occurred) VALUES ($1)", time.Now())
	db.Table(tableName).Create(&Ping{Occurred: time.Now()})
}

func pingFunc(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		defer registerPing(db)
		// r := db.QueryRow("SELECT occurred FROM ping_timestamp ORDER BY id DESC LIMIT 1")
		var ping Ping
		var pings []Ping
		db.Table(tableName).Find(&pings)
		db.Table(tableName).Last(&ping)
		// var ping Ping
		c.JSON(200, gin.H{
			// "message": ping.Occurred,
			"all":  pings,
			"last": ping,
		})
	}
}

func main() {

	r := gin.Default()
	api := r.Group("/api")
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("failed to connect database")
	}
	api.GET("/ping", pingFunc(db))

	r.Run()
}

package main

import (
	"final-project-golang-fga-hacktiv8/config"
	"final-project-golang-fga-hacktiv8/router"
)

func main() {
	config.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
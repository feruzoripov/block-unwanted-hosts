// internal/app/app.go
package app

import (
	"log"

	"github.com/feruzoripov/block-unwanted-hosts/internal/app/blocking"
)

// Run initializes and runs your application.
func Run() {
	log.Println("Fetching and blocking websites...")

	err := blocking.FetchAndBlockWebsites()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	log.Println("Done!")
}

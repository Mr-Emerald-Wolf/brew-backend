package migrations

import (
	"fmt"
	"log"

	"github.com/mr-emerald-wolf/brew-backend/internal/database"
	"github.com/mr-emerald-wolf/brew-backend/internal/domain"
)

func RunMigrations() {
	log.Println("Running Migrations")

	db := database.DB

	err := db.AutoMigrate(&domain.User{})

	if err != nil {
		fmt.Println("Could not migrate Account")
		return
	}

	log.Println("🚀 Migrations completed")
}

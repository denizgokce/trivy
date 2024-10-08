package db

import (
	"database/sql"
	"log"
	"user-service/internal/models"

	"github.com/bxcodec/faker/v3"
)

func SeedUsers(db *sql.DB) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		log.Fatalf("Failed to count users: %v", err)
	}

	if count == 0 {
		log.Println("Seeding users...")

		// Generate random users using faker
		var users []models.User
		for i := 0; i < 10; i++ { // Change the number to generate more or fewer users
			user := models.User{
				Name:     faker.Name(),
				Email:    faker.Email(),
				Password: faker.Password(),
			}
			users = append(users, user)
		}

		for _, user := range users {
			_, err := db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
			if err != nil {
				log.Fatalf("Failed to seed user: %v", err)
			}
		}

		log.Println("Users seeded successfully.")
	} else {
		log.Println("Users already exist, skipping seeding.")
	}
}

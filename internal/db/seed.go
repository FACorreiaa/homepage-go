package db

import (
	"context"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdminUser(ctx context.Context, q *Queries) error {
	email := getEnvOrDefault("ADMIN_EMAIL", "fernandocorreia316@gmail.com")
	password := getEnvOrDefault("ADMIN_PASSWORD", "StudioAdmin123!")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return q.CreateUser(ctx, CreateUserParams{
		ID:            uuid.NewString(),
		Name:          "Fernando Correia",
		Email:         email,
		PasswordHash:  string(hash),
		ProjectStatus: "System Administrator",
	})
}

func getEnvOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

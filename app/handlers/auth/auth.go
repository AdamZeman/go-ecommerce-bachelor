package auth

import (
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"log"
	"os"
)

const (
	key    = "hahaha"
	MaxAge = 86400 * 7
	IsProd = false
)

func NewAuth() *sessions.CookieStore {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	//defaultUrl := os.Getenv("DEFAULT_URL")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, "http://localhost:4000/auth/google/callback", "email", "profile"),
		//google.New(googleClientId, googleClientSecret, "http://localhost:4000/auth/google/callback"),
	)

	return store

}

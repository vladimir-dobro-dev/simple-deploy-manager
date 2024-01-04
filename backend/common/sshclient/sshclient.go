package sshclient

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/melbahja/goph"
)

func New() *goph.Client {
	godotenv.Load()
	sshUser := os.Getenv("SSH_USER")
	sshPassword := os.Getenv("SSH_PASSWORD")
	client, err := goph.New(sshUser, "127.0.0.1", goph.Password(sshPassword))
	if err != nil {
		log.Fatal(err)
	}

	return client
	// defer client.Close()

	// out, err := client.Run("touch file.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(string(out))
}

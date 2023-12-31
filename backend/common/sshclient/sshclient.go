package sshclient

import (
	"log"

	"github.com/melbahja/goph"
)

func Connect() *goph.Client {
	client, err := goph.New("devs", "127.0.0.1", goph.Password("1234"))
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

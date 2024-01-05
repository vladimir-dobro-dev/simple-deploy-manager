package sshclient

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/melbahja/goph"
)

var sshClient *goph.Client

type ConnectData struct {
	Address  string
	User     string
	Password string
	Port     string
}

func New(data ConnectData) (c *goph.Client, err error) {
	if sshClient != nil {
		return sshClient, nil
	}

	var address, user, password, port string
	e := godotenv.Load()
	if e == nil {
		address = os.Getenv("SSH_ADDRESS")
		user = os.Getenv("SSH_USER")
		password = os.Getenv("SSH_PASSWORD")
		port = os.Getenv("SSH_PORT")

	} else {
		address = data.Address
		user = data.User
		password = data.Password
		port = data.Port
	}

	c, err = goph.New(user, address+":"+port, goph.Password(password))

	return
}

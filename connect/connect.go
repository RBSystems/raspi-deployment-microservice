package connect

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/byuoitav/raspi-deployment-microservice/passwords"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
)

//runs a command on a remote host
//abstracts building a session running the command
func RunCommand(hostname, command string) error {

	session, err := connect(hostname)
	if err != nil {
		return err
	}

	log.Printf("[ssh] running the following command on %s: %s", hostname, command)
	return session.Run(command)
}

func connect(hostname string) (*ssh.Session, error) {

	log.Printf("%s", color.HiGreenString("[ssh] connecting to %s", hostname))

	//get password
	password, err := passwords.GetPassword(hostname)
	if err != nil {
		msg := fmt.Sprintf("password for %s not found: %s", hostname, err.Error())
		log.Printf("%s", color.HiRedString("[connect] %s", msg))
		return &ssh.Session{}, errors.New(msg)
	}

	//build configuration
	var sshConfig = &ssh.ClientConfig{
		User: os.Getenv("PI_SSH_USERNAME"),
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}

	//build connection
	connection, err := ssh.Dial("tcp", hostname+":22", sshConfig)
	if err != nil {
		msg := fmt.Sprintf("error dialing %s: %s", hostname, err.Error())
		log.Printf("%s", color.HiRedString("[ssh] %s", msg))
		return &ssh.Session{}, errors.New(msg)
	}

	log.Printf("[ssh] TCP connection established to %s", hostname)
	defer connection.Close()

	//start session
	session, err := connection.NewSession()
	if err != nil {
		msg := fmt.Sprintf("error starting session on %s: %s", hostname, err.Error())
		log.Printf("%s", color.HiRedString("[ssh] %s", msg))
		return &ssh.Session{}, errors.New(msg)
	}

	log.Printf("[ssh] session established with %s", hostname)
	return session, nil
}

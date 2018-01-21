package helpers

import (
	"log"

	"github.com/fatih/color"

	"golang.org/x/crypto/ssh"
)

//@param active - true indicates monitoring the contact points, false indicates not monitoring the contact points
func UpdateContactState(hostname string, active bool) error {

	state := "disabling"
	if active {
		state = "enabling"
	}

	log.Printf("%s", color.HiGreenString("[helpers] %s contact point monitoring on %s...", state, hostname))

	if active {
		return ssh.RunCommand("sudo systemctl enable contacts && sudo systemctl start contacts")
	}

	return ssh.RunCommand("sudo systemctl stop contacts && sudo systemctl disable contacts")
}

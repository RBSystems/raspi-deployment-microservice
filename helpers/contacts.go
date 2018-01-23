package helpers

import (
	"log"

	"github.com/byuoitav/raspi-deployment-microservice/connect"
	"github.com/fatih/color"
)

const ENABLE = "sudo systemctl enable contacts && sudo systemctl start contacts"
const DISABLE = "sudo systemctl stop contacts && sudo systemctl disable contacts"

//@param active - true indicates monitoring the contact points, false indicates not monitoring the contact points
func UpdateContactState(hostname string, active bool) error {

	state := "disabling"
	if active {
		state = "enabling"
	}

	log.Printf("%s", color.HiGreenString("[helpers] %s contact point monitoring on %s...", state, hostname))

	if active {
		return connect.RunCommand(hostname, ENABLE)
	}

	return connect.RunCommand(hostname, DISABLE)
}

#!/bin/bash 

echo "Second boot."

	wait 45
	sudo chvt 2

	until $(sudo usermod -aG docker pi); do
		curl -sSL https://get.docker.com -k | sh
		wait
	done
	echo "Added user pi to the docker group"

# get environment variables
until $PI_HOSTNAME; do 
	curl http://sandbag.byu.edu:2000/deploy/$(hostname)
	source /etc/environment
done

printf "\nrecieved env. variables\n"

# maria db setup
until $(curl https://raw.githubusercontent.com/byuoitav/raspi-deployment-microservice/master/mariadb-setup.sh > /tmp/mariadb-setup.sh); do
	echo "Trying again."
done
chmod +x /tmp/mariadb-setup.sh

# salt setup
until $(curl https://raw.githubusercontent.com/byuoitav/raspi-deployment-microservice/master/salt-setup.sh > /tmp/salt-setup.sh); do
	echo "Trying again."
done
chmod +x /tmp/salt-setup.sh

until [ -d "/etc/salt/" ]; do
	/tmp/salt-setup.sh
done

# docker 
until [ $(docker ps -q | wc -l) -gt 1 ]; do
	echo "Waiting for docker containers to download"
	sleep 10
done

echo "Removing symlink to startup script."
sudo rm /usr/lib/systemd/system/default.target.wants/first-boot.service

clear
printf "\n\tSetup complete. Please wait for me to reboot...\n"
sleep 30

sudo sh -c "reboot"


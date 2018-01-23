package passwords

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/byuoitav/authmiddleware/bearertoken"
	"github.com/byuoitav/pi-credentials-microservice/structs"
	"github.com/fatih/color"
)

const PASSWORD_LENGTH = 512

//returns the password of the given host
//for now we rely on the bearer-token-microservice to ensure security
//TODO add ADFS authentication
func GetPassword(hostname string) (string, error) {

	//build client
	var client http.Client

	//build request
	url := fmt.Sprintf("%s/devices/%s", os.Getenv("RASPI_CRED_MICROSERVICE_ADDRESS"), hostname)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		msg := fmt.Sprintf("unable to build request: %s", err.Error())
		log.Printf("%s", color.HiRedString("[passwords] %s", msg))
		return "", errors.New(msg)
	}

	//set bearer token
	err = SetToken(req)
	if err != nil {
		return "", err
	}

	//DO IT
	resp, err := client.Do(req)
	if err != nil {
		msg := fmt.Sprintf("unable to complete request: %s", err.Error())
		log.Printf("%s", color.HiRedString("[passwords] %s", msg))
		return "", errors.New(msg)
	}

	//read response body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		msg := fmt.Sprintf("unable to read response body: %s", err.Error())
		log.Printf("%s", color.HiRedString("[passwords] %s", msg))
		return "", errors.New(msg)
	}

	//get the dang value
	var entry structs.Entry
	err = json.Unmarshal(body, &entry)
	if err != nil {
		msg := fmt.Sprintf("unable to read response body: %s", err.Error())
		log.Printf("%s", color.HiRedString("[passwords] %s", msg))
		return "", errors.New(msg)
	}

	return entry.Password, nil
}

func GenerateRandomPassword() (string, error) {
	bytes := make([]byte, PASSWORD_LENGTH)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

//sets the bearer token
func SetToken(request *http.Request) error {

	if len(os.Getenv("LOCAL_ENVIRONMENT")) == 0 {

		log.Printf("[helpers] setting bearer token...")

		token, err := bearertoken.GetToken()
		if err != nil {
			msg := fmt.Sprintf("cannot get bearer token: %s", err.Error())
			log.Printf("%s", color.HiRedString("[passwords] %s", msg))
			return errors.New(msg)
		}

		request.Header.Set("Authorization", "Bearer "+token.Token)
	}

	return nil
}

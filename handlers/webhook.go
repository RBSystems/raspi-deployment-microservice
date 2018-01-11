package handlers

import (
	"net/http"

	"github.com/byuoitav/raspi-deployment-microservice/helpers"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func WebhookDeployment(context echo.Context) error {

	deviceClass := context.Param("class")
	deploymentType := context.Param("designation")

	response, err := helpers.ScheduleDeployment(deviceClass, deploymentType)
	if err != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, err.Error())
		return nil
	}

	jsonresp.New(context.Response(), http.StatusOK, response)
	return nil
}

func WebhookDeviceFirstTime(context echo.Context) error {

	response, err := helpers.DeployDevice(context.Param("hostname"), true)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func WebhookDevice(context echo.Context) error {
	response, err := helpers.DeployDevice(context.Param("hostname"), false)
	if err != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, err.Error())
		return nil
	}

	jsonresp.New(context.Response(), http.StatusOK, response)
	return nil
}

func EnableContacts(context echo.Context) error {

	err := helpers.UpdateContactState(context.Param("hostname"), true)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"Response": "Failed to set state"})
	}

	return context.JSON(http.StatusOK, map[string]string{"Response": "Success"})
}

func DisableContacts(context echo.Context) error {

	err := helpers.UpdateContactState(context.Param("hostname"), false)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"Response": "Failed to set state"})
	}

	return context.JSON(http.StatusOK, map[string]string{"Response": "Success"})
}

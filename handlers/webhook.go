package handlers

import (
	"fmt"
	"net/http"

	"github.com/byuoitav/raspi-deployment-microservice/helpers"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func WebhookDeployment(context echo.Context) error {

	class := context.Param("class")
	designation := context.Param("designation")

	err := helpers.Deploy(class, designation)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, fmt.Sprintf("deployment to %s %s devices started", designation, class))
}

func DisableDeploymentsByBranch(context echo.Context) error {
	branch := context.Param("branch")
	helpers.HoldDeployment(branch, true)
	return context.String(http.StatusOK, fmt.Sprintf("Disabled %s deployments", branch))
}

func EnableDeploymentsByBranch(context echo.Context) error {
	branch := context.Param("branch")
	helpers.HoldDeployment(branch, false)
	return context.String(http.StatusOK, fmt.Sprintf("Enabled %s deployments", branch))
}

func WebhookDevice(context echo.Context) error {
	response, err := helpers.DeployDevice(context.Param("hostname"))
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

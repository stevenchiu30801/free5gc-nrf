/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

import (
	"net/http"

	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi"
	"free5gc/lib/openapi/models"
	"free5gc/src/nrf/logger"
	"free5gc/src/nrf/producer"
	"github.com/gin-gonic/gin"
)

// GetNFInstances - Retrieves a collection of NF Instances
func HTTPGetNFInstances(c *gin.Context) {

	req := http_wrapper.NewRequest(c.Request, nil)
	req.Query = c.Request.URL.Query()

	httpResponse := producer.HandleGetNFInstancesRequest(req)

	responseBody, err := openapi.Serialize(httpResponse.Body, "application/json")
	if err != nil {
		logger.ManagementLog.Warnln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(httpResponse.Status, "application/json", responseBody)
	}
}

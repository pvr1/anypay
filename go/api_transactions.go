/*
 * AnyPay
 *
 * This the AnyPay service targeted towards, parents with children doing payments and russian oligarks
 *
 * API version: 1.0.0
 * Contact: per.von.rosen@swedbank.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTransaction - Get a details of a specific transaction
func GetTransaction(c *gin.Context) {
	c.String(http.StatusOK, "There you got your specific transaction\n")
}

// GetTransactions - Get a collection of transactions
func GetTransactions(c *gin.Context) {
	c.String(http.StatusOK, "Get list of Transactions originating from FX Orders and Payments\n")
}

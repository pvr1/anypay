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

// AddPayment - Add a new payment
func AddPayment(c *gin.Context) {
	c.JSON(http.StatusOK, "Payment order added\n")
}

// DeletePayment - Deletes a payment-instruction not settled yet
func DeletePayment(c *gin.Context) {
	c.JSON(http.StatusOK, "Payment order deleted\n")
}

// GetPayment - Get a specific payment-instruction with details
func GetPayment(c *gin.Context) {
	c.JSON(http.StatusOK, "Get specific Payment order\n")
}

// GetPayments - Get all non-settled payment-instructions IDs that the login has privileges to get
func GetPayments(c *gin.Context) {
	c.JSON(http.StatusOK, "Get list of Payment orders\n")
}

// UpdatePayment - Update an existing payment-instruction not settled yet
func UpdatePayment(c *gin.Context) {
	c.JSON(http.StatusOK, "Payment order updated\n")
}
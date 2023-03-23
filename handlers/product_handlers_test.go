package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddProduct(t *testing.T) {
	// Set up the request payload
	productPayload := `{"name": "Test Product", "price": 1000, "description": "Test product description", "quantity": 10}`
	req, err := http.NewRequest("POST", "/api/products", bytes.NewBufferString(productPayload))
	require.NoError(t, err)

	// Create the response recorder
	rr := httptest.NewRecorder()

	// Call the AddProduct handler
	AddProduct(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Parse the response JSON
	var resp map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	require.NoError(t, err)

	// Check the response data
	assert.Equal(t, "Test Product", resp["name"])
	assert.Equal(t, float64(1000), resp["price"])
	assert.Equal(t, "Test product description", resp["description"])
	assert.Equal(t, float64(10), resp["quantity"])
}

func TestListProducts(t *testing.T) {
	// Create a sample product
	productPayload := `{"name": "Sample Product", "price": 90000, "description": "Sample product description", "quantity": 5}`
	req, err := http.NewRequest("POST", "/api/products", bytes.NewBufferString(productPayload))
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	AddProduct(rr, req)

	// Set up the request
	req, err = http.NewRequest("GET", "/api/products", nil)
	require.NoError(t, err)

	// Create the response recorder
	rr = httptest.NewRecorder()

	// Call the ListProducts handler
	ListProducts(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Parse the response JSON
	var resp []map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	require.NoError(t, err)

	// Check the response data
	assert.NotEmpty(t, resp)
}

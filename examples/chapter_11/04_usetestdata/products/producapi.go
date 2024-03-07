// Package products provides a client for the external service API
package products

import "net/http"

// Client is a client for the external service API.
// It uses the base URL and the HTTP client to make requests
type Client struct {
	baseURL string
	client  *http.Client
}

// New constructor: creates a new client for the external service API.
//
// - baseURL is the base URL of the external service
//
// - client is the HTTP client to use
func New(baseURL string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{baseURL: baseURL, client: client}
}

// Product is a product model.
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// RangeInList is a range of prices in the list.
type RangeInList struct {
	From float64 `json:"from"`
	To   float64 `json:"to"`
}

// ProductList is a list of products with total count and range of prices.
type ProductList struct {
	Products []Product   `json:"items"`
	Total    int         `json:"total"`
	Range    RangeInList `json:"range"`
}

// ClientAPI is an interface for the external service API.
type ClientAPI interface {
	GetProduct(int) (*Product, error)
	GetProductList(from, to float64) (ProductList, error)
}

// GetProduct returns a product by its ID.
func (c *Client) GetProduct(id int) (*Product, error) {
	// ... implementation (HTTP request to the external service)
	return &Product{ID: id}, nil
}

// GetProductList returns a list of products with prices in the range from-to.
func (c *Client) GetProductList(from, to float64) (ProductList, error) {
	// ... implementation (HTTP request to the external service)
	return ProductList{
		Products: []Product{},
		Total:    0,
		Range:    RangeInList{From: from, To: to},
	}, nil
}

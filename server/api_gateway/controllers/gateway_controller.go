package controllers

import (
	"api_gateway/configs"
	"bytes"
	"fmt"
	"log"
	"net/http"

	"api_gateway/models"
	"encoding/json"

	// "io/ioutil"
	"io"

	"github.com/gofiber/fiber/v2"
)

// HealthCheck
// func HealthCheck(c *fiber.Ctx) error {
// 	return c.SendString("API Gateway is up and running!")
// }

func formatRespBody(resp *http.Response, c *fiber.Ctx) ([]byte, error) {

	// body, err := ioutil.ReadAll(resp.Body)  // depricated
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, c.Status(http.StatusInternalServerError).SendString("Error reading response from service")
	}

	// Pretty print the response JSON
	var formattedBody map[string]interface{}
	if err := json.Unmarshal(body, &formattedBody); err != nil {
		log.Printf("Error unmarshaling response body: %v", err)
		return nil, c.Status(http.StatusInternalServerError).SendString("Error unmarshaling response from service")
	}

	// Marshal the formatted body with indentation
	prettyBody, err := json.MarshalIndent(formattedBody, "", "    ")
	if err != nil {
		log.Printf("Error marshalling formatted body: %v", err)
		return nil, c.Status(http.StatusInternalServerError).SendString("Error formatting response")
	}

	return prettyBody, nil
}

// ProxyRequest handles the proxying of requests to microservices
func ProxyRequest(c *fiber.Ctx) error {

	var resp *http.Response
	var err error

	// Get the service name from the route parameter
	service := c.Params("service")

	// fmt.Println(c.AllParams())

	// Determine which service to forward the request to based on the service parameter
	switch service {
	case "file":
		fileHandler()
	case "game":
		resp, err = gameHandler(c)
	case "sales":
		saleHandler()
	case "user":
		userHandler()
	default:
		return c.Status(http.StatusNotFound).SendString("Service not found")
	}

	prettyBody, _ := formatRespBody(resp, c)

	// can send the non formatted body and it will work well. just not formatted in postman
	// body, _ := io.ReadAll(resp.Body)

	// Defer closing the response body only after verifying `resp` is not nil
	defer resp.Body.Close()

	if err == nil {
		// Return the formatted JSON response back to the client
		return c.Status(resp.StatusCode).Send(prettyBody)
	} else {
		return c.Status(http.StatusBadRequest).SendString("Bad Request")
	}
}

func gameHandler(c *fiber.Ctx) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", configs.GAME_SERVICE_URL, c.Params("*"))

	var resp *http.Response
	var err error

	var game models.Game

	switch c.Method() {
	case "GET":
		resp, err = http.Get(url)
		if err != nil {
			return nil, c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error forwarding GET request to %s", "game"))
		}
	case "POST":
		if err := c.BodyParser(&game); err != nil {
			return nil, c.Status(http.StatusBadRequest).SendString("Invalid request body")
		}

		jsonData, err := json.Marshal(game)
		if err != nil {
			log.Printf("Error serializing request body: %v", err)
			return nil, c.Status(http.StatusInternalServerError).SendString("Error serializing request body")
		}

		resp, err = http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error forwarding POST request to %s", "game"))
		}
	case "PUT":
		rawBody := c.Body()
		httpClient := &http.Client{}
		putReq, _ := http.NewRequest("PUT", url, bytes.NewBuffer(rawBody))
		putReq.Header.Set("Content-Type", "application/json")
		resp, err = httpClient.Do(putReq)
		if err != nil {
			return nil, c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error forwarding PUT request to %s", "game"))
		}
	case "DELETE":
		httpClient := &http.Client{}
		deleteReq, _ := http.NewRequest("DELETE", url, nil)
		resp, err = httpClient.Do(deleteReq)
		if err != nil {
			return nil, c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error forwarding DELETE request to %s", "game"))
		}
	default:
		return nil, c.Status(http.StatusMethodNotAllowed).SendString("Method not allowed")
	}

	return resp, nil
}

func fileHandler() {

}

func userHandler() {

}

func saleHandler() {

}

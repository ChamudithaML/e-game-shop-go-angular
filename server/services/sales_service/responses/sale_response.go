package responses

// import (
// 	"sales_service/models"
// )

type SalesResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Titles  []string `json:"titles"`
	Stocks  []int    `json:"stocks"`
	Sales   []int    `json:"sales"`
}

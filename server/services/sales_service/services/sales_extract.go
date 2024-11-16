package services

import (
	"sales_service/models"
)

func ExtractNames(sales []models.Sales) []string {
	titles := []string{}

	for _, item := range sales {
		titles = append(titles, item.Title)
	}

	return titles
}

func ExtractStock(sales []models.Sales) []int {
	stocks := []int{}

	for _, item := range sales {
		stocks = append(stocks, item.Stock)
	}

	return stocks
}

func ExtractSales(sales []models.Sales) []int {
	sales_amount := []int{}

	for _, item := range sales {
		sales_amount = append(sales_amount, item.Stock)
	}

	return sales_amount
}

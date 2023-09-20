package main

import (
	"fmt"
)

type Size struct {
	ID    int
	Name  string
	Price int
}

type Product struct {
	ID    int
	Name  string
	Sizes []Size
}

type CardProducts struct {
	ProductID int
	SizeID    int
	Quantity  int
}

type Card struct {
	Products []CardProducts
}

type Client struct {
	ID   int
	Name string
	Card Card
}

func main() {
	clients := []Client{
		{
			ID:   1,
			Name: "Shohruh",
			Card: Card{
				Products: []CardProducts{
					{
						ProductID: 1,
						SizeID:    1,
						Quantity:  9,
					},
					{
						ProductID: 2,
						SizeID:    2,
						Quantity:  3,
					},
				},
			},
		},
		{
			ID:   2,
			Name: "Islom",
			Card: Card{
				Products: []CardProducts{
					{
						ProductID: 1,
						SizeID:    1,
						Quantity:  1,
					},
					{
						ProductID: 2,
						SizeID:    2,
						Quantity:  4,
					},
				},
			},
		},
	}

	products := []Product{
		{
			ID:   1,
			Name: "Cola",
			Sizes: []Size{
				{
					ID:    1,
					Name:  "25sm",
					Price: 5000,
				},
				{
					ID:    2,
					Name:  "30sm",
					Price: 6000,
				},
			},
		},
		{
			ID:   2,
			Name: "Fanta",
			Sizes: []Size{
				{
					ID:    1,
					Name:  "25sm",
					Price: 4000,
				},
				{
					ID:    2,
					Name:  "30sm",
					Price: 4500,
				},
			},
		},
	}

	// / sum of product prices
	for _, client := range clients {
		totalPrice := 0
		for _, product := range client.Card.Products {
			productID := product.ProductID
			sizeID := product.SizeID
			quantity := product.Quantity

			productInfo := getProductByID(products, productID)
			if productInfo != nil {
				sizeInfo := getSizeByID(productInfo.Sizes, sizeID)
				if sizeInfo != nil {
					price := sizeInfo.Price
					totalPrice += price * quantity
				}
			}
		}

		fmt.Printf("%s - %d\n", client.Name, totalPrice)
	}

	productQuantities := make(map[int]int)

	for _, client := range clients {
		for _, product := range client.Card.Products {
			productID := product.ProductID
			quantity := product.Quantity

			productQuantities[productID] += quantity
		}
	}

	mostAddedProductID := findMaxQuantityProductID(productQuantities)

	mostAddedProduct := getProductByID(products, mostAddedProductID)
	if mostAddedProduct != nil {
		fmt.Printf("%s - %d\n", mostAddedProduct.Name, productQuantities[mostAddedProductID])
	}
}

func getProductByID(products []Product, id int) *Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func getSizeByID(sizes []Size, id int) *Size {
	for _, size := range sizes {
		if size.ID == id {
			return &size
		}
	}
	return nil
}

func findMaxQuantityProductID(productQuantities map[int]int) int {
	maxQuantity := 0
	maxQuantityProductID := 0

	for productID, quantity := range productQuantities {
		if quantity > maxQuantity {
			maxQuantity = quantity
			maxQuantityProductID = productID
		}
	}

	return maxQuantityProductID
}

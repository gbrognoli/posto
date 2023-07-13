package main

import (
    "fmt"
    "net/http"
    "strings"
)

func main() {
    // Get the current level of the diesel tank.
    tankLevel, err := getTankLevel()
    if err != nil {
        fmt.Println(err)
        return
    }

    // If the tank level is 5000 liters or more, send a quote request to the suppliers.
    if tankLevel >= 5000 {
        suppliers := []string{"supplier1", "supplier2", "supplier3"}
        for _, supplier := range suppliers {
            http.Post("https://www.supplier.com/quote", nil, nil)
        }
    }

    // Wait for a response from the suppliers.
    responses := make(chan string)
    go func() {
        for response := range responses {
            if response != "" {
                fmt.Println("Received quote from", response)
            }
        }
    }()

    // Get the best price per liter.
    bestPrice := 0.0
    for _, supplier := range suppliers {
        response := <-responses
        if response != "" {
            price, err := strconv.ParseFloat(response, 64)
            if err != nil {
                fmt.Println(err)
                return
            }
            if price < bestPrice {
                bestPrice = price
            }
        }
    }

    // If there is a best price, place an order with the supplier.
    if bestPrice > 0 {
        order := fmt.Sprintf("Order %d liters of diesel at %f per liter", tankLevel, bestPrice)
        http.Post("https://www.supplier.com/order", nil, order)
    }
}

// Get the current level of the diesel tank.
func getTankLevel() (int, error) {
    // TODO: Implement this function.
    return 0, nil
}

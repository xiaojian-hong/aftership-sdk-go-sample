package main

import (
	"context"
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
)

type sampleClient struct {
	client *aftership.Client
}

func main() {
	fmt.Print("Enter API Key: ")
	var key string
	_, err := fmt.Scanln(&key)

	client, err := aftership.NewClient(aftership.Config{
		APIKey: key,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get couriers
	result, err := client.GetCouriers(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Total %d couriers", result.Total)
	fmt.Println()

	for _, courier := range result.Couriers {
		fmt.Println(courier.Name)
	}
}

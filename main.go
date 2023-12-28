package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/blockfrost/blockfrost-go"
)

func main() {
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{
			ProjectID:   "previewvtppJPo0y2Caj1rPOjpuQt016NwyixVB",
			Server:      "https://cardano-preview.blockfrost.io/api/v0/",
			MaxRoutines: 0,
		},
	)

	info, err := api.Info(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("API Info:\n\tUrl: %s\n\tVersion: %s\n", info.Url, info.Version)

	address := "addr_test1qp4ptse290n6m8tmm9npwglsqhuzkjnz3dc2zzxjz3dnjm0vk4vnf84uffk4sgr79n3gycjl0l2rxvg20c6mxey5ydfqcem9c8"
	balance, err := queryAddressBalance(api, context.TODO(), address)
	if err != nil {
		log.Fatalf("Error querying address balance: %v", err)
	}

	fmt.Printf("Balance of address %s is %s ADA\n", address, balance)

}

// queryAddressBalance retrieves the balance of the given address using the Blockfrost API.
func queryAddressBalance(api blockfrost.APIClient, ctx context.Context, address string) (string, error) {
	details, err := api.AddressDetails(ctx, address)
	if err != nil {
		return "", err
	}

	receivedSum := sumAmounts(details.ReceivedSum)
	sentSum := sumAmounts(details.SentSum)

	// Calculate the balance (received - sent)
	balance := receivedSum - sentSum

	// Assuming you want the balance in ADA (1 ADA = 1,000,000 lovelace)
	adaBalance := float64(balance) / 1000000

	return fmt.Sprintf("%.6f", adaBalance), nil
}

// sumAmounts calculates the total amount from a slice of AddressAmount.
func sumAmounts(amounts []blockfrost.AddressAmount) int64 {
	var sum int64 = 0
	for _, amount := range amounts {
		// Convert the quantity from string to int64
		value, err := strconv.ParseInt(amount.Quantity, 10, 64)
		if err == nil {
			sum += value
		}
	}
	return sum
}

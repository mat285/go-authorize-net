package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	goauthorizenet "github.com/mat285/go-authorize-net"
	"github.com/mat285/go-authorize-net/config"
	"github.com/mat285/go-authorize-net/transactions"
)

func main() {
	cfg, err := config.ReadFile("_config/dev.yml")
	if err != nil {
		panic(err)
	}

	client := goauthorizenet.New(*cfg)

	data, err := os.ReadFile("_testdata/createTransactionRequest.json")
	if err != nil {
		panic(err)
	}
	var req transactions.CreateTransactionRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		panic(err)
	}

	res, err := client.Transactions().CreateTransaction(context.Background(), req)
	if err != nil {
		panic(err)
	}
	out, _ := json.MarshalIndent(res, "", " ")
	fmt.Println(string(out))
}

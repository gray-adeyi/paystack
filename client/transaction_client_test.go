package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/gray-adeyi/paystack/models"
	_ "github.com/joho/godotenv/autoload"
)

func getTransactionClient(t *testing.T) *TransactionClient {
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		t.Error("unable to retrieve secret key from environmental variable required to run test")
	}
	return NewTransactionClient(WithSecretKey(secretKey))
}

func TestCanInitialize(t *testing.T) {
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		t.Error("unable to retrieve secret key from environmental variable required to run test")
	}
	client := getTransactionClient(t)
	var r models.Response[any]
	if err := client.Initialize(context.TODO(), 20000, "adeyigbenga027@gmail.com", &r); err != nil {
		t.Error(err)
	}
	fmt.Println(r)
}

func TestCanInitializeWithOptionalParameters(t *testing.T) {
	client := getTransactionClient(t)
	var r models.Response[any]
	if err := client.Initialize(context.TODO(), 20000, "<email>", &r,
		WithOptionalParameter("metadata", "{\"ref_id\":\"pot-5085072209\"}"),
	); err != nil {
		t.Error(err)
	}
	fmt.Println(r)
}

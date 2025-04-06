package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

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
	r, err := client.Initialize(context.TODO(), 20000, "adeyigbenga027@gmail.com")
	if err != nil {
		t.Error(err)
	}
	g := make(map[string]interface{})
	err = json.Unmarshal(r.Data, &g)
	if err != nil {
		t.Errorf("Error in client: %v", err)
	}
	fmt.Println(g)
}

func TestCanInitializeWithOptionalParameters(t *testing.T) {
	client := getTransactionClient(t)
	r, err := client.Initialize(context.TODO(), 20000, "<email>",
		WithOptionalParameter("metadata", "{\"ref_id\":\"pot-5085072209\"}"),
	)
	if err != nil {
		t.Error(err)
	}
	g := make(map[string]interface{})
	err = json.Unmarshal(r.Data, &g)
	if err != nil {
		t.Errorf("Error in client: %v", err)
	}
	fmt.Println(g)
}

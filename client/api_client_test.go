package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

func TestAPIClient(t *testing.T) {
	client := NewPaystackClient(WithSecretKey("<paystack-secret-key>"))
	r, err := client.Transactions.Initialize(context.TODO(), 1000, "<dummy-email>", WithOptionalParameter("currency", "NGN"))
	if err != nil {
		t.Errorf("Error in client: %v", err)
	}
	fmt.Println(r)
	g := make(map[string]interface{})
	err = json.Unmarshal(r.Data, &g)
	if err != nil {
		t.Errorf("Error in client: %v", err)
	}
	fmt.Println(g)
}


var (
	Err1 = errors.New("this is error")
	Err2 = errors.New("this is error")
)

func TestErrIsSame(t *testing.T){
	if ! errors.Is(Err1, Err2){
		t.Error("Err1 is not Err2")
	}
}
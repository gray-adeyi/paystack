package paystack

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCanInitialize(t *testing.T) {
	client := NewTransactionClient(WithSecretKey("sk_test_628850d539b080a5fbf5d3bfd4b35d15ac6d071f"))
	r, err := client.Initialize(20000, "adeyigbenga027@gmail.com")
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
	client := NewTransactionClient(WithSecretKey("<paystack-secret-key>"))
	r, err := client.Initialize(20000, "<email>",
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

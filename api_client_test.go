package paystack

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAPIClient(t *testing.T) {
	client := NewAPIClient(WithSecretKey("<paystack-secret-key>"))
	r, err := client.Transactions.Initialize(1000, "<dummy-email>", WithOptionalParameter("currency", "NGN"))
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

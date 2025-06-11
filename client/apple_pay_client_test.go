package client

import (
	"fmt"
	"testing"

	"context"

	"github.com/gray-adeyi/paystack/models"
)

func TestApplePayClient(t *testing.T) {
	client := NewApplePayClient(WithSecretKey("sk_test_628850d539b080a5fbf5d3bfd4b35d15ac6d071f"))
	var r models.Response[any]
	if err := client.All(context.TODO(), &r); err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v", r)
}

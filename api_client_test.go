package paystack

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/gray-adeyi/paystack/models"
)

func TestAPIClient(t *testing.T) {
	client := NewPaystackClient(WithSecretKey("<paystack-secret-key>"))
	var response models.Response[any]
	if err := client.Transactions.Initialize(context.TODO(), 1000, "<dummy-email>", &response, WithOptionalParameter("currency", "NGN")); err != nil {
		t.Errorf("Error in client: %v", err)
	}
	fmt.Printf("%#v", response)
}

var (
	Err1 = errors.New("this is error")
	Err2 = errors.New("this is error")
)

func TestErrIsSame(t *testing.T) {
	if !errors.Is(Err1, Err2) {
		t.Error("Err1 is not Err2")
	}
}

package paystack

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gray-adeyi/paystack/models"
	_ "github.com/joho/godotenv/autoload"
)

func TestCanCreateNewProductClient(t *testing.T) {
	productClient := NewProductClient()
	have := reflect.TypeOf(productClient)
	want := reflect.TypeOf(&ProductClient{})
	if !(want == have) {
		t.Errorf("NewProductClient is not creating a ProductClient. want: %v have: %v", want, have)
	}
}

func getProductClient(t *testing.T) *ProductClient {
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		t.Error("unable to retrieve secret key from environmental variable required to run test")
	}
	return NewProductClient(WithSecretKey(secretKey))
}

func getTestServer(t *testing.T, endpointPath string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		if req.URL.String() != endpointPath {
			t.Errorf("APICall to the wrong endpont. want: %s got: %s", endpointPath, req.URL.String())
		}
		// Send response to be tested
		rw.Write([]byte(`OK`))
	}))
}

func TestCanCreate(t *testing.T) {
	productClient := getProductClient(t)
	var resp models.Response[any]
	if err := productClient.Create(context.TODO(), "test product", "test description", 1000, "NGN", &resp); err != nil {
		t.Errorf("an error occured while calling productClient.Create. err: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("productClient.Create returned wrong response. want status code: %d, got status code: %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestCanCreateMocked(t *testing.T) {
	testServer := getTestServer(t, "/product")
	defer testServer.Close()
	productClient := getProductClient(t)
	productClient.baseUrl = testServer.URL
	var resp models.Response[any]
	if err := productClient.Create(context.TODO(), "test product", "test description", 1000, "NGN"); err != nil {
		t.Errorf("an error occured while calling productClient.Create. err: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("productClient.Create returned wrong response. want status code: %d, got status code: %d", http.StatusOK, resp.StatusCode)
	}
	if !bytes.Equal(resp.Raw, []byte("OK")) {
		t.Errorf("productClient.Create returned wrong Data")
	}
}

func TestCanAll(t *testing.T) {
	productClient := getProductClient(t)
	var resp models.Response[any]
	if err := productClient.All(context.TODO(), &resp); err != nil {
		t.Errorf("an error occured while calling productClient.All. err: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("productClient.All returned wrong respons. want status code: %d, got status code: %d", http.StatusOK, resp.StatusCode)
	}
}

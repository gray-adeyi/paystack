# Paystack

[![Go Reference](https://pkg.go.dev/badge/github.com/gray-adeyi/paystack.svg)](https://pkg.go.dev/github.com/gray-adeyi/paystack)
[![Go Report Card](https://goreportcard.com/badge/github.com/gray-adeyi/paystack)](https://goreportcard.com/report/github.com/gray-adeyi/paystack)

A 3rd party client package for [Paystack](https://paystack.com) in GO.

## Usage
Using `paystack` in your go project is quite simple. All you need to do is import it into your project and create an
`APIClient` as shown below. The `APIClient` is all you need for interacting with [Paystack](https://paystack.com/) in 
your Go project. it has fields on with which are dedicated clients tailored to interfacing with endpoints groups 
provided by [Paystack](https://paystack.com/). As seen in the example below, the `APIClient.Transactions` field 
provides associated functions to interfacing with [Paystack's](https://paystack.com/) Transaction resource. 
e.g., `All` in the example below retrieves all the transactions in you integration.
See the comprehensive list below or package reference for fields on the `APIClient`. It is worth mentioning that all
the associated functions of the fields on the `APIClient` return a `Response` and an `error`.
The `Response` struct contains a `StatusCode` and a `Data` field. The `StatusCode` is the http status code returned 
from making an http request to [Paystack](https://paystack.com/) and the `Data` is a json serializable slice of 
byte containing the response data returned from [Paystack](https://paystack.com/) from making the request. 
The `paystack` package does not provide structs for deserializing the `Data`, hence you're free to deserialize 
the `Data` as you wish. The simplest way to deserialize the data is shown below.
```go
package main

import (
	p "github.com/gray-adeyi/paystack"
	"encoding/json"
	"fmt"
)

func main(){
	client := p.NewAPIClient(p.WithSecretKey("<your-secret-key>"))
	resp, err := client.Transactions.All()
	if err != nil {
		panic(err)
    }
	
	var data map[string]interface{}
	
	if err = json.Unmarshal(resp.Data, data); err != nil {
		panic(err)
    }
	fmt.Println(data)
}
```

## APIClient fields
| Fields                               |
|--------------------------------------|
| `APIClient.Tansactions`              |
| `APIClient.TansactionSplits`         |
| `APIClient.Terminals`                |
| `APIClient.Customers`                |
| `APIClient.DedicatedVirtualAccounts` |
| `APIClient.ApplePay`                 |
| `APIClient.SubAccounts`              |
| `APIClient.Plans`                    |
| `APIClient.Subscriptions`            |
| `APIClient.Products`                 |
| `APIClient.PaymentPages`             |
| `APIClient.PaymentRequests`          |
| `APIClient.Settlements`              |
| `APIClient.TransferControl`          |
| `APIClient.TransferRecipients`       |
| `APIClient.Transfers`                |
| `APIClient.BulkCharges`              |
| `APIClient.Integration`              |
| `APIClient.Charges`                  |
| `APIClient.Disputes`                 |
| `APIClient.Refunds`                  |
| `APIClient.Verification`             |
| `APIClient.Miscellaneus`             |

## Package Limitations
While the `paystack` package has an 100% coverage of the endpoints currently provided by 
[Paystack](https://paystack.com/), it does not have a sufficient test coverage.
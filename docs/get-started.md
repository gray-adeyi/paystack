# Get Started


## Installation

```bash
$ go get -U github.com/gray-adeyi/paystack
```

## Usage
```go
package main

import (
	p "github.com/gray-adeyi/paystack"
	"github.com/gray-adeyi/models"
	"fmt"
	"log"
)

func main(){
	// Init a client to interact with paystack
	client := p.NewClient(p.WithSecretKey("<your-secret-key>"))
	
	// Retrieving all the transactions in you integration 
	response models.Response[[]models.Transaction]
	if err := client.Transactions.All(context.Background(), &response); err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
```
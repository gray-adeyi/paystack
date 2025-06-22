# Examples

## Creating a customer on you integration

```go
package main
import (
		"context"
		"fmt"

		p "github.com/gray-adeyi/paystack"
		"github.com/gray-adeyi/paystack/models"
	)

	func main() {
		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))

		var response models.Response[models.Customer]
		if err := client.Customers.Create(
			context.TODO(),
			"johndoe@example.com",
			"John",
			"Doe", 
			&response,
		); err != nil {
			panic(err)
		}

		fmt.Println(response)

		// With optional parameters
		err := client.Customers.Create(
			context.TODO(),
			"johndoe@example.com",
			"John",
			"Doe", 
			&response, 
			p.WithOptionalPayload("phone","+2348123456789"),
		)
	}
```

> Examples are provided in the docs of the client methods

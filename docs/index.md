![PyPaystack2 logo](assets/paystack.svg)
<center>**A golang client library for Paystack**</center>
<hr/>
Documentation: [https://github.com/gray-adeyi/paystack](https://gray-adeyi.github.io/paystack/)

Source Code: [https://gray-adeyi.github.io/paystack/](https://github.com/gray-adeyi/paystack)
<hr/>
Paystack is a client library in Go for integrating with paystack. It aims at being 
developer friendly and easy to use.

## Disclaimer
This project is an open-source client library for Paystack. It is not officially endorsed or affiliated with Paystack. 
All trademarks and company names belong to their respective owners.


## Installation

```bash
$ go get -U github.com/gray-adeyi/paystack
```

## Example
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

	var data map[string]any

	if err = json.Unmarshal(resp.Data, &data); err != nil {
		panic(err)
    }
	fmt.Println(data)
}
```

## License

This project is licensed under the terms of the MIT license.

## Buy me a coffee

[https://www.buymeacoffee.com/jigani](https://www.buymeacoffee.com/jigani)
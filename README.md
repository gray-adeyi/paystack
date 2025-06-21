# Paystack Client (GO)

[![Go Reference](https://pkg.go.dev/badge/github.com/gray-adeyi/paystack.svg)](https://pkg.go.dev/github.com/gray-adeyi/paystack)
[![Go Report Card](https://goreportcard.com/badge/github.com/gray-adeyi/paystack)](https://goreportcard.com/report/github.com/gray-adeyi/paystack)

An open source client for [Paystack](https://paystack.com) in GO.

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

See [Documentation](#Documentation) for more.

## Documentation

See [Documentation](https://gray-adeyi.github.io/paystack/) for more on this package.

## Disclaimer

This project is an Open Source client library for [Paystack](https://paystack.com/). It is not officially endorsed or
affiliated with [Paystack](https://paystack.com/). All trademarks and company names belong to their respective owners.

## Contributions

Thank you for being interested in contributing to `github.com/gray-adeyi/paystack`.
There are many ways you can contribute to the project:

- [Star on GitHub](https://github.com/gray-adeyi/paystack/)
- Try `github.com/gray-adeyi/paystack` and [report bugs/issues you find](https://github.com/gray-adeyi/paystack/issues/new)
- [Buy me a coffee](https://www.buymeacoffee.com/jigani)

## Other Related Projects

| Name                                                                               | Language              | Functionality                                                                    |
|------------------------------------------------------------------------------------|-----------------------|----------------------------------------------------------------------------------|
| [Paystack CLI](https://pypi.org/project/paystack-cli/)                             | Python                | A command line app for interacting with paystack APIs                            |
| [paystack](https://github.com/gray-adeyi/paystack)                                 | Go (This project)                   | A client library for integration paystack in go                                  |
| [@gray-adeyi/paystack-sdk](https://www.npmjs.com/package/@gray-adeyi/paystack-sdk) | Typescript/Javascript | A client library for integrating paystack in Javascript runtimes (Node,Deno,Bun) |
| [paystack](https://pub.dev/packages/paystack)                                      | Dart                  | A client library for integration paystack in Dart                                |
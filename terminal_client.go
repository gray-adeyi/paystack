package paystack

import (
	"context"
	"fmt"
	"github.com/gray-adeyi/paystack/enum"
	"net/http"
)

// TerminalClient interacts with endpoints related to paystack Terminal resource that allows you to
// build delightful in-person payment experiences.
type TerminalClient struct {
	*restClient
}

// NewTerminalClient creates a TerminalClient
func NewTerminalClient(options ...ClientOptions) *TerminalClient {
	client := NewClient(options...)
	return client.Terminals
}

// SendEvent lets you send an event from your application to the Paystack Terminal
//
// Default response: models.Response[models.TerminalEventData]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[models.TerminalEventData]
// 		data := map[string]any{"id": 7895939, "reference": 4634337895939}
//		if err := client.Terminals.SendEvent(context.TODO(),"30",event.TerminalEventInvoice,enum.TerminalEventActionProcess,data, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TerminalClient) SendEvent(ctx context.Context, terminalId string, eventType enum.TerminalEvent, action enum.TerminalEventAction, data, response any) error {
	payload := map[string]any{
		"type":   eventType,
		"action": action,
		"data":   data,
	}

	return t.APICall(ctx, http.MethodPost, fmt.Sprintf("/terminal/%s/event", terminalId), payload, response)
}

// EventStatus lets you check the status of an event sent to the Terminal
//
// Default response: models.Response[models.TerminalEventStatusData]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[models.TerminalEventStatusData]
//		if err := client.Terminals.EventStatus(context.TODO(),"30","616d721e8c5cd40a0cdd54a6", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TerminalClient) EventStatus(ctx context.Context, terminalId string, eventId string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/terminal/%s/event/%s", terminalId, eventId), nil, response)
}

// TerminalStatus lets you check the availability of a Terminal before sending an event to it
//
// Default response: models.Response[models.TerminalStatusData]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[models.TerminalStatusData]
//		if err := client.Terminals.EventStatus(context.TODO(),"30", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TerminalClient) TerminalStatus(ctx context.Context, terminalId string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/terminal/%s/presence", terminalId), nil, response)
}

// All lets you retrieve the Terminals available on your Integration
//
// Default response: models.Response[[]models.Terminal]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[[]models.Terminal]
//		if err := client.Terminals.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Terminals.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/terminal/
func (t *TerminalClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/terminal", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you get the details of a Terminal
//
// Default response: models.Response[models.Terminal]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[models.Terminal]
//		if err := client.Terminals.FetchOne(context.TODO(),"<terminalId>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TerminalClient) FetchOne(ctx context.Context, terminalId string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/terminal/%s", terminalId), nil, response)
}

// Update lets you update the details of a Terminal
//
// Default response: models.Response[struc{}]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[struct{}]
//		if err := client.Terminals.Update(context.TODO(),"<terminalId>", "New Terminal","somewhere on earth", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TerminalClient) Update(ctx context.Context, terminalId string, name string, address string, response any) error {
	payload := map[string]any{
		"name":    name,
		"address": address,
	}

	return t.APICall(ctx, http.MethodPut, fmt.Sprintf("/terminal/%s", terminalId), payload, response)
}

// Commission lets you activate your debug device by linking it to your Integration
//
// Default response: models.Response[struc{}]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[struct{}]
//		if err := client.Terminals.Commission(context.TODO(),"<serialNumber>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TerminalClient) Commission(ctx context.Context, serialNumber string, response any) error {
	payload := map[string]any{
		"serial_number": serialNumber,
	}

	return t.APICall(ctx, http.MethodPost, "/terminal/commission_device", payload, response)
}

// Decommission lets you unlink your debug device from your Integration
//
// Default response: models.Response[struc{}]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[struct{}]
//		if err := client.Terminals.Decommission(context.TODO(),"<serialNumber>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TerminalClient) Decommission(ctx context.Context, serialNumber string, response any) error {
	payload := map[string]string{
		"serial_number": serialNumber,
	}

	return t.APICall(ctx, http.MethodPost, "/terminal/decommission_device", payload, response)
}

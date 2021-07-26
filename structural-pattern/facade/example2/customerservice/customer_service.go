package customerservice

import "go-design-pattern/structural-pattern/facade/example2/gofood"

type TicketType string

const (
	Gofood TicketType = "gofood"
	GoRide TicketType = "goride"
	GoSend TicketType = "gosend"
)

type Ticket struct {
	UserEmail string
	Type      string
	Message   string
}

type Response struct {
	Status  string
	Message string
}

func HandleTicket(t Ticket) Response {

	switch t.Type {
	case "gofood":
		response := gofood.CheckMessage(t.Message)
		return Response{Status: response}
	case "goride":
	case "gosend":

	}

	return Response{
		Status: "customer service is busy",
	}
}

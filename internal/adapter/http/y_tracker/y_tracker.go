package ytracker

import tracker "github.com/dvsnin/yandex-tracker-go"

type Client interface {
	// GetTicket - get Yandex.Tracker ticket by ticket keys
	GetTicket(ticketKey string) (ticket tracker.Ticket, err error)
	// PatchTicket - patch Yandex.Tracker ticket by ticket key
	PatchTicket(ticketKey string, body map[string]string) (ticket tracker.Ticket, err error)
	// GetTicketComments - get Yandex.Tracker ticket comments by ticket key
	GetTicketComments(ticketKey string) (comments tracker.TicketComments, err error)
	// Myself - get information about the current Yandex.Tracker user
	Myself() (user *tracker.User, err error)
}

type Adapter struct {
	client Client
}

func New(client Client) *Adapter {
	return &Adapter{
		client: client,
	}
}

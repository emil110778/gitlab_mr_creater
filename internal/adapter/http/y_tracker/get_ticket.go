package ytracker

import (
	"github.com/emildeev/gitlab_helper/internal/adapter/http/y_tracker/port"
	ytrackercore "github.com/emildeev/gitlab_helper/internal/core/y_tracker"
)

func (adapter *Adapter) GetTicket(ticketKey string) (ticket ytrackercore.Ticket, err error) {
	rawTicket, err := adapter.client.GetTicket(ticketKey)
	if err != nil {
		return ticket, err
	}
	return port.GetGetTicketResponse(rawTicket), nil
}

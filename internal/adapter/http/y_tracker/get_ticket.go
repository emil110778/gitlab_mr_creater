package ytracker

import (
	"github.com/emil110778/gitlab_mr_creator/internal/adapter/http/y_tracker/port"
	ytrackercore "github.com/emil110778/gitlab_mr_creator/internal/core/y_tracker"
)

func (adapter *Adapter) GetTicket(ticketKey string) (ticket ytrackercore.Ticket, err error) {
	rawTicket, err := adapter.client.GetTicket(ticketKey)
	if err != nil {
		return ticket, err
	}
	return port.GetGetTicketResponse(rawTicket), nil
}

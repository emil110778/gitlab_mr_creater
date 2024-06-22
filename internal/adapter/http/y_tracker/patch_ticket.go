package ytracker

import (
	"github.com/emil110778/gitlab_mr_creator/internal/adapter/http/y_tracker/port"
	ytrackercore "github.com/emil110778/gitlab_mr_creator/internal/core/y_tracker"
)

func (adapter *Adapter) PatchTicket(ticketKey string, ticket ytrackercore.TicketPatch) (err error) {
	request := port.GetPatchTicketRequest(ticket)
	_, err = adapter.client.PatchTicket(ticketKey, request)
	if err != nil {
		return err
	}
	return nil
}

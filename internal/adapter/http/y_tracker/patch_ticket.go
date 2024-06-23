package ytracker

import (
	"github.com/emildeev/gitlab_helper/internal/adapter/http/y_tracker/port"
	ytrackercore "github.com/emildeev/gitlab_helper/internal/core/y_tracker"
)

func (adapter *Adapter) PatchTicket(ticketKey string, ticket ytrackercore.TicketPatch) (err error) {
	request := port.GetPatchTicketRequest(ticket)
	_, err = adapter.client.PatchTicket(ticketKey, request)
	if err != nil {
		return err
	}
	return nil
}

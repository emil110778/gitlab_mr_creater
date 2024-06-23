package ytracker

import ytrackercore "github.com/emildeev/gitlab_helper/internal/core/y_tracker"

type adapterI interface {
	GetTicket(ticketKey string) (ticket ytrackercore.Ticket, err error)
	PatchTicket(ticketKey string, ticket ytrackercore.TicketPatch) (err error)
}

type Service struct {
	adapter adapterI
}

func New(adapter adapterI) *Service {
	return &Service{
		adapter: adapter,
	}
}

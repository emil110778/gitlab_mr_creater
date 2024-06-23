package ytracker

import ytrackercore "github.com/emil110778/gitlab_mr_creator/internal/core/y_tracker"

func (service *Service) GetTicket(ticketKey string) (ticket ytrackercore.Ticket, err error) {
	return service.adapter.GetTicket(ticketKey)
}

func (service *Service) SetMR(ticketKey string, mr string) (err error) {
	return service.adapter.PatchTicket(ticketKey, ytrackercore.TicketPatch{MR: &mr})
}

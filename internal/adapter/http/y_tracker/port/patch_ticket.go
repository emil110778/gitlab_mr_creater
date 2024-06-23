package port

import ytrackercore "github.com/emil110778/gitlab_mr_creator/internal/core/y_tracker"

func GetPatchTicketRequest(ticket ytrackercore.TicketPatch) map[string]string {
	request := make(map[string]string)

	if ticket.Title != nil {
		request[titleKey] = *ticket.Title
	}
	if ticket.Description != nil {
		request[descriptionKey] = *ticket.Description
	}
	if ticket.MR != nil {
		request[mrKey] = *ticket.MR
	}

	return request
}

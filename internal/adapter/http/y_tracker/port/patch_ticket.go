package port

import ytrackercore "github.com/emildeev/gitlab_helper/internal/core/y_tracker"

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

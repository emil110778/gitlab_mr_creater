package port

import ytrackercore "github.com/emildeev/gitlab_helper/internal/core/y_tracker"

func GetGetTicketResponse(response map[string]any) ytrackercore.Ticket {
	self, _ := response[selfKey].(string)
	id, _ := response[idKey].(string)
	key, _ := response[keyKey].(string)
	title, _ := response[titleKey].(string)
	description, _ := response[descriptionKey].(string)
	mr, _ := response[mrKey].(string)

	return ytrackercore.Ticket{
		Self:        self,
		ID:          id,
		Key:         key,
		Title:       title,
		Description: description,
		MR:          mr,
	}
}

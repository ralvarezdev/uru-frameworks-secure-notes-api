package versions

type (
	// ListUserNoteVersionsRequest is the request DTO to list user note versions
	ListUserNoteVersionsRequest struct {
		NoteID int64 `json:"note_id"`
	}

	// ListUserNoteVersionsResponse is the response DTO to list user note versions
	ListUserNoteVersionsResponse struct {
		NoteVersionsID []int64 `json:"note_versions_id"`
	}
)

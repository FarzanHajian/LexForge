package httpapi

type UserUpdateRequest struct {
	StudyItemsPerSession int `json:"studyItemsPerSession"`
}

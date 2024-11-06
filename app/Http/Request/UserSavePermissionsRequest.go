package Request

type UserSavePermissionsRequest struct {
	Request

	Permissions []int `json:"permissions" validate:"required"`
}

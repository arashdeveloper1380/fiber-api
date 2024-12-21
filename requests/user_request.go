package requests

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required,min=3,max=20"`
	Email   string `json:"email" validate:"required,min=3,max=100"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}

package dto

type RegisterUserRequest struct {
	Name            string `json:"name"`
	CellNumber      string `json:"cell_number"`
	Email           string `json:"email"`
	NationalCode    string `json:"national_code"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type RegisterUserResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	CellNumber   string `json:"cell_number"`
	Email        string `json:"email"`
	NationalCode string `json:"national_code"`
	Token        string `json:"token"`
}

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

type LoginUserRequest struct {
	CellNumber string `json:"cell_number"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type LoginUserResponse struct {
	ID         uint   `json:"id"`
	Email      string `json:"email"`
	CellNumber string `json:"cell_number"`
	Name       string `json:"name"`
	Token      string `json:"token"`
}

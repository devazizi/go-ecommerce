package dto

type RegisterUserRequest struct {
	Name         string `json:"name"`
	CellNumber   string `json:"cell_number"`
	Email        string `json:"email"`
	NationalCode string `json:"national_code"`
}

type RegisterUserResponse struct {
	Name         string `json:"name"`
	CellNumber   string `json:"cell_number"`
	Email        string `json:"email"`
	NationalCode string `json:"national_code"`
	Token        string `json:"token"`
}

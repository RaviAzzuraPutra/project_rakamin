package request

type RegisterRequest struct {
	Nama     *string `json:"nama"`
	Email    *string `json:"email"`
	NoTelp   *string `json:"notelp"`
	Password *string `json:"password"`
	IsAdmin  *bool   `json:"is_admin"`
}

type LoginRequest struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

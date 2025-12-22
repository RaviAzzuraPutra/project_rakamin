package request

type RegisterRequest struct {
	Nama     *string `form:"nama"`
	Email    *string `form:"email"`
	NoTelp   *string `form:"notelp"`
	Password *string `form:"password"`
	IsAdmin  *bool   `form:"is_admin"`
}

type LoginRequest struct {
	Email    *string `form:"email"`
	Password *string `form:"password"`
}

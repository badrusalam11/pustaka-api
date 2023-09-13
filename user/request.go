package user

type SignupRequest struct {
	Email    string `binding:"required"`
	Password string `bidning:"required"`
}

type LoginRequest struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

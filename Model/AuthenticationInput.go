package Model

type AuthenticationInput struct {
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Role     string `form:"role" json:"role" binding:"required"`
}

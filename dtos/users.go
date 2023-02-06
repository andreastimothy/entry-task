package dtos

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginWithGoogleRequest struct {
	Token string `json:"token" binding:"required"`
}

type LoginWithGoogleResponse struct {
	Email string `json:"email"`
}

type JwtToken struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
}

package schemas

type ProductCreate struct {
	Name              string `form:"name" binding:"required"`
	Description       string `form:"description"`
	PriceCents        int64  `form:"price_cents" binding:"required"`
	Stock             int    `form:"stock" binding:"required"`
	ResponsibleUserID *uint  `form:"responsible_user_id"`
}

type ProductUpdate struct {
	Name              *string `form:"name"`
	Description       *string `form:"description"`
	PriceCents        *int64  `form:"price_cents"`
	Stock             *int    `form:"stock"`
	ResponsibleUserID *uint   `form:"responsible_user_id"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

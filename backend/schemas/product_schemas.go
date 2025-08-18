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

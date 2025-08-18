package databases

import (
	"backend/configs"
	models "backend/model"
	"fmt"
)

func AutoMigrate() {
	err := configs.DB.AutoMigrate(
		&models.Product{}, &models.User{},
	)
	if err != nil {
		errorLog := fmt.Sprintf("Gagal Auto Migrate: %s", err.Error())
		panic(errorLog)
	}
}

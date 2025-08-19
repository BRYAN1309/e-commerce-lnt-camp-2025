package handlers

import (
	"backend/repository"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// GET /api/products/export
func ExportProductsToExcel(c *gin.Context) {
	products, err := repository.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	f := excelize.NewFile()
	sheet := "Products"
	f.NewSheet(sheet)

	// Set headers
	headers := []string{"ID", "Name", "Description", "Price (cents)", "Stock", "ImageURL", "ResponsibleUserID", "CreatedAt"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Fill data
	for idx, p := range products {
		row := idx + 2 // start from row 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), p.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), p.Name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), p.Description)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), p.PriceCents)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), p.Stock)

		if p.ImageURL != nil {
			f.SetCellValue(sheet, fmt.Sprintf("F%d", row), *p.ImageURL)
		} else {
			f.SetCellValue(sheet, fmt.Sprintf("F%d", row), "")
		}

		if p.ResponsibleUserID != nil {
			f.SetCellValue(sheet, fmt.Sprintf("G%d", row), *p.ResponsibleUserID)
		} else {
			f.SetCellValue(sheet, fmt.Sprintf("G%d", row), "")
		}

		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), p.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	idx, err := f.GetSheetIndex(sheet)
	if err == nil {
		f.SetActiveSheet(idx)
	}

	// Generate filename with timestamp
	fileName := fmt.Sprintf("products_%s.xlsx", time.Now().Format("20060102_150405"))

	// Serve file as download
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("File-Name", fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	_ = f.Write(c.Writer)
}

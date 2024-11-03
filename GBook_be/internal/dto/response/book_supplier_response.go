package response

import (
	"GBook_be/internal/models"
	"time"
)

type BookSupplierResponse struct {
	ID          string          `json:"id"`
	SupplyPrice float64         `json:"supply_price"`
	SupplyDate  time.Time       `json:"supply_date"`
	Book        models.Book     `json:"book"`
	Supplier    models.Supplier `json:"supplier"`
}

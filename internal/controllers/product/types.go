package product

import (
	"github.com/bozd4g/comparator/internal/services/products"
)

type Controller struct {
	service products.Servicer
}

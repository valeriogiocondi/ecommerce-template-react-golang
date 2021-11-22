package product

import (
	"fmt"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"

	BO "products/model"
)

func GetProductList(db *gorm.DB, limit int, offset int) ([]BO.Product, error) {

	var products []BO.Product
	result := db.Table("products").Limit(limit).Offset(offset).Find(&products)

	return products, errorGorm(result.Error)
}

func GetProductById(db *gorm.DB, productId uuid.UUID) (BO.Product, error) {

	var product BO.Product
	result := db.Table("products").First(&product, "uuid = ?", productId)

	return product, errorGorm(result.Error)
}

func CreateNewProduct(db *gorm.DB, newProduct BO.Product) (bool, error) {

	// TODO
	// uuid "github.com/google/uuid"
	// newProduct.InternalId = uuid.New()
	// result := db.Table("products").Create(&newProduct)

	// return (result.Error == nil), errorGorm(result.Error)

	return false, errorGorm(nil)
}

func errorGorm(err error) error {

	if err != nil {
		return fmt.Errorf("GORM - %s", err)
	}

	return nil
}

package mysql

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"

	ConnectionMysql "products/database/mysql/_connection"
	ProductMysql "products/database/mysql/product"
	BO "products/model"
)

// Definition
type getProductList_type func(db *gorm.DB, limit int, offset int) ([]BO.Product, error)
type getProductById_type func(db *gorm.DB, productId uuid.UUID) (BO.Product, error)
type createNewProduct_type func(db *gorm.DB, newProduct BO.Product) (bool, error)

type ProductFactory struct {
	GetProductList   getProductList_type
	GetProductById   getProductById_type
	CreateNewProduct createNewProduct_type
}

// Export
var Connection = ConnectionMysql.Init
var Product = ProductFactory{
	GetProductList:   ProductMysql.GetProductList,
	GetProductById:   ProductMysql.GetProductById,
	CreateNewProduct: ProductMysql.CreateNewProduct,
}

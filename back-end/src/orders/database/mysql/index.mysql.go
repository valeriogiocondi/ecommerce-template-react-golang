package mysql

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"

	ConnectionMysql "orders/database/mysql/_connection"
	OrderMysql "orders/database/mysql/order"
	ProductMysql "orders/database/mysql/product"
	BO "orders/model"
)

// Definition
type getOrderListByCustomerId_type func(db *gorm.DB, customerId uuid.UUID, limit int, offset int) ([]BO.Order, error)
type getOrderById_type func(db *gorm.DB, orderId uuid.UUID) (BO.Order, error)
type createNewOrder_type func(db *gorm.DB, newOrder BO.Order) (bool, error)
type deleteOrder_type func(db *gorm.DB, order BO.Order) (bool, error)

type createNewProductByOrder_type func(db *gorm.DB, newProduct BO.Product) (bool, error)
type getProductListByOrderId_type func(db *gorm.DB, productId uuid.UUID) ([]BO.Product, error)
type deleteProduct_type func(db *gorm.DB, product BO.Product) (bool, error)

type OrderFactory struct {
	GetOrderListByCustomerId getOrderListByCustomerId_type
	GetOrderById             getOrderById_type
	CreateNewOrder           createNewOrder_type
	DeleteOrder              deleteOrder_type
}

type ProductFactory struct {
	GetProductListByOrderId getProductListByOrderId_type
	CreateNewProductByOrder createNewProductByOrder_type
	DeleteProduct           deleteProduct_type
}

// Export
var Connection = ConnectionMysql.Init
var Order = OrderFactory{
	GetOrderListByCustomerId: OrderMysql.GetOrderListByCustomerId,
	GetOrderById:             OrderMysql.GetOrderById,
	CreateNewOrder:           OrderMysql.CreateNewOrder,
	DeleteOrder:              OrderMysql.DeleteOrder,
}
var Product = ProductFactory{
	GetProductListByOrderId: ProductMysql.GetProductListByOrderId,
	CreateNewProductByOrder: ProductMysql.CreateNewProductByOrder,
	DeleteProduct:           ProductMysql.DeleteProduct,
}

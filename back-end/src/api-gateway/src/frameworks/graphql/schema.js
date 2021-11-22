const { gql } = require('apollo-server')
// const schemaGraphql = require('./schemaGraphql.graphql')

const typeDefs = gql`
type Query {
    
    # Products
    queryProductList (params: paginationRequest): ProductListResponse
    queryProductById (params: byIdRequest): ProductResponse
    
    # Customers
    queryCustomerById (params: byIdRequest): CustomerResponse
    
    # Orders
    queryOrderListByCustomerId (params: paginationByIdRequest): OrderListResponse
    queryOrderById (params: byIdRequest): OrderResponse
    queryProductListByOrderId (params: paginationByIdRequest): OrderProductListResponse
}

type Mutation { 
    
    # Customers
    mutationNewCustomer(params: CustomerInput): CustomerResponse 
    
    # Orders
    mutationNewOrder(params: OrderInput): OrderResponse 
}   


#
# REQUEST
#

# Query 

input paginationRequest {
    authToken:  String
    limit:      Int!
    offset:     Int!
}

input byIdRequest {
    authToken:  String
    id:         ID!
}

input paginationByIdRequest {
    authToken:  String
    id:         ID!
    limit:      Int!
    offset:     Int!
}

# Mutation 

input CustomerInput {
    firebaseId:  String!
    firstName:   String!
    lastName:    String!
    email:       String!
    tel:         String
    address:     String
    num:         String
    cap:         String
    city:        String
    state:       String
}

input OrderInput {
    customerId:          String!
    promotionId:         String
    discountPercentage:  Int
    products:            [OrderProductInput]!
}

input OrderProductInput {
    originalPrice:        Float
    discountPercentage:   Int
    price:                Float
}


#
# RESPONSE
#

interface ResponsePrototype {
    errors: [String]
}

interface Response {
    token: String
}

type ProductListResponse implements Response {
    token: String
    productList: [Product]!
}

type ProductResponse implements Response {
    token: String
    product: Product!
}

type CustomerResponse implements Response {
    token: String
    customer: Customer!
}

type OrderListResponse implements Response {
    token: String
    orderList: [Order]!
}

type OrderResponse implements Response {
    token: String
    order: Order!
}

type OrderProductListResponse implements Response {
    token:        String
    orderProductList:  [OrderProduct]!
}


#
# TYPES
#

type Product {
    id:     String!
    name:   String!
    price:  Float!
}

type Promotion {
    id:                  String!
    productId:           String!
    name:                String!
    discountPercentage:  Int
}

type Customer {
    id:          String!
    firebaseId:  String!
    firstName:   String!
    lastName:    String!
    email:       String!
    tel:         String
    address:     String
    num:         String
    cap:         String
    city:        String
    state:       String
}

type Order {
    id:                  String!
    customerId:          String!
    promotionId:         String
    originalPrice:       Float
    discountPercentage:  Int
    totalPrice:          Float
    products:            [OrderProduct]!
}

type OrderProduct {
    id:                   String!
    orderId:              String!
    originalPrice:        Float
    discountPercentage:   Int
    price:                Float
}
`;

module.exports = typeDefs
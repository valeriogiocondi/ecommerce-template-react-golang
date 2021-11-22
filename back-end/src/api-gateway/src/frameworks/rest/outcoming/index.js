const serviceREST = require('../../services/serviceRest');

module.exports = {

  products: {

    getProductList: async (requestDTO) => {

      try {
        
        if (!requestDTO || typeof requestDTO.offset === undefined || typeof requestDTO.limit === undefined)
          throw new Error("DTO undefined");

        var options = {
          host: process.env.PRODUCTS_HOST,
          port: process.env.PRODUCTS_PORT,
          path: `/product?offset=${requestDTO.offset}&limit=${requestDTO.limit}`,
          method: 'GET',
          headers: { 'Content-Type': 'application/json' }
        };
        return await serviceREST.request(options).then(result => result);

      } catch (error) {
        
        console.error(error)
      }
    },  

    getProductById: async (requestDTO) => {

      try {

        var options = {
          host: process.env.PRODUCTS_HOST,
          port: process.env.PRODUCTS_PORT,
          path: `/product/${requestDTO.id}`,
          method: 'GET',
          headers: { 'Content-Type': 'application/json' }
        };
        return await serviceREST.request(options).then(result => result);

      } catch (error) {
        
        console.error(error)
      }
    },
  },

  customers: {
    
    getCustomerById: async (requestDTO) => {

      try {

        var options = {
          host: process.env.CUSTOMERS_HOST,
          port: process.env.CUSTOMERS_PORT,
          path: `/customer/${requestDTO.id}`,
          method: 'GET',
          headers: { 'Content-Type': 'application/json' }
        };
        return await serviceREST.request(options).then(result => result);

      } catch (error) {
        
        console.error(error)
      }
    },

    createNewCustomer: async (requestDTO) => {

      try {
        
        requestDTO = JSON.stringify(requestDTO);

        var options = {
          host: process.env.CUSTOMERS_HOST,
          port: process.env.CUSTOMERS_PORT,
          path: '/customer',
          method: 'POST',
          headers: { 
            'Content-Type': 'application/json', 
            'Content-Length': Buffer.byteLength(requestDTO)
          },
          body: requestDTO
        };
        return await serviceREST.request(options).then(result => result);

      } catch (error) {
        
        console.error(error)
      }
    },
  },

  orders: {

    getOrderListByCustomerId: async (requestDTO) => {

      try {

        var options = {
          host: process.env.ORDERS_HOST,
          port: process.env.ORDERS_PORT,
          path: `/order/customer/${requestDTO.id}/offset/${requestDTO.offset}/limit/${requestDTO.limit}`,
          method: 'GET',
          headers: { 'Content-Type': 'application/json' }
        };
        return await serviceREST.request(options).then(result => result);

      } catch (error) {
        
        console.error(error)
      }
    },

    getOrderById: async (requestDTO) => {

      try {

        var options = {
          host: process.env.ORDERS_HOST,
          port: process.env.ORDERS_PORT,
          path: `/order/${requestDTO.id}`,
          method: 'GET',
          headers: { 'Content-Type': 'application/json' }
        };
        return await serviceREST.request(options).then(result => result);

      } catch (error) {
        
        console.error(error)
      }
    },

    getProductListByOrderId: async (requestDTO) => {

      try {

        var options = {
          host: process.env.ORDERS_HOST,
          port: process.env.ORDERS_PORT,
          path: `/order/${requestDTO.id}/products`,
          method: 'GET',
          headers: { 'Content-Type': 'application/json' }
        };
        return await serviceREST.request(options).then(result => result);

      } catch (error) {
        
        console.error(error)
      }
    },

    createNewOrder: async (requestDTO) => {

      try {
        
        requestDTO = JSON.stringify(requestDTO);

        var options = {
          host: process.env.ORDERS_HOST,
          port: process.env.ORDERS_PORT,
          path: '/order',
          method: 'POST',
          headers: { 
            'Content-Type': 'application/json', 
            'Content-Length': requestDTO.length
          },
          body: requestDTO
        };
        return await serviceREST.request(options).then(result => result);

      } catch (error) {
        
        console.error(error)
      }
    },
  },
};
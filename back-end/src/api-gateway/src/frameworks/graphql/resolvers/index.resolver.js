const outcomingREST = require('../../rest/outcoming');
const AuthController = require('../../../controller/AuthController');

/* 
 * TO SEE SCHEMA
 *  -> express-server/src/frameworks/graphql/resolvers/index.js
 *
 *  TO SEE QUERY (RELAY)
 *  -> react-client/src/_model/relay/query/*.ts
 * 
 */

module.exports = () => {

    return {
        
        Query: {
            
            /* 
            *  LOGIN
            async loginAuth(parent, args, context, info) {

                return AuthController.login(args);
            },
            async loginVerify(parent, args, context, info) {

                return AuthController.verify(args);
            },
            async logout(parent, args, context, info) {

                return AuthController.logout(args);
            },
            */

            /* 
            *  PRODUCTS
            */
           async queryProductList(parent, args, context, info) {

                if (!context)
                    return {
                        token: '123',
                    }
               
               const res = await outcomingREST.products.getProductList(args.params);
               return { 
                   productList: res.Payload,
                   token: context.user.accessToken,
               };
            },
            async queryProductById(parent, args, context, info) {
                
                const res = await outcomingREST.products.getProductById(args.params);
                return { 
                    product: res.Payload,
                    token: context.user.accessToken,
                };
            },


            /* 
            *  CUSTOMERS
            */
            async queryCustomerById(parent, args, context, info) {

                const res = await outcomingREST.customers.getCustomerById(args.params);
                return { 
                    customer: res.Payload,
                    token: context.user.accessToken,
                };
            },
            

            /* 
            *  ORDERS
            */
            async queryOrderListByCustomerId(parent, args, context, info) {
                
                const res = await outcomingREST.orders.getOrderListByCustomerId(args.params);
                return { 
                    orderList: res.Payload,
                    token: context.user.accessToken,
                };
            },
            async queryOrderById(parent, args, context, info) {
                
                const res = await outcomingREST.orders.getOrderById(args.params);
                return { 
                    order: res.Payload,
                    token: context.user.accessToken,
                };
            },
            async queryProductListByOrderId(parent, args, context, info) {
                
                const res = await outcomingREST.orders.getProductListByOrderId(args.params);
                return { 
                    orderProductList: res.Payload,
                    token: context.user.accessToken,
                };
            },
        },

        Mutation: {
            
            /* 
            *  CUSTOMERS
            */
            async mutationNewCustomer(parent, args, context, info) {
                const res = await outcomingREST.customers.createNewCustomer(args.params);
                return { 
                    customer: res.Payload,
                    token: context.user.accessToken,
                };
            },


            /* 
            *  ORDERS
            */
            async mutationNewOrder(parent, args, context, info) {
                const res = await outcomingREST.orders.createNewOrder(args.params);
                return { 
                    order: res.Payload,
                    token: context.user.accessToken,
                };
            },
        }
    }
};
const IAuth = require('../../../application/contracts/IAuth');
const ModelAuth = require('../../../entities/sequelize/AuthModel');
const AuthResponseBO = require('../../../entities/response/bo/AuthResponseBO');

class AuthPersistence extends IAuth {

    async get(requestBO) {

        return new Promise((resolve, reject) => {

            let criteria = { 
                where: {
                    userId: requestBO.userId
                },
                order: [ [ 'date_create', 'DESC' ] ]
            };

            ModelAuth.findOne(criteria).then((result) => {

                if (result)
                    resolve(new AuthResponseBO(result));
                    
                resolve({});
                
            }).catch((err) => {
            
                console.error(err);
                reject(err);
            });  
        });  
    };

    async getRefreshTokenDb(requestBO) {

        return new Promise((resolve, reject) => {

            let criteria = { 
                where: {
                    // userId: requestBO.userId,
                    accessToken: requestBO.accessToken
                },
                order: [ [ 'date_create', 'DESC' ] ]
            };

            ModelAuth.findOne(criteria).then((result) => {

                if (result)
                    resolve(new AuthResponseBO(result));
                    
                resolve({});
                
            }).catch((err) => {
            
                console.error(err);
                reject(err);
            });  
        });  
    };

    async insert(requestBO) {

        return new Promise((resolve, reject) => {

            ModelAuth.create(requestBO).then((result) => {

                resolve(result);
                
            }).catch((err) => {
            
                console.error(err);
                reject(err);
            });  
        });  
    };
};

module.exports = new AuthPersistence();
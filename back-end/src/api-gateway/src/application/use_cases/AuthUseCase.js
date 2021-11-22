const AuthPersistence = require('../../frameworks/persistance/sequelize/AuthPersistence');

class Auth {

    constructor() {

    }
    
    async get(requestBO) {
        
        return AuthPersistence.get(requestBO);
    }

    async getRefreshTokenDb(requestBO) {
        
        return AuthPersistence.getRefreshTokenDb(requestBO);
    }
    
    async insert(requestBO) {
        
        return AuthPersistence.insert(requestBO);
    }
};

module.exports = new Auth();
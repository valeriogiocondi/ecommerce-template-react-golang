module.exports = class AuthRequestBO {

    constructor(params) {

        this.id = params.id;
        this.uuid = params.uuid;
        this.userId = params.userId;
        this.accessToken = params.accessToken;
        this.refreshToken = params.refreshToken;
    }
};
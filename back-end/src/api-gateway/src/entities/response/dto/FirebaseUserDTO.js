module.exports = class FirebaseUserDTO {

    constructor(params) {

        this.id = params.id;
        this.firstName = params.firstName;
        this.lastName = params.lastName;
        this.email = params.email;
        this.accessToken = params.accessToken;
        this.refreshToken = params.refreshToken;
    }
};
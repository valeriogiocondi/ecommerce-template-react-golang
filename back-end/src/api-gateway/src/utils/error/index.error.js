module.exports = class CustomError {

    constructor(service, code, message) {
        this.type = 'error';
        this.service = service;
        this.code = code;
        this.message = message;
    }
};
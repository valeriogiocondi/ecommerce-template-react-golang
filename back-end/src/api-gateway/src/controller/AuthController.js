// const jwt = require('json-web-token')
const jsonwebtoken = require('jsonwebtoken')
const { v4: uuidv4 } = require('uuid');
const firebaseService = require('../frameworks/firebase');
const AuthUseCase = require("../application/use_cases/AuthUseCase");
const AuthRequestBO = require("../entities/request/bo/AuthRequestBO");
const AuthRequestDTOMapper = require('../entities/request/mappers/AuthRequestDTOMapper');
const AuthResponseDTOMapper = require('../entities/response/mappers/AuthResponseDTOMapper');
const outcomingREST = require('../frameworks/rest/outcoming');
const FirebaseUserDTO = require('../entities/response/dto/FirebaseUserDTO');


process.env.ACCESS_TOKEN_LIFE = "1d";
process.env.ACCESS_TOKEN_LIFE = "60";

class AuthController {
    
  check = async (token) => (token && token !== 'Bearer null' && token.startsWith('Bearer ')) ? await firebaseService.verifyToken(token.split('Bearer ')[1]) : false

  async login(requestDTO) {
    
    if (!requestDTO || !requestDTO.username || !requestDTO.password)
      // res.status(500).send({ error: 'Server error!' })
      return '';
    
    return await firebaseService.login(requestDTO.username, requestDTO.password)
      .then((res) => {

        if (res.type === 'error') throw res;
        
        AuthUseCase.insert(
          new AuthRequestBO({ 
            userId: res.userId,
            accessToken: res.accessToken,
            refreshToken: res.refreshToken,
          }) 
        );
        return res;
      })
      .catch(e => e)
  }

  async create(requestDTO) {

    if (!requestDTO)
      // res.status(500).send({ error: 'Server error!' })
      return '';

    return await firebaseService.create(requestDTO.email, requestDTO.password)
      .then((res) => {

        if (res.type === 'error') throw res;
        
        try {
          
          // request to customer microservice
          requestDTO.firebaseId = res.user.uid;
          outcomingREST.customers.create(requestDTO);
          return res.user.token;

        } catch (e) {
          
          console.error(e)
          // TODO
          // Test
          // Handle rollback - delete new user from Firebase
          firebaseService.deleteCurrentUser().then().catch();

          return new CustomError('misc', e.code, e.message) 
        }
      })
      .catch(e => e)
  }

  // Return User - if accessToken has expired, then refresh and update User
  async getUser(accessToken) {
    
    if (!accessToken) 
      return null;
      // return new CustozmError('firebase', 0, 'Token revoked');
    
    const res = await this.verify(accessToken);

    const refreshTokenFromDB = this.getRefreshTokenDb(accessToken);

    if (res.type === 'error') {

      const errHandler = {
        'auth/id-token-revoked': () => {
          // Token has been revoked. Inform the user to reauthenticate or signOut() the user.
          this.logout();
          // res.status(500).send({ error: 'Server error!' })
          return '';
        },
        'auth/argument-error': async () => {
          // Decoding Firebase ID token failed. Make sure you passed the entire string JWT which represents an ID token.
          // res.status(500).send({ error: 'Server error!' })
          return '';
        },
        'auth/id-token-expired': async () => {
          // Token has expired. Get a fresh ID token from your client app and try again.
          return (await this.getNewToken(refreshTokenFromDB)).access_token;
        },
      }
      accessToken = await errHandler[res.code]();

      if (!accessToken) 
        return null;
    }

    return new FirebaseUserDTO({
      email:           res.email,
      accessToken:     accessToken,
      refreshToken:    refreshTokenFromDB,
    });
  }

  // Return User
  async verify(accessToken) {
    
    if (!accessToken) return null;    

    return await firebaseService.verifyIdToken(accessToken)
      .then(async (res) => {

        if (res.type === 'error') throw res;
        return res;
      })
      .catch(e => e)
  }
    
  async getNewToken(refreshToken) {
    
    return await firebaseService.getNewToken(refreshToken)
      .then((res) => {
        
        if (res.type === 'error') throw res; 

        AuthUseCase.insert(
          new AuthRequestBO({ 
            userId: res.userId,
            accessToken: res.accessToken,
            refreshToken: res.refreshToken,
          }) 
        );

        return res;
      })
      .catch(e => e)
  }

  async getRefreshTokenDb(accessToken) {
    
    const requestBO = new AuthRequestBO({ accessToken: accessToken });
    AuthUseCase.getRefreshTokenDb(requestBO);
  }

  async updateEmail(requestDTO) {
    
    if (!requestDTO && Object.keys(requestDTO).length === 0)
      // res.status(500).send({ error: 'Server error!' })
      return '';
    
    return await firebaseService.updateEmail(requestDTO)
      .then((res) => {
        if (res.type === 'error') throw res;
      })
      .catch(e => e)
  }
  
  async updatePassword(requestDTO) {
    
    if (!requestDTO && Object.keys(requestDTO).length === 0)
      // res.status(500).send({ error: 'Server error!' })
      return '';
    
    return await firebaseService.updatePassword(requestDTO)
      .then((res) => {
        if (res.type === 'error') throw res;
      })
      .catch(e => e)
  }

  async logout() {
    
    return await firebaseService.logout()
      .then((res) => {
        
        if (res.type === 'error') throw res;
      })
      .catch(e => e)
  }
}

module.exports = new AuthController();
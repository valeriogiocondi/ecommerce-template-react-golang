/* 
 *
 *    https://firebase.google.com/docs/auth/admin/errors
 *
 *    https://console.firebase.google.com
 *
 */
'use strict';
require('dotenv').config();
const CustomError = require('../../utils/error/index.error');
const firebaseApp = require('firebase');
const firebaseAdmin = require('firebase-admin')
const serviceAccount = require("./serviceAccountKey.json");
const serviceREST = require('../services/serviceRest');

class FirebaseProxy {

  constructor() { }

  init() {
    
    try {
      
      firebaseApp.initializeApp({
        apiKey:              process.env.FIREBASE_API_KEY,
        authDomain:          process.env.FIREBASE_AUTH_DOMAIN,
        projectId:           process.env.FIREBASE_PROJECT_ID,
        storageBucket:       process.env.FIREBASE_STORAGE_BUCKET,
        messagingSenderId:   process.env.FIREBASE_MESSAGING_SENDER_ID,
        appId:               process.env.FIREBASE_APP_ID,
      });
      firebaseAdmin.initializeApp({
        credential: firebaseAdmin.credential.cert(serviceAccount)
      });
      
    } catch (e) {
      
      console.error(e);
    }
  }

  async create(username, password) {

    return await
      firebaseApp.auth().createUserWithEmailAndPassword(username, password)
        .then(result => result)
        .catch(e => new CustomError('firebase', e.code, e.message) )
  }

  // TODO
  // Test
  async deleteCurrentUser() {

    return await
      firebaseApp.auth().currentUser.delete()
        .then(result => result)
        .catch(e => new CustomError('firebase', e.code, e.message) )
  }

  // TODO
  // Use another service
  async sendEmailVerification() {

    return await
      firebaseApp.auth().currentUser.sendEmailVerification()
        .then(result => result)
        .catch(e => new CustomError('firebase', e.code, e.message) )
  }

  // TODO
  // Test
  async updateEmail(newEmail) {

    // send link by email

    // console.log(firebaseApp.auth().currentUser)
    // var cred = firebase.auth.EmailAuthProvider.credential(
      //   email,
      //   password
      // );
      // firebaseApp.auth().reauthenticateWithCredential(

    return await
      firebaseApp.auth().updateEmail(newEmail)
        .then(result => {console.log(result)})
        .catch(e => new CustomError('firebase', e.code, e.message) )
  }
  
  // TODO
  // Test
  async updatePassword(requestDTO) {

    // send link by email

    if (!firebaseApp.auth().currentUser) {
      
      const credential = firebaseApp.auth().EmailAuthProvider.credential(requestDTO.email, requestDTO.oldPassword)
      firebaseApp.auth().reauthenticateWithCredential(credential)
        .then(result => result)
        .catch(e => new CustomError('firebase', e.code, e.message) );

      console.log(firebaseApp.auth().currentUser)
    }

    return await
      firebaseApp.auth().updatePassword(requestDTO.newPassword)
        .then(result => {console.log(result); return result})
        .catch(e => new CustomError('firebase', e.code, e.message) )
  }

  async login(username, password) {
  
    return await 
      firebaseApp.auth().signInWithEmailAndPassword(username, password)
        .then(( {user} ) => {

          return {
            userId: user.uid,
            firstName: user.firstName,
            lastName: user.lastName,
            email: user.email,
            accessToken: user._lat,
            refreshToken: user.refreshToken,
          }
        })
        .catch(e => new CustomError('firebase', e.code, e.message) )
  }
  
  //  Token verification: return DecodedIdToken - user information
  async verifyIdToken(token) {
  
    return await 
      firebaseAdmin.auth().verifyIdToken(token)
        .then(result => result)
        .catch(e => new CustomError('firebase', e.code, e.message) )
  }

  async getNewToken(refreshToken) {

    try {

      var options = {
        host: 'securetoken.googleapis.com',
        path: `/v1/token?key=${process.env.FIREBASE_API_KEY}&grant_type=refresh_token&refresh_token=${refreshToken}`,
        method: 'POST',
        headers: { 'Content-Type': 'application/json' }
      };
      const result =  await serviceREST.requestSSL(options).then(res => res);

      return {
        userId: result.user_id,
        accessToken: result.access_token,
        refreshToken: result.refresh_token,
      }

    } catch (e) { return e }
  } 

  async logout() {
  
    return await 
    firebaseApp.auth().signOut()
        .then(result => result)
        .catch(e => new CustomError('firebase', e.code, e.message) )
  } 
}


// SINGLETON
// Export an instance of the class directly
const firebaseProxy = new FirebaseProxy();
firebaseProxy.init();

module.exports = firebaseProxy;
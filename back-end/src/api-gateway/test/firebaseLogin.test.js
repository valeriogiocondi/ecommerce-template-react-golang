var assert = require('assert');
const AuthController = require('../src/controller/AuthController');


// describe('Array', function() {

//   describe('#indexOf()', function() {
//     it('should return -1 when the value is not present', function() {
//       assert.equal([1, 2, 3].indexOf(4), -1);
//     });
//   });

// });


describe('Firebase', function() {

  it('login-failed', function() {

    let credential = { username: "test.test@gmail.co", password: "12345" };

    AuthController.login(credential.username, credential.password)
      .then((user) => {

        console.log(user)
    });

  });

  it('login-passed', function() {

    let credential = { username: "test.test@gmail.com", password: "123456" };

    AuthController.login(credential.username, credential.password)
      .then((user) => {

        console.log(user)
    });

  });
});

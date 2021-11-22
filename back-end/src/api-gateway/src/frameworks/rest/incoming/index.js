const { ApolloServer } = require('apollo-server-express');
const resolvers = require('../../graphql/resolvers/index.resolver');
const schema = require('../../graphql/schema');
const outcomingREST = require('../outcoming');
const AuthController = require('../../../controller/AuthController');

module.exports = (app) => {

  /* 
   *  AUTH Middleware
   *
   */
  app.use(async (req, res, next) => {
    
    if (
      req.originalUrl.startsWith('/graphql') ||
      req.originalUrl.startsWith('/api') ||
      req.originalUrl.startsWith('/auth') 
    ) {
      
      // GraphQL POST without Authorization
      // if (
      //   req.originalUrl.startsWith('/graphql') &&
      //   req.method == 'POST' &&
      //   !await AuthController.check(req.headers.authorization) 
      // )
      //   return res.status(401).send({ message: '401 Unauthorized' });

      next();
    
    } else {

      // Unauthorized
      return res.status(401).send({ message: '401 Unauthorized' });
    }
    
  });

  /* 
   *  Apollo-Server GraphQL Middleware
   *
   */
  const apolloMiddleware = new ApolloServer({
      typeDefs: schema,
      resolvers: resolvers(),
      context: async ({ req }) => {
        
        const token = req.headers.authorization || '';

        if (!token || !token.startsWith('Bearer ')) throw new Error('token invalid');

        // Try to retrieve a user with the token
        const user = await AuthController.getUser( token.split("Bearer ").pop() );
        // console.log(user)
        
        // optionally block the user
        // we could also check user roles/permissions here
        // if (!user) throw new AuthenticationError('you must be logged in');
        if (!user) throw new Error('you must be logged in');
        
        // add the user to the context
        return { user };
      },
  });
  apolloMiddleware.applyMiddleware({ app, path: '/graphql' });
    
  /* 
   *  AUTH
   *
   */
  app.post('/auth/login', async (req, res) => {
  
    const requestDTO = {
      username: req.body.email,
      password: req.body.password
    };
    result = await AuthController.login(requestDTO);
    res.send(result);
  });

  app.post('/auth/reset-email', async (req, res) => {
  
    const requestDTO = {
      email: req.body.email,
      oldPassword: req.body.oldPassword,
      newPassword: req.body.newPassword,
    };
    result = await AuthController.updatePassword(requestDTO);
    res.send(result);
  });

  app.post('/auth/update-email', async (req, res) => {
  
    const requestDTO = {
      email: req.body.email,
      oldPassword: req.body.oldPassword,
      newPassword: req.body.newPassword,
    };
    result = await AuthController.updateEmail(requestDTO);
    res.send(result);
  });

  app.post('/auth/create', async (req, res) => {
  
    const requestDTO = {
      firebaseId: req.body.firebaseId,
      firstName:  req.body.firstName,
      lastName:   req.body.lastName,
      email:      req.body.email,
      password:   req.body.password,
      tel:        req.body.tel,
      address:    req.body.address,
      num:        req.body.num,
      cap:        req.body.cap,
      city:       req.body.city,
      state:      req.body.state
    };
    result = await AuthController.create(requestDTO);
    res.send(result);
  });
  
};
const { DataTypes } = require("sequelize");
const database = require("../../config/database");

const AuthbModel = database.sequelize.instance.define('Auth', {
    id: {
      type: DataTypes.INTEGER,
      primaryKey: true,
      autoIncrement: true,
      allowNull: false
    },
    userId: {
      type: DataTypes.STRING,
      allowNull: false
    },  
    accessToken: {
      type: DataTypes.STRING,
      allowNull: false
    },  
    refreshToken: {
      type: DataTypes.STRING,
      allowNull: false
    },  
    date_create: {
      type: DataTypes.STRING,
      allowNull: true
    },
    date_edit: {
      type: DataTypes.STRING,
      allowNull: true
    }
  }, 
  { 
    timestamps: false,
    freezeTableName: true,
    tableName: 'AUTH'
  }
);

// `sequelize.define` also returns the model
module.exports = AuthbModel;
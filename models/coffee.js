'use strict';
const {
  Model
} = require('sequelize');
module.exports = (sequelize, DataTypes) => {
  class Coffee extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
    }
  }
  Coffee.init({
    name: DataTypes.STRING,
    roastery: DataTypes.STRING,
    processing: DataTypes.STRING,
    roast: DataTypes.STRING,
    t_notes: DataTypes.STRING,
    estate: DataTypes.STRING,
    type: DataTypes.STRING,
    link: DataTypes.STRING,
    tags: DataTypes.STRING,
    availablity: DataTypes.STRING,
    published: DataTypes.BOOLEAN
  }, {
    sequelize,
    modelName: 'Coffee',
  });
  return Coffee;
};
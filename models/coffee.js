module.exports = (sequelize, Sequelize) => {
    const Coffee = sequelize.define("coffee", {
      name: {
        type: Sequelize.STRING
      },
      description: {
        type: Sequelize.STRING
      },
      published: {
        type: Sequelize.BOOLEAN
      }
    });
    return Coffee;
  };
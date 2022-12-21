module.exports = (sequelize, Sequelize) => {
    const Coffee = sequelize.define("coffee", {
      name: {
        type: Sequelize.STRING
      },
      roastery: {
        type: Sequelize.STRING
      },
      processing: {
        type: Sequelize.STRING
      },
      roast: {
        type: Sequelize.STRING
      },
      t_notes: {
        type: Sequelize.STRING
      },
      estate: {
        type: Sequelize.STRING
      },
      type: {
        type: Sequelize.STRING
      },
      link: {
        type: Sequelize.STRING
      },
      tags: {
        type: Sequelize.STRING
      },
      availablity: {
        type: Sequelize.STRING
      },
      published: {
        type: Sequelize.BOOLEAN
      }
    });
    return Coffee;
  };
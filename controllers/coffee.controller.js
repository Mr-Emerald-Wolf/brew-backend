const db = require("../models");
const Coffee = db.coffee;
const Op = db.Sequelize.Op;
// Create and Save a new coffee
exports.create = (req, res) => {
    // Validate request
    if (!req.body.title) {
        res.status(400).send({
            message: "Content can not be empty!"
        });
        return;
    }
    // Create a coffee
    const Coffee = {
        name: req.body.title,
        description: req.body.description,
        published: req.body.published ? req.body.published : false
    };
    // Save coffee in the database
    Coffee.create(coffee)
        .then(data => {
            res.send(data);
        })
        .catch(err => {
            res.status(500).send({
                message:
                    err.message || "Some error occurred while creating the coffee."
            });
        });
};
// Retrieve all coffees from the database.
exports.findAll = (req, res) => {
    const title = req.query.title;
    var condition = title ? { title: { [Op.iLike]: `%${title}%` } } : null;
    Coffee.findAll({ where: condition })
      .then(data => {
        res.send(data);
      })
      .catch(err => {
        res.status(500).send({
          message:
            err.message || "Some error occurred while retrieving coffees."
        });
      });
  };
// Find a single coffee with an id
exports.findOne = (req, res) => {

};
// Update a coffee by the id in the request
exports.update = (req, res) => {

};
// Delete a coffee with the specified id in the request
exports.delete = (req, res) => {

};
// Delete all coffees from the database.
exports.deleteAll = (req, res) => {

};
// Find all published coffees
exports.findAllPublished = (req, res) => {

};
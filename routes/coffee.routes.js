module.exports = app => {
    const coffee = require("../controllers/coffee.controller.js");
    var router = require("express").Router();
    // Create a new Tutorial
    router.post("/", coffee.create);
    // Retrieve all Tutorials
    router.get("/", coffee.findAll);
    app.use('/api/coffee', router);
};
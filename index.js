const express = require('express')
require('dotenv').config()
const Sequelize = require("sequelize");
const cors = require('cors')
const bodyParser = require('body-parser')
const app = express()
var corsOptions = {
    origin: "http://localhost:8081"
};

app.use(cors(corsOptions));
// parse requests of content-type - application/json
app.use(bodyParser.json())
app.use(
    bodyParser.urlencoded({
        extended: true,
    })
)

const db = require("./models");

db.sequelize.sync()
    .then(() => {
        console.log("Synced db.");
    })
    .catch((err) => {
        console.log("Failed to sync db: " + err.message);
    });

app.get('*', (req, res) => res.status(404).send({
    message: 'Nothing to see here',
}));

require("./routes/tutorial.routes")(app);
require("./routes/coffee.routes")(app);

// set port, listen for requests
const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}.`);
});
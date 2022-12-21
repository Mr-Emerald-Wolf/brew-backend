const express = require('express')
require('dotenv').config()
const cors = require('cors')
const bodyParser = require('body-parser')



const app = express()
var corsOptions = {
    origin: "http://localhost:3002"
};

app.use(cors());

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

const routes = require("./routes/index");
app.use('/',routes)

// set port, listen for requests
const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}.`);
});
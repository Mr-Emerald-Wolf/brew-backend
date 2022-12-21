const express = require('express')
const router = express.Router()

const user = require('./user')
const coffee = require('./coffee')

router.use('/users', user)
router.use('/coffee', coffee)

module.exports = router
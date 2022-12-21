const express = require('express')
const router = express.Router()
const { getUser ,findUser,createUser ,updateUser, deleteUser, loginUser} = require('../controllers/userController')
// const { refreshToken } = require('../controllers/RefreshToken')
const { verifyToken } = require("../middleware/VerifyToken");

router.post('/createUser',createUser)
router.get('/getUser' ,getUser)
router.get('/findUser', verifyToken, findUser)
router.put('/updateUser',updateUser)
router.delete('/deleteUser',deleteUser)
router.post('/login', loginUser)
router.get('/token');
 
module.exports = router
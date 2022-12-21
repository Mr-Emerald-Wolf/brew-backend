const { User } = require('../models')
const bcrypt = require('bcrypt')
const jwt = require("jsonwebtoken");

exports.getUser = async (req, res) => {
    try {
        const data = await User.findAll()
        res.send({
            status: true,
            data: data
        })
    }
    catch (err) {
        res.send({
            status: false,
            message: err
        })
    }
};

exports.createUser = async (req, res) => {
    const { firstName, lastName, email, password, role } = req.body;
    // if(password !== password) return res.status(400).json({msg: "Password and Confirm Password do not match"});
    const salt = await bcrypt.genSalt();
    const hashPassword = await bcrypt.hash(password, salt);
    try {
        const data = await User.create({
            firstName: firstName,
            lastName: lastName,
            email: email,
            password: hashPassword,
            role: role
        })

        const accessToken = jwt.sign({ firstName, role, email }, process.env.ACCESS_TOKEN_SECRET);

        res.send({
            status: true,
            data: data,
            accessToken: accessToken
        })
    }
    catch (err) {
        res.send({
            status: false,
            message: err
        })
    }
}

exports.findUser = async (req, res) => {
    try {
        const email = req.email
        const findUser = await User.findOne({ where: { email: email } })
        if (findUser) {
            res.send({
                status: true,
                data: findUser
            })
        } else {
            res.send({
                status: true,
                message: 'User not found'
            })
        }
    }
    catch (err) {
        res.send({
            status: false,
            message: err
        })
    }
}

exports.updateUser = async (req, res) => {
    try {
        const id = req.query.id
        const body = req.body
        const findUser = await User.findOne({ where: { id: id } })
        if (findUser) {
            const data = await User.update(body, { where: { id: id } })
            if (data) {
                res.send({
                    status: true,
                    message: 'User updated'
                })
            }
            else {
                res.send({
                    status: true,
                    message: 'User not updated'
                })
            }
        } else {
            res.send({
                status: true,
                message: 'User not found'
            })
        }
    }
    catch (err) {
        res.send({
            status: false,
            message: err
        })
    }
}

exports.deleteUser = async (req, res) => {
    try {
        const id = req.query.id
        const findUser = await User.findOne({ where: { id: id } })
        if (findUser) {
            const data = await User.destroy({ where: { id: id } })
            if (data) {
                res.send({
                    status: true,
                    message: 'User deleted'
                })
            } else {
                res.send({
                    status: true,
                    message: 'User not deleted'
                })
            }
        } else {
            res.send({
                status: true,
                message: 'User not found'
            })
        }
    }
    catch (err) {
        res.send({
            status: false,
            message: err
        })
    }
}
exports.loginUser = async (req, res) => {
    try {
        const user = await User.findAll({
            where: {
                email: req.body.email
            }
        });
        const match = await bcrypt.compare(req.body.password, user[0].password);
        if (!match) return res.status(400).json({ msg: "Wrong Password" });
        const userId = user[0].id;
        const name = user[0].firstName;
        const email = user[0].email;
        const role = user[0].role;
        const accessToken = jwt.sign({ userId, name, role, email }, process.env.ACCESS_TOKEN_SECRET);

        res.json({ accessToken });
    } catch (error) {
        console.log(error);
        res.status(404).json({ msg: "Email not found" });
    }
}
exports.userLogout = async (req, res) => {
    const refreshToken = req.cookies.refreshToken;
    if (!refreshToken) return res.sendStatus(204);
    const user = await User.findAll({
        where: {
            refresh_token: refreshToken
        }
    });
    if (!user[0]) return res.sendStatus(204);
    const userId = user[0].id;
    await User.update({ refresh_token: null }, {
        where: {
            id: userId
        }
    });
    res.clearCookie('refreshToken');
    return res.sendStatus(200);
}
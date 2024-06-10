const express = require('express');
const { calcularMatriz } = require('../controllers/matriz.controller');

const router = express.Router();

router.post('/', calcularMatriz);

module.exports = router;

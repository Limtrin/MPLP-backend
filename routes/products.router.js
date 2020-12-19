const Router = require('express')
const router = new Router()
const productsController = require('../controllers/products.controller')

router.get('/products', productsController.getProduct)
router.post('/products', productsController.createProduct)

module.exports = router
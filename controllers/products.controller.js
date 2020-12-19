const db = require('../db')

class ProductsController {
    async getProduct(req, res) {
        const products = await db.query('SELECT * FROM products')
        res.json(products.rows)
    }

    async createProduct(req, res) {
        const {name, price, instock} = req.body
        const newProduct = await db.query(`INSERT INTO products (name, price, instock) values ($1, $2, $3) RETURNING *`,[name, price, instock])

        res.json(newProduct.rows[0])
    }
}

module.exports = new ProductsController()
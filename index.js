const express = require('express')
const productsRouter = require('./routes/products.router')

const PORT = process.env.PORT || 8080

const app = express()

app.use(express.json())
app.use('/', productsRouter)

app.listen(PORT, () => console.log(`server start on port :${PORT}`))

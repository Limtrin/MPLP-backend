const express = require('express')

const app = express()
const hostname = '127.0.0.1'

var port = process.env.port || 3000

app.get('/', (req, res) => {
    res.send('Погнали!')
})

app.listen(port, hostname, () => {
    console.log('running port' + port)
})
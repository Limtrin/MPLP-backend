const express = require('express')

const app = express()

var port = process.env.port || 3000

app.get('/', (req, res) => {
    res.send('Погнали!')
})

app.listen(port, () => {
    console.log('running port' + port)
})
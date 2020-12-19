const Pool = require('pg').Pool

const pool = new Pool({
    host: 'app-85ee1782-8a3e-4e03-a5e7-04361eaa98d0-do-user-8435387-0.b.db.ondigitalocean.com',
    user: 'db',
    password: 'xt9hhj9lav1nlehv',
    port: '25060',
    database: 'db',
    ssl: {
        rejectUnauthorized: false
    }
})

module.exports = pool
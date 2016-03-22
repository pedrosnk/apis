'use strict'

let express = require('express')
let app = express();

app.get('/healthcheck', (req, res) => {
  res.send({status: 'ok'})
})

module.exports = {
  start() {
    app.listen(8888, () => console.log('listening to port 8888') )
  }
}

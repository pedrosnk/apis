'use strict'

var express = require('express')
var bodyParser = require('body-parser')
var app = express()
var customcode = require('../cc_runner/customcode')
var registry = require('./registry')

app.use(bodyParser.json())
app.post('/cc-run', (req, res) => {
  let body = req.body
  let response = customcode(body.code, body.context)
  res.send({response: response})
})

app.post('/cc-registry', (req, res) => {
  let body = req.body
  let response = registry(body)
  res.send({response: response})
})

app.listen(5000, () => {
  console.log('listening to port 5000')
})


var customcode = require('./customcode')
var args = process.argv.slice(2)

console.info(customcode(args[0], args[1]))

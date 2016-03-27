
var vm = require('vm')

var args = process.argv.slice(2)
var scriptCode = `
  response = ${args[0]}
`
var script = new vm.Script(scriptCode)
var sandbox
if (args[1] !== undefined) {
  sandbox = JSON.parse(args[1])
} else {
  sandbox = {}
}

sandbox.response = ''
var context = vm.createContext(sandbox)
script.runInContext(context)

console.info(sandbox.response)

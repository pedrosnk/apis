
var vm = require('vm')

module.exports = (code, context) => {
  var scriptCode = `
    response = ${code}
  `
  var script = new vm.Script(scriptCode)
  var sandbox
  if (typeof context == "object") {
    sandbox = context 
  }else if (typeof context == "string") {
    sandbox = JSON.parse(context)
  } else {
    sandbox = {}
  }

  sandbox.response = ''
  var context = vm.createContext(sandbox)
  script.runInContext(context)

  return context.response
}

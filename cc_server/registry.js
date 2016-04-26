var vm = require('vm')
var leftPadScript = `response = function() {
		var leftpad = (str, len) => {
			str = String(str)
			var i = -1
			len = len - str.length
			while (++i < len) {
				str = '#' + str
			}
			return str
		}
		return leftpad(message, 10)
}()`

var leftPadCompiled = new vm.Script(leftPadScript)

module.exports = function(context){
  let sandbox
  if(typeof context === 'object') {
    sandbox = context
  } else if (typeof context === 'string') {
    sandbox = JSON.parse(context)
  } else {
    sandbox = {}
  }

  sandbox.response = ''
  var ctx = vm.createContext(sandbox)
  script.runInContext(ctx)

  return ctx.response
}

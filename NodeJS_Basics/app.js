// get the public objects from 4-names.js 
const names = require("./4-names");
const sayHi = require("./5-utils");
const flavor = require("./6-alternatives-flavor")


sayHi(", This code is imported from another module")
sayHi(`${names.dance}` + " <= this also too")
sayHi(flavor.first_var)
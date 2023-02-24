// export 1 object at a time
module.exports.first_var = "1st"
module.exports.second_var = {2:"2nd", 3:"3rd"}

const num1 = 5
const num2 = 10

function addValues(){
    console.log(`the sum is ${num1 + num2}`)
}

addValues()
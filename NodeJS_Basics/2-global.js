#!/home/new-user/.node/bin/node
console.log(__dirname)
console.log("Running infinite code")
setInterval( ()=>{
    console.log("Hello")
}, 1000) // run infinite code, between a period time
// in this case, It's 1 second (1000 miliseconds) 
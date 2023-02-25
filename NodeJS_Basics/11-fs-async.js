// dealing with file system: Fs
// asynchronously: nonblocking <- prefered
// synchronously: blocking

const {readFile, writeFile} = require("fs")

readFile("./content/first.txt","utf8",(err,result)=>{
    // (err,result)=>{} is a callback
    // callback is the function that excute at the end of another synchronous function
    // the synchronous function will always return 2 options error(err) and result(result)
    if (err){
        console.log(`Error while Reading file first.txt: ${err}`)
        return
    }
    const first = result
    readFile("./content/second.txt","utf8",(err, result)=>{
        if (err){
            console.log(`Error while Reading file first.txt: ${err}`)
            return
        }
        const second = result
        writeFile(
            "./content/result-sync.txt",
            `\nSynchronous result:\n${first}\n${second}`,
            {flag: "a"},
            (err,result)=>{
                if (err){
                    console.log(`erro ${err}`)
                    return
                }
                console.log(result)
            })
    })
})
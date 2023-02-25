// dealing with file system: Fs
// asynchronously: nonblocking <- prefered
// or synchronously: blocking

const {readFileSync, writeFileSync} = require("fs")

const first = readFileSync("./content/first.txt","utf8")
const second = readFileSync("./content/second.txt","utf8")

var test_write_path = './content/result-sync.txt'

writeFileSync(
    test_write_path,
    `\nHere is the content of 2 files:\n${first}\n${second}`,
    {flag: "a"} // append text
)

const test_written_file = readFileSync(test_write_path, "utf8") // utf8 is default
console.log(test_written_file)
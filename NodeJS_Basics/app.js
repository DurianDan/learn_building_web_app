const path = require("path")

var os_path_seperator = path.sep

const normalize_by_os_filePath = path.join("/content///","test_folder","test_text.txt")
console.log(normalize_by_os_filePath) // return directory
console.log(path.basename(normalize_by_os_filePath)) // lskdjlf
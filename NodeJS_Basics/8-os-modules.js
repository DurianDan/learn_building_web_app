const os_module = require("os")

// get cpus info 
var cpus_count = os_module.cpus().length
var cpus_info = os_module.cpus()

// get info of the current desktop user 
var user_info = os_module.userInfo()

// get uptime of the syste...
// the time this computer has been turned on
// in seconds
var live_time = os_module.uptime()

// get system infos

const systemInfo = {
    name: os_module.type(), // get platform name like: windows, linux, mac...
    release: os_module.release(), // get kernel version
    totalMem: os_module.totalmem(),
    freeMem: os_module.freemem()
}

console.log(systemInfo)
const http =  require("http")

const server = http.createServer((request, response)=>{
    // this object is the endpoint of an app
    // this is where the website communicate with the world!!!

    // when the client requesting a path in the url,
    // throught the listening port (5000)
    if (request.url == "/"){
        // has to define clear/unoverlapped conditions...
        // ..for each response
        // because, at some condition, response has been ended
        // and in other met condition, the response is still called
        // creating errors 
        response.end("your are at the front page\n")
    }else if(request.url == "/about"){
        response.end("This is an about page")
    }else{
        response.end(`
        <h1>Opps!</h1>
        <p>
        This is the default message,
        when the page you are requesting does not exist
        </p>
        <a href="/">home page</a>
        `)
    }
})

server.listen(5000) // tell the server to listen to port 5000
// get to the url localhost:5000
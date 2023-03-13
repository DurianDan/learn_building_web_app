package utils 

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, X interface{}){
	if body,err := io.ReadAll(r.Body); err == nil{
		// try to read the request body
		// and parse it into the variable "body"
		if err := json.Unmarshal(([]byte(body)), X); err != nil{
			// read the JSON data from the variable "body"
			// turn it into GO data structure
			// pass it into the interface X
			return	
		}
	}
}
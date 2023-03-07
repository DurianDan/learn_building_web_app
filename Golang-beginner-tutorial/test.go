package main

import "fmt"

func main(){
	var testmap = []map[string]string{
		{"name":"Huy", "age":"23"},
		{"name":"Lan", "age":"25"},
		{"name":"Hu", "age":"27"},
	}

	for _,map_ele := range testmap {
		for k,v := range map_ele {
			fmt.Printf("key is %v (type %T)\n",k,k)
			fmt.Printf("value is %v (type %T)\n",v,v)
		}
		fmt.Println("_________________________")
	}
}
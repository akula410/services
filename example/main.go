package main

import (
	"fmt"
	"github.com/akula410/services"
)
func main() {
	var get = services.Get("https://dokamir.com.ua/")
	get.
		Param("param_1", "value_1").
		Param("param_2", "value_2").
		Header("header_1", "value_1").
		Send()

	fmt.Println(string(get.GetResponseBody()))
	fmt.Println(get.GetHeader("Content-Type"))
	fmt.Println(get.GetHeaders())
	fmt.Println(get.GetCookie("language"))
}

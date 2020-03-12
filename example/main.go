package main

import "services"
func main() {
	var get = services.Get("http://request.com")
	get.
		Param("param_1", "value_1").
		Param("param_2", "value_2").
		Header("header_1", "value_1").
		Send()

	var get2 = services.Get("http://request.com")
	get2.Send()
}

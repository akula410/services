package services

type get struct {
	abstract
}

func Get(url string)*get{
	var obj = &get{}
	obj.setUrl(url)
	obj.setMethod("GET")
	return obj
}



package services

type post struct {
	abstract
}

func Post(url string)*post{
	var obj = &post{}
	obj.setUrl(url)
	obj.setMethod("POST")
	return obj
}
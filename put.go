package services

type put struct {
	abstract
}

func Put(url string)*put{
	var obj = &put{}
	obj.setUrl(url)
	obj.setMethod("PUT")
	return obj
}
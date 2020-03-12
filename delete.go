package services
type _delete struct {
	abstract
}

func Delete(url string)*_delete{
	var obj = &_delete{}
	obj.setUrl(url)
	obj.setMethod("DELETE")
	return obj
}
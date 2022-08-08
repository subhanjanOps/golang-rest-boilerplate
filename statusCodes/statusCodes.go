package statuscodes

type StatusCode struct {
	Code    int32
	Message string
}

type ResponseBody struct {
	Status StatusCode
	Data   map[string]interface{}
}

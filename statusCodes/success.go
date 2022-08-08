package statuscodes

var SuccessCodes = map[string]*StatusCode{
	"200": {
		Code:    200,
		Message: "Successful",
	},
	"201": {
		Code:    201,
		Message: "Created",
	},
}

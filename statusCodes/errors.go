package statuscodes

var ErrorCodes = map[string]*StatusCode{
	"401": {
		Code:    401,
		Message: "Unauthorized",
	},
	"400": {
		Code:    400,
		Message: "Bad Request",
	},
	"403": {
		Code:    403,
		Message: "Forbidden",
	},
	"500": {
		Code:    500,
		Message: "Internal server Error",
	},
}

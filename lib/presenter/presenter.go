package presenter

const (
	RESPONSE_SUCCESS_GET      = "Success Get Data"
	RESPONSE_SUCCESS_GET_LIST = "Success Get List Data"
	RESPONSE_SUCCESS_CREATE   = "Success Create Data"
	RESPONSE_SUCCESS_UPDATE   = "Success Update Data"
	RESPONSE_SUCCESS_DELETE   = "Success Delete Data"
	RESPONSE_SUCCESS_VOID     = "Success Void Data"

	//make sure to add param name before this constant
	RESPONSE_ERROR_INVALID_PARAM_STRING = "can't be empty"
	//make sure to add param name before this constant
	RESPONSE_ERROR_INVALID_PARAM_NUMBER = "should be a number"
	//make sure to add param name before this constant
	RESPONSE_ERROR_INVALID_PARAM_NUMBER_NOT_EMPTY = "should be a number and can't be empty"
)

//used to present default response
type Default struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` //won't show if data is empty / null
}

//used to present response with metadata
type List struct {
	Message string      `json:"msg"`
	Meta    Meta        `json:"meta"`
	Data    interface{} `json:"data"`
}

type Meta struct {
	TotalData int `json:"total_data"` // total data without pagination
	TotalPage int `json:"total_page"` // total page based on per_page / limit
	Page      int `json:"page"`       // current page
	PerPage   int `json:"per_page"`   // total data per_page
}

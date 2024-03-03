package constant

import "fmt"

const (
	// StatusSuccessText for
	StatusSuccessText = "success"
	// StatusCreatedText for
	StatusCreatedText = "created"
	// StatusConflictText for
	StatusConflictText = "conflict"
	// StatusInternalServerErrorText for
	StatusInternalServerErrorText = "internal_server_error"
	// StatusBadRequestText for
	StatusBadRequestText = "bad_request"
	// StatusNotFoundText for
	StatusNotFoundText = "not_found"
	// StatusUnprocessableEntityText for
	StatusUnprocessableEntityText = "unprocessable_entity"
	// StatusUnauthorized for
	StatusUnauthorized = "unauthorized"
	// StatusForbidden for
	StatusForbidden = "forbidden"

	StatusInvalidParam = "invalid_parameter_value"

	StatusInvalidParamFormat = "invalid_parameter_format"

	StatusInvalidDateRange = "invalid_date_range_parameter"

	StatusInvalidPagination = "invalid_pagination_parameter"

	StatusInvalidToken = "invalid_token"

	StatusTokenExpired = "token_expired"

	StatusUserDisabled = "user_disable"

	StatusAccessDenied = "access_denied"
)

const (
	MessageSuccessInsert = "success insert auto text"

	MessageSuccessUpdate = "success update auto text"

	MessageSuccessDelete = "success delete auto text"

	MessageSuccess = "success get auto texts"

	MessageNotFound = "auto text not found"

	MessagConflict = "name already exist"

	MessageErrorSort = "sort parameter is required"
)

var (
	ErrNoData   = fmt.Errorf("no data")
	ErrNoResult = fmt.Errorf("no result")
	ErrConflict = fmt.Errorf("name already exist")
	ErrSort     = fmt.Errorf("sort parameter is required")
)

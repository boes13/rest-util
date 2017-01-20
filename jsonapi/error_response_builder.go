package jsonapi

import (
	"errors"
	"reflect"
	"strconv"
	"net/http"
)

type ErrorResponse struct {
	// Information of jsonapi specification used.
	Jsonapi JsonapiVersion `json:"jsonapi"`

	// List of JsonError encountered during endpoint execution.
	Errors []JsonError `json:"errors"`
}

type JsonError struct {
	// A unique identifier for this particular occurrence of the problem.
	Id string `json:"id"`

	Links []ErrorLink `json:"links"`

	// The HTTP status code applicable to this problem, expressed as a string value.
	Status string `json:"status"`

	// An application-specific error code, expressed as a string value.
	Code string `json:"code"`

	// A short, human-readable summary of the problem that SHOULD NOT change
	// from occurrence to occurrence of the problem,
	// except for purposes of localization.
	Title string `json:"title"`

	// A human-readable explanation specific to this occurrence of the problem.
	// Like title, this field’s value can be localized.
	Detail string `json:"detail"`

	// An object containing information about the source of the error.
	Source ErrorSource `json:"source"`

	// A meta object containing non-standard meta-information about the error.
	meta interface{} `json:"meta"`
}

type ErrorLink struct {
	// A link that leads to further details about this particular occurrence of the problem.
	About string `json:"about"`
}

type ErrorSource struct {
	// A JSON Pointer [RFC6901] to the associated entity in the request document.
	Pointer string `json:"pointer"`

	// A string indicating which URI query parameter caused the error.
	Parameter string `json:"parameter"`
}

// CreateErrorResponse create new object of ErrorResponse type with zero JsonError.
// It returns ErrorResponse object reference.
func CreateErrorResponse() *ErrorResponse {
	var errResponse = ErrorResponse{}
	errResponse.Errors = make([]JsonError, 0)
	errResponse.Jsonapi = JsonapiVersion{Version: "1.0"}
	return &errResponse
}

// AddError inserts new JsonError to ErrorResponse object.
// Parameters: id, links, status, code, title, detail, source, meta (refer to type documentation for more info).
// Parameter meta, non-standard information, has to be of type struct.
// It returns nil if the operation successful, or else it will return error.
func (resp *ErrorResponse) AddError(id string, links []ErrorLink, status string, code string, title string, detail string, source *ErrorSource, meta interface{}) error {
	jsonError := JsonError{
		Id:     id,
		Links:  links,
		Status: status,
		Code:   code,
		Title:  title,
		Detail: detail,
		Source: *source,
	}

	if meta != nil {
		err := jsonError.addMeta(meta)
		if err != nil {
			return err
		}
	}

	resp.Errors = append(resp.Errors, jsonError)
	return nil
}

// CreateErrorLinks create slice of ErrorLink based on input urls.
// It return slice of ErrorLink
func CreateErrorLinks(urls []string) []ErrorLink {
	errorLinks := make([]ErrorLink, 0)
	for _, url := range urls {
		link := ErrorLink{url}
		errorLinks = append(errorLinks, link)
	}
	return errorLinks
}

// CreateErrorSource creates new ErrorSource object.
// It returns reference to ErrorSource object.
func CreateErrorSource(pointer string, parameter string) *ErrorSource {
	errSource := ErrorSource{pointer, parameter}
	return &errSource
}

// Add meta, non-standard information, to JsonError object.
// It requires argument of type struct or else it will return error.
// It returns nil if successful.
func (err *JsonError) addMeta(meta interface{}) error {
	kind := reflect.ValueOf(meta).Kind()
	if kind != reflect.Struct {
		return errors.New("Argument meta should be of type struct!")
	}
	err.meta = meta
	return nil
}

// Create simple http error response with only JsonapiVersion, ErrorSource, http status & title, and detail error.
// Returns ErrorResponse object reference.
func CreateSimpleHttpErrorResponse(errorUrl string, errorParameter string, httpError int, detailError string) *ErrorResponse {
	errResponse := CreateErrorResponse()
	errSource := CreateErrorSource(errorUrl, errorParameter)
	errResponse.AddError("", nil, strconv.Itoa(httpError), "", http.StatusText(httpError), detailError, errSource, nil)
	return errResponse
}

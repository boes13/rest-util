// Package jsonapi provides utility to build json response based on jsonapi.org spec 1.0 easily.
// It does not fully comply to the spec yet.
package jsonapi

import (
	"errors"
	"reflect"
)

type DataResponse struct {
	// Information of jsonapi specification used.
	Jsonapi JsonapiVersion `json:"jsonapi"`

	// A meta object containing non-standard meta-information about the response data.
	Meta interface{} `json:"meta"`

	// Document's primary data. It can be a single data or collection of data.
	Data interface{} `json:"data"`

	// A links object related to the primary data.
	Links ResponseLinks `json:"links"`

	// An array of resource objects that are related to the primary data and/or each other.
	//Included []interface{} `json:"included"`
}

type JsonapiVersion struct {
	// Version of jsonapi specification used.
	Version string `json:"version"`
}

type ResponseLinks struct {
	// The link that generated the current response document.
	Self string `json:"self"`

	// A related resource link when the primary data represents a resource relationship.
	Related string `json:"related"`

	// The first page of data.
	First string `json:"first"`

	// The last page of data
	Last string `json:"last"`

	// The previous page of data.
	Prev string `json:"prev"`

	// The next page of data.
	Next string `json:"next"`
}

// CreateDataResponse creates empty data response.
// It returns DataResponse object reference.
func CreateDataResponse() *DataResponse {
	var dataResponse = DataResponse{}
	dataResponse.Jsonapi = JsonapiVersion{Version: "1.0"}
	return &dataResponse
}

// SetMeta sets the meta, non-standard information, to DataResponse object.
// It requires argument of type struct or else it will return error.
// It returns nil if successful.
func (resp *DataResponse) SetMeta(meta interface{}) error {
	kind := reflect.ValueOf(meta).Kind()
	if kind != reflect.Struct {
		return errors.New("Argument meta should be of type struct!")
	}
	resp.Meta = meta
	return nil
}

// SetData sets the primary data of DataResponse object.
// The parameter data should either be struct, map, or array/slice of struct.
// It returns error when the parameter type is invalid or else nil.
func (resp *DataResponse) SetData(data interface{}) error {
	kind := reflect.ValueOf(data).Kind()
	errMessage := "Argument data should be of type struct or array/slice of struct!"

	switch {
	case kind == reflect.Struct:
		resp.Data = data
	case kind == reflect.Map:
		resp.Data = data
	case kind == reflect.Array || kind == reflect.Slice:
		value := reflect.ValueOf(data)
		if value.Len() > 0 {
			if value.Index(0).Kind() == reflect.Struct {
				resp.Data = data
			} else {
				return errors.New(errMessage)
			}
		} else {
			resp.Data = data
		}
	default:
		return errors.New(errMessage)
	}
	return nil
}

// SetLinks sets ResponseLinks of object DataResponse.
func (resp *DataResponse) SetLinks(self string, related string, first string, last string, prev string, next string){
	resp.Links.Self = self
	resp.Links.Related = related
	resp.Links.First = first
	resp.Links.Last = last
	resp.Links.Prev = prev
	resp.Links.Next = next
}

// SetIncluded sets Included field of object DataResponse.
// Included is an array of resource objects (struct) that are related to the primary data and/or each other.
// It returns error if type of any element of parameter included is not struct.
// If no error, it returns nil.
//func (resp *DataResponse) SetIncluded(included []interface{}) error {
//	errMessage := "Parameter included should be of type slice of struct!"
//	// parameter validation
//	for _, inc := range included {
//		kind := reflect.ValueOf(inc).Kind()
//		if kind != reflect.Struct {
//			return errors.New(errMessage)
//		}
//	}
//	resp.Included = included
//	return nil
//}
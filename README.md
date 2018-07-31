# rest-util
## package
* http: validate string parameter value to integer and date, and returns its value based on respective type.
* jsonapi: build json response (success or error) based on jsonapi specification

## documentation
We use standard godoc as our code documentation tool. To view it, please follow these steps :
1. Open your terminal, head to this cloned repo (boes13/rest-util)
2. run `godoc -http=:6060` (this will trigger godoc at port 6060)
3. Open your browser, and hit `http://127.0.0.1:6060/pkg/github.com/boes13/rest-util/`

## usage
### http
```
func UserCheckHandler(w http.ResponseWriter, r *http.Request) {
  // parse parameter from request
  uId := r.FormValue("user_id")

  // validate parameter value
  userId, err := http.ValidatePositiveNumber(uId, false)
  if err != nil {
    log.Fatalf("Error validating user_id, error: %s\n", err.Error())
  }
  ....
}
```
### json
```
type meta struct {
	id string
	name string
  ....
}
type data struct {
	id      string
	name    string
	address string
	phone   string
  ....
}

func UserCheckHandler(w http.ResponseWriter, r *http.Request) {
  ....
  // error response
  if err != nil {
    errResp := CreateErrorResponse()

    // setup parameters to add JsonError object to ErrorResponse
    urls := []string {
      "http://host1.com/some_end_point",
      "http://host2.com/another_end_point"
    }
	  links := CreateErrorLinks(urls)
    errSource := CreateErrorSource(r.URL.Path, "user_id")
    err2 := errResp.AddError("some_id", links, "some_status", "some_code", "some_title", "some_detail", errSource, any{id:"meta_id", name:"meta_name",})
	  if err2 != nil {
      log.Fatalln("Found error:", err2)
    }

    // you may add other errors by repeating from // setup parameters to add JsonError object to ErrorResponse
    ....
    json.NewEncoder(w).Encode(errResp)
    return
  } else {
    ....
    // success response
    dataResponse := CreateDataResponse()
    err := dataResponse.SetMeta(meta{id:"some_id", name:"some_name",})
    if err != nil {
      log.Fatalln("Error adding meta to DataResponse:", err)
    }

    // when response data is single
    data1 := data{id:"my_id",address:"my_address",name:"my_name",phone:"my_phone",}
    err = dataResponse.SetData(data1)
    if err != nil {
      log.Fatalln("Error setting data:", err)
    }

    // when response data is array or slice
    dataSlice := make([]data, 0)
    dataSlice = append(dataSlice, data1)
    err = dataResponse.SetData(dataSlice)
    if err != nil {
      log.Fatalln("Error setting data:", err)
    }
    json.NewEncoder(w).Encode(dataResponse)
    return
  }
}
```

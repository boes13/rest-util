# rest-util
## package
* http: validate string parameter value to integer and date, and returns its value based on respective type.
* jsonapi: build json response (success or error) based on jsonapi specification

## documentation
refer to each function documentation.

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

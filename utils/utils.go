package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
Function name:
"CheckEnvironmentVariable".

Description:
The main task of the function is to check the value of the environment variable.
*/
func CheckEnvironmentVariable(key string) string{
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("The \"%s\" environment variable does not exist.", key)
	}
	return value
}

/*
Function name:
"Response".

Function description:
The main task of the function is to create and configure a response with JSON.
*/
func Response(responseWriter http.ResponseWriter, statusCode int, information interface{}) {
	response, err := json.Marshal(information)
	if err != nil {
		log.Println(err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, _ = responseWriter.Write([]byte(err.Error()))
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	_, _ = responseWriter.Write([]byte(response))
}

/*
Function name:
"ResponseWithError".

Function description:
The main task of the function is to create a response with detailed information about the error.
*/
func ResponseWithError(responseWriter http.ResponseWriter, responseStatus int, errorMessage string) {
	Response(responseWriter, responseStatus, map[string]string{"STATUS": "ERROR", "DESCRIPTION": errorMessage})
}

/*
Function name:
"ResponseWithSuccess".

Function description:
The main task of the function is to create a response with success message.
*/
func ResponseWithSuccess(responseWriter http.ResponseWriter, responseStatus int, successMessage string) {
	Response(responseWriter, responseStatus, map[string]string{"STATUS": "SUCCESS", "DESCRIPTION": successMessage})
}

/*
Function name:
"ConvertIntArrayToString".

Function description:
The main task of the function is to convert int array to string separated by ','.
*/
func ConvertIntArrayToString(input []int) string {
	if len(input) == 0 {
		return ""
	}
	estimate := len(input) * 4
	b := make([]byte, 0, estimate)
	for _, n := range input {
		b = strconv.AppendInt(b, int64(n), 10)
		b = append(b, ',')
	}
	b = b[:len(b)-1]
	return string(b)
}

/*
Function name:
"ConvertStringArrayToString".

Function description:
The main task of the function is to convert string array to string separated by ','.
*/
func ConvertStringArrayToString(input []string) string {
	var output strings.Builder
	for i := 0; i < len(input); i++ {
		output.WriteString("'")
		output.WriteString(input[i])
		output.WriteString("'")
		if i < len(input) - 1 {
			output.WriteString(",")
		}
	}
	return output.String()
}

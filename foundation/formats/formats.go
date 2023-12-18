package formats

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

var Format string

func GetFormat(r *http.Request) {

	Format = r.URL.Query()["format"][0]

}

func SetFormat(data interface{}) []byte {

	var apiOutput []byte
	Format = "json"
	if Format == "json" {
		output, _ := json.Marshal(data)
		apiOutput = output
	} else if Format == "xml" {
		output, _ := xml.Marshal(data)
		apiOutput = output
	}
	return apiOutput
}

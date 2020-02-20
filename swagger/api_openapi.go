/*
 * AnyPay
 *
 * This the AnyPay service targeted towards, parents with children doing payments and russian oligarks
 *
 * API version: 1.0.0
 * Contact: pr.von.rosen@swedbank.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/* Openapi_yaml - here we can retrieve the openapi3.0 aka "swagger file" for auto-provision*/
func Openapi_yaml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/yaml; charset=UTF-8")
	file, err := os.Open("/openapi.yaml")
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not find openapi file")
		return
	}

	filecontent, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not read openapi file after loading it")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(filecontent)

}

/* Openapi_json - here we can retrieve the openapi3.0 aka "swagger file" for auto-provision*/
func Openapi_json(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	file, err := os.Open("/openapi.json")
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not find openapi file")
		return
	}

	filecontent, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not read openapi file after loading it")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(filecontent)

}

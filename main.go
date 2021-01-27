package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prime-factors/model"
	"github.com/prime-factors/service"
	"github.com/prime-factors/utility/request"
	"github.com/prime-factors/utility/response"
)

var BIG_NUMBERS = []int{4231123, 2341231, 4121231, 2311237, 1312333}
var MEDIUM_NUMBERS = []int{31123, 41231, 21231, 11237, 12333}
var SMALL_NUMBERS = []int{224, 133, 247, 363, 609}

func main() {
	router := httprouter.New()

	router.GET("/", IndexHandler)
	router.POST("/prime-factors", PrimeFactorsHandler)
	router.ServeFiles("/static/*filepath", http.Dir("app"))

	log.Println("Listen at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func PrimeFactorsHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	form, err := getRequestParam(r)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	var p model.PrimesChecking
	if form.IsGoroutine {
		p = service.PrimesCheckWithGoroutine(getNumbers(form))
	} else {
		p = service.PrimesCheckWithoutGoroutine(getNumbers(form))
	}

	response.WriteSuccess(w, p, "ok")
}

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	t, _ := template.ParseFiles("app/index.html")
	t.Execute(w, nil)
}

func getRequestParam(r *http.Request) (request.Form, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return request.Form{}, err
	}

	var form request.Form
	err = json.Unmarshal(b, &form)
	if err != nil {
		return request.Form{}, err
	}

	err = form.Validate()
	if err != nil {
		return request.Form{}, err
	}

	return form, nil
}

func getNumbers(form request.Form) []int {
	if form.ExampleType == "small" {
		return SMALL_NUMBERS
	} else if form.ExampleType == "medium" {
		return MEDIUM_NUMBERS
	} else if form.ExampleType == "big" {
		return BIG_NUMBERS
	} else {
		return form.Numbers
	}
}

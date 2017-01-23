package app

import (
	"net/http"
	"strings"
)

func SetUp() {
	runHTTP()
}

func runHTTP() {
	port := ":1234"
	// listen := fmt.Sprintf("0.0.0.0:%d", port)
	// http.Handle("/", router.NewMux())
	// http.Handle("/ro", handler http.Handler)
	http.HandleFunc("/v1/members", memberController)
	http.ListenAndServe(port, nil)
}

func memberController(w http.ResponseWriter, r *http.Request) {
	// some how, get the members from Miranda, save it as a go model and save it.
	// make a membersController for this
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	println(city)
}

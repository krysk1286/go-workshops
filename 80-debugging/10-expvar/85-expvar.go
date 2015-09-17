// Demo of the expvar package. You register metrics by creating NewT, then
// update it.

// You can access the exposed metrics via HTTP at /debug/vars, you'll get a JSON
// object with your exposed variables and some pre defined system ones.

// You can use monitoring system such as Nagios and OpenNMS to monitor the
// system and plot the change of data over time.

// po uruchomieniu wywołaj "curl http://localhost:8080?user=placel"
// "curl http://localhost:8080/debug/vars | python -m json.tool".

// dodajmy interfejs do tych danych :)
//
package main

import (
	"expvar"
	"fmt"
	"io"
	"net/http"
)

// Two metrics, these are exposed by "magic" :)
// Number of calls to our server.
var numCalls = expvar.NewInt("num_calls")

// Last user.
var lastUser = expvar.NewString("last_user")

func HelloServer(w http.ResponseWriter, req *http.Request) {
	user := req.FormValue("user")

	// Update metrics
	numCalls.Add(1)
	lastUser.Set(user)

	msg := fmt.Sprintf("G'day %s\n", user)
	io.WriteString(w, msg)
}

func main() {
	http.HandleFunc("/", HelloServer)
	panic(http.ListenAndServe(":8080", nil))
}

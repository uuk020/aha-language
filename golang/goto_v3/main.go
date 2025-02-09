package main

import (
	"flag"
	"fmt"
	"net/http"
)

const AddForm = `
<html>
<body>
	<form method="POST" action="/add">
		URL: <input type="text" name="url">
		<input type="submit" value="Add">
	</form>
<body>
</html>
`

var (
	listenAddr = flag.String("http", ":8080", "http listen address")
	dataFile = flag.String("file", "store.json", "data store file name")
	hostname = flag.String("host", "localhost:8080", "host name and port")
)

var store *URLStore2

func main()  {
	flag.Parse()
	store = NewURLStore2(*dataFile)
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(*listenAddr, nil)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := store.Get2(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		fmt.Fprint(w, AddForm)
		return
	}
	key := store.Put2(url)
	fmt.Fprintf(w, "http://%s/%s", *hostname, key)
}

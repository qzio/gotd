package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

var epoch = time.Unix(0, 0).Format(time.RFC1123)

var noCacheHeaders = map[string]string{
	"Expires":         epoch,
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

func main() {

	var host, port string
	flag.StringVar(&port, "p", "3000", "port to listen to, default to 3000")
	flag.StringVar(&host, "h", "", "host to listen to, default to nil (0.0.0.0)")
	flag.Parse()

	fmt.Printf("start %s\n", epoch)

	fs := http.FileServer(http.Dir("./"))
	i := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range noCacheHeaders {
			w.Header().Add(k, v)
		}

		fmt.Printf("%d %s %s\n", i, r.Method, r.RequestURI)
		i++

		for k, v := range r.Header {
			fmt.Printf("%s:%s\n", k, v)
		}
		b, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("err reading body: %v\n", err)
		} else {
			fmt.Printf("body: \n%s\n", string(b))
		}
		fs.ServeHTTP(w, r)
	})

	fmt.Printf("listning on %s:%s", host, port)
	if e := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil); e != nil {
		panic(e)
	}
}

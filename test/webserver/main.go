package main

import (
	"fmt"
	"github.com/stretchrcom/goweb"
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/signature"
	"github.com/stretchrcom/tracer"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func absoluteUrlForRequest(request *http.Request) string {
	return fmt.Sprintf("http://%s%s", request.Host, request.RequestURI)
}

func main() {

	// handle every path
	goweb.Map("***", func(c context.Context) error {

		method := c.HttpRequest().Method
		requestUrl := absoluteUrlForRequest(c.HttpRequest())
		privateKey := "PRIVATE"
		body, bodyErr := ioutil.ReadAll(c.HttpRequest().Body)

		fmt.Printf("URL: %s", requestUrl)

		if bodyErr != nil {
			goweb.Respond.With(c, 500, []byte(fmt.Sprintf("Balls! Couldn't read the body because %s", bodyErr)))
			return nil
		}

		t := tracer.New(tracer.LevelEverything)

		valid, _ := signature.ValidateSignatureWithTrace(method, requestUrl, string(body), privateKey, t)

		// return the tracing
		c.HttpResponseWriter().Header().Set("Content Type", "text/html")
		c.HttpResponseWriter().Write([]byte("<html>"))
		c.HttpResponseWriter().Write([]byte("<head>"))
		c.HttpResponseWriter().Write([]byte("<title>Signature</title>"))

		c.HttpResponseWriter().Write([]byte("<link href='http://reset5.googlecode.com/hg/reset.min.css' rel='stylesheet' />"))

		c.HttpResponseWriter().Write([]byte("<style>"))
		c.HttpResponseWriter().Write([]byte(".page{padding:10px}"))
		c.HttpResponseWriter().Write([]byte(".status{padding:10px}"))

		if valid {
			c.HttpResponseWriter().Write([]byte(".status{background-color:green}"))
		} else {
			c.HttpResponseWriter().Write([]byte(".status{background-color:red}"))
		}
		c.HttpResponseWriter().Write([]byte("</style>"))

		c.HttpResponseWriter().Write([]byte("</head>"))
		c.HttpResponseWriter().Write([]byte("<body>"))
		c.HttpResponseWriter().Write([]byte("<div class='page'>"))
		c.HttpResponseWriter().Write([]byte("<div class='status'></div>"))
		c.HttpResponseWriter().Write([]byte("<pre>"))

		for _, trace := range t.Data() {
			c.HttpResponseWriter().Write([]byte(trace.Data))
			c.HttpResponseWriter().Write([]byte("\n"))
		}

		c.HttpResponseWriter().Write([]byte("</pre>"))
		c.HttpResponseWriter().Write([]byte("</div>"))

		c.HttpResponseWriter().Write([]byte("</body>"))
		c.HttpResponseWriter().Write([]byte("</html>"))

		return nil

	})

	fmt.Println("Signature test webserver")
	fmt.Println("by Mat Ryer and Tyler Bunnell")
	fmt.Println(" ")
	fmt.Println("Start making signed request to: http://localhost:8080/")

	/*

	   START OF WEB SERVER CODE

	*/
	Address := ":8080"

	log.Print("Goweb 2")
	log.Print("by Mat Ryer and Tyler Bunnell")
	log.Print(" ")
	log.Print("Starting Goweb powered server...")

	// make a http server using the goweb.DefaultHttpHandler()
	s := &http.Server{
		Addr:           "Address",
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	listener, listenErr := net.Listen("tcp", Address)

	log.Printf("  visit: %s", Address)

	if listenErr != nil {
		log.Fatalf("Could not listen: %s", listenErr)
	}

	log.Println("Also try some of these routes:")
	log.Printf("%s", goweb.DefaultHttpHandler())

	go func() {
		for _ = range c {

			// sig is a ^C, handle it

			// stop the HTTP server
			log.Print("Stopping the server...")
			listener.Close()

			/*
			   Tidy up and tear down
			*/
			log.Print("Tearing down...")

			// TODO: tidy code up here

			log.Fatal("Finished - bye bye.  ;-)")

		}
	}()

	// begin the server
	log.Fatalf("Error in Serve: %s", s.Serve(listener))

	/*

	   END OF WEB SERVER CODE

	*/

}

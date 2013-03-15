package main

import (
	"fmt"
	"github.com/stretchrcom/goweb/goweb"
	"github.com/stretchrcom/signature"
	"github.com/stretchrcom/tracer"
	"io/ioutil"
)

func main() {

	// handle every path
	goweb.MapFunc("/", func(c *goweb.Context) {

		method := c.Request.Method
		requestUrl := c.Request.URL.String()
		privateKey := "PRIVATE"
		body, bodyErr := ioutil.ReadAll(c.Request.Body)

		if bodyErr != nil {
			c.RespondWithErrorMessage(fmt.Sprintf("Balls! Couldn't read the body because %s", bodyErr), 500)
			return
		}

		t := tracer.New(tracer.LevelEverything)

		valid, _ := signature.ValidateSignatureWithTrace(method, requestUrl, string(body), privateKey, t)

		// return the tracing
		c.ResponseWriter.Header().Set("Content Type", "text/html")
		c.ResponseWriter.Write([]byte("<html>"))
		c.ResponseWriter.Write([]byte("<head>"))
		c.ResponseWriter.Write([]byte("<title>Signature</title>"))

		c.ResponseWriter.Write([]byte("<link href='http://reset5.googlecode.com/hg/reset.min.css' rel='stylesheet' />"))

		c.ResponseWriter.Write([]byte("<style>"))
		c.ResponseWriter.Write([]byte(".page{padding:10px}"))
		c.ResponseWriter.Write([]byte(".status{padding:10px}"))

		if valid {
			c.ResponseWriter.Write([]byte(".status{background-color:green}"))
		} else {
			c.ResponseWriter.Write([]byte(".status{background-color:red}"))
		}
		c.ResponseWriter.Write([]byte("</style>"))

		c.ResponseWriter.Write([]byte("</head>"))
		c.ResponseWriter.Write([]byte("<body>"))
		c.ResponseWriter.Write([]byte("<div class='page'>"))
		c.ResponseWriter.Write([]byte("<div class='status'></div>"))
		c.ResponseWriter.Write([]byte("<pre>"))

		for _, trace := range t.Data() {
			c.ResponseWriter.Write([]byte(trace.Data))
			c.ResponseWriter.Write([]byte("\n"))
		}

		c.ResponseWriter.Write([]byte("</pre>"))
		c.ResponseWriter.Write([]byte("</div>"))

		c.ResponseWriter.Write([]byte("</body>"))
		c.ResponseWriter.Write([]byte("</html>"))

	})

	fmt.Println("Signature test webserver")
	fmt.Println("by Mat Ryer and Tyler Bunnell")
	fmt.Println(" ")
	fmt.Println("Start making signed request to: http://localhost:8080/")

	// start the server
	goweb.ConfigureDefaultFormatters()
	goweb.ListenAndServe(":8080")

}

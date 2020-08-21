package main

import (
	"fmt"
	"time"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	router := routing.New()

	router.Get("/", getHangler)

	panic(fasthttp.ListenAndServe(":8080", router.HandleRequest))
}

func getHangler(c *routing.Context) error {
	start := time.Now()
	fmt.Fprintf(c, "Hello, world!")
	duration := time.Since(start)
	fmt.Println(duration)
	return nil
}

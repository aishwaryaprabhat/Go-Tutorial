package main

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	router := routing.New()

	router.Get("/", getHangler)

	panic(fasthttp.ListenAndServe(":8080", router.HandleRequest))
}

func getHangler(c *routing.Context) error {
	fmt.Fprintf(c, "Hello, world!")
	return nil
}

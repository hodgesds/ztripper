package main

import (
	"fmt"
	"github.com/hodgesds/ztripper"
	"net/http"
)

func main() {
	t := &http.Transport{}

	zReqTripper, err := ztripper.NewZmqTripper(ztripper.NewByteSerializer())
	if err != nil {
		panic(err)
	}
	defer zReqTripper.Destroy()

	t.RegisterProtocol("zmq", zReqTripper)

	c := &http.Client{Transport: t}

	res, err := c.Get("zmq://127.0.0.1:8001/foo/bar/baz?qux=1")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

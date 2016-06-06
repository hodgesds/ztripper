# ZMQ RoundTripper support for go

# Example
Run the server.py (requires pyzmq) and then run the following example:

```go
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

    // register the zmq uri to the transport
	t.RegisterProtocol("zmq", zReqTripper)

	c := &http.Client{Transport: t}

	res, err := c.Get("zmq://127.0.0.1:8001/foo/bar/baz?qux=1")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
```

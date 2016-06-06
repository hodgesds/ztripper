# ZMQ RoundTripper support for go
ztripper is used for providing a *`RoundTripper`* interface that uses ZMQ as a transport. A [RoundTripper](https://golang.org/pkg/net/http/#RoundTripper) is described as:
>RoundTripper is an interface representing the ability to execute a single HTTP transaction, obtaining the Response for a given Request.

# Why
After looking into HTTP2 support worked in go 1.6+ I decided it would be fun to try to make something similar work with ZMQ. The interesting result is that you can use a regular [http.Client](https://golang.org/pkg/net/http/#Client) to make ZMQ requests. 

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

	// register the zmq scheme to the transport
	// XXX: different schemes for different socket types?
	t.RegisterProtocol("zmq", zReqTripper)

	c := &http.Client{Transport: t}

	res, err := c.Get("zmq://127.0.0.1:8001/foo/bar/baz?qux=1")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
```

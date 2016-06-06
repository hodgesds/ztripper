package ztripper

import (
	"github.com/zeromq/goczmq"
	"net/http"
)

type ZmqTripper struct {
	serializer Serializer
	channelers map[string]*goczmq.Channeler
}

func NewZmqTripper(
	serializer Serializer,
) (*ZmqTripper, error) {
	return &ZmqTripper{
		channelers: map[string]*goczmq.Channeler{},
		serializer: serializer,
	}, nil
}

func (zrt *ZmqTripper) Destroy() {
	for _, channeler := range zrt.channelers {
		channeler.Destroy()
	}
	zrt.channelers = map[string]*goczmq.Channeler{}
}

func (zrt *ZmqTripper) RoundTrip(
	req *http.Request,
) (*http.Response, error) {
	var channeler *goczmq.Channeler
	var ok bool

	key := req.URL.Scheme + req.URL.Host
	tcpHost := "tcp://" + req.URL.Host

	// check for cached connection
	channeler, ok = zrt.channelers[key]
	if !ok {
		channeler = zrt.getChanneler(key, tcpHost)
		zrt.channelers[key] = channeler
	}

	//serialize request for transport
	reqBytes, err := zrt.serializer.Marshal(req)
	if err != nil {
		return nil, err
	}

	channeler.SendChan <- [][]byte{reqBytes}

	replyBytes := <-channeler.RecvChan

	res, err := zrt.serializer.Unmarshal(replyBytes[0])
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (zrt ZmqTripper) getChanneler(key, tcpHost string) *goczmq.Channeler {
	switch key {
	case "REQ":
		return goczmq.NewReqChanneler(tcpHost)

	case "PUSH":
		return goczmq.NewPushChanneler(tcpHost)

	case "PULL":
		return goczmq.NewPullChanneler(tcpHost)

	case "ROUTER":
		return goczmq.NewRouterChanneler(tcpHost)

	case "DEALER":
		return goczmq.NewDealerChanneler(tcpHost)

	default:
		return goczmq.NewReqChanneler(tcpHost)
	}
}

package ztripper

import (
	"bufio"
	"bytes"
	"net/http"
)

type Serializer interface {
	Marshal(*http.Request) ([]byte, error)
	Unmarshal([]byte) (*http.Response, error)
}

type ByteSerializer struct{}

func NewByteSerializer() *ByteSerializer {
	return &ByteSerializer{}
}

func (bs ByteSerializer) Marshal(req *http.Request) ([]byte, error) {
	var buf bytes.Buffer

	// other options:
	// https://golang.org/pkg/net/http/httputil/#DumpRequest
	// net/http/httputil.DumpRequest(req, true)

	err := req.Write(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (bs ByteSerializer) Unmarshal(b []byte) (*http.Response, error) {
	buf := bytes.NewBuffer(b)

	reader := bufio.NewReader(buf)

	return http.ReadResponse(reader, nil)
}

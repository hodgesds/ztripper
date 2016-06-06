import zmq
import time
import sys

port = "8001"
if len(sys.argv) > 1:
    port =  sys.argv[1]
    int(port)

context = zmq.Context()
socket = context.socket(zmq.REP)
uri = "tcp://127.0.0.1:%s" % port
socket.bind(uri)


reply = """HTTP/1.1 200
Location: %s
Server: Zmq Rep
Content-Length: %d

%s"""

while True:
    #  Wait for next request from client
    message = socket.recv()
    print "Received request:\n%s" % message
    body =  "<h1>Hello World</h1>"
    res = reply % (uri, len(body), body)
    print "Sending reply:\n%s" % res
    socket.send(res)

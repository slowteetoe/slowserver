Simple test server to make sure that an http client is using read timeout correctly.  This will simulate a slow http server, sending chunked data every X seconds for Y seconds total. (configurable)

Usage
====
go run main.go

Then in a separate window run something like ngrok to listen to the TCP connection on port 6666.
  sudo ngrep -W byline -d lo0 port 6666

Finally, request http://localhost:6666 using the client with which you want to check the behavior


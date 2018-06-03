# http-server-sent-events
A simple Golang HTTP server supporting server sent events to client. This would be the backbone to prototype **reactive** Go HTTP clients.

### Usage:

Make a HTTP GET request on localhost:8080/foo. This would trigger corresponding HTTP handler to emit a stream of integer numbers after a set delay of 5 millisecond each. 

### Future plan:

Planning to write a reactive client framework for Golang complying to the [reactive](http://reactivex.io/) manifesto.

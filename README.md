# netswi
Network switcher for HTTP

# Usage

```go
client := netswi.NewClientBindedToIp("127.0.0.1")
```
or 

```go
client := &http.Client{Transport: netswi.NewTransportBindedToIP("127.0.0.1")}
```

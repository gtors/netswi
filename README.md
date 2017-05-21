# netswi
Network switcher for HTTP

# Usage

```go
client, err := netswi.NewClientBindedToIP("127.0.0.1")
```
or 

```go
if trans, err := netswi.NewTransportBindedToIP("127.0.0.1"); err != nil {
  // handle error
} else {
  client := &http.Client{Transport: trans}
}
```

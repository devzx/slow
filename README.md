# Slow
Takes the output of commands like `kubectl logs` and `stern`, and outputs them ...slowly

### Build
From source
```sh
$ git clone git@github.dns.ad.zopa.com:zopaUK/slow.git
$ cd slow
$ go build slow.go
$ mv slow /usr/local/bin
```
Docker
```sh
$ git clone git@github.dns.ad.zopa.com:zopaUK/slow.git
$ cd slow
$ docker build -t slow .
```
### Flags
`-wait int`

Time to wait in seconds between printing output (default 1)
### Examples
Binary

`kubectl logs <pod> | slow`

`stern <pod> | slow -wait 2`

Docker

`kubectl logs <pod> | docker run -i slow -wait 1  -`

`stern <pod> | slow -wait 2`

Prebuilt image

`kubectl logs <pod> | docker run -i --rm devzx/slow -`

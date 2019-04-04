# Slow

#### Takes the output of commands like `kubectl logs` and `stern`, and outputs them ...slowly

### Build

```sh
git clone git@github.dns.ad.zopa.com:zopaUK/slow.git
cd slow
go build slow.go
mv slow /usr/local/bin
```

### Flags
`-wait int`

time to wait in seconds between printing output

### Examples
`kubectl logs <some nginx pod> | slow`

`stern <some nginx pod> | slow -wait 2`

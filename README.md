# smsexpress

[Go](https://golang.org) package to quickly send messages via various SMS providers.

View [docs](https://godoc.org/github.com/samora/smsexpress)

## Features

* Attempt sequentially sending SMS via each registered provider until successful.
* Easily plugin new providers by implementing `smsexpress.Provider` interface.

## Supported providers (Ghana)

* [SMSGH](http://www.smsgh.com)
* [Nasara Mobile](http://www.nasaramobile.com)
* ... contributions?


## License

MIT
# Go Wasm with Spin

[Spin](https://developer.fermyon.com/spin/index)
_Spin is a framework for building and running event-driven microservice applications with WebAssembly (Wasm) components._

This sample application illustrates writing multiple microservices with WASM in Golang.

## Code organization

* [bar-http](bar-http/main.go) - Is a sample code that can respond to http calls.  The code is compiled to WASM.
  It is a very simple miroservice http with endpoint `/bar`.  


## Build

To build the application with spin.

``` shell
$ export SPIN_APP_APP_TEXT=1234
$ export SPIN_APP_USER_TOKEN=abc45789
$ spin build --up
```
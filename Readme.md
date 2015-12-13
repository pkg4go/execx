
[![Build status][travis-img]][travis-url]
[![License][license-img]][license-url]
[![GoDoc][doc-img]][doc-url]

### execx

* A tiny wrapper of `os/exec`.

### APIs

* Run: `Run(name string, arg ...string) (out string, err error)`

```go
out, err := Run("ls", "-a")
```

### License
MIT

[travis-img]: https://img.shields.io/travis/pkg4go/execx.svg?style=flat-square
[travis-url]: https://travis-ci.org/pkg4go/execx
[license-img]: http://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[doc-img]: http://img.shields.io/badge/GoDoc-reference-blue.svg?style=flat-square
[doc-url]: http://godoc.org/github.com/pkg4go/execx

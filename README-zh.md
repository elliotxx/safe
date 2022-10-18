<div align = "center">
<p>
    <img width="160" src="https://github.com/elliotxx/safe/blob/master/golang-logo.png?sanitize=true">
</p>
<h2>一个安全使用 Goroutine 的 Golang 库</h2>
<a title="Go Reference" target="_blank" href="https://pkg.go.dev/github.com/elliotxx/safe"><img src="https://pkg.go.dev/badge/github.com/elliotxx/safe.svg"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/elliotxx/safe"><img src="https://goreportcard.com/badge/github.com/elliotxx/safe?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/github/elliotxx/safe?branch=master"><img src="https://img.shields.io/coveralls/github/elliotxx/safe/master"></a>
<a title="Code Size" target="_blank" href="https://github.com/elliotxx/safe"><img src="https://img.shields.io/github/languages/code-size/elliotxx/safe.svg?style=flat-square"></a>
<br>
<a title="GitHub release" target="_blank" href="https://github.com/elliotxx/safe/releases"><img src="https://img.shields.io/github/release/elliotxx/safe.svg"></a>
<a title="License" target="_blank" href="https://github.com/elliotxx/safe/blob/master/LICENSE"><img src="https://img.shields.io/github/license/elliotxx/safe"></a>
<a title="GitHub Commits" target="_blank" href="https://github.com/elliotxx/safe/commits/master"><img src="https://img.shields.io/github/commit-activity/m/elliotxx/safe.svg?style=flat-square"></a>
<a title="Last Commit" target="_blank" href="https://github.com/elliotxx/safe/commits/master"><img src="https://img.shields.io/github/last-commit/elliotxx/safe.svg?style=flat-square&color=FF9900"></a>
</p>
</div>

`elliotxx/safe` 是一个安全使用 Goroutine 的 Golang 库，启发于 [safe](https://pkg.go.dev/github.com/traefik/traefik/v2@v2.9.1/pkg/safe)!

## 📜 语言

[English](https://github.com/elliotxx/safe/blob/master/README.md) | [简体中文](https://github.com/elliotxx/safe/blob/master/README-zh.md)

## ⚡ 使用

```
go get -u github.com/elliotxx/safe
```

## 📚 样例

```go
package main

import "github.com/elliotxx/safe"

func doSomething() {
	panic("BOOM!")
}

func customRecoverHandler(r interface{}) {
	log.Printf("Recovered from %v", r)
}

func main() {
	// Go starts a recoverable goroutine.
	safe.Go(
		doSomething,
	)

	// GoR starts a recoverable goroutine using given recover handler.
	safe.GoR(
		doSomething,
		customRecoverHandler
	)

	// DefaultHandleCrash simply catches a crash with the default recover handler.
	go func() {
		defer safe.DefaultHandleCrash()
		doSomething()
	}()

	// HandleCrash catches a crash with the custom recover handlers.
	go func() {
		defer safe.HandleCrash(customRecoverHandler)
		doSomething()
	}()
}
```

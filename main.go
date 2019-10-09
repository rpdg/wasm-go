package main

import (
	"syscall/js"
	"crypto/md5"
    "encoding/hex"
)

var c chan bool

func init() {
	println("Initialize WebAssembly!")
	c = make(chan bool)
}

func main() {
	println("WASM Initialized")

	// register functions
	registerCallbacks()

	<- c
}


func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("sayMsg", js.FuncOf(sayMsg))
	js.Global().Set("calMD5", js.FuncOf(calMD5))
}


func sayMsg(this js.Value, args []js.Value) interface{} {
	alert := js.Global().Get("alert");
	alert.Invoke(args[0].String())
	return nil
}

func calMD5(this js.Value, args []js.Value) interface{} {
	hasher := md5.New()
    hasher.Write([]byte(args[0].String()))
	res := hex.EncodeToString(hasher.Sum(nil))
	return js.ValueOf(res)
}


// https://tutorialedge.net/golang/go-webassembly-tutorial/
// https://stackoverflow.com/questions/56398142/is-it-possible-to-explicitly-call-an-exported-go-webassembly-function-from-js
func add(this js.Value, i []js.Value) interface{} {
	sum := i[0].Int() + i[1].Int()
    return js.ValueOf(sum)
}


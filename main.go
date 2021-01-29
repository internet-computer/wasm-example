package main

import (
	"strconv"
	"syscall/js"
)

func add(_ js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	v1, _ := strconv.Atoi(value1)

	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	v2, _ := strconv.Atoi(value2)

	result := v1 + v2
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", result)
	return result
}

func subtract(_ js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	v1, _ := strconv.Atoi(value1)

	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	v2, _ := strconv.Atoi(value2)

	result := v1 - v2
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", result)
	return result
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	println("Initialized WASM Go.")
	<-c
}

package main

import "fmt"

var xyz int

func main() {
	fmt.Println(5, 4, 3, 2)
	fmt.Println(5.4, 5.2, 5.123)
	fmt.Println(true, false)
	fmt.Println("Miłosz")
	fmt.Println("Miłosz Gałązka")

	var x string
	fmt.Println(0x01)
	fmt.Println(0x02)
	fmt.Println(0x03)
	fmt.Println(0x04)
	fmt.Println(0x05)
	fmt.Println(0x06)
	fmt.Println(0x07)
	fmt.Println(0x08)
	fmt.Println(0x09)
	fmt.Println(0x0A)
	fmt.Println(0x0B)
	fmt.Println(0x0C)
	fmt.Println(0x0D)
	fmt.Println(0x0E)
	fmt.Println(0x0F)
	fmt.Println(0x11)
	fmt.Println(0x19)
	fmt.Println(0x32)
	fmt.Println(0x64)

	fmt.Printf("%q\n", x)

	var height int
	fmt.Println(height)

	var isOn bool
	fmt.Println(isOn)

	var brightness float64
	fmt.Println(brightness)

	var svariable string
	fmt.Printf("s (%T): %q\n", svariable, svariable)

	var isLiquid bool
	_ = isLiquid

	var v1 int
	var v2 int8
	var v3 int16
	var v4 int32
	var v5 int64
	var v6 float32
	var v7 float64
	var v8 complex64
	var v9 complex128
	var v91 bool
	var v92 string
	var v93 rune
	var v94 byte

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)
	fmt.Println(v6)
	fmt.Println(v7)
	fmt.Println(v8)
	fmt.Println(v9)
	fmt.Println(v91)
	fmt.Println(v92)
	fmt.Println(v93)
	fmt.Println(v94)

	var (
		isActive bool
		delta    int
	)

	fmt.Println(isActive)
	fmt.Println(delta)

	fmt.Println(after)
	var after int
}

package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {

	pathUrl := "/learn-golang-basic/array.go"

	fmt.Println(path.Base(pathUrl))
	fmt.Println(path.Dir(pathUrl))
	fmt.Println(path.Ext(pathUrl))
	fmt.Println(path.Join("/", "learn-golang-basic", "array.go"))

	fmt.Println(filepath.Base(pathUrl))
	fmt.Println(filepath.Dir(pathUrl))
	fmt.Println(filepath.Ext(pathUrl))
	fmt.Println(filepath.Join("/", "learn-golang-basic", "array.go"))
	fmt.Println(filepath.IsAbs(pathUrl))
	fmt.Println(filepath.IsLocal(pathUrl))

}

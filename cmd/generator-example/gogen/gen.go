package main

import "os"

func main() {
	f, _ := os.Create("hello_gen.go")
	f.WriteString("package main \n\nimport \"fmt\"\n\nfunc sayHello() { fmt.Println(\"Привет из генераора!\")}")
	f.Close()
}

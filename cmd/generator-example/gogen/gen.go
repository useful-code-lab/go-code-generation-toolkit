package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var name = flag.String("name", "Gopher", "имя для приветствия")
	flag.Parse()
	fmt.Println("Привет,", *name, "!")
	f, _ := os.Create("hello_gen.go")
	f.WriteString("package main \n\nimport \"fmt\"\n\nfunc sayHello() { fmt.Println(\"Привет из генератора!\")}")
	f.Close()
}

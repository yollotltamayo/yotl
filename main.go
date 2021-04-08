package main

import (
    "fmt"
    "os"
    "os/user"
    "yotl/repl"
)

func main() {
    user ,err :=  user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("hola %s bienvenido a SNELLY!\n" , user.Username)
    fmt.Printf("vamos es momento de programar ;) \n")
    repl.Start(os.Stdin,os.Stdout)
}


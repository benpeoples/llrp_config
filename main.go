package main

import (
  "github.com/abiosoft/ishell"
  //"fmt"
  // "flag"
  // "log"
  // "encoding/hex"
  // "os"
  // "time"
)

func discovery(c *ishell.Context) {
  c.Println("Discovery")
}

func device_label(c *ishell.Context) {
  c.Println("device_label")
}



func main() {
  shell := ishell.New()

  shell.Println("LLRP Configuration Tool")

  shell.AddCmd(&ishell.Cmd{
        Name: "d",
        Help: "Run discovery",
        Func: discovery,
    })

  shell.AddCmd(&ishell.Cmd{
        Name: "l",
        Help: "Get or Set Device Label",
        Func: device_label,
    })

    shell.Run()

}

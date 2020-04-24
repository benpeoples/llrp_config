package main

import (
  "github.com/abiosoft/ishell"
  //"fmt"
  // "flag"
  // "log"
  // "encoding/hex"
  // "os"
  // "time"
  "net"
  "regexp"
)

var active_interface int = 0;

func discovery(c *ishell.Context) {
  c.Println("Discovery")
}

func device_label(c *ishell.Context) {
  c.Println("device_label")
}

func factory_defaults(c *ishell.Context) {
  c.Println("factory_defaults")
}

func identify_device(c *ishell.Context) {
  c.Println("identify_device")
}

func reset_device(c *ishell.Context) {
  c.Println("reset_device")
}

func dhcp_mode(c *ishell.Context) {
  c.Println("dhcp_mode")
}

func current_address(c *ishell.Context) {
  c.Println("current_address")
}

func static_address(c *ishell.Context) {
  c.Println("static_address")
}

func default_route(c *ishell.Context) {
  c.Println("default_route")
}

func print_device_table(c *ishell.Context) {
  c.Println("print_device_table")
}

func refresh(c *ishell.Context) {
  c.Println("refresh")
}

func list_interfaces(c *ishell.Context) {
  c.Println("list_interfaces")

  interfaces, _ := net.Interfaces()

  for _, iface := range interfaces {
    // \b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b
    addr, _ := iface.Addrs()
    address := ""
    var validIP = regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
    for _, this_address := range addr {
        address = this_address.String();
      }


    c.Println(iface.Index, "|", iface.MTU, "|", iface.Name, "|", iface.HardwareAddr, "|", address);

  }

}

func select_interface(c *ishell.Context) {
  choice := c.MultiChoice([]string{
        "Golangers",
        "Go programmers",
        "Gophers",
        "Goers",
    }, "What are Go programmers called ?")
    if choice == 2 {
        c.Println("You got it!")
    } else {
        c.Println("Sorry, you're wrong.")
    }
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
    Help: "Get or Set Device Label: l # [label]",
    Func: device_label,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "factory",
    Help: "Reset Device Settings to Factory Defaults: factory #",
    Func: factory_defaults,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "i",
    Help: "Get or Set Device Identify Status: i # [on/off]",
    Func: identify_device,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "reset",
    Help: "Reset (Reboot) Device: reset #",
    Func: reset_device,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "reboot",
    Help: "Alias of reset. Reset (Reboot) Device: reboot #",
    Func: reset_device,
  })


  shell.AddCmd(&ishell.Cmd{
    Name: "dhcp",
    Help: "Get or Set DHCP Mode: dhcp # [on/off]",
    Func: dhcp_mode,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "c",
    Help: "Get current IP address: c #",
    Func: current_address,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "setip",
    Help: "Get or Set Static IP: setip # [dotted quad ip] [dotted quad netmask]",
    Func: static_address,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "gw",
    Help: "Get or Set Device Gateway: gw # [dotted quad gateway ip]",
    Func: default_route,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "p",
    Help: "Print device table",
    Func: print_device_table,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "r",
    Help: "Refresh device table",
    Func: refresh,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "li",
    Help: "List Computer IP interfaces",
    Func: list_interfaces,
  })

  shell.AddCmd(&ishell.Cmd{
    Name: "si",
    Help: "Select Computer IP interface",
    Func: select_interface,
  })

  shell.Run()

}

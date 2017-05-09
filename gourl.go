
package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()

  app.Commands = []cli.Command{
    {
      Name:    "get",
      Aliases: []string{"g"},
      Usage:   "make an HTTP GET request",
      Action:  func(c *cli.Context) error {
        fmt.Println("GET: ", c.Args().First())
       return nil
      },
    },
    {
      Name:    "post",
      Aliases: []string{"p"},
      Usage:   "make an HTTP POST request",
      Action:  func(c *cli.Context) error {
        fmt.Println("POST: ", c.Args().First())
        return nil
      },
    },
  }

  app.Run(os.Args)
}

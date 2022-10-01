package main

import "gin-dasar/routers"

func main() {
  var port = ":8080"

  routers.StartServer().Run(port)
}

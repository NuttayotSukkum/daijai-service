package main

import "daijai-service/routers"

func main() {
	e := routers.ProjectRouter()

	routers.Execute(e)
}

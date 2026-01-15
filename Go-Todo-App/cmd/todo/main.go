package main

import (
	"fmt"

	"github.com/Aparneesh-Patil/GO-Todo-CLI/internal/app"
)

func main() {
	fmt.Println(`
___________        .___        ____    .__          __   
\__    ___/___   __| _/____   |    |   |__| _______/  |_ 
  |    | /  _ \ / __ |/  _ \  |    |   |  |/  ___/\   __\
  |    |(  <_> ) /_/ (  <_> ) |    |___|  |\___ \  |  |  
  |____| \____/\____ |\____/  |_______ \__/____  > |__|  
                    \/                \/       \/         `)

	fmt.Print("Welcome to the Todo List App! ")
	app.Run()
}

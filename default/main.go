package main
 
import(
"flag"
"fmt"
)

func main(){
	configClient := flag.String("client", "false", "a string")
	flag.Parse()
	fmt.Println("client:", *configClient)
}
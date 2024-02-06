package chapter2

import (
	"goAction/chapter2/search"
	"log"
	"os"
)

func init()  {
	log.SetOutput(os.Stdout)
}

func main()  {
	search.Run()
}
package main


import(
	"fmt"
	//"reflect"
)

type Animanl struct{
	Spiecies string
	Origin string
}
//go use Composition there is no inheritance
type Bird struct{
	Animanl
	Fly int
}
func main(){

	birds := Bird{}
	birds.Spiecies = "Eagle"
	birds.Origin = "Africa"
	birds.Fly = 2323
	fmt.Println(birds.Origin)
}
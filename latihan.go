package main
import "fmt"
func faktorial(i int) int{
    if i == 1 || i == 0{
        return 1
    }else{
        return i*faktorial(i-1)
    }

}

func main(){
    var i int
    fmt.Scan(&i)
    fmt.Print(faktorial(i))
}
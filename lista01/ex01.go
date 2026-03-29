package main
import ("fmt")

func main() {
 var x float32
 var y float32
 var z float32
 var media float32
 fmt.Scan(&x)
 fmt.Scan(&y)
 fmt.Scan(&z)
 media = (x + y + z) / 3
fmt.Printf("MEDIA = %.2f\n", media)
if media>=6 {
	fmt.Println("APROVADO")
}else {
	fmt.Println("REPROVADO")
}
}
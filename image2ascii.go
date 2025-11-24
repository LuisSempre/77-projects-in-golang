//converter imagem para ascii art
package main
import (
	"fmt"
	"os"
	"github.com/qeesung/image2ascii/convert"
)
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run mascote.go <image_path>")
		return
	}
	imagePath := os.Args[1]
	convertOptions := convert.DefaultOptions
	converter := convert.NewImageConverter()
	fmt.Println(converter.ImageFile2ASCIIString(imagePath, &convertOptions))
}
// Para executar o codigo, salve-o em um file chamado mascote.go e use o comando:
// go run mascote.go caminho/para/sua/image.jpg
// install: go get -u github.com/qeesung/image2ascii/convert


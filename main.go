package main 
import (
	"fmt"
	"encoding/base64"
	"io"
	"os"
	"strings"
)
const data = `SGVsbG8gYmFzZTY0IGVuY29kZXIuIFRoYW5rcyBVdGFoIEdvIE1lZXR1cCE=`

func main() {
	var r io.Reader
	r = strings.NewReader(data)
	r = base64.NewDecoder(base64.StdEncoding, r)
	// r, _ = gzip,Newreader(r)
	b, err := io.Copy(os.Stdout, r)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(b)

	// fmt.Println("Hello World!")
}
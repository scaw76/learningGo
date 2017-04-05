package reading

import (
	"encoding/base64"
	"io"
	"strings"
)
func ReadMyData(data string) io.Reader {
	var r io.Reader
	r = strings.NewReader(data)
	r = base64.NewDecoder(base64.StdEncoding, r)
	// r, _ = gzip,Newreader(r)
	//b, err := io.Copy(os.Stdout, r)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(b)

	// fmt.Println("Hello World!")

	return r
}
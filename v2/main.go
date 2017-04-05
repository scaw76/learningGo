package main 


import (
	"github.com/scaw76/learningGo/v2/reading"
	"github.com/scaw76/learningGo/v2/web"

)
const data = `SGVsbG8gYmFzZTY0IGVuY29kZXIuIFRoYW5rcyBVdGFoIEdvIE1lZXR1cCE=`

func main() {
	var g web.Greeting
	r := reading.ReadMyData(data)

	g.Run(r)

}
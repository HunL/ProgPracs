package main

import (
	. "fmt"
	"github.com/skip2/go-qrcode"
	"image/color"
)

func main() {
	Println("main...")
	// qrcode.WriteFile("李杰辉很帅！", qrcode.Medium, 256, "./name_qrcode.png")
	qr, err := qrcode.New("ljiehui0826@gmail.com", qrcode.Medium)
	if err != nil {
		Println("error!")
	} else {
		qr.BackgroundColor = color.RGBA{50,205,50,255}
		qr.ForegroundColor = color.Black
		qr.WriteFile(256, "xxx.png")
	}
}
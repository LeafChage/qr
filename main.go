package main

import (
    "flag"
    "image/png"
    "os"
    "github.com/boombuler/barcode"
    "github.com/boombuler/barcode/qr"
    "os/exec"
)

func main() {
    var(
        w int
        h int
        filePath string
    )
    flag.IntVar(&w, "w", 200, "int width")
    flag.IntVar(&h, "h", 200, "int height")
    flag.StringVar(&filePath, "o", "~/tmp/go_qrcode.png", "string output file path")

    flag.Parse()

    qrCode, _ := qr.Encode(flag.Arg(0), qr.M, qr.Auto)
    qrCode, _ = barcode.Scale(qrCode, w, h)
    file, _ := os.Create(filePath)
    defer file.Close()
    png.Encode(file, qrCode)

    err := exec.Command("open", filePath).Start()
    if err != nil {
        panic(err)
    }
}




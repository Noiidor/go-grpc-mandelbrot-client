package main

import (
	"bytes"
	"context"
	"image/png"
	"log"
	"os"

	pb "github.com/Noiidor/go-grpc-mandelbrot-client/internal/proto"
	"github.com/Noiidor/go-grpc-mandelbrot-client/pkg/client"
)

func main() {
	client, err := client.NewClient()
	if err != nil {
		log.Fatalf("Cant create gRPC client: %v", err)
	}

	mandelSettings := pb.MandelbrotSettings{
		Width:   2000,
		Height:  1000,
		Zoom:    1487573,
		CenterX: -0.226266648,
		CenterY: 1.11617444,
	}

	img, err := client.GetImage(context.Background(), &mandelSettings)
	if err != nil {
		log.Fatalf("Cannot get image: %v", err)
	}

	imgBytes := img.ImageContent

	buffer := bytes.NewBuffer(imgBytes)

	decodedImg, err := png.Decode(buffer)
	if err != nil {
		log.Fatalf("Error during decoding image buffer: %v", err)
	}

	imgFile, err := os.Create("decoded.png")
	if err != nil {
		log.Fatalf("Cannot create image file: %v", err)
	}
	defer imgFile.Close()

	if err = png.Encode(imgFile, decodedImg); err != nil {
		log.Fatalf("Cannot encode image bytes into file: %v", err)
	}

	log.Print("Image received successfully")

}

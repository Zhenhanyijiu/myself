package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
	"github.com/rwcarlsen/goexif/exif"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
)

func readFile(inPath string) ([]byte, error) {
	inf, err := os.Open(inPath)
	if err != nil {
		return nil, err
	}
	defer inf.Close()
	return ioutil.ReadAll(inf)
}

func writeFile(outPth string, des []byte) (int, error) {
	outf, err := os.Create(outPth)
	if err != nil {
		return 0, err
	}
	defer outf.Close()
	return outf.Write(des)
}
func main() {
	inPath := flag.String("in", "", "input file")
	outPath := flag.String("out", "", "output file")
	flag.Parse()
	if *inPath == "" || *outPath == "" {
		flag.Usage()
		return
	}
	out, err := readFile(*inPath)
	if err != nil {
		fmt.Printf("read file error:%v\n", err)
		return
	}
	imgReader := bytes.NewReader(out)
	img, str, err := image.Decode(imgReader)
	if err != nil {
		fmt.Printf("image decode error:%v\n", err)
		return
	}
	fmt.Printf("image decode :%v\n", str)
	//adjust photo

	//rotation
	imgReader.Seek(0, 0)
	x, err := exif.Decode(imgReader)
	//loading EXIF sub-IFD: exif: sub-IFD ExifIFDPointer decode failed: zero length tag value
	//has exif infomation
	if err != nil {
		fmt.Printf("exif.Decode,rotation failed error:%v\n", err)
	}
	if x != nil {
		fmt.Printf("exif.Decode Exif(%p)\n", x)
		orient, err := x.Get(exif.Orientation)
		//orient.String() == "1" or "0", dont need reverseRotation
		if err == nil {
			fmt.Printf("rotation value:%v\n", orient.String())
			if orient != nil && orient.String() != "1" && orient.String() != "0" {
				img = reverseOrientation(img, orient.String())
				fmt.Printf("picture had orientation.......\n")
			}
		}
	}
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	fmt.Printf("====width:%v, hight:%v\n", w, h)
	//before resize ,adjust picture first
	if n, tag := min(w, h); n < 150 {
		fmt.Printf("=======min(w, h)<150,error, tag(%v)\n", tag)
		return
	}
	if n, tag := min(w, h); n > 720 {
		var imgNew image.Image
		switch tag {
		case "w<=h":
			imgNew = resize.Resize(720, 0, img, resize.Lanczos3)
		case "w>h":
			imgNew = resize.Resize(0, 720, img, resize.Lanczos3)
		default:
			fmt.Printf("========error\n")
			return
		}
		buf := bytes.Buffer{}
		if err := jpeg.Encode(&buf, imgNew, &jpeg.Options{100}); err != nil {
			fmt.Printf("jpeg encode :%v\n", err)
			return
		}
		fmt.Printf("=======width:%v, hight:%v\n", imgNew.Bounds().Dx(), imgNew.Bounds().Dy())
		n, err := writeFile(*outPath, buf.Bytes())
		fmt.Printf("writeFile :n(%v),err(%v)\n", n, err)
		return
	}
	fmt.Printf("=======min(w,h)>=150&&min(w,h)<=720\n")
}
func reverseOrientation(img image.Image, o string) *image.NRGBA {
	fmt.Printf("execute reverseOrientation(%v)\n", o)
	switch o {
	case "1":
		return imaging.Clone(img)
	case "2":
		return imaging.FlipV(img)
	case "3":
		return imaging.Rotate180(img)
	case "4":
		return imaging.Rotate180(imaging.FlipV(img))
	case "5":
		return imaging.Rotate270(imaging.FlipV(img))
	case "6":
		return imaging.Rotate270(img)
	case "7":
		return imaging.Rotate90(imaging.FlipV(img))
	case "8":
		return imaging.Rotate90(img)
	}
	return imaging.Clone(img)
}

func min(w, h int) (int, string) {
	if w <= h {
		return w, "w<=h"
	}
	return h, "w>h"
}

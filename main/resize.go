package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"time"
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
	//img, _, err := image.Decode(imgReader)
	img, err := imaging.Decode(imgReader, imaging.AutoOrientation(true))
	if err != nil {
		fmt.Printf("image decode error:%v\n", err)
		return
	}
	fmt.Printf("image decode :%v\n")
	//adjust photo

	//rotation
	/***************************
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
	*************************/
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
			imgNew = imaging.Resize(img, 720, 0, imaging.NearestNeighbor)
		case "w>h":
			imgNew = imaging.Resize(img, 0, 720, imaging.NearestNeighbor)
		default:
			fmt.Printf("========error\n")
			return
		}
		buf := bytes.Buffer{}
		//if err := jpeg.Encode(&buf, imgNew, &jpeg.Options{100}); err != nil {
		//	fmt.Printf("jpeg encode :%v\n", err)
		//	return
		//}
		format, err := imaging.FormatFromExtension("jpg")
		if err != nil {
			fmt.Printf("FormatFromExtension error :%v\n", err)
			return
		}
		err = imaging.Encode(&buf, imgNew, format)
		if err != nil {
			fmt.Printf("imaging.Encode error :%v\n", err)
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
func RGBToImage(rgb []byte, w int, h int) (img []byte, err error) {
	begin := time.Now()
	defer func() {
		fmt.Printf("RGBToImg End need time(%v ms)\n", time.Since(begin).Nanoseconds()/1e6)
	}()
	if len(rgb) == 0 || w < 1 || h < 1 {
		return nil, errors.New("RGBToImg param error")
	}
	rect := image.Rect(0, 0, w, h)
	rgba := image.NewRGBA(rect)
	fmt.Printf("RGBToImg: source size(%v bytes)\n", len(rgb))
	rd := bytes.NewReader(rgb)
	buf := make([]byte, 3)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			io.ReadFull(rd, buf)
			y1, cb, cr := color.RGBToYCbCr(buf[0], buf[1], buf[2])
			//r, g, b, a := c.RGBA()
			ycbcr := color.YCbCr{
				Y:  y1,
				Cb: cb,
				Cr: cr,
			}
			//col:= color.Color{r, g, b, a}
			rgba.Set(x, y, ycbcr)
		}
	}
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, rgba, &jpeg.Options{100}); err != nil {
		return nil, err
	}
	fmt.Printf("RGBToImg: to jpg size(%v bytes)\n", len(buffer.Bytes()))
	return buffer.Bytes(), nil
}

func IamgeToRGBByExif(imgSrc []byte) (rgb []byte, w int, h int, err error) {
	begin := time.Now()
	defer func() {
		fmt.Printf("ImgToRGB End need time(%v ms)\n", time.Since(begin).Nanoseconds()/1e6)
	}()
	if len(imgSrc) == 0 {
		return nil, 0, 0, errors.New("ImgToRGB param image is null")
	}
	imgsReader := bytes.NewReader(imgSrc)
	conf, format, err := image.DecodeConfig(imgsReader)
	if err != nil {
		return nil, 0, 0, err
	}
	fmt.Printf("ImgToRGB, photo-format=%v, len(imageSrc)=%v,,before rotation(w:%v, h:%v)\n", format, len(imgSrc), conf.Width, conf.Height)
	imgsReader.Seek(0, 0)
	var imgs image.Image
	if format == "png" {
		imgs, err = png.Decode(imgsReader)
		if err != nil {
			return nil, 0, 0, err
		}
	} else {
		imgs, err = jpeg.Decode(imgsReader)
		if err != nil {
			return nil, 0, 0, err
		}
		//rotation
		imgsReader.Seek(0, 0)
		x, err := exif.Decode(imgsReader)
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
					imgs = reverseOrientation(imgs, orient.String())
					fmt.Printf("picture had orientation.......\n")
				}
			}
		}
	}
	bounds := imgs.Bounds()
	//m := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	m := image.NewRGBA(bounds)
	//draw.Draw(m, m.Bounds(), imgs, bounds.Min, draw.Src)
	draw.Draw(m, bounds, imgs, bounds.Min, draw.Src)
	//draw.Draw(m, bounds, imgs, bounds.Min, draw.Over)
	//var buf bytes.Buffer
	pixLen := len(m.Pix)
	rgb = make([]byte, pixLen*3/4)
	k := 0
	for i := 0; i < pixLen; i += 4 {
		rgb[k] = m.Pix[i]
		rgb[k+1] = m.Pix[i+1]
		rgb[k+2] = m.Pix[i+2]
		k += 3
	}
	//rgb = buf.Bytes()
	fmt.Printf("ImgToRGB, len(rgb):%v,  width:%v,  height:%v\n", len(rgb), bounds.Max.X, bounds.Max.Y)
	return rgb, bounds.Max.X, bounds.Max.Y, nil
}
func RGBToImageGoRoutine(rgb []byte, w int, h int) (img []byte, err error) {
	begin := time.Now()
	defer func() {
		fmt.Printf("RGBToImg End need time(%v ms)\n", time.Since(begin).Nanoseconds()/1e6)
	}()
	if len(rgb) == 0 || w < 1 || h < 1 {
		return nil, errors.New("RGBToImg param error")
	}
	rect := image.Rect(0, 0, w, h)
	rgba := image.NewRGBA(rect)
	fmt.Printf("RGBToImg: source size(%v bytes)\n", len(rgb))
	//rd := bytes.NewReader(rgb)
	//buf := make([]byte, 3)
	//for y := 0; y < h; y++ {
	//	for x := 0; x < w; x++ {
	//		io.ReadFull(rd, buf)
	//		y1, cb, cr := color.RGBToYCbCr(buf[0], buf[1], buf[2])
	//		//r, g, b, a := c.RGBA()
	//		ycbcr := color.YCbCr{
	//			Y:  y1,
	//			Cb: cb,
	//			Cr: cr,
	//		}
	//		//col:= color.Color{r, g, b, a}
	//		rgba.Set(x, y, ycbcr)
	//	}
	//}
	//
	len1 := 3 * w

	parallel(0, h, func(ints <-chan int) {
		for y := range ints {
			len2 := len1 * y
			//buf := make([]byte, 3)
			for x := 0; x < w; x++ {
				//io.ReadFull(rd, buf)
				st := 3*x + len2
				buf := rgb[st : st+3]
				y1, cb, cr := color.RGBToYCbCr(buf[0], buf[1], buf[2])
				//r, g, b, a := c.RGBA()
				ycbcr := color.YCbCr{
					Y:  y1,
					Cb: cb,
					Cr: cr,
				}
				//col:= color.Color{r, g, b, a}
				rgba.Set(x, y, ycbcr)
			}
		}
	})
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, rgba, &jpeg.Options{100}); err != nil {
		return nil, err
	}
	fmt.Printf("RGBToImg: to jpg size(%v bytes)\n", len(buffer.Bytes()))
	return buffer.Bytes(), nil
}

func parallel(start, stop int, f func(<-chan int)) {
	count := stop - start
	if count < 1 {
		return
	}

	procs := runtime.GOMAXPROCS(0)
	if procs > count {
		procs = count
	}

	c := make(chan int, count)
	for i := start; i < stop; i++ {
		c <- i
	}
	close(c)

	var wg sync.WaitGroup
	for i := 0; i < procs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f(c)
		}()
	}
	wg.Wait()
}

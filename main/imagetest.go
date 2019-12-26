package main

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
	"image"
	"image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"os"
)

func ReadOrientation(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("failed to open file, err: ", err)
		return 0
	}
	defer file.Close()

	x, err := exif.Decode(file)
	if err != nil {
		fmt.Println("failed to decode file, err: ", err)
		return 0
	}

	orientation, err := x.Get(exif.Orientation)
	if err != nil {
		fmt.Println("failed to get orientation, err: ", err)
		return 0
	}
	orientVal, err := orientation.Int(0)
	if err != nil {
		fmt.Println("failed to convert type of orientation, err: ", err)
		return 0
	}

	fmt.Println("the value of photo orientation is :", orientVal)
	return orientVal
}

func ReadOrientationExif(filename string) (*exif.Exif, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("failed to open file, err: ", err)
		return nil, 0
	}
	defer file.Close()

	x, err := exif.Decode(file)
	if err != nil {
		fmt.Println("failed to decode file, err: ", err)
		return nil, 0
	}

	orientation, err := x.Get(exif.Orientation)
	if err != nil {
		fmt.Println("failed to get orientation, err: ", err)
		return nil, 0
	}
	orientVal, err := orientation.Int(0)
	if err != nil {
		fmt.Println("failed to convert type of orientation, err: ", err)
		return nil, 0
	}

	//ff, err := os.Create("rgb.tst")
	//if err != nil {
	//	return nil, 0
	//}
	//defer ff.Close()
	//_, err = ff.Write(x.Raw)
	//if err != nil {
	//	return nil, 0
	//}
	fmt.Println("the value of photo orientation is :", orientVal)
	return x, orientVal
}
func main() {
	fn := "chenfen_female_40_L.jpg"
	//fn := "yangyanan_male_20_L.jpg"
	//fn := "20.jpg"
	//fn := "777.jpg"
	f, err := os.Open(fn)
	if err != nil {

		fmt.Println("ahdghjas:", err)
		return
	}
	defer f.Close()
	con, name, err := image.DecodeConfig(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("====w:%v, h:%v ,na:%v,err:%v\n", con.Width, con.Height, name, err)
	orientation := ReadOrientation(fn)
	fmt.Printf("===orientation:%v\n", orientation)

	fmt.Println("##################")
	x, _ := ReadOrientationExif(fn)
	fd := bytes.NewReader(x.Raw)
	con, name, err = image.DecodeConfig(fd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("====w:%v, h:%v ,na:%v,err:%v\n", con.Width, con.Height, name, err)
}

// ReadImage makes a copy of image (jpg,png or gif) and applies
// all necessary operation to reverse its orientation to 1
// The result is a image with corrected orientation and without
// exif data.
func ReadImage(fpath string) *image.Image {
	var img image.Image
	var err error
	// deal with image
	ifile, err := os.Open(fpath)
	if err != nil {
		//logrus.Warnf("could not open file for image transformation: %s", fpath)
		return nil
	}
	defer ifile.Close()
	filetype, err := GetSuffix(fpath)
	if err != nil {
		return nil
	}
	if filetype == "jpg" {
		img, err = jpeg.Decode(ifile)
		if err != nil {
			return nil
		}
	} else if filetype == "png" {
		img, err = png.Decode(ifile)
		if err != nil {
			return nil
		}
	} else if filetype == "gif" {
		img, err = gif.Decode(ifile)
		if err != nil {
			return nil
		}
	}
	// deal with exif
	efile, err := os.Open(fpath)
	if err != nil {
		//logrus.Warnf("could not open file for exif decoder: %s", fpath)
	}
	defer efile.Close()
	x, err := exif.Decode(efile)
	if err != nil {
		if x == nil {
			// ignore - image exif data has been already stripped
		}
		//logrus.Errorf("failed reading exif data in [%s]: %s", fpath, err.Error())
	}
	if x != nil {
		orient, _ := x.Get(exif.Orientation)
		if orient != nil {
			//logrus.Infof("%s had orientation %s", fpath, orient.String())
			img = reverseOrientation(img, orient.String())
		} else {
			//logrus.Warnf("%s had no orientation - implying 1", fpath)
			img = reverseOrientation(img, "1")
		}
		imaging.Save(img, fpath)
	}
	return &img
}
func GetSuffix(s string) (string, error) {
	return "jpg", nil
}

// reverseOrientation amply`s what ever operation is necessary to transform given orientation
// to the orientation 1
func reverseOrientation(img image.Image, o string) *image.NRGBA {
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
	//logrus.Errorf("unknown orientation %s, expect 1-8", o)
	return imaging.Clone(img)
}

package temp

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"net/http"
	"reflect"
	"time"
	"unsafe"
)

func Float32ToByte(fSlice []float32) ([]byte, error) {
	length := len(fSlice)
	if length == 0 {
		return nil, errors.New("param is null")
	}
	bys := make([]byte, 4)
	var buf bytes.Buffer
	for i := 0; i < length; i++ {
		bits := math.Float32bits(fSlice[i])
		binary.LittleEndian.PutUint32(bys, bits)
		buf.Write(bys)
	}
	return buf.Bytes(), nil
}

func ByteToFloat32(bys []byte) ([]float32, error) {
	length := len(bys)
	if length == 0 || (length%4) != 0 {
		//LOG.Error("param is null\n")
		return nil, errors.New("param is null")
	}
	rd := bytes.NewReader(bys)
	byte4 := make([]byte, 4)
	float32Num := length / 4
	fSlice := make([]float32, float32Num)
	for i := 0; i < float32Num; i++ {
		io.ReadFull(rd, byte4)
		//fmt.Println("byte4", byte4)
		tn := binary.LittleEndian.Uint32(byte4)
		fSlice[i] = math.Float32frombits(tn)
	}
	return fSlice, nil
}

// Flt2binary 将float slice转变为byte slice
//func Flt2binary(features []float32) ([]byte, error) {
//	buf := new(bytes.Buffer)
//	for idx := 0; idx < len(features); idx++ {
//		err := binary.Write(buf, binary.LittleEndian, features[idx])
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	return buf.Bytes(), nil
//}
//
//// Binary2flt 将byte slice转变为float slice
//func Binary2flt(features []byte) ([]float32, error) {
//	r := bytes.NewReader(features)
//	resVal := []float32{}
//	var feature float32
//	var err error
//	for {
//		err = binary.Read(r, binary.LittleEndian, &feature)
//		if err != nil {
//			break
//		}
//		resVal = append(resVal, feature)
//	}
//	if err != io.EOF {
//		return nil, err
//	}
//	return resVal, nil
//}

func ByteToFloat(bys []byte) ([]float32, error) {
	length := len(bys)
	if length == 0 || (length%4) != 0 {
		//LOG.Error("param is null\n")
		return nil, errors.New("param is null")
	}
	rd := bytes.NewReader(bys)
	byte4 := make([]byte, 4)
	float32Num := length / 4
	fSlice := make([]float32, float32Num)
	for i := 0; i < float32Num; i++ {
		io.ReadFull(rd, byte4)
		//fmt.Println("byte4", byte4)
		tn := binary.LittleEndian.Uint32(byte4)
		//fSlice[i] = math.Float32frombits(tn)
		fSlice = append(fSlice, math.Float32frombits(tn))
	}
	return fSlice, nil
}

//func ImgToRGB(imgSrc []byte) (rgb []byte, w int, h int, err error) {
//	if len(imgSrc) == 0 {
//		return nil, 0, 0, errors.New("ImgToRGB param error")
//	}
//	imgsReader := bytes.NewReader(imgSrc)
//	imgs, err := jpeg.Decode(imgsReader)
//	if err != nil {
//		return nil, 0, 0, err
//	}
//	bounds := imgs.Bounds()
//	m := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
//	draw.Draw(m, m.Bounds(), imgs, bounds.Min, draw.Src)
//	var buf bytes.Buffer
//	pixLen := len(m.Pix)
//	buf.Grow(pixLen * 3 / 4)
//	//buf.Grow(len(m.Pix) * 3 / 4)
//	//fmt.Println(">>>", pixLen)
//	//for i := 0; i < len(m.Pix); i += 4 {
//	for i := 0; i < pixLen; i += 4 {
//		buf.Write(m.Pix[i : i+3])
//	}
//	rgb = buf.Bytes()
//	return rgb, bounds.Max.X, bounds.Max.Y, nil
//}

func ImgToRGB(imgSrc []byte) (rgb []byte, w int, h int, err error) {
	//begin := time.Now()
	//defer func() {
	//	fmt.Printf("ImgToRGB End need time(%v ms)\n", time.Since(begin).Nanoseconds()/1e6)
	//}()
	if len(imgSrc) == 0 {
		return nil, 0, 0, errors.New("ImgToRGB param error")
	}
	//fmt.Printf("ImgToRGB, len(imageSrc)=%v\n", len(imgSrc))
	imgsReader := bytes.NewReader(imgSrc)
	_, a, err := image.DecodeConfig(imgsReader)
	//LOG.Infof("ImgToRGB, photo-format=%v\n", a)
	if err != nil {
		return nil, 0, 0, err
	}
	imgsReader.Seek(0, 0)
	//LOG.Infof("ImgToRGB, len(imageSrc)=%v\n", len(imgSrc))
	var imgs image.Image
	if a == "png" {
		imgs, err = png.Decode(imgsReader)
		if err != nil {
			return nil, 0, 0, err
		}
	} else {
		imgs, err = jpeg.Decode(imgsReader)
		if err != nil {
			return nil, 0, 0, err
		}
	}
	bounds := imgs.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(m, m.Bounds(), imgs, bounds.Min, draw.Src)
	var buf bytes.Buffer
	pixLen := len(m.Pix)
	buf.Grow(pixLen * 3 / 4)
	for i := 0; i < pixLen; i += 4 {
		buf.Write(m.Pix[i : i+3])
	}
	rgb = buf.Bytes()
	//fmt.Printf("ImgToRGB, len(rgb):%v,  width:%v,  height:%v\n", len(rgb), bounds.Max.X, bounds.Max.Y)
	return rgb, bounds.Max.X, bounds.Max.Y, nil
}

func ImgToRGBTmp(imgSrc []byte) (rgb []byte, w int, h int, err error) {
	//begin := time.Now()
	//defer func() {
	//	fmt.Printf("ImgToRGB End need time(%v ms)\n", time.Since(begin).Nanoseconds()/1e6)
	//}()
	if len(imgSrc) == 0 {
		return nil, 0, 0, errors.New("ImgToRGB param error")
	}
	//LOG.Infof("ImgToRGB, len(imageSrc)=%v\n", len(imgSrc))
	imgsReader := bytes.NewReader(imgSrc)
	_, format, err := image.DecodeConfig(imgsReader)
	//LOG.Infof("ImgToRGB, photo-format=%v\n", format)
	if err != nil {
		return nil, 0, 0, err
	}
	imgsReader.Seek(0, 0)
	//LOG.Infof("ImgToRGB, len(imageSrc)=%v\n", len(imgSrc))
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
	}
	bounds := imgs.Bounds()
	//m := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	m := image.NewRGBA(bounds)
	//draw.Draw(m, m.Bounds(), imgs, bounds.Min, draw.Src)
	draw.Draw(m, bounds, imgs, bounds.Min, draw.Src)
	//var buf bytes.Buffer
	pixLen := len(m.Pix)
	//fmt.Printf("=====%v\n", pixLen)
	//buf.Grow(pixLen * 3 / 4)
	tmp := make([]byte, pixLen*3/4)
	//tmp := []byte{}
	//buf.Grow(pixLen)
	//jNum := pixLen / 4
	k := 0
	for i := 0; i < pixLen; i += 4 {
		//buf.Write(m.Pix[i : i+3])
		//tmp = append(tmp, m.Pix[i:i+3]...)
		//k := 3 * j
		tmp[k] = m.Pix[i]
		tmp[k+1] = m.Pix[i+1]
		tmp[k+2] = m.Pix[i+2]
		k += 3
	}
	//rgb = buf.Bytes()
	//LOG.Debugf("ImgToRGB, len(rgb):%v,  width:%v,  height:%v\n", len(tmp), bounds.Max.X, bounds.Max.Y)
	return tmp, bounds.Max.X, bounds.Max.Y, nil
}

type photo struct {
	Image  []byte //`json:"image" binding:"gt=0"`
	UserID string `json:"userID" `
}
type req struct {
	Tm      time.Time  `json:"tm" binding:"gt"`
	Flg     bool       `json:"flg" binding:"required"`
	Url     string     `json:"url" binding:"min=2,required"`
	Num     int        `json:"num" binding:"eq=7"`
	Img     photo      `json:"img"`
	Urls    []string   `json:"urls" binding:"required"`
	Persons [][]string `json:"persons" binding:"gt=0,dive,len=1,dive,required"`
	Imgs    []string   //`json:"imgs" binding:"gt=0,dive,required"`
	//Images  []photo    `json:"images" binding:"gt=0,dive"` //gt=0,dive,dive,required//`json:"images" binding:"required"`
	//Score  float32  `json:"score" binding:"required"`
	//Num    int      `json:"num" binding:"required"`
}

type server struct {
}

func (s *server) BindCheck(c *gin.Context) {
	reqInfo := req{}
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		fmt.Printf(">>>>>>>>>>Error(%v)\n", err)
		c.JSON(200, gin.H{"errmsg": err.Error()})
		//return
	}
	fmt.Println("===reqInfo:", reqInfo)
	out, _ := json.Marshal(&reqInfo)
	fmt.Printf("===out:%v\n", string(out))
	c.JSON(200, gin.H{"text": "hello!!!"})
}

//////////////////////
// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func Append() {
	str := []string{"good", "morning"}
	str = append(str, str...)
	//fmt.Printf("%v,%v\n", str, len(str))
}

func Slice() {
	str := make([]string, 2)
	//str = append(str, "good", "morning")
	str[0] = "good"
	str[1] = "morning"
}

// Flt2binary 将float slice转变为byte slice
func Flt2binary(features []float32) ([]byte, error) {
	buf := new(bytes.Buffer)
	for idx := 0; idx < len(features); idx++ {
		err := binary.Write(buf, binary.LittleEndian, features[idx])
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

// Binary2flt 将byte slice转变为float slice
func Binary2flt(features []byte) ([]float32, error) {
	r := bytes.NewReader(features)
	resVal := []float32{}
	var feature float32
	var err error
	for {
		err = binary.Read(r, binary.LittleEndian, &feature)
		if err != nil {
			break
		}
		resVal = append(resVal, feature)
	}
	if err != io.EOF {
		return nil, err
	}
	return resVal, nil
}

func StringToBytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

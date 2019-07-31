package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
	"gocv.io/x/gocv"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Changeable interface {
	Set(x, y int, c color.Color)
}

func process(imgPath string) bool {
	src := gocv.IMRead(imgPath, gocv.IMReadFlag(1))
	img, _ := src.ToImage()
	// 获取rgba并将黑色替换为白色
	for i := 0; i < src.Cols(); i++ {
		for j := 0; j < src.Rows(); j++ {
			cl := img.At(i, j)
			r, g, b, _ := cl.RGBA()
			//fmt.Println(r, g, b, a)
			if r == 13107 && g == 13107 && b == 13107 {
				img.(Changeable).Set(i, j, color.RGBA{255, 255, 255, 255})
			}
		}
	}
	mat, _ := gocv.ImageToMatRGBA(img)
	// 双层获取目录
	proPath := filepath.Dir(imgPath)
	proPath = filepath.Dir(proPath)
	err := os.MkdirAll(proPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	b := gocv.IMWrite(filepath.Join(proPath, "process", filepath.Base(imgPath)), mat)
	return b
	// opencv处理图片
	//window1 := gocv.NewWindow("1")
	//window2 := gocv.NewWindow("2")
	//window3 := gocv.NewWindow("3")
	//window4 := gocv.NewWindow("4")
	//
	//window1.IMShow(mat)
	//
	//gray := gocv.NewMat()
	//gocv.CvtColor(mat, &gray, gocv.ColorBGRAToGray)
	//window2.IMShow(gray)
	//
	//binaryMat := gocv.NewMat()
	//gocv.Threshold(gray,&binaryMat,100,255,gocv.ThresholdBinaryInv | gocv.ThresholdOtsu)
	//window3.IMShow(binaryMat)
	//
	//verticalSize := binaryMat.Rows() / 20
	//verticalKernel := gocv.GetStructuringElement(gocv.MorphRect, image.Point{1,verticalSize})
	//gocv.Erode(binaryMat,&binaryMat,verticalKernel)
	//gocv.Dilate(binaryMat,&binaryMat,verticalKernel)
	//window4.IMShow(binaryMat)
	//
	//
	//gocv.WaitKey(0)
	//window1.Close()
	//window2.Close()
	//window3.Close()
	//window4.Close()
}

func main() {
	dirPath := "/home/batty/桌面/verify/raw"
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range dir {
		process(filepath.Join(dirPath, file.Name()))
	}

	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("mycode")
	proPath := "/home/batty/桌面/verify/process"
	dir, err = ioutil.ReadDir(proPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range dir {
		client.SetImage(filepath.Join(proPath, file.Name()))
		text, _ := client.Text()
		fmt.Println(text)
	}




}

package main

import (
	"os"
	"strconv"
	"fmt"
	"image"
	"gocv.io/x/gocv"
)


func getCameras(cam_nbs []int) []*gocv.VideoCapture {

	var cams []*gocv.VideoCapture

	for i := 0; i < len(cam_nbs); i++ {
		cam_nb := cam_nbs[i]
		cam, _ := gocv.OpenVideoCapture(cam_nb)
		cams = append(cams, cam)
	}

	return cams
}


func main() {

	var cam_nbs []int
	var mag int
	nb_args := len(os.Args)

	if nb_args != 1 {
		mag = 1
		for i := 1; i < nb_args; i++ {
			cam_nb, _ := strconv.Atoi(os.Args[i])
			cam_nbs = append(cam_nbs, cam_nb)
		}
	} else {
		fmt.Println("CMD [CAM #1] [CAM #2] ...")
		return
	}

	cams := getCameras(cam_nbs)

	var windows []*gocv.Window
	var img_ori, img_crop, img_mag []gocv.Mat
	var rects []image.Rectangle
	nb_cams := len(cams)

	for i := 0; i < nb_cams; i++ {
		s := strconv.Itoa(i)
		windows = append(windows, gocv.NewWindow(s))
		img_ori = append(img_ori, gocv.NewMat())
		img_crop = append(img_crop, gocv.NewMat())
		img_mag = append(img_mag, gocv.NewMat())

		defer cams[i].Close()
		defer windows[i].Close()
	}

	// mag := 2
	for i := 0; i < nb_cams; i++ {
		cams[i].Read(&img_ori[i])
		rects = append(rects, gocv.SelectROI("Select ROI", img_ori[i]))
		windows[i].ResizeWindow(mag * rects[i].Dx(), mag * rects[i].Dy())
	}

	for {
		for i := 0; i < nb_cams; i++ {
			cams[i].Read(&img_ori[i])

			img_crop[i] = img_ori[i].Region(rects[i])
			gocv.Resize(img_crop[i], &img_mag[i], image.Point{}, float64(mag), float64(mag), 1)

			windows[i].IMShow(img_mag[i])
			k := windows[i].WaitKey(1)
			if k == 113 { // 'q'
				goto QUIT_FOR
			}
		}
	}
	QUIT_FOR:
}

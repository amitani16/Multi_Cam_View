package main

import (
	"os"
	"strconv"
	"fmt"
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
	nb_args := len(os.Args)

	if nb_args != 1 {
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
	var img     []gocv.Mat
	nb_cams := len(cams)

	for i := 0; i < nb_cams; i++ {
		s := strconv.Itoa(i)
		windows = append(windows, gocv.NewWindow(s))
		img = append(img, gocv.NewMat())

		defer cams[i].Close()
		defer windows[i].Close()
	}

	for {
		for i := 0; i < nb_cams; i++ {
			cams[i].Read(&img[i])
			windows[i].IMShow(img[i])
			k := windows[i].WaitKey(1)
			if k == 113 { // 'q'
				goto FOR_QUIT
			}
		}
	}
	FOR_QUIT:
}

package main

import "github.com/poi5305/go-otto-opencv-wapper/otto_wrapper"

func main() {
	o := otto_wrapper.NewOttoFunctions()
	vm := o.GetVM()
	vm.Run(`
		var video = cvOpenVideoCapture(0);
		var window = cvNewWindow("Test GOCV with OTTO")
		var mat = cvNewMat()
		while(true) {
			cvVideoCaptureRead(video, mat)
			cvWindowIMShow(window, mat)
			cvWindowWaitKey(window, 1)
		}
	`)
}

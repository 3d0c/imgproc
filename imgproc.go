package imgproc

//#cgo pkg-config: --libs-only-L opencv
//#cgo CFLAGS: -Wno-error=unused-function
//#cgo LDFLAGS: -lopencv_imgproc -lopencv_core -lopencv_highgui
//#include "cv_handler.h"
import "C"

import (
	"log"
	"unsafe"
)

type Options struct {
	Base    *Source
	Scale   *Scale
	Crop    *Roi
	Format  string
	Method  int
	Quality int
}

func Proc(o *Options) []byte {
	if o.Crop != nil && o.Scale != nil {
		// if both options selected, crop will be first, the scale size will be calculated from cropped dimension
		roi := o.Crop.Calc(o.Base.Size())
		zoom := o.Scale.Size(&Dimension{roi.Width, roi.Height})

		return resize(o, zoom, roi)
	}

	if o.Crop != nil {
		return resize(o, nil, o.Crop.Calc(o.Base.Size()))
	}

	if o.Scale != nil {
		return resize(o, o.Scale.Size(o.Base.Size()), nil)
	}

	return resize(o, o.Base.Size(), nil)
}

func resize(o *Options, zoom *Dimension, roi *Rect) []byte {
	var data []byte

	result := C.resizer(
		(*C.Blob)(unsafe.Pointer(blobptr(o.Base))),
		(*C.PixelDim)(unsafe.Pointer(zoom)),
		C.int(o.Quality), C.int(o.Method), C.CString("."+o.Format),
		(*C.CvRect)(initCvRect(roi)),
	)

	if result != nil {
		length := result.length
		data = C.GoBytes(unsafe.Pointer(result.data), C.int(length))

		C.free(unsafe.Pointer(result.data))
		C.free(unsafe.Pointer(result))
	}

	return (data)
}

func blobptr(s *Source) *C.Blob {
	if s == nil {
		return nil
	}

	return &C.Blob{
		data:   (*C.uchar)(unsafe.Pointer(&s.Blob[0])),
		length: C.uint(s.BlobLen),
	}
}

func initCvRect(roi *Rect) *C.CvRect {
	if roi == nil {
		return nil
	}

	if roi.Width == 0 || roi.Height == 0 {
		log.Println("Wrong roi init for crop action, should contain width and height")
		return nil
	}

	return &C.CvRect{C.int(roi.X), C.int(roi.Y), C.int(roi.Width), C.int(roi.Height)}
}

package imgproc

import (
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	MAX_RATIO = 6
	MIN_RATIO = 0.08
)

type Dimension struct {
	Width  int
	Height int
}

type Scale struct {
	width  int
	height int
	maxdim float64
}

func NewScale(i interface{}) *Scale {
	v := i.(string)
	if v == "" {
		return nil
	}

	parts := strings.Split(v, "x")

	switch len(parts) {
	case 1:
		f, err := strconv.ParseFloat(parts[0], 32)
		if err != nil {
			log.Printf("Illegal scale option '%v'\n", v)
			return nil
		}

		return &Scale{maxdim: f}

	case 2:
		w, _ := strconv.Atoi(parts[0])
		h, _ := strconv.Atoi(parts[1])

		return &Scale{
			width:  w,
			height: h,
		}

	default:
		log.Println("Illegal scale option '%v'\n", v)
		return nil
	}

	log.Printf("Illegal size value = '%v'\n", v)

	return nil
}

func (this *Scale) Size(src *Dimension) *Dimension {
	var ratio, dstRatio float64

	ratio = float64(src.Width) / float64(src.Height)

	if this.width > 0 && this.height > 0 {
		dstRatio = float64(this.width) / float64(this.height)
	}

	if ratio > MAX_RATIO || ratio < MIN_RATIO {
		log.Printf("Aspect ratio is %.4f. Seems unreasonable to scale.\n", ratio)
		return nil
	}

	if this.height > 0 && this.width > 0 && dstRatio < ratio {
		this.height = 0
	}

	if this.height > 0 && this.width > 0 && dstRatio >= ratio {
		this.width = 0
	}

	if this.height > 0 && this.width == 0 {
		return &Dimension{
			Width:  int(math.Ceil(float64(this.height) * ratio)),
			Height: this.height,
		}
	}

	if this.height == 0 && this.width > 0 {
		return &Dimension{
			Width:  this.width,
			Height: int(math.Ceil(float64(this.width) / ratio)),
		}
	}

	if this.maxdim > 0 && this.maxdim < 1 {
		return &Dimension{
			Width:  int(math.Ceil(float64(src.Width) * this.maxdim)),
			Height: int(math.Ceil(float64(src.Height) * this.maxdim)),
		}
	}

	if this.maxdim >= 1 && src.Width > src.Height {
		this.width = int(this.maxdim)
		return this.Size(src)
	}

	if this.maxdim >= 1 && src.Width < src.Height {
		this.height = int(this.maxdim)
		return this.Size(src)
	}

	return &Dimension{Width: src.Width, Height: src.Height}
}

func (this Dimension) String() string {
	return strconv.Itoa(this.Width) + "x" + strconv.Itoa(this.Height)
}

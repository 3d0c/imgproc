package imgproc

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"mime"
	"os"
	"reflect"
)

type Source struct {
	BlobLen int
	Blob    []byte
	Imgcfg  image.Config
	imgtype string
}

func NewSource(i interface{}) *Source {
	source := &Source{}

	switch i.(type) {
	case string:
		return source.fromFile(i.(string))

	case []byte:
		return source.fromBytes(i.([]byte))

	default:
		log.Println("Unsupported type:", reflect.TypeOf(i))
	}

	return nil
}

func (this *Source) fromFile(filepath string) *Source {
	var err error

	r, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return nil
	}

	defer r.Close()

	this.Blob, err = ioutil.ReadAll(r)
	if err != nil {
		log.Println("Unable to read file:", filepath)
		return nil
	}

	this.BlobLen = len(this.Blob)

	this.Imgcfg, this.imgtype, err = image.DecodeConfig(bytes.NewReader(this.Blob))
	if err != nil {
		log.Printf("Unable to DecodeConfig() file %v. %v\n", filepath, err)
		return nil
	}

	return this
}

func (this *Source) fromBytes(b []byte) *Source {
	var err error

	this.Blob = b

	this.Imgcfg, this.imgtype, err = image.DecodeConfig(bytes.NewReader(this.Blob))
	if err != nil {
		log.Printf("Unable to DecodeConfig(). %v\n", err)
		return nil
	} else {
		log.Println("this is a", this.imgtype)
	}

	this.BlobLen = len(this.Blob)

	return this
}

func (this *Source) Config() *image.Config {
	return &this.Imgcfg
}

func (this *Source) Type() string {
	return this.imgtype
}

func (this *Source) Mime() string {
	return mime.TypeByExtension("." + this.Type())
}

func (this *Source) Size() *Dimension {
	return &Dimension{Width: this.Imgcfg.Width, Height: this.Imgcfg.Height}
}

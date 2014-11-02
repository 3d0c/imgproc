### Usage

```go
	package controllers

	import (
		img "github.com/3d0c/imgproc"
	)
	
	... some stuff ...
	

	base := img.NewSource(AppConfig.StorageFilePath(params["id"]))
	if base == nil {
		return http.StatusNotFound, []byte{}
	}

	target := &img.Options{
		Base:    base,
		// options.Scale is just a string like x800 or 800x or 0.5 
		Scale:   img.NewScale(options.Scale),
		Method:  3,
		Quality: options.Quality,
	}

	target.Format = typ.Format()

	w.Header().Set("Content-Type", typ.ContentType())

	// serve our resized image
	return http.StatusOK, img.Proc(target)

```

Here is a list of supported options for `Scale`:

Prototype: ([0-9]+x) or (x[0-9]+) or ([0-9]+) or (0.[0-9]+)
E.g.:  
	- `800x` scale to width 800px, height will be calculated  
	- `x600` scale to height 600px, width will be calculated  
	- `640` maximum dimension is 640px, e.g. original 1024x768 pixel image will be scaled to 640x480, same option applied for 900x1600 image results 360x640  
	- `0.5` 50% of original dimensions, e.g. 1024x768 = 512x384

`Quality`: (int) 0-100  

`Format`: (string) "png" or "jpg"

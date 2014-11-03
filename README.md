### Usage

#### Resize

Prepare `Options`

```go
	base := imgproc.NewSource("../some/path/xxx.jpg")
	if base == nil {
		return 
	}
	
	// Resize image to 50% of original
	target := &imgproc.Options{
		Base:    base,
		Scale:   imgproc.NewScale("0.5"),
		Method:  3,
		Format:  "jpg",
		Quality: 80,
	}
```

Run `Proc` function. If everything is ok, it will return resized image.

```go
	b := imgproc.Proc(target)
```

Here is a list of supported options for `Scale`:

Prototype:  `([0-9]+x) or (x[0-9]+) or ([0-9]+) or (0.[0-9]+)`  
E.g.:  

- `800x` scale to width 800px, height will be calculated  
- `x600` scale to height 600px, width will be calculated  
- `640` maximum dimension is 640px, e.g. original 1024x768 pixel image will be scaled to 640x480, same option applied for 900x1600 image results 360x640  
- `0.5` 50% of original dimensions, e.g. 1024x768 = 512x384

#### Crop

```go
	
	// Crop 100x100 pixel from center
	target := &imgproc.Options{
		Base:    base,
		Crop:    imgproc.NewRoi("center,100,100"),
		Method:  3,
		Format:  "jpg",
		Quality: 100,
	}
```

`NewRoi` supported options:

- `X,Y,width,height`
- `center|left|right|bleft|bright,width,height`

#### Crop and Resize

```go
	// Crop and resize, crop will be first
	target := &imgproc.Options{
		Base:    base,
		Crop:    imgproc.NewRoi("center,100,100"),
		Scale:   imgproc.NewScale("0.5"),
		Method:  3,
		Format:  "jpg",
		Quality: 100,
	}
```


package internal

import "image"

// This struct can contain additional information such as download error.
type DownloadedImage struct {
	Img image.Image
	Url string
}

package downloader

import (
	"crypto/tls"
	"errors"
	"fmt"
	"image"
	"math"
	"net"
	"net/http"
	"time"
)

var transport = &http.Transport{
	DialContext: (&net.Dialer{
		// In production, this would be limited to something like: 10 * time.Second
		Timeout: 0,
	}).DialContext,
	// Disabled SSL verification for faster downloads.
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}

var client = &http.Client{
	Transport: transport,
	// In production, this would be limited to something like: 10 * time.Second
	Timeout: 0,
}

func DownloadImage(url string) (img image.Image, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("downloading %s failed with error: %s", url, err.Error()))
	}

	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("downloading %s returned invalid response code: %d", url, resp.StatusCode))
		return nil, err
	}

	img, _, err = image.Decode(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("decoding %s failed with error: %s", url, err.Error()))
	}
	return img, nil
}

// This method can be used if retry is desired.
func DownloadImageWithRetry(url string, retry int) (img image.Image, err error) {
	if retry < 0 {
		err := errors.New(fmt.Sprintf("download %s retry failed", url))
		return nil, err
	}

	img, err = DownloadImage(url)
	if nil != err {
		// Progressive delay
		delay := time.Duration(math.Round(float64(60)/float64(retry*2))) * time.Second
		time.Sleep(delay)
		return DownloadImageWithRetry(url, retry-1)
	}
	return img, nil
}

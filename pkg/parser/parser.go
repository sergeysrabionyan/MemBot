package parser

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"net/url"
	"strings"
	"sync"
)

var waitGroup sync.WaitGroup

const (
	coordinateX = 115
	coordinateY = 295
)

func GetImageUrls(word string) ([]string, error) {
	urlChan := make(chan string, 5)
	x, y := getImageParseStartCoordinates()
	waitGroup.Add(5)
	searchUrl, err := prepareSearchUrl(word)
	if err != nil {
		return nil, err
	}
	for i := 0; i < 5; i++ {
		go parseImages(x, y, urlChan, searchUrl)
		x += 250
	}
	waitGroup.Wait()
	close(urlChan)
	var urls []string
	for imageUrl := range urlChan {
		urls = append(urls, imageUrl)
	}
	return urls, nil
}

func GetRandomImageUrl(word string) (string, error) {
	searchUrl, err := prepareSearchUrl(word)
	if err != nil {
		return "", err
	}
	//Todo реализовать алгоритм вычисления рандомных координат картинки
	randomX := 115.00
	randomY := 295.00
	return parseImage(randomX, randomY, searchUrl)
}

func prepareSearchUrl(word string) (string, error) {
	var Url *url.URL
	Url, err := url.Parse(GoogleImagesUrl)
	if err != nil {
		return "", err
	}
	Url.Path += GoogleSearchUrl
	parameters := url.Values{}
	//TODO заменить +мем на динамичный параметр/срез параметров
	parameters.Add("q", word+" "+"мем")
	parameters.Add("tbm", "isch")
	parameters.Add("sclient", "img")
	parameters.Add("start", "0")
	parameters.Add("source", "lnms")
	Url.RawQuery = parameters.Encode()

	return Url.String(), nil
}

func parseImage(x float64, y float64, url string) (string, error) {
	var capturedUrl string
	var b1 []byte

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	err := chromedp.Run(ctx,
		chromedp.Emulate(device.Reset),
		chromedp.EmulateViewport(1972, 2000),
		chromedp.Navigate(url),
		chromedp.MouseClickXY(x, y),
		chromedp.CaptureScreenshot(&b1),
		chromedp.Evaluate(`let ele = document.elementFromPoint(1507, 305); ele ? ele.getAttribute('src') : ''`, &capturedUrl),
	)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if strings.Contains(capturedUrl, "data:image/") {
		return "", nil
	}
	return capturedUrl, nil
}

func parseImages(x float64, y float64, urlChan chan<- string, url string) {
	defer waitGroup.Done()
	imageUrl, _ := parseImage(x, y, url)
	urlChan <- imageUrl
}

func getImageParseStartCoordinates() (float64, float64) {
	return coordinateX, coordinateY
}

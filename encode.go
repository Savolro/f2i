package main

import (
	"fmt"
	"image"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/Savolro/f2i/nrgba"
	"golang.org/x/image/bmp"
)

// Encode encodes file into multiple bmp images of given dimensions
func Encode(path string, dir string, width int, height int) error {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("opening file %s: %w", path, err)
	}
	count := 0
	// TODO: add concurency here
	for {
		buf := make([]byte, width*height*4)
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		// TODO: adjust number of zeros to the number of images accordingly
		imgFileName := fmt.Sprintf("%s/img_%09d.bmp", dir, count)
		imgF, err := os.Create(imgFileName)
		if err != nil {
			return fmt.Errorf("creating file: %w", err)
		}
		img, err := nrgba.Encode(buf, width, height)
		if err != nil {
			imgF.Close()
			return fmt.Errorf("encoding data to image: %w", err)
		}
		err = bmp.Encode(imgF, img)
		imgF.Close()
		if err != nil {
			return fmt.Errorf("writing image to file: %w", err)
		}
		count++
	}
	return nil
}

// Decode decodes images in directory into a file
func Decode(path string, dir string) error {
	fileNames, err := imageFileNames(dir)
	if err != nil {
		return fmt.Errorf("retrieving image file nanes: %w", err)
	}
	destF, err := os.Create(path)
	defer destF.Close()
	if err != nil {
		return fmt.Errorf("creating file %s: %w", path, err)
	}

	for i, name := range fileNames {
		imgF, err := os.Open(name)
		if err != nil {
			return fmt.Errorf("opening image file %s: %w", name, err)
		}
		img, err := bmp.Decode(imgF)
		if err != nil {
			imgF.Close()
			return fmt.Errorf("decoding BMP image %s: %w", name, err)
		}
		nrgbaImg, ok := img.(*image.NRGBA)
		if !ok {
			return fmt.Errorf("image of unexpected type %s", reflect.TypeOf(img))
		}
		trimZeros := i == len(fileNames)-1
		data, err := nrgba.Decode(nrgbaImg, trimZeros)
		if err != nil {
			return fmt.Errorf("decoding image %s: %w", name, err)
		}
		destF.Write(data)
	}

	return nil
}

func imageFileNames(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, fmt.Errorf("opening directory %s: %w", dir, err)
	}
	fis, err := f.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("reading file directory %s: %w", dir, err)
	}
	var names []string
	for _, finfo := range fis {
		name := finfo.Name()
		// TODO: Check full name pattern
		if strings.HasSuffix(name, ".bmp") {
			names = append(names, dir+"/"+finfo.Name())
		}
	}
	sort.Strings(names)
	return names, nil
}

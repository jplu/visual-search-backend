package uc

import (
	"bytes"
	"github.mpi-internal.com/leboncoin-lab/visual-search-backend/domain"
	"io"
	"mime/multipart"
	"net/http"
)

const (
	// JPEGMime is the content-type related to a .jpg or .jpeg picture
	JPEGMime = "image/jpeg"
)

func (i Interactor) GetSimilarImages(fd multipart.File) (domain.Neighbors, error) {
	var err error
	var ct string
	var features []float32

	// Detect and validate content type
	if ct, err = detectContentType(fd); err != nil {
		return nil, err
	}

	if ct != JPEGMime {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(fd); err != nil {
		return nil, err
	}

	if features, err = i.Image.GetFeatures(buf.Bytes()); err != nil {
		return nil, err
	}

	return i.Ann.GetTopK(features, 10)
}

func detectContentType(r io.ReadSeeker) (string, error) {
	var err error
	buf := make([]byte, 512)
	_, err = r.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}
	if _, err = r.Seek(0, 0); err != nil {
		return "", err
	}
	return http.DetectContentType(buf), nil
}
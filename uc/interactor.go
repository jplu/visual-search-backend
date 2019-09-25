package uc

import (
	"github.com/jplu/visual-search-backend/domain"
	"mime/multipart"
)

//LogicHandler is the interface that must implement the struct holding the usecases
type LogicHandler interface {
	GetSimilarImages(fd multipart.File) (domain.Neighbors, error)
}

// Interactor is a struct holding the various interfaces to use in the usecases
type Interactor struct {
	Image Image
	Ann   Ann
}

// NewInteractor will create a new interactor
func NewInteractor(i Image, a Ann) Interactor {
	return Interactor {
		Image: i,
		Ann: a,
	}
}

// Image is an interface to manipulate images
type Image interface {
	GetFeatures(img []byte) ([]float32, error)
}

// Ann is an interface to manipulate approximate nearest neighbors frameworks
type Ann interface {
	GetTopK(featurizedImage []float32, k uint64) (domain.Neighbors, error)
}

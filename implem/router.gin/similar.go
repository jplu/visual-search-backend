package router

import (
	"github.com/gin-gonic/gin"
	"github.mpi-internal.com/leboncoin-lab/visual-search-backend/domain"
	"mime/multipart"
	"net/http"
)

// Similar allows to retrieve similar images
func (r Router) Similar(c *gin.Context) {
	var err error
	var fd multipart.File
	var fh *multipart.FileHeader
	var neighbors domain.Neighbors

	if fh, err = c.FormFile("picture"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if fd, err = fh.Open(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if neighbors, err = r.interactor.GetSimilarImages(fd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, neighbors)
}

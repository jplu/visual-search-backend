package faiss

import (
	"context"
	"github.com/jplu/visual-search-backend/domain"
	"github.com/jplu/visual-search-backend/uc"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

// Faiss type
type Faiss struct {
	client FaissServiceClient
}

// NewFaissClient creates a new gRPC client
func NewFaissClient(addr string, port int) (uc.Ann, error) {
	servingAddress := addr + ":" + strconv.Itoa(port)
	conn, err := grpc.Dial(servingAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Cannot connect to the grpc server: %v\n", err)
	}

	client := NewFaissServiceClient(conn)

	return Faiss{client: client}, nil
}

func (f Faiss) GetTopK(featurizedImage []float32, k uint64) (domain.Neighbors, error) {
	vec := Vector{FloatVal: featurizedImage}
	req := SearchRequest{Vector: &vec, TopK: k}
	tt, err := f.client.Search(context.Background(), &req)

	if err != nil {
		log.Fatalf("%v.HeartBeat(_) = _, %v", f.client, err)
	}

	return toDomainNeighbors(tt.GetNeighbors())
}

func toDomainNeighbors(neighbors []*Neighbor) (domain.Neighbors, error) {
	newNeighbors := make([]domain.Neighbor, 0)

	for _, neighbor := range neighbors {
		newNeighbors = append(newNeighbors, domain.Neighbor{ID: neighbor.Id, Score: neighbor.Score})
	}

	return newNeighbors, nil
}

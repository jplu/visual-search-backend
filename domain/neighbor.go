package domain

// Neighbor is the format that represents a neighbor in FAISS
type Neighbor struct {
	ID  uint64    `json:"id"`
	Score float32 `json:"score"`
}

// Neighbors represents a list of Neighbor
type Neighbors []Neighbor
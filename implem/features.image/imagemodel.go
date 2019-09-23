package image

import (
	"context"
	tf "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"
	pb "github.mpi-internal.com/leboncoin-lab/visual-search-backend/proto/tensorflow/serving"
	"github.mpi-internal.com/leboncoin-lab/visual-search-backend/uc"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

// ImageModel type
type ImageModel struct {
	client pb.PredictionServiceClient
}

// NewImageModelClient creates a new gRPC client
func NewImageModelClient(addr string, port int) (uc.Image, error) {
	servingAddress := addr + ":" + strconv.Itoa(port)
	log.Println(servingAddress)
	conn, err := grpc.Dial(servingAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Cannot connect to the grpc server: %v\n", err)
	}

	client := pb.NewPredictionServiceClient(conn)

	return &ImageModel{client: client}, nil
}

// GetFeatures returns the features corresponding to the given image
func (im ImageModel) GetFeatures(img []byte) ([]float32, error) {
	request := &pb.PredictRequest {
		ModelSpec: &pb.ModelSpec {
			Name:          "resnet18",
			SignatureName: "serving_default",
		},
		Inputs: map[string]*tf.TensorProto {
			"image": &tf.TensorProto {
				Dtype: tf.DataType_DT_STRING,
				TensorShape: &tf.TensorShapeProto {
					Dim: []*tf.TensorShapeProto_Dim {
						&tf.TensorShapeProto_Dim {
							Size: int64(1),
						},
						&tf.TensorShapeProto_Dim {
							Size: int64(1),
						},
					},
				},
				StringVal: [][]byte{img},
			},
		},
	}
	resp, err := im.client.Predict(context.Background(), request)
	if err != nil {
		log.Fatalln(err)
	}

	tp, ok := resp.Outputs["features"]

	if !ok {
		log.Fatalln(err)
	}

	return tp.FloatVal, nil
}
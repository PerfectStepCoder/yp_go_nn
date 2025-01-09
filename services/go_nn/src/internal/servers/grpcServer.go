package servers

import (
	"context"
	"fmt"
	//"fmt"
	"log"
	"net"

	"github.com/PerfectStepCoder/yp_go_nn/src/internal/engine"
	pb "github.com/PerfectStepCoder/yp_go_nn/src/internal/proto/gen"
	"google.golang.org/grpc"
	//"github.com/PerfectStepCoder/yp_go_nn/src/internal/storage"
	//"google.golang.org/grpc/metadata"
	//"google.golang.org/protobuf/types/known/emptypb"
)

// ServerGRPC поддерживает все необходимые методы сервера.
type ServerGRPC struct {
	nn *engine.OnnxNeuralNetwork
	pb.UnimplementedClassifyNNServer
	server *grpc.Server 
}

// CreateTask - отправляем изображение для классификации.
func (s *ServerGRPC) CreateOneTask(ctx context.Context, in *pb.TaskOneRequest) (*pb.TaskOneResponse, error) {
	var response pb.TaskOneResponse

	image, err := engine.BytesToFloat32Matrix(in.Image, int(in.Height), int(in.Width))

	if err != nil {
		return nil, err
	}

	labelClassNames, err := s.nn.Detect(image)

	response.TaskUID = in.TaskUID
	response.ClassName = labelClassNames[0]

	return &response, nil
}

// CreateBatchTask - отправляем изображение для классификации.
func (s *ServerGRPC) CreateBatchTask(ctx context.Context, in *pb.TaskBatchRequest) (*pb.TaskBatchResponse, error) {
	var response pb.TaskBatchResponse

	images := make([][]float32, len(in.Images))

	for i , image := range in.Images {
		img, err := engine.BytesToFloat32Slice(image)
		if err != nil {
			fmt.Println(err)
		}
		images[i] = img
	}
	
	labelClassNames, err := s.nn.Detect(images)
	if err != nil {
		fmt.Println(err)
	}

	response.TaskUID = in.TaskUID
	response.ClassNames = labelClassNames

	return &response, nil
}

// CreateBatchCodeTask - отправляем батч изображений для классификации.
func (s *ServerGRPC) CreateBatchCodeTask(ctx context.Context, in *pb.TaskBatchRequest) (*pb.TaskBatchCodeResponse, error) {
	var response pb.TaskBatchCodeResponse

	images := make([][]float32, len(in.Images))

	for i , image := range in.Images {
		img, err := engine.BytesToFloat32Slice(image)
		if err != nil {
			fmt.Println(err)
		}
		images[i] = img
	}
	
	labelClassCodes, err := s.nn.DetectCode(images)
	if err != nil {
		fmt.Println(err)
	}
	
	response.TaskUID = in.TaskUID
	response.ClassCodes = engine.IntToInt32Slice(labelClassCodes)

	return &response, nil
}

func NewServerGRPC(nn *engine.OnnxNeuralNetwork) (*ServerGRPC, error) {
	
	return &ServerGRPC{
		nn: nn,
	}, nil
}

func (s *ServerGRPC) Start(addr string) error {
	
	// Cоздаём gRPC-сервер без зарегистрированной службы
	s.server = grpc.NewServer()
		
	// Регистрируем сервис
	pb.RegisterClassifyNNServer(s.server, s)

	listen, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalln(err)
	}

	return s.server.Serve(listen)
}

func (s *ServerGRPC) Stop(ctx context.Context) error {
	// Остановка сервера с плавным завершением
	s.server.GracefulStop()
	return nil
}

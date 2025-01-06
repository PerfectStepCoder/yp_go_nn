package engine

import (
	"fmt"
	"log"
	"os"
	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"gorgonia.org/tensor"
)

type AgentNetWork struct {
	ModelFile []byte
	ModelBackend *gorgonnx.Graph
	ModelOnnx *onnx.Model
}

func NewAgentNetWork(pathToonnxFile string) *AgentNetWork {
	//agentRealUID, _ := uuid.Parse(agentUID)
	// Открываем файл модели
	modelFile, err := os.ReadFile(pathToonnxFile)
	if err != nil {
		log.Fatalf("Ошибка при открытии модели: %v", err)
	}

	// Загружаем модель ONNX
	backend := gorgonnx.NewGraph()
	modelOnnx := onnx.NewModel(backend)

	if err := modelOnnx.UnmarshalBinary(modelFile); err != nil {
		log.Fatalf("Ошибка при загрузке модели: %v", err)
	}

	return &AgentNetWork{
		ModelFile: modelFile,
		ModelOnnx: modelOnnx,
		ModelBackend: backend,
	}
}

func (a *AgentNetWork) Predict(inputData []float32) {
	
	fmt.Println("inputData:", inputData)

	input := tensor.New(tensor.WithShape(1, 1, 28, 28), tensor.WithBacking(inputData))
	
	// Running inference
	err := a.ModelOnnx.SetInput(0, input)
	if err != nil {
		log.Fatalf("Error setting input: %v", err)
	}

	err = a.ModelBackend.Run()
	if err != nil {
		log.Fatalf("Error running inference: %v", err)
	}

	// Processing output
	outputs, err := a.ModelOnnx.GetOutputTensors()
	if err != nil {
		log.Fatalf("Error getting output: %v", err)
	}

	fmt.Println("Outputs: ", outputs)

}

func (a *AgentNetWork) Close() {
	//a.ModelFile.Close();
}

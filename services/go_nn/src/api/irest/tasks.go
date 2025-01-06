package httpp // http protocol

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"github.com/PerfectStepCoder/yp_go_nn/src/internal/engine"

	//"github.com/PerfectStepCoder/yp_go_nn/src/internal/models"
	"github.com/PerfectStepCoder/yp_go_nn/src/internal/security"
	"github.com/PerfectStepCoder/yp_go_nn/src/internal/storage"
	//"github.com/go-chi/chi/v5"
)

// @Summary      Task
// @Description  Create Task with one image
// @Tags         Tasks
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true "File image"
// @Security     BearerAuth
// @Success      200   body  models.TaskOneRequest true "Task one image"
// @Failure      401  {string}  string  "Unauthorized"
// @Router       /tasks/one  [post]
func TaskOneHandler(mainStorage storage.Storage, nn *engine.OnnxNeuralNetwork) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// Извлечение пользователя из контекста
		_, ok := req.Context().Value(OperatorKey).(*security.OperatorJWT)
		if !ok {
			http.Error(res, "Operator not found in context", http.StatusUnauthorized)
			return
		}

		// Попытка извлечь файл из запроса
		file, fileHeader, err := req.FormFile("file")
		if err != nil {
			http.Error(res, "Did not catch the file: " + err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close() // Закрыть файл после обработки

		// Вывод информации о загруженном файле
		fmt.Fprintf(res, "Загружен файл: %s (размер: %d байт)\n", fileHeader.Filename, fileHeader.Size)


		// Cериализуем ответ сервера
		// jsonResp, err := json.Marshal(result)
		// if err != nil {
		// 	http.Error(res, "Error writing response", http.StatusInternalServerError)
		// 	return
		// }

		//res.Write(jsonResp)

		res.WriteHeader(http.StatusOK)

	}
}

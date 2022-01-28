/*
 * Model Server
 *
 * This is API spec of model server
 *
 * API version: 0.0.1
 * Contact: hello@instill.tech
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package handlers

import (
	"archive/zip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/instill-ai/model-backend/configs"
	"github.com/instill-ai/model-backend/internal-protogen-go/model"
	"github.com/instill-ai/model-backend/internal/triton"
	"github.com/instill-ai/model-backend/pkg/db"
	"github.com/instill-ai/model-backend/pkg/models"
)

type ServiceHandlers struct{}

func _isTritonServerReady() bool {
	serverLiveResponse := triton.ServerLiveRequest(triton.TritonClient)
	if serverLiveResponse == nil {
		return false
	}
	fmt.Printf("Triton Health - Live: %v\n", serverLiveResponse.Live)
	if !serverLiveResponse.Live {
		return false
	}

	serverReadyResponse := triton.ServerReadyRequest(triton.TritonClient)
	fmt.Printf("Triton Health - Ready: %v\n", serverReadyResponse.Ready)
	return serverReadyResponse.Ready
}

func _createModelResponse(modelInDB models.Model, versions []models.Version) *model.CreateModelResponse {
	var mRes model.CreateModelResponse
	mRes.Name = modelInDB.Name
	mRes.Id = modelInDB.Id
	mRes.Optimized = modelInDB.Optimized
	mRes.Description = modelInDB.Description
	mRes.Framework = modelInDB.Framework
	mRes.CreatedAt = &model.Timestamp{Timestamp: timestamppb.New(modelInDB.CreatedAt)}
	mRes.UpdatedAt = &model.Timestamp{Timestamp: timestamppb.New(modelInDB.UpdatedAt)}
	mRes.Organization = modelInDB.Organization
	mRes.Icon = modelInDB.Icon
	mRes.Type = modelInDB.Type
	mRes.Visibility = modelInDB.Visibility
	var vers []*model.ModelVersion
	for i := 0; i < len(versions); i++ {
		vers = append(vers, &model.ModelVersion{
			Version:     versions[i].Version,
			ModelId:     versions[i].ModelId,
			Description: versions[i].Description,
			CreatedAt:   &model.Timestamp{Timestamp: timestamppb.New(versions[i].CreatedAt)},
			UpdatedAt:   &model.Timestamp{Timestamp: timestamppb.New(versions[i].UpdatedAt)},
		})
	}
	mRes.Versions = vers
	return &mRes
}

//writeToFp takes in a file pointer and byte array and writes the byte array into the file
//returns error if pointer is nil or error in writing to file
func _writeToFp(fp *os.File, data []byte) error {
	w := 0
	n := len(data)
	for {

		nw, err := fp.Write(data[w:])
		if err != nil {
			return err
		}
		w += nw
		if nw >= n {
			return nil
		}
	}
}

func _unzip(filePath string, dstDir string, dstName string) (string, bool) {
	archive, err := zip.OpenReader(filePath)
	if err != nil {
		fmt.Println("Error when open zip file ", err)
		return "", false
	}
	defer archive.Close()
	topDir := true
	var modelDirName string
	for _, f := range archive.File {
		filePath := filepath.Join(dstDir, f.Name)
		fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dstDir)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return "", false
		}
		if f.FileInfo().IsDir() {
			if topDir {
				topDir = false
				modelDirName = f.Name
			}
			fmt.Println("creating directory... ", filePath)
			_ = os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return "", false
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return "", false
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return "", false
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return "", false
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	if string(modelDirName[len(modelDirName)-1]) == "/" {
		modelDirName = modelDirName[:len(modelDirName)-1]
	}
	return modelDirName, true
}

func _saveFile(stream model.Model_CreateModelServer) (outFile string, createdModel *models.Model, err error) {
	firstChunk := true
	var fp *os.File

	var fileData *model.CreateModelRequest

	var tmpFile string

	newModel := &models.Model{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Icon:         "",
		Organization: "domain@instill.tech",
		Author:       "local-user@instill.tech",
	}

	for {
		fileData, err = stream.Recv() //ignoring the data  TO-Do save files received
		if err != nil {
			fmt.Println("????? err", err)
			if err == io.EOF {
				break
			}

			err = errors.Wrapf(err,
				"failed unexpectadely while reading chunks from stream")
			return "", nil, err
		}

		if firstChunk { //first chunk contains file name
			newModel.Name = fileData.Name
			newModel.Description = fileData.Description
			newModel.Optimized = fileData.Optimized
			newModel.Type = fileData.Type
			newModel.Framework = fileData.Framework
			newModel.Visibility = fileData.Visibility
			if fileData.Filename != "" {
				tmpFile = path.Join("/tmp", filepath.Base(fileData.Filename))
				fp, err = os.Create(tmpFile)
				if err != nil {
					return "", nil, err
				}
				defer fp.Close()
			} else {
				return "", nil, errors.Errorf("No filename")
			}
			firstChunk = false
		}
		err = _writeToFp(fp, fileData.Content)
		if err != nil {
			return "", nil, err
		}
	}
	return tmpFile, newModel, nil
}

func _savePredictFile(stream model.Model_PredictModelServer) (imageFile string, modelId string, version int, modelType string, err error) {
	firstChunk := true
	var fp *os.File

	var fileData *model.PredictModelRequest

	var tmpFile string

	for {
		fileData, err = stream.Recv() //ignoring the data  TO-Do save files received
		if err != nil {
			if err == io.EOF {
				break
			}

			err = errors.Wrapf(err,
				"failed unexpectadely while reading chunks from stream")
			return "", "", -1, "", nil
		}

		if firstChunk { //first chunk contains file name
			modelId = fileData.ModelId
			version = int(fileData.ModelVersion)
			modelType = fileData.ModelType

			tmpFile = path.Join("/tmp/", uuid.New().String())
			fp, err = os.Create(tmpFile)
			if err != nil {
				return "", "", -1, "", err
			}
			defer fp.Close()

			firstChunk = false
		}
		err = _writeToFp(fp, fileData.Content)
		if err != nil {
			return "", "", -1, "", err
		}
	}
	return tmpFile, modelId, version, modelType, nil
}

func _makeError(statusCode codes.Code, title string, detail string, duration float64) error {
	err := &models.Error{
		Status: int32(statusCode),
		Title:  title,
		Detail: detail,
	}
	data, _ := json.Marshal(err)
	return status.Error(statusCode, string(data))
}

// AddModel - upload a model to the model server
func (s *ServiceHandlers) CreateModel(stream model.Model_CreateModelServer) (err error) {
	start := time.Now()

	tmpFile, newModel, err := _saveFile(stream)
	if err != nil {
		return _makeError(400, "Save File Error", err.Error(), float64(time.Since(start).Milliseconds()))
	}

	// extract zip file from tmp to models directory
	modelDirName, isOk := _unzip(tmpFile, configs.Config.TritonServer.ModelStore, newModel.Id)
	if !isOk {
		return _makeError(400, "Save File Error", "Could not extract zip file", float64(time.Since(start).Milliseconds()))
	}
	newModel.Id = modelDirName

	result := db.DB.Create(&newModel)
	if result.Error != nil {
		return _makeError(400, "Add Model Error", result.Error.Error(), float64(time.Since(start).Milliseconds()))
	}

	newVersion := models.Version{
		Version:     1,
		ModelId:     newModel.Id,
		Description: newModel.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Status:      "offline",
		Metadata:    models.JSONB{},
	}
	result = db.DB.Create(&newVersion)
	if result.Error != nil {
		return _makeError(500, "Add Model Error", result.Error.Error(), float64(time.Since(start).Milliseconds()))
	}

	var modelInDB models.Model
	result = db.DB.Model(&models.Model{}).Where("id", newModel.Id).First(&modelInDB)
	if result.Error != nil {
		return _makeError(500, "Add Model Error", result.Error.Error(), float64(time.Since(start).Milliseconds()))
	}

	resp := _createModelResponse(modelInDB, []models.Version{newVersion})

	err = stream.SendAndClose(resp)
	if err != nil {
		return _makeError(500, "Add Model Error", err.Error(), float64(time.Since(start).Milliseconds()))
	}

	return
}

func (s *ServiceHandlers) LoadModel(ctx context.Context, in *model.LoadModelRequest) (*model.LoadModelResponse, error) {
	fmt.Println("Load model ", in)

	if !_isTritonServerReady() {
		return &model.LoadModelResponse{}, nil
	}
	loadModelResponse := triton.LoadModelRequest(triton.TritonClient, in.ModelId)
	fmt.Println(loadModelResponse)

	return &model.LoadModelResponse{}, nil
}

func (s *ServiceHandlers) UnloadModel(ctx context.Context, in *model.UnloadModelRequest) (*model.UnloadModelResponse, error) {
	fmt.Println("UnloadModel model ", in)

	if !_isTritonServerReady() {
		return &model.UnloadModelResponse{}, nil
	}
	unloadModelResponse := triton.UnloadModelRequest(triton.TritonClient, in.ModelId)
	fmt.Println(unloadModelResponse)

	return &model.UnloadModelResponse{}, nil
}

func (s *ServiceHandlers) ListModels(ctx context.Context, in *model.ListModelRequest) (*model.ListModelResponse, error) {
	fmt.Println("ListModels model ", in)

	if !_isTritonServerReady() {
		return &model.ListModelResponse{}, nil
	}

	listModelsResponse := triton.ListModelsRequest(triton.TritonClient)
	fmt.Println("listModelsResponse ", listModelsResponse)

	var resModels []*model.CreateModelResponse
	models := listModelsResponse.Models
	for i := 0; i < len(models); i++ {
		md := model.CreateModelResponse{
			Id:   models[i].Name,
			Name: models[i].Name,
		}
		resModels = append(resModels, &md)
	}
	fmt.Println("resModels ", resModels)
	return &model.ListModelResponse{Models: resModels}, nil
}

func (s *ServiceHandlers) PredictModel(stream model.Model_PredictModelServer) error {
	start := time.Now()

	if !_isTritonServerReady() {
		return _makeError(500, "PredictModel", "Triton Server not ready yet", float64(time.Since(start).Milliseconds()))
	}

	imageFile, modelId, version, modelType, err := _savePredictFile(stream)

	modelMetadataResponse := triton.ModelMetadataRequest(triton.TritonClient, modelId, fmt.Sprint(version))
	if modelMetadataResponse == nil {
		return _makeError(400, "PredictModel", "Model not found", float64(time.Since(start).Milliseconds()))
	}

	modelConfigResponse := triton.ModelConfigRequest(triton.TritonClient, modelId, fmt.Sprint(version))
	if modelMetadataResponse == nil {
		return _makeError(400, "PredictModel", "Model config not found", float64(time.Since(start).Milliseconds()))
	}

	if err != nil {
		return _makeError(500, "PredictModel", "Could not save file", float64(time.Since(start).Milliseconds()))
	}

	fmt.Println(modelMetadataResponse)

	// err = stream.SendAndClose(&model.PredictModelResponse{})

	input, err := triton.Preprocess(modelType, imageFile, modelMetadataResponse, modelConfigResponse)
	if err != nil {
		return _makeError(400, "PredictModel", err.Error(), float64(time.Since(start).Milliseconds()))
	}

	// /* We use a simple model that takes 2 input tensors of 16 integers
	// each and returns 2 output tensors of 16 integers each. One
	// output tensor is the element-wise sum of the inputs and one
	// output is the element-wise difference. */
	inferResponse := triton.ModelInferRequest(triton.TritonClient, modelType, input, modelId, fmt.Sprint(version), modelMetadataResponse, modelConfigResponse)

	// /* We expect there to be 2 results (each with batch-size 1). Walk
	// over all 16 result elements and print the sum and difference
	// calculated by the model. */
	postprocessResponse := triton.Postprocess(modelType, inferResponse, modelMetadataResponse)

	err = stream.SendAndClose(&model.PredictModelResponse{Data: postprocessResponse})

	return err
}

func HandleUploadOutput(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	contentType := r.Header.Get("Content-Type")

	if contentType == "multipart/form-data" {
	} else {
		w.Header().Add("Content-Type", "application/json+problem")
		w.WriteHeader(405)
	}
}

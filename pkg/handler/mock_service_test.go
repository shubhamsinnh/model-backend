// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/instill-ai/model-backend/pkg/service (interfaces: Service)

// Package handler_test is a generated GoMock package.
package handler_test

import (
	context "context"
	reflect "reflect"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	redis "github.com/go-redis/redis/v9"
	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	resource "github.com/instill-ai/model-backend/internal/resource"
	acl "github.com/instill-ai/model-backend/pkg/acl"
	datamodel "github.com/instill-ai/model-backend/pkg/datamodel"
	repository "github.com/instill-ai/model-backend/pkg/repository"
	service "github.com/instill-ai/model-backend/pkg/service"
	utils "github.com/instill-ai/model-backend/pkg/utils"
	taskv1alpha "github.com/instill-ai/protogen-go/common/task/v1alpha"
	mgmtv1beta "github.com/instill-ai/protogen-go/core/mgmt/v1beta"
	modelv1alpha "github.com/instill-ai/protogen-go/model/model/v1alpha"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AuthenticateUser mocks base method.
func (m *MockService) AuthenticateUser(arg0 context.Context, arg1 bool) (*service.AuthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateUser", arg0, arg1)
	ret0, _ := ret[0].(*service.AuthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateUser indicates an expected call of AuthenticateUser.
func (mr *MockServiceMockRecorder) AuthenticateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateUser", reflect.TypeOf((*MockService)(nil).AuthenticateUser), arg0, arg1)
}

// ConvertOwnerNameToPermalink mocks base method.
func (m *MockService) ConvertOwnerNameToPermalink(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertOwnerNameToPermalink", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertOwnerNameToPermalink indicates an expected call of ConvertOwnerNameToPermalink.
func (mr *MockServiceMockRecorder) ConvertOwnerNameToPermalink(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertOwnerNameToPermalink", reflect.TypeOf((*MockService)(nil).ConvertOwnerNameToPermalink), arg0)
}

// ConvertOwnerPermalinkToName mocks base method.
func (m *MockService) ConvertOwnerPermalinkToName(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertOwnerPermalinkToName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertOwnerPermalinkToName indicates an expected call of ConvertOwnerPermalinkToName.
func (mr *MockServiceMockRecorder) ConvertOwnerPermalinkToName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertOwnerPermalinkToName", reflect.TypeOf((*MockService)(nil).ConvertOwnerPermalinkToName), arg0)
}

// ConvertRepositoryNameToRscName mocks base method.
func (m *MockService) ConvertRepositoryNameToRscName(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertRepositoryNameToRscName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertRepositoryNameToRscName indicates an expected call of ConvertRepositoryNameToRscName.
func (mr *MockServiceMockRecorder) ConvertRepositoryNameToRscName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertRepositoryNameToRscName", reflect.TypeOf((*MockService)(nil).ConvertRepositoryNameToRscName), arg0)
}

// CreateModelPrediction mocks base method.
func (m *MockService) CreateModelPrediction(arg0 context.Context, arg1 *datamodel.ModelPrediction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateModelPrediction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateModelPrediction indicates an expected call of CreateModelPrediction.
func (mr *MockServiceMockRecorder) CreateModelPrediction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateModelPrediction", reflect.TypeOf((*MockService)(nil).CreateModelPrediction), arg0, arg1)
}

// CreateModelVersionAdmin mocks base method.
func (m *MockService) CreateModelVersionAdmin(arg0 context.Context, arg1 *datamodel.ModelVersion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateModelVersionAdmin", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateModelVersionAdmin indicates an expected call of CreateModelVersionAdmin.
func (mr *MockServiceMockRecorder) CreateModelVersionAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateModelVersionAdmin", reflect.TypeOf((*MockService)(nil).CreateModelVersionAdmin), arg0, arg1)
}

// CreateNamespaceModel mocks base method.
func (m *MockService) CreateNamespaceModel(arg0 context.Context, arg1 resource.Namespace, arg2 *service.AuthUser, arg3 *datamodel.Model) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamespaceModel", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNamespaceModel indicates an expected call of CreateNamespaceModel.
func (mr *MockServiceMockRecorder) CreateNamespaceModel(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespaceModel", reflect.TypeOf((*MockService)(nil).CreateNamespaceModel), arg0, arg1, arg2, arg3)
}

// DBToPBModel mocks base method.
func (m *MockService) DBToPBModel(arg0 context.Context, arg1 *datamodel.ModelDefinition, arg2 *datamodel.Model) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBToPBModel", arg0, arg1, arg2)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DBToPBModel indicates an expected call of DBToPBModel.
func (mr *MockServiceMockRecorder) DBToPBModel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBToPBModel", reflect.TypeOf((*MockService)(nil).DBToPBModel), arg0, arg1, arg2)
}

// DBToPBModelDefinition mocks base method.
func (m *MockService) DBToPBModelDefinition(arg0 context.Context, arg1 *datamodel.ModelDefinition) (*modelv1alpha.ModelDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBToPBModelDefinition", arg0, arg1)
	ret0, _ := ret[0].(*modelv1alpha.ModelDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DBToPBModelDefinition indicates an expected call of DBToPBModelDefinition.
func (mr *MockServiceMockRecorder) DBToPBModelDefinition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBToPBModelDefinition", reflect.TypeOf((*MockService)(nil).DBToPBModelDefinition), arg0, arg1)
}

// DBToPBModelDefinitions mocks base method.
func (m *MockService) DBToPBModelDefinitions(arg0 context.Context, arg1 []*datamodel.ModelDefinition) ([]*modelv1alpha.ModelDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBToPBModelDefinitions", arg0, arg1)
	ret0, _ := ret[0].([]*modelv1alpha.ModelDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DBToPBModelDefinitions indicates an expected call of DBToPBModelDefinitions.
func (mr *MockServiceMockRecorder) DBToPBModelDefinitions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBToPBModelDefinitions", reflect.TypeOf((*MockService)(nil).DBToPBModelDefinitions), arg0, arg1)
}

// DBToPBModels mocks base method.
func (m *MockService) DBToPBModels(arg0 context.Context, arg1 []*datamodel.Model) ([]*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBToPBModels", arg0, arg1)
	ret0, _ := ret[0].([]*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DBToPBModels indicates an expected call of DBToPBModels.
func (mr *MockServiceMockRecorder) DBToPBModels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBToPBModels", reflect.TypeOf((*MockService)(nil).DBToPBModels), arg0, arg1)
}

// DeleteModelVersionAdmin mocks base method.
func (m *MockService) DeleteModelVersionAdmin(arg0 context.Context, arg1 *datamodel.ModelVersion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteModelVersionAdmin", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteModelVersionAdmin indicates an expected call of DeleteModelVersionAdmin.
func (mr *MockServiceMockRecorder) DeleteModelVersionAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteModelVersionAdmin", reflect.TypeOf((*MockService)(nil).DeleteModelVersionAdmin), arg0, arg1)
}

// DeleteNamespaceModelByID mocks base method.
func (m *MockService) DeleteNamespaceModelByID(arg0 context.Context, arg1 resource.Namespace, arg2 *service.AuthUser, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamespaceModelByID", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNamespaceModelByID indicates an expected call of DeleteNamespaceModelByID.
func (mr *MockServiceMockRecorder) DeleteNamespaceModelByID(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespaceModelByID", reflect.TypeOf((*MockService)(nil).DeleteNamespaceModelByID), arg0, arg1, arg2, arg3)
}

// GetACLClient mocks base method.
func (m *MockService) GetACLClient() *acl.ACLClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetACLClient")
	ret0, _ := ret[0].(*acl.ACLClient)
	return ret0
}

// GetACLClient indicates an expected call of GetACLClient.
func (mr *MockServiceMockRecorder) GetACLClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetACLClient", reflect.TypeOf((*MockService)(nil).GetACLClient))
}

// GetMgmtPrivateServiceClient mocks base method.
func (m *MockService) GetMgmtPrivateServiceClient() mgmtv1beta.MgmtPrivateServiceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMgmtPrivateServiceClient")
	ret0, _ := ret[0].(mgmtv1beta.MgmtPrivateServiceClient)
	return ret0
}

// GetMgmtPrivateServiceClient indicates an expected call of GetMgmtPrivateServiceClient.
func (mr *MockServiceMockRecorder) GetMgmtPrivateServiceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMgmtPrivateServiceClient", reflect.TypeOf((*MockService)(nil).GetMgmtPrivateServiceClient))
}

// GetModelByIDAdmin mocks base method.
func (m *MockService) GetModelByIDAdmin(arg0 context.Context, arg1 resource.Namespace, arg2 string, arg3 modelv1alpha.View) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelByIDAdmin", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelByIDAdmin indicates an expected call of GetModelByIDAdmin.
func (mr *MockServiceMockRecorder) GetModelByIDAdmin(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelByIDAdmin", reflect.TypeOf((*MockService)(nil).GetModelByIDAdmin), arg0, arg1, arg2, arg3)
}

// GetModelByUID mocks base method.
func (m *MockService) GetModelByUID(arg0 context.Context, arg1 *service.AuthUser, arg2 uuid.UUID, arg3 modelv1alpha.View) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelByUID", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelByUID indicates an expected call of GetModelByUID.
func (mr *MockServiceMockRecorder) GetModelByUID(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelByUID", reflect.TypeOf((*MockService)(nil).GetModelByUID), arg0, arg1, arg2, arg3)
}

// GetModelByUIDAdmin mocks base method.
func (m *MockService) GetModelByUIDAdmin(arg0 context.Context, arg1 uuid.UUID, arg2 modelv1alpha.View) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelByUIDAdmin", arg0, arg1, arg2)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelByUIDAdmin indicates an expected call of GetModelByUIDAdmin.
func (mr *MockServiceMockRecorder) GetModelByUIDAdmin(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelByUIDAdmin", reflect.TypeOf((*MockService)(nil).GetModelByUIDAdmin), arg0, arg1, arg2)
}

// GetModelDefinition mocks base method.
func (m *MockService) GetModelDefinition(arg0 context.Context, arg1 string) (*modelv1alpha.ModelDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelDefinition", arg0, arg1)
	ret0, _ := ret[0].(*modelv1alpha.ModelDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelDefinition indicates an expected call of GetModelDefinition.
func (mr *MockServiceMockRecorder) GetModelDefinition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelDefinition", reflect.TypeOf((*MockService)(nil).GetModelDefinition), arg0, arg1)
}

// GetModelDefinitionByUID mocks base method.
func (m *MockService) GetModelDefinitionByUID(arg0 context.Context, arg1 uuid.UUID) (*modelv1alpha.ModelDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelDefinitionByUID", arg0, arg1)
	ret0, _ := ret[0].(*modelv1alpha.ModelDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelDefinitionByUID indicates an expected call of GetModelDefinitionByUID.
func (mr *MockServiceMockRecorder) GetModelDefinitionByUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelDefinitionByUID", reflect.TypeOf((*MockService)(nil).GetModelDefinitionByUID), arg0, arg1)
}

// GetModelVersionAdmin mocks base method.
func (m *MockService) GetModelVersionAdmin(arg0 context.Context, arg1 uuid.UUID, arg2 string) (*datamodel.ModelVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelVersionAdmin", arg0, arg1, arg2)
	ret0, _ := ret[0].(*datamodel.ModelVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelVersionAdmin indicates an expected call of GetModelVersionAdmin.
func (mr *MockServiceMockRecorder) GetModelVersionAdmin(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelVersionAdmin", reflect.TypeOf((*MockService)(nil).GetModelVersionAdmin), arg0, arg1, arg2)
}

// GetNamespaceModelByID mocks base method.
func (m *MockService) GetNamespaceModelByID(arg0 context.Context, arg1 resource.Namespace, arg2 *service.AuthUser, arg3 string, arg4 modelv1alpha.View) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceModelByID", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespaceModelByID indicates an expected call of GetNamespaceModelByID.
func (mr *MockServiceMockRecorder) GetNamespaceModelByID(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceModelByID", reflect.TypeOf((*MockService)(nil).GetNamespaceModelByID), arg0, arg1, arg2, arg3, arg4)
}

// GetOperation mocks base method.
func (m *MockService) GetOperation(arg0 context.Context, arg1 string) (*longrunningpb.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperation", arg0, arg1)
	ret0, _ := ret[0].(*longrunningpb.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperation indicates an expected call of GetOperation.
func (mr *MockServiceMockRecorder) GetOperation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockService)(nil).GetOperation), arg0, arg1)
}

// GetRedisClient mocks base method.
func (m *MockService) GetRedisClient() *redis.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRedisClient")
	ret0, _ := ret[0].(*redis.Client)
	return ret0
}

// GetRedisClient indicates an expected call of GetRedisClient.
func (mr *MockServiceMockRecorder) GetRedisClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRedisClient", reflect.TypeOf((*MockService)(nil).GetRedisClient))
}

// GetRepository mocks base method.
func (m *MockService) GetRepository() repository.Repository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepository")
	ret0, _ := ret[0].(repository.Repository)
	return ret0
}

// GetRepository indicates an expected call of GetRepository.
func (mr *MockServiceMockRecorder) GetRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepository", reflect.TypeOf((*MockService)(nil).GetRepository))
}

// GetRscNamespaceAndNameID mocks base method.
func (m *MockService) GetRscNamespaceAndNameID(arg0 string) (resource.Namespace, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRscNamespaceAndNameID", arg0)
	ret0, _ := ret[0].(resource.Namespace)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRscNamespaceAndNameID indicates an expected call of GetRscNamespaceAndNameID.
func (mr *MockServiceMockRecorder) GetRscNamespaceAndNameID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRscNamespaceAndNameID", reflect.TypeOf((*MockService)(nil).GetRscNamespaceAndNameID), arg0)
}

// GetRscNamespaceAndPermalinkUID mocks base method.
func (m *MockService) GetRscNamespaceAndPermalinkUID(arg0 string) (resource.Namespace, uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRscNamespaceAndPermalinkUID", arg0)
	ret0, _ := ret[0].(resource.Namespace)
	ret1, _ := ret[1].(uuid.UUID)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRscNamespaceAndPermalinkUID indicates an expected call of GetRscNamespaceAndPermalinkUID.
func (mr *MockServiceMockRecorder) GetRscNamespaceAndPermalinkUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRscNamespaceAndPermalinkUID", reflect.TypeOf((*MockService)(nil).GetRscNamespaceAndPermalinkUID), arg0)
}

// ListModelDefinitions mocks base method.
func (m *MockService) ListModelDefinitions(arg0 context.Context, arg1 modelv1alpha.View, arg2 int32, arg3 string) ([]*modelv1alpha.ModelDefinition, int32, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModelDefinitions", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*modelv1alpha.ModelDefinition)
	ret1, _ := ret[1].(int32)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListModelDefinitions indicates an expected call of ListModelDefinitions.
func (mr *MockServiceMockRecorder) ListModelDefinitions(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModelDefinitions", reflect.TypeOf((*MockService)(nil).ListModelDefinitions), arg0, arg1, arg2, arg3)
}

// ListModels mocks base method.
func (m *MockService) ListModels(arg0 context.Context, arg1 *service.AuthUser, arg2 int32, arg3 string, arg4 modelv1alpha.View, arg5 *modelv1alpha.Model_Visibility, arg6 bool) ([]*modelv1alpha.Model, int32, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModels", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].([]*modelv1alpha.Model)
	ret1, _ := ret[1].(int32)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListModels indicates an expected call of ListModels.
func (mr *MockServiceMockRecorder) ListModels(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModels", reflect.TypeOf((*MockService)(nil).ListModels), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// ListModelsAdmin mocks base method.
func (m *MockService) ListModelsAdmin(arg0 context.Context, arg1 int32, arg2 string, arg3 modelv1alpha.View, arg4 bool) ([]*modelv1alpha.Model, int32, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModelsAdmin", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*modelv1alpha.Model)
	ret1, _ := ret[1].(int32)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListModelsAdmin indicates an expected call of ListModelsAdmin.
func (mr *MockServiceMockRecorder) ListModelsAdmin(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModelsAdmin", reflect.TypeOf((*MockService)(nil).ListModelsAdmin), arg0, arg1, arg2, arg3, arg4)
}

// ListNamespaceModelVersions mocks base method.
func (m *MockService) ListNamespaceModelVersions(arg0 context.Context, arg1 resource.Namespace, arg2 *service.AuthUser, arg3, arg4 int32, arg5 string) ([]*modelv1alpha.ModelVersion, int32, int32, int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespaceModelVersions", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*modelv1alpha.ModelVersion)
	ret1, _ := ret[1].(int32)
	ret2, _ := ret[2].(int32)
	ret3, _ := ret[3].(int32)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// ListNamespaceModelVersions indicates an expected call of ListNamespaceModelVersions.
func (mr *MockServiceMockRecorder) ListNamespaceModelVersions(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaceModelVersions", reflect.TypeOf((*MockService)(nil).ListNamespaceModelVersions), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListNamespaceModels mocks base method.
func (m *MockService) ListNamespaceModels(arg0 context.Context, arg1 resource.Namespace, arg2 *service.AuthUser, arg3 int32, arg4 string, arg5 modelv1alpha.View, arg6 bool) ([]*modelv1alpha.Model, int32, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespaceModels", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].([]*modelv1alpha.Model)
	ret1, _ := ret[1].(int32)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListNamespaceModels indicates an expected call of ListNamespaceModels.
func (mr *MockServiceMockRecorder) ListNamespaceModels(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaceModels", reflect.TypeOf((*MockService)(nil).ListNamespaceModels), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// PBToDBModel mocks base method.
func (m *MockService) PBToDBModel(arg0 context.Context, arg1 resource.Namespace, arg2 *modelv1alpha.Model) *datamodel.Model {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PBToDBModel", arg0, arg1, arg2)
	ret0, _ := ret[0].(*datamodel.Model)
	return ret0
}

// PBToDBModel indicates an expected call of PBToDBModel.
func (mr *MockServiceMockRecorder) PBToDBModel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PBToDBModel", reflect.TypeOf((*MockService)(nil).PBToDBModel), arg0, arg1, arg2)
}

// RenameNamespaceModelByID mocks base method.
func (m *MockService) RenameNamespaceModelByID(arg0 context.Context, arg1 resource.Namespace, arg2 *service.AuthUser, arg3, arg4 string) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameNamespaceModelByID", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenameNamespaceModelByID indicates an expected call of RenameNamespaceModelByID.
func (mr *MockServiceMockRecorder) RenameNamespaceModelByID(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameNamespaceModelByID", reflect.TypeOf((*MockService)(nil).RenameNamespaceModelByID), arg0, arg1, arg2, arg3, arg4)
}

// TriggerAsyncNamespaceModelByID mocks base method.
func (m *MockService) TriggerAsyncNamespaceModelByID(arg0 context.Context, arg1 resource.Namespace, arg2 *service.AuthUser, arg3, arg4 string, arg5 []byte, arg6 taskv1alpha.Task, arg7 string) (*longrunningpb.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerAsyncNamespaceModelByID", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(*longrunningpb.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerAsyncNamespaceModelByID indicates an expected call of TriggerAsyncNamespaceModelByID.
func (mr *MockServiceMockRecorder) TriggerAsyncNamespaceModelByID(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerAsyncNamespaceModelByID", reflect.TypeOf((*MockService)(nil).TriggerAsyncNamespaceModelByID), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// TriggerNamespaceModelByID mocks base method.
func (m *MockService) TriggerNamespaceModelByID(arg0 context.Context, arg1 resource.Namespace, arg2 *service.AuthUser, arg3, arg4 string, arg5 []byte, arg6 taskv1alpha.Task, arg7 string) ([]*modelv1alpha.TaskOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerNamespaceModelByID", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].([]*modelv1alpha.TaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerNamespaceModelByID indicates an expected call of TriggerNamespaceModelByID.
func (mr *MockServiceMockRecorder) TriggerNamespaceModelByID(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerNamespaceModelByID", reflect.TypeOf((*MockService)(nil).TriggerNamespaceModelByID), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// UpdateModelInstanceAdmin mocks base method.
func (m *MockService) UpdateModelInstanceAdmin(arg0 context.Context, arg1 resource.Namespace, arg2, arg3, arg4 string, arg5 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateModelInstanceAdmin", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateModelInstanceAdmin indicates an expected call of UpdateModelInstanceAdmin.
func (mr *MockServiceMockRecorder) UpdateModelInstanceAdmin(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateModelInstanceAdmin", reflect.TypeOf((*MockService)(nil).UpdateModelInstanceAdmin), arg0, arg1, arg2, arg3, arg4, arg5)
}

// UpdateNamespaceModelByID mocks base method.
func (m *MockService) UpdateNamespaceModelByID(arg0 context.Context, arg1 resource.Namespace, arg2 *service.AuthUser, arg3 string, arg4 *modelv1alpha.Model) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNamespaceModelByID", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNamespaceModelByID indicates an expected call of UpdateNamespaceModelByID.
func (mr *MockServiceMockRecorder) UpdateNamespaceModelByID(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNamespaceModelByID", reflect.TypeOf((*MockService)(nil).UpdateNamespaceModelByID), arg0, arg1, arg2, arg3, arg4)
}

// WatchModel mocks base method.
func (m *MockService) WatchModel(arg0 context.Context, arg1 resource.Namespace, arg2 *service.AuthUser, arg3, arg4 string) (*modelv1alpha.State, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchModel", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*modelv1alpha.State)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// WatchModel indicates an expected call of WatchModel.
func (mr *MockServiceMockRecorder) WatchModel(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchModel", reflect.TypeOf((*MockService)(nil).WatchModel), arg0, arg1, arg2, arg3, arg4)
}

// WriteNewDataPoint mocks base method.
func (m *MockService) WriteNewDataPoint(arg0 context.Context, arg1 *utils.UsageMetricData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteNewDataPoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteNewDataPoint indicates an expected call of WriteNewDataPoint.
func (mr *MockServiceMockRecorder) WriteNewDataPoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteNewDataPoint", reflect.TypeOf((*MockService)(nil).WriteNewDataPoint), arg0, arg1)
}

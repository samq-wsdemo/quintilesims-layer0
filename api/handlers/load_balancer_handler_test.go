package handlers

import (
	"testing"

	"github.com/emicklei/go-restful"
	"github.com/golang/mock/gomock"
	"github.com/quintilesims/layer0/api/logic/mock_logic"
	"github.com/quintilesims/layer0/common/errors"
	"github.com/quintilesims/layer0/common/models"
	"github.com/quintilesims/layer0/common/testutils"
	"github.com/quintilesims/layer0/common/types"
)

func TestListLoadBalancers(t *testing.T) {
	loadBalancers := []*models.LoadBalancerSummary{
		{
			LoadBalancerID: "some_id_1",
		},
		{
			LoadBalancerID: "some_id_2",
		},
	}

	testCases := []HandlerTestCase{
		{
			Name:    "Should return loadBalancers from logic layer",
			Request: &TestRequest{},
			Setup: func(ctrl *gomock.Controller) interface{} {
				logicMock := mock_logic.NewMockLoadBalancerLogic(ctrl)
				logicMock.EXPECT().
					ListLoadBalancers().
					Return(loadBalancers, nil)

				mockJob := mock_logic.NewMockJobLogic(ctrl)
				return NewLoadBalancerHandler(logicMock, mockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.ListLoadBalancers(req, resp)

				var response []*models.LoadBalancerSummary
				read(&response)

				reporter.AssertEqual(response, loadBalancers)
			},
		},
		{
			Name:    "Should propagate ListLoadBalancers error",
			Request: &TestRequest{},
			Setup: func(ctrl *gomock.Controller) interface{} {
				logicMock := mock_logic.NewMockLoadBalancerLogic(ctrl)
				logicMock.EXPECT().
					ListLoadBalancers().
					Return(nil, errors.Newf(errors.UnexpectedError, "some error"))

				mockJob := mock_logic.NewMockJobLogic(ctrl)
				return NewLoadBalancerHandler(logicMock, mockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.ListLoadBalancers(req, resp)

				var response *models.ServerError
				read(&response)

				reporter.AssertEqual(response.ErrorCode, int64(errors.UnexpectedError))
			},
		},
	}

	RunHandlerTestCases(t, testCases)
}

func TestGetLoadBalancer(t *testing.T) {
	loadBalancer := &models.LoadBalancer{
		LoadBalancerID: "some_id",
	}

	testCases := []HandlerTestCase{
		{
			Name: "Should call GetLoadBalancer with proper params",
			Request: &TestRequest{
				Parameters: map[string]string{"id": "some_id"},
			},
			Setup: func(ctrl *gomock.Controller) interface{} {
				logicMock := mock_logic.NewMockLoadBalancerLogic(ctrl)
				logicMock.EXPECT().
					GetLoadBalancer("some_id").
					Return(loadBalancer, nil)

				mockJob := mock_logic.NewMockJobLogic(ctrl)
				return NewLoadBalancerHandler(logicMock, mockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.GetLoadBalancer(req, resp)
			},
		},
		{
			Name: "Should return loadBalancer from logic layer",
			Request: &TestRequest{
				Parameters: map[string]string{"id": "some_id"},
			},
			Setup: func(ctrl *gomock.Controller) interface{} {
				logicMock := mock_logic.NewMockLoadBalancerLogic(ctrl)
				logicMock.EXPECT().
					GetLoadBalancer(gomock.Any()).
					Return(loadBalancer, nil)

				mockJob := mock_logic.NewMockJobLogic(ctrl)
				return NewLoadBalancerHandler(logicMock, mockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.GetLoadBalancer(req, resp)

				var response *models.LoadBalancer
				read(&response)

				reporter.AssertEqual(response, loadBalancer)
			},
		},
		{
			Name:    "Should return MissingParameter error with no id",
			Request: &TestRequest{},
			Setup: func(ctrl *gomock.Controller) interface{} {
				logicMock := mock_logic.NewMockLoadBalancerLogic(ctrl)
				mockJob := mock_logic.NewMockJobLogic(ctrl)
				return NewLoadBalancerHandler(logicMock, mockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.GetLoadBalancer(req, resp)

				var response *models.ServerError
				read(&response)

				reporter.AssertEqual(response.ErrorCode, int64(errors.MissingParameter))
			},
		},
		{
			Name: "Should propagate GetLoadBalancer error",
			Request: &TestRequest{
				Parameters: map[string]string{"id": "some_id"},
			},
			Setup: func(ctrl *gomock.Controller) interface{} {
				logicMock := mock_logic.NewMockLoadBalancerLogic(ctrl)
				logicMock.EXPECT().
					GetLoadBalancer(gomock.Any()).
					Return(nil, errors.Newf(errors.UnexpectedError, "some error"))

				mockJob := mock_logic.NewMockJobLogic(ctrl)
				return NewLoadBalancerHandler(logicMock, mockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.GetLoadBalancer(req, resp)

				var response *models.ServerError
				read(&response)

				reporter.AssertEqual(response.ErrorCode, int64(errors.UnexpectedError))
			},
		},
	}

	RunHandlerTestCases(t, testCases)
}

func TestCreateLoadBalancer(t *testing.T) {
	request := models.CreateLoadBalancerRequest{
		LoadBalancerName: "lb_name",
		EnvironmentID:    "envid",
		IsPublic:         true,
		Ports:            []models.Port{},
	}

	testCases := []HandlerTestCase{
		{
			Name: "Should call logic:CreateLoadBalancer with correct params",
			Request: &TestRequest{
				Body: request,
			},
			Setup: func(ctrl *gomock.Controller) interface{} {
				mockLB := mock_logic.NewMockLoadBalancerLogic(ctrl)
				mockLB.EXPECT().
					CreateLoadBalancer(request).
					Return(nil, nil)

				mockJob := mock_logic.NewMockJobLogic(ctrl)
				return NewLoadBalancerHandler(mockLB, mockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.CreateLoadBalancer(req, resp)
			},
		},
		{
			Name: "Should propagate CreateLoadBalancer error",
			Request: &TestRequest{
				Body: request,
			},
			Setup: func(ctrl *gomock.Controller) interface{} {
				mockLB := mock_logic.NewMockLoadBalancerLogic(ctrl)
				mockLB.EXPECT().
					CreateLoadBalancer(request).
					Return(nil, errors.Newf(errors.UnexpectedError, "some error"))

				mockJob := mock_logic.NewMockJobLogic(ctrl)
				return NewLoadBalancerHandler(mockLB, mockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.CreateLoadBalancer(req, resp)

				var response *models.ServerError
				read(&response)

				reporter.AssertEqual(int64(errors.UnexpectedError), response.ErrorCode)
			},
		},
	}

	RunHandlerTestCases(t, testCases)
}

func TestDeleteLoadBalancer(t *testing.T) {
	testCases := []HandlerTestCase{
		{
			Name: "Should call CreateJob with correct params",
			Request: &TestRequest{
				Parameters: map[string]string{"id": "some_id"},
			},
			Setup: func(ctrl *gomock.Controller) interface{} {
				mockLB := mock_logic.NewMockLoadBalancerLogic(ctrl)

				MockJob := mock_logic.NewMockJobLogic(ctrl)
				MockJob.EXPECT().
					CreateJob(types.DeleteLoadBalancerJob, "some_id").
					Return(&models.Job{}, nil)

				return NewLoadBalancerHandler(mockLB, MockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.DeleteLoadBalancer(req, resp)
			},
		},
		{
			Name: "Should set Location and X-Jobid headers",
			Request: &TestRequest{
				Parameters: map[string]string{"id": "some_id"},
			},
			Setup: func(ctrl *gomock.Controller) interface{} {
				mockLB := mock_logic.NewMockLoadBalancerLogic(ctrl)

				MockJob := mock_logic.NewMockJobLogic(ctrl)
				MockJob.EXPECT().
					CreateJob(gomock.Any(), gomock.Any()).
					Return(&models.Job{JobID: "job_id"}, nil)

				return NewLoadBalancerHandler(mockLB, MockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.DeleteLoadBalancer(req, resp)

				header := resp.Header()
				reporter.AssertInSlice("/job/job_id", header["Location"])
				reporter.AssertInSlice("job_id", header["X-Jobid"])
			},
		},
		{
			Name: "Should propagate CreateJob error",
			Request: &TestRequest{
				Parameters: map[string]string{"id": "some_id"},
			},
			Setup: func(ctrl *gomock.Controller) interface{} {
				mockLB := mock_logic.NewMockLoadBalancerLogic(ctrl)

				MockJob := mock_logic.NewMockJobLogic(ctrl)
				MockJob.EXPECT().
					CreateJob(gomock.Any(), gomock.Any()).
					Return(nil, errors.Newf(errors.UnexpectedError, "some error"))

				return NewLoadBalancerHandler(mockLB, MockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.DeleteLoadBalancer(req, resp)

				var response *models.ServerError
				read(&response)

				reporter.AssertEqual(int64(errors.UnexpectedError), response.ErrorCode)
			},
		},
		{
			Name:    "Should return MissingParameter error with no id",
			Request: &TestRequest{},
			Setup: func(ctrl *gomock.Controller) interface{} {
				mockLB := mock_logic.NewMockLoadBalancerLogic(ctrl)
				MockJob := mock_logic.NewMockJobLogic(ctrl)

				return NewLoadBalancerHandler(mockLB, MockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.DeleteLoadBalancer(req, resp)

				var response *models.ServerError
				read(&response)

				reporter.AssertEqual(int64(errors.MissingParameter), response.ErrorCode)
			},
		},
	}

	RunHandlerTestCases(t, testCases)
}

func TestUpdateLoadBalancerHealthCheck(t *testing.T) {
	request := models.UpdateLoadBalancerHealthCheckRequest{
		models.HealthCheck{
			Target:             "TCP:80",
			Interval:           30,
			Timeout:            5,
			HealthyThreshold:   2,
			UnhealthyThreshold: 2,
		},
	}

	testCases := []HandlerTestCase{
		{
			Name: "Should call UpdateLoadBalancerHealthCheck with correct params",
			Request: &TestRequest{
				Parameters: map[string]string{"id": "some_id"},
				Body:       request,
			},
			Setup: func(ctrl *gomock.Controller) interface{} {
				mockLogic := mock_logic.NewMockLoadBalancerLogic(ctrl)
				mockJob := mock_logic.NewMockJobLogic(ctrl)

				mockLogic.EXPECT().
					UpdateLoadBalancerHealthCheck("some_id", request.HealthCheck)

				return NewLoadBalancerHandler(mockLogic, mockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.UpdateLoadBalancerHealthCheck(req, resp)
			},
		},
	}

	RunHandlerTestCases(t, testCases)
}

func TestUpdateLoadBalancerIdleTimeout(t *testing.T) {
	request := models.UpdateLoadBalancerIdleTimeoutRequest{IdleTimeout: 60}

	testCases := []HandlerTestCase{
		{
			Name: "Should call UpdateLoadBalancerIdleTimeout with correct params",
			Request: &TestRequest{
				Parameters: map[string]string{"id": "some_id"},
				Body:       request,
			},
			Setup: func(ctrl *gomock.Controller) interface{} {
				mockLogic := mock_logic.NewMockLoadBalancerLogic(ctrl)
				mockJob := mock_logic.NewMockJobLogic(ctrl)

				mockLogic.EXPECT().
					UpdateLoadBalancerIdleTimeout("some_id", request.IdleTimeout)

				return NewLoadBalancerHandler(mockLogic, mockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.UpdateLoadBalancerIdleTimeout(req, resp)
			},
		},
	}

	RunHandlerTestCases(t, testCases)
}

func TestUpdateLoadBalancerCrossZone(t *testing.T) {
	request := models.UpdateLoadBalancerCrossZoneRequest{CrossZone: true}

	testCases := []HandlerTestCase{
		{
			Name: "Should call UpdateLoadBalancerCrossZone with correct params",
			Request: &TestRequest{
				Parameters: map[string]string{"id": "some_id"},
				Body:       request,
			},
			Setup: func(ctrl *gomock.Controller) interface{} {
				mockLogic := mock_logic.NewMockLoadBalancerLogic(ctrl)
				mockJob := mock_logic.NewMockJobLogic(ctrl)

				mockLogic.EXPECT().
					UpdateLoadBalancerCrossZone("some_id", request.CrossZone)

				return NewLoadBalancerHandler(mockLogic, mockJob)
			},
			Run: func(reporter *testutils.Reporter, target interface{}, req *restful.Request, resp *restful.Response, read Readf) {
				handler := target.(*LoadBalancerHandler)
				handler.UpdateLoadBalancerCrossZone(req, resp)
			},
		},
	}

	RunHandlerTestCases(t, testCases)
}

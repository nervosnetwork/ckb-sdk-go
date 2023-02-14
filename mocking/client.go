package mocking

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type MockJsonResponse struct {
	mock.Mock
}

type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type jsonrpcMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *jsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

func (c *MockClient) newMessage(method string, paramsIn ...interface{}) (*jsonrpcMessage, error) {
	msg := &jsonrpcMessage{Version: "2.0", ID: strconv.AppendUint(nil, uint64(1), 10), Method: method}
	if paramsIn != nil { // prevent sending "params":null
		var err error
		if msg.Params, err = json.Marshal(paramsIn); err != nil {
			return nil, err
		}
	}
	return msg, nil
}

func newMockJsonResp() *MockJsonResponse { return &MockJsonResponse{} }

type MockClient struct {
	idCounter uint32
	mockJson  *MockJsonResponse
	types.GenericRPCClient
}

func (c *MockClient) LoadMockingTestFromFile(t *testing.T, method string, paramsIn ...interface{}) *MockClient {
	var sb strings.Builder
	sb.WriteString("mocking/")
	sb.WriteString(method)
	sb.WriteString("/")
	prefix := sb.String()
	response, err := os.ReadFile(prefix + "response.json")
	assert.NoError(t, err)
	request, err := os.ReadFile(prefix + "request.json")
	assert.NoError(t, err)
	requestJson, err := c.newMessage(method, paramsIn...)
	assert.NoError(t, err)
	requstJsonMsg, err := json.Marshal(requestJson)
	assert.NoError(t, err)
	// Test for request json equivalent
	require.JSONEqf(t, string(request), string(requstJsonMsg), "failed on json equivalent test: %s")
	// If last passed, use the generated one(for passing testify validation)
	c.mockJson.On("mockJsonRPC", requstJsonMsg).Return(response)
	return c
}

func (c *MockClient) Expect(msg *jsonrpcMessage, jsonData []byte) *MockClient {
	c.mockJson.On("mockJsonRPC", msg).Return(jsonData)
	return c
}

func (m *MockJsonResponse) mockJsonRPC(msg []byte) []byte {
	args := m.Called(msg)
	return args.Get(0).([]byte)
}

func DialContext(ctx context.Context, rawurl string) (*MockClient, error) {
	_, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	c := &MockClient{
		idCounter: 0,
		mockJson:  newMockJsonResp(),
	}
	return c, nil
}

func (c *MockClient) GetRPCJsonMessage(method string, args ...interface{}) (*jsonrpcMessage, error) {
	msg, err := c.newMessage(method, args...)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (c *MockClient) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	if result != nil && reflect.TypeOf(result).Kind() != reflect.Ptr {
		return fmt.Errorf("call result parameter must be pointer or nil interface: %v", result)
	}
	msg, err := c.newMessage(method, args...)
	if err != nil {
		return err
	}
	requstJsonMsg, err := json.Marshal(msg)
	// We use []byte here to avoid json eval on testify
	resp := c.mockJson.mockJsonRPC(requstJsonMsg)
	var rpcResp *jsonrpcMessage
	json.Unmarshal(resp, &rpcResp)
	return json.Unmarshal(rpcResp.Result, &result)
}

func (c *MockClient) Close() {

}

func (c *MockClient) LoadMockingTestFromFilePatched(t *testing.T, method string, actualMethod string, args ...interface{}) *MockClient {
	var sb strings.Builder
	sb.WriteString("mocking/")
	sb.WriteString(method)
	sb.WriteString("/")
	prefix := sb.String()
	response, err := os.ReadFile(prefix + "response.json")
	assert.NoError(t, err)
	request, err := os.ReadFile(prefix + "request.json")
	assert.NoError(t, err)
	requestJson, err := c.newMessage(actualMethod, args...)
	assert.NoError(t, err)
	requstJsonMsg, err := json.Marshal(requestJson)
	assert.NoError(t, err)
	// Test for request json equivalent
	require.JSONEqf(t, string(request), string(requstJsonMsg), "failed on json equivalent test: %s")
	// If last passed, use the generated one(for passing testify validation)
	c.mockJson.On("mockJsonRPC", requstJsonMsg).Return(response)
	return c
}

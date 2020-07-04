// Code generated by jrpc. DO NOT EDIT.

package deviceregistrydef

import (
	context "context"
	"testing"

	taxi "github.com/jakewright/home-automation/libraries/go/taxi"
)

// DeviceRegistry is the public interface of this service
type DeviceRegistry interface {
	GetDevice(ctx context.Context, body *GetDeviceRequest) *GetDeviceFuture
	ListDevices(ctx context.Context, body *ListDevicesRequest) *ListDevicesFuture
	GetRoom(ctx context.Context, body *GetRoomRequest) *GetRoomFuture
	ListRooms(ctx context.Context, body *ListRoomsRequest) *ListRoomsFuture
}

// GetDeviceFuture represents an in-flight GetDevice request
type GetDeviceFuture struct {
	done <-chan struct{}
	rsp  *GetDeviceResponse
	err  error
}

// Wait blocks until the response is ready
func (f *GetDeviceFuture) Wait() (*GetDeviceResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// ListDevicesFuture represents an in-flight ListDevices request
type ListDevicesFuture struct {
	done <-chan struct{}
	rsp  *ListDevicesResponse
	err  error
}

// Wait blocks until the response is ready
func (f *ListDevicesFuture) Wait() (*ListDevicesResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// GetRoomFuture represents an in-flight GetRoom request
type GetRoomFuture struct {
	done <-chan struct{}
	rsp  *GetRoomResponse
	err  error
}

// Wait blocks until the response is ready
func (f *GetRoomFuture) Wait() (*GetRoomResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// ListRoomsFuture represents an in-flight ListRooms request
type ListRoomsFuture struct {
	done <-chan struct{}
	rsp  *ListRoomsResponse
	err  error
}

// Wait blocks until the response is ready
func (f *ListRoomsFuture) Wait() (*ListRoomsResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// DeviceRegistryClient makes requests to this service
type DeviceRegistryClient struct {
	dispatcher taxi.Dispatcher
}

// Compile-time assertion that the client implements the interface
var _ DeviceRegistry = (*DeviceRegistryClient)(nil)

// NewDeviceRegistryClient returns a new client
func NewDeviceRegistryClient(d taxi.Dispatcher) *DeviceRegistryClient {
	return &DeviceRegistryClient{
		dispatcher: d,
	}
}

// GetDevice dispatches an RPC to the service
func (c *DeviceRegistryClient) GetDevice(ctx context.Context, body *GetDeviceRequest) *GetDeviceFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "service.device-registry/device",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &GetDeviceFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// ListDevices dispatches an RPC to the service
func (c *DeviceRegistryClient) ListDevices(ctx context.Context, body *ListDevicesRequest) *ListDevicesFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "service.device-registry/devices",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &ListDevicesFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// GetRoom dispatches an RPC to the service
func (c *DeviceRegistryClient) GetRoom(ctx context.Context, body *GetRoomRequest) *GetRoomFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "service.device-registry/room",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &GetRoomFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// ListRooms dispatches an RPC to the service
func (c *DeviceRegistryClient) ListRooms(ctx context.Context, body *ListRoomsRequest) *ListRoomsFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "service.device-registry/rooms",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &ListRoomsFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// MockDeviceRegistryClient can be used in tests
type MockDeviceRegistryClient struct {
	dispatcher *taxi.MockClient
}

// Compile-time assertion that the mock client implements the interface
var _ DeviceRegistry = (*MockDeviceRegistryClient)(nil)

// NewMockDeviceRegistryClient returns a new mock client
func NewMockDeviceRegistryClient(ctx context.Context, t *testing.T) *MockDeviceRegistryClient {
	f := taxi.NewTestFixture(t)

	return &MockDeviceRegistryClient{
		dispatcher: &taxi.MockClient{Handler: f},
	}
}

// GetDevice dispatches an RPC to the mock client
func (c *MockDeviceRegistryClient) GetDevice(ctx context.Context, body *GetDeviceRequest) *GetDeviceFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "service.device-registry/device",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &GetDeviceFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// ListDevices dispatches an RPC to the mock client
func (c *MockDeviceRegistryClient) ListDevices(ctx context.Context, body *ListDevicesRequest) *ListDevicesFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "service.device-registry/devices",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &ListDevicesFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// GetRoom dispatches an RPC to the mock client
func (c *MockDeviceRegistryClient) GetRoom(ctx context.Context, body *GetRoomRequest) *GetRoomFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "service.device-registry/room",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &GetRoomFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// ListRooms dispatches an RPC to the mock client
func (c *MockDeviceRegistryClient) ListRooms(ctx context.Context, body *ListRoomsRequest) *ListRoomsFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "service.device-registry/rooms",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &ListRoomsFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

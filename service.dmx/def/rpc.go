// Code generated by jrpc. DO NOT EDIT.

package dmxdef

import (
	context "context"

	rpc "github.com/jakewright/home-automation/libraries/go/rpc"
)

// Do performs the request
func (m *GetDeviceRequest) Request() *rpc.Request {
	return &rpc.Request{
		Method: "GET",
		URL:    "service.dmx/device",
		Body:   m,
	}
}

func (m *GetDeviceRequest) Do(ctx context.Context) (*GetDeviceResponse, error) {
	rsp := &GetDeviceResponse{}
	_, err := rpc.Do(ctx, m.Request(), rsp)
	return rsp, err
}

// Do performs the request
func (m *UpdateDeviceRequest) Request() *rpc.Request {
	return &rpc.Request{
		Method: "PATCH",
		URL:    "service.dmx/device",
		Body:   m,
	}
}

func (m *UpdateDeviceRequest) Do(ctx context.Context) (*UpdateDeviceResponse, error) {
	rsp := &UpdateDeviceResponse{}
	_, err := rpc.Do(ctx, m.Request(), rsp)
	return rsp, err
}

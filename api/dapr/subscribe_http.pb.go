// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http 0.1.0

package dapr

import (
	context "context"
	go_restful "github.com/emicklei/go-restful"
	errors "github.com/tkeel-io/kit/errors"
	result "github.com/tkeel-io/kit/result"
	protojson "google.golang.org/protobuf/encoding/protojson"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
)

import transportHTTP "github.com/tkeel-io/kit/transport/http"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the tkeel package it is being compiled against.
// import package.context.http.anypb.result.protojson.go_restful.errors.emptypb.

var (
	_ = protojson.MarshalOptions{}
	_ = anypb.Any{}
	_ = emptypb.Empty{}
)

type SubscribeHTTPServer interface {
	GetSubscribe(context.Context, *emptypb.Empty) (*ListTopicSubscriptionsResponse, error)
}

type SubscribeHTTPHandler struct {
	srv SubscribeHTTPServer
}

func newSubscribeHTTPHandler(s SubscribeHTTPServer) *SubscribeHTTPHandler {
	return &SubscribeHTTPHandler{srv: s}
}

func (h *SubscribeHTTPHandler) GetSubscribe(req *go_restful.Request, resp *go_restful.Response) {
	in := emptypb.Empty{}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.GetSubscribe(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(httpCode, tErr.Message, out), "application/json")
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.WriteAsJson(out.GetSubscriptions())
}

func RegisterSubscribeHTTPServer(container *go_restful.Container, srv SubscribeHTTPServer) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "/dapr" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.ApiVersion("/dapr")
		ws.Path("/dapr").Produces(go_restful.MIME_JSON)
		container.Add(ws)
	}

	handler := newSubscribeHTTPHandler(srv)
	ws.Route(ws.GET("/subscribe").
		To(handler.GetSubscribe))
}

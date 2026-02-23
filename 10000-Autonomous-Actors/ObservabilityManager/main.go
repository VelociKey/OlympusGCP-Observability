package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	observabilityv1 "OlympusGCP-Observability/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/observability/v1"
	"OlympusGCP-Observability/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/observability/v1/observabilityv1connect"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ObservabilityServer struct{}

func (s *ObservabilityServer) WriteLog(ctx context.Context, req *connect.Request[observabilityv1.WriteLogRequest]) (*connect.Response[observabilityv1.WriteLogResponse], error) {
	slog.Info("WriteLog", "log", req.Msg.LogName, "severity", req.Msg.Severity, "msg", req.Msg.Message)
	return connect.NewResponse(&observabilityv1.WriteLogResponse{}), nil
}

func (s *ObservabilityServer) RecordMetric(ctx context.Context, req *connect.Request[observabilityv1.RecordMetricRequest]) (*connect.Response[observabilityv1.RecordMetricResponse], error) {
	slog.Info("RecordMetric", "type", req.Msg.MetricType, "value", req.Msg.Value)
	return connect.NewResponse(&observabilityv1.RecordMetricResponse{}), nil
}

func (s *ObservabilityServer) StartSpan(ctx context.Context, req *connect.Request[observabilityv1.StartSpanRequest]) (*connect.Response[observabilityv1.StartSpanResponse], error) {
	slog.Info("StartSpan", "name", req.Msg.Name, "trace", req.Msg.TraceId)
	return connect.NewResponse(&observabilityv1.StartSpanResponse{SpanId: "span-123"}), nil
}

func (s *ObservabilityServer) EndSpan(ctx context.Context, req *connect.Request[observabilityv1.EndSpanRequest]) (*connect.Response[observabilityv1.EndSpanResponse], error) {
	slog.Info("EndSpan", "span", req.Msg.SpanId)
	return connect.NewResponse(&observabilityv1.EndSpanResponse{}), nil
}

func main() {
	server := &ObservabilityServer{}
	mux := http.NewServeMux()
	path, handler := observabilityv1connect.NewObservabilityServiceHandler(server)
	mux.Handle(path, handler)

	port := "8097" // From genesis.json
	slog.Info("ObservabilityManager starting", "port", port)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("Server failed", "error", err)
	}
}

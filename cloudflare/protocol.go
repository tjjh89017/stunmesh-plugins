package main

import "context"

// These types mirror the stunmesh-go exec plugin protocol (JSON over
// stdin/stdout), documented at https://docs.stunmesh.dev/plugins/exec-protocol.

const (
	OpSet = "set"
	OpGet = "get"
)

// Store is the key/value backend the request handler talks to.
type Store interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

// ExecRequest is the JSON request read from stdin.
type ExecRequest struct {
	Action string `json:"action"`
	Key    string `json:"key"`
	Value  string `json:"value,omitempty"`
}

// ExecResponse is the JSON response written to stdout.
type ExecResponse struct {
	Success bool   `json:"success"`
	Value   string `json:"value,omitempty"`
	Error   string `json:"error,omitempty"`
}

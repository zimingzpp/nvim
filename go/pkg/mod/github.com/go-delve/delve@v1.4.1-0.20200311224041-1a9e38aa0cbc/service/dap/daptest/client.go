// Package daptest provides a sample client with utilities
// for DAP mode testing.
package daptest

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"testing"

	"github.com/google/go-dap"
)

// Client is a debugger service client that uses Debug Adaptor Protocol.
// It does not (yet?) implement service.Client interface.
// All client methods are synchronous.
type Client struct {
	conn   net.Conn
	reader *bufio.Reader
	// seq is used to track the sequence number of each
	// requests that the client sends to the server
	seq int
}

// NewClient creates a new Client over a TCP connection.
// Call Close() to close the connection.
func NewClient(addr string) *Client {
	fmt.Println("Connecting to server at:", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	c := &Client{conn: conn, reader: bufio.NewReader(conn)}
	c.seq = 1 // match VS Code numbering
	return c
}

// Close closes the client connection.
func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) send(request dap.Message) {
	dap.WriteProtocolMessage(c.conn, request)
}

func (c *Client) ReadMessage() (dap.Message, error) {
	return dap.ReadProtocolMessage(c.reader)
}

func (c *Client) expectReadProtocolMessage(t *testing.T) dap.Message {
	t.Helper()
	m, err := dap.ReadProtocolMessage(c.reader)
	if err != nil {
		t.Error(err)
	}
	return m
}

func (c *Client) ExpectDisconnectResponse(t *testing.T) *dap.DisconnectResponse {
	t.Helper()
	return c.expectReadProtocolMessage(t).(*dap.DisconnectResponse)
}

func (c *Client) ExpectErrorResponse(t *testing.T) *dap.ErrorResponse {
	t.Helper()
	return c.expectReadProtocolMessage(t).(*dap.ErrorResponse)
}

func (c *Client) ExpectContinueResponse(t *testing.T) *dap.ContinueResponse {
	t.Helper()
	return c.expectReadProtocolMessage(t).(*dap.ContinueResponse)
}

func (c *Client) ExpectTerminatedEvent(t *testing.T) *dap.TerminatedEvent {
	t.Helper()
	return c.expectReadProtocolMessage(t).(*dap.TerminatedEvent)
}

func (c *Client) ExpectInitializeResponse(t *testing.T) *dap.InitializeResponse {
	t.Helper()
	initResp := c.expectReadProtocolMessage(t).(*dap.InitializeResponse)
	if !initResp.Body.SupportsConfigurationDoneRequest {
		t.Errorf("got %#v, want SupportsConfigurationDoneRequest=true", initResp)
	}
	return initResp
}

func (c *Client) ExpectInitializedEvent(t *testing.T) *dap.InitializedEvent {
	t.Helper()
	return c.expectReadProtocolMessage(t).(*dap.InitializedEvent)
}

func (c *Client) ExpectLaunchResponse(t *testing.T) *dap.LaunchResponse {
	t.Helper()
	return c.expectReadProtocolMessage(t).(*dap.LaunchResponse)
}

func (c *Client) ExpectSetExceptionBreakpointsResponse(t *testing.T) *dap.SetExceptionBreakpointsResponse {
	t.Helper()
	return c.expectReadProtocolMessage(t).(*dap.SetExceptionBreakpointsResponse)
}

func (c *Client) ExpectSetBreakpointsResponse(t *testing.T) *dap.SetBreakpointsResponse {
	t.Helper()
	return c.expectReadProtocolMessage(t).(*dap.SetBreakpointsResponse)
}

func (c *Client) ExpectStoppedEvent(t *testing.T) *dap.StoppedEvent {
	t.Helper()
	return c.expectReadProtocolMessage(t).(*dap.StoppedEvent)
}

func (c *Client) ExpectConfigurationDoneResponse(t *testing.T) *dap.ConfigurationDoneResponse {
	t.Helper()
	return c.expectReadProtocolMessage(t).(*dap.ConfigurationDoneResponse)
}

func (c *Client) ExpectThreadsResponse(t *testing.T) *dap.ThreadsResponse {
	t.Helper()
	return c.expectReadProtocolMessage(t).(*dap.ThreadsResponse)
}

func (c *Client) ExpectStackTraceResponse(t *testing.T) *dap.StackTraceResponse {
	t.Helper()
	return c.expectReadProtocolMessage(t).(*dap.StackTraceResponse)
}

// InitializeRequest sends an 'initialize' request.
func (c *Client) InitializeRequest() {
	request := &dap.InitializeRequest{Request: *c.newRequest("initialize")}
	request.Arguments = dap.InitializeRequestArguments{
		AdapterID:                    "go",
		PathFormat:                   "path",
		LinesStartAt1:                true,
		ColumnsStartAt1:              true,
		SupportsVariableType:         true,
		SupportsVariablePaging:       true,
		SupportsRunInTerminalRequest: true,
		Locale:                       "en-us",
	}
	c.send(request)
}

// LaunchRequest sends a 'launch' request with the specified args.
func (c *Client) LaunchRequest(mode string, program string, stopOnEntry bool) {
	request := &dap.LaunchRequest{Request: *c.newRequest("launch")}
	request.Arguments = map[string]interface{}{
		"request":     "launch",
		"mode":        mode,
		"program":     program,
		"stopOnEntry": stopOnEntry,
	}
	c.send(request)
}

// LaunchRequestWithArgs takes a map of untyped implementation-specific
// arguments to send a 'launch' request. This version can be used to
// test for values of unexpected types or unspecified values.
func (c *Client) LaunchRequestWithArgs(arguments map[string]interface{}) {
	request := &dap.LaunchRequest{Request: *c.newRequest("launch")}
	request.Arguments = arguments
	c.send(request)
}

// DisconnectRequest sends a 'disconnect' request.
func (c *Client) DisconnectRequest() {
	request := &dap.DisconnectRequest{Request: *c.newRequest("disconnect")}
	c.send(request)
}

// SetBreakpointsRequest sends a 'setBreakpoints' request.
func (c *Client) SetBreakpointsRequest(file string, lines []int) {
	request := &dap.SetBreakpointsRequest{Request: *c.newRequest("setBreakpoints")}
	request.Arguments = dap.SetBreakpointsArguments{
		Source: dap.Source{
			Name: filepath.Base(file),
			Path: file,
		},
		Breakpoints: make([]dap.SourceBreakpoint, len(lines)),
		//sourceModified: false,
	}
	for i, l := range lines {
		request.Arguments.Breakpoints[i].Line = l
	}
	c.send(request)
}

// SetExceptionBreakpointsRequest sends a 'setExceptionBreakpoints' request.
func (c *Client) SetExceptionBreakpointsRequest() {
	request := &dap.SetBreakpointsRequest{Request: *c.newRequest("setExceptionBreakpoints")}
	c.send(request)
}

// ConfigurationDoneRequest sends a 'configurationDone' request.
func (c *Client) ConfigurationDoneRequest() {
	request := &dap.ConfigurationDoneRequest{Request: *c.newRequest("configurationDone")}
	c.send(request)
}

// ContinueRequest sends a 'continue' request.
func (c *Client) ContinueRequest(thread int) {
	request := &dap.ContinueRequest{Request: *c.newRequest("continue")}
	request.Arguments.ThreadId = thread
	c.send(request)
}

// ThreadsRequest sends a 'threads' request.
func (c *Client) ThreadsRequest() {
	request := &dap.ThreadsRequest{Request: *c.newRequest("threads")}
	c.send(request)
}

// StackTraceRequest sends a 'stackTrace' request.
func (c *Client) StackTraceRequest() {
	request := &dap.StackTraceRequest{Request: *c.newRequest("stackTrace")}
	c.send(request)
}

// UnknownRequest triggers dap.DecodeProtocolMessageFieldError.
func (c *Client) UnknownRequest() {
	request := c.newRequest("unknown")
	c.send(request)
}

// UnknownEvent triggers dap.DecodeProtocolMessageFieldError.
func (c *Client) UnknownEvent() {
	event := &dap.Event{}
	event.Type = "event"
	event.Seq = -1
	event.Event = "unknown"
	c.send(event)
}

// KnownEvent passes decode checks, but delve has no 'case' to
// handle it. This behaves the same way a new request type
// added to go-dap, but not to delve.
func (c *Client) KnownEvent() {
	event := &dap.Event{}
	event.Type = "event"
	event.Seq = -1
	event.Event = "terminated"
	c.send(event)
}

func (c *Client) newRequest(command string) *dap.Request {
	request := &dap.Request{}
	request.Type = "request"
	request.Command = command
	request.Seq = c.seq
	c.seq++
	return request
}

// Code generated by protoc-gen-twirp v5.3.0, DO NOT EDIT.
// source: rpc/service.proto

/*
Package trackdata is a generated twirp stub package.
This code was generated with github.com/twitchtv/twirp/protoc-gen-twirp v5.3.0.

It is generated from these files:
	rpc/service.proto
*/
package trackdata

import bytes "bytes"
import strings "strings"
import context "context"
import fmt "fmt"
import ioutil "io/ioutil"
import http "net/http"

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Imports only used by utility functions:
import io "io"
import strconv "strconv"
import json "encoding/json"
import url "net/url"
import bufio "bufio"
import binary "encoding/binary"

// ==========================
// TrackDataService Interface
// ==========================

type TrackDataService interface {
	StreamTrackData(ctx context.Context, in *UserTrack) (<-chan TrackChunkOrError, error)

	UploadTrackData(ctx context.Context, in *TrackUpload) (*TrackServerId, error)
}

// ================================
// TrackDataService Protobuf Client
// ================================

type trackDataServiceProtobufClient struct {
	client HTTPClient
	urls   [2]string
}

// NewTrackDataServiceProtobufClient creates a Protobuf client that implements the TrackDataService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewTrackDataServiceProtobufClient(addr string, client HTTPClient) TrackDataService {
	prefix := urlBase(addr) + TrackDataServicePathPrefix
	urls := [2]string{
		prefix + "StreamTrackData",
		prefix + "UploadTrackData",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &trackDataServiceProtobufClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &trackDataServiceProtobufClient{
		client: client,
		urls:   urls,
	}
}

func (c *trackDataServiceProtobufClient) StreamTrackData(ctx context.Context, in *UserTrack) (<-chan TrackChunkOrError, error) {
	ctx = ctxsetters.WithPackageName(ctx, "resonate.api.trackdata")
	ctx = ctxsetters.WithServiceName(ctx, "TrackDataService")
	ctx = ctxsetters.WithMethodName(ctx, "StreamTrackData")
	reqBodyBytes, err := proto.Marshal(in)
	if err != nil {
		return nil, clientError("failed to marshal proto request", err)
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)
	if err = ctx.Err(); err != nil {
		return nil, clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, c.urls[0], reqBody, "application/protobuf")
	if err != nil {
		return nil, clientError("could not build request", err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, clientError("failed to do request", err)
	}
	if err = ctx.Err(); err != nil {
		return nil, clientError("aborted because context was done", err)
	}
	if resp.StatusCode != 200 {
		return nil, errorFromResponse(resp)
	}

	respStream := make(chan TrackChunkOrError)
	go func() {
		defer func() {
			resp.Body.Close()
			close(respStream)
		}()
		reader := protoStreamReader{
			r:       bufio.NewReader(resp.Body),
			maxSize: 1 << 21, // 1GB
		}
		out := new(TrackChunk)
		for {
			if err = reader.Read(out); err != nil {
				if err == io.EOF {
					return
				}
				respStream <- TrackChunkOrError{Err: err}
				return
			}
			respStream <- TrackChunkOrError{Msg: out}
		}
	}()
	return respStream, nil
}

func (c *trackDataServiceProtobufClient) UploadTrackData(ctx context.Context, in *TrackUpload) (*TrackServerId, error) {
	ctx = ctxsetters.WithPackageName(ctx, "resonate.api.trackdata")
	ctx = ctxsetters.WithServiceName(ctx, "TrackDataService")
	ctx = ctxsetters.WithMethodName(ctx, "UploadTrackData")
	out := new(TrackServerId)
	err := doProtobufRequest(ctx, c.client, c.urls[1], in, out)
	return out, err
}

// ============================
// TrackDataService JSON Client
// ============================

type trackDataServiceJSONClient struct {
	client HTTPClient
	urls   [2]string
}

// NewTrackDataServiceJSONClient creates a JSON client that implements the TrackDataService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewTrackDataServiceJSONClient(addr string, client HTTPClient) TrackDataService {
	prefix := urlBase(addr) + TrackDataServicePathPrefix
	urls := [2]string{
		prefix + "StreamTrackData",
		prefix + "UploadTrackData",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &trackDataServiceJSONClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &trackDataServiceJSONClient{
		client: client,
		urls:   urls,
	}
}

func (c *trackDataServiceJSONClient) StreamTrackData(ctx context.Context, in *UserTrack) (<-chan TrackChunkOrError, error) {
	ctx = ctxsetters.WithPackageName(ctx, "resonate.api.trackdata")
	ctx = ctxsetters.WithServiceName(ctx, "TrackDataService")
	ctx = ctxsetters.WithMethodName(ctx, "StreamTrackData")
	reqBodyBytes, err := proto.Marshal(in)
	if err != nil {
		return nil, clientError("failed to marshal proto request", err)
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)
	if err = ctx.Err(); err != nil {
		return nil, clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, c.urls[0], reqBody, "application/json")
	if err != nil {
		return nil, clientError("could not build request", err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, clientError("failed to do request", err)
	}
	if err = ctx.Err(); err != nil {
		return nil, clientError("aborted because context was done", err)
	}
	if resp.StatusCode != 200 {
		return nil, errorFromResponse(resp)
	}

	respStream := make(chan TrackChunkOrError)
	go func() {
		defer func() {
			resp.Body.Close()
			close(respStream)
		}()
		reader, err := newJSONStreamReader(resp.Body)
		if err != nil {
			respStream <- TrackChunkOrError{Err: err}
			return
		}
		out := new(TrackChunk)
		for {
			if err = reader.Read(out); err != nil {
				if err == io.EOF {
					return
				}
				respStream <- TrackChunkOrError{Err: err}
				return
			}
			respStream <- TrackChunkOrError{Msg: out}
		}
	}()
	return respStream, nil
}

func (c *trackDataServiceJSONClient) UploadTrackData(ctx context.Context, in *TrackUpload) (*TrackServerId, error) {
	ctx = ctxsetters.WithPackageName(ctx, "resonate.api.trackdata")
	ctx = ctxsetters.WithServiceName(ctx, "TrackDataService")
	ctx = ctxsetters.WithMethodName(ctx, "UploadTrackData")
	out := new(TrackServerId)
	err := doJSONRequest(ctx, c.client, c.urls[1], in, out)
	return out, err
}

// ===============================
// TrackDataService Server Handler
// ===============================

type trackDataServiceServer struct {
	TrackDataService
	hooks *twirp.ServerHooks
}

func NewTrackDataServiceServer(svc TrackDataService, hooks *twirp.ServerHooks) TwirpServer {
	return &trackDataServiceServer{
		TrackDataService: svc,
		hooks:            hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *trackDataServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// TrackDataServicePathPrefix is used for all URL paths on a twirp TrackDataService server.
// Requests are always: POST TrackDataServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const TrackDataServicePathPrefix = "/twirp/resonate.api.trackdata.TrackDataService/"

func (s *trackDataServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "resonate.api.trackdata")
	ctx = ctxsetters.WithServiceName(ctx, "TrackDataService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/twirp/resonate.api.trackdata.TrackDataService/StreamTrackData":
		s.serveStreamTrackData(ctx, resp, req)
		return
	case "/twirp/resonate.api.trackdata.TrackDataService/UploadTrackData":
		s.serveUploadTrackData(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *trackDataServiceServer) serveStreamTrackData(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveStreamTrackDataJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveStreamTrackDataProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *trackDataServiceServer) serveStreamTrackDataJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
}

func (s *trackDataServiceServer) serveStreamTrackDataProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "StreamTrackData")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		err = wrapErr(err, "failed to read request body")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}
	reqContent := new(UserTrack)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request proto")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent <-chan TrackChunkOrError
	respFlusher, canFlush := resp.(http.Flusher)
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				if canFlush {
					respFlusher.Flush()
				}
				panic(r)
			}
		}()
		respContent, err = s.StreamTrackData(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil UserTrack and nil error while calling StreamTrackData. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)
	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(http.StatusOK)

	// Prepare trailer
	trailer := proto.NewBuffer(nil)
	_ = trailer.EncodeVarint((2 << 3) | 2) // field tag
	writeTrailer := func(err error) {
		defer func() {
			_, writeErr := resp.Write(trailer.Bytes())
			if writeErr != nil {
				// Ignored, for the same reason as in the writeError func
				_ = writeErr
			}
		}()
		if err == io.EOF {
			trailer.EncodeStringBytes("EOF")
			return
		}
		// Write trailer as json-encoded twirp err
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		statusCode := twirp.ServerHTTPStatusFromErrorCode(twerr.Code())
		ctx = ctxsetters.WithStatusCode(ctx, statusCode)
		ctx = callError(ctx, s.hooks, twerr)
		if encodeErr := trailer.EncodeStringBytes(string(marshalErrorToJSON(twerr))); encodeErr != nil {
			_ = trailer.EncodeStringBytes("{\"code\":\"" + string(twirp.Internal) + "\",\"msg\":\"There was an error but it could not be serialized into JSON\"}") // fallback
		}
	}

	messages := proto.NewBuffer(nil)
	for {
		var msg *TrackChunk
		select {
		case <-ctx.Done():
			return
		case msgOrErr, open := <-respContent:
			if !open {
				writeTrailer(io.EOF)
				return
			}
			if msgOrErr.Err != nil {
				writeTrailer(msgOrErr.Err)
				return
			}
			msg = msgOrErr.Msg
		}

		messages.Reset()
		_ = messages.EncodeVarint((1 << 3) | 2) // field tag
		err = messages.EncodeMessage(msg)
		if err != nil {
			err = wrapErr(err, "failed to marshal proto message")
			writeTrailer(err)
			return
		}

		_, err = resp.Write(messages.Bytes())
		if err != nil {
			err = wrapErr(err, "failed to send proto message")
			writeTrailer(err) // likely to fail on write, but try anyway to ensure ctx gets error code set for responseSent hook
			return
		}

		if canFlush {
			// TODO: Come up with a batching scheme to improve performance under high load
			//       and/or provide a hook for the respStream to control flushing the response.
			//       Flushing after each message dramatically reduces high-load throughput --
			//       difference can be more than 10x based on initial experiments
			respFlusher.Flush()
		}

		// TODO: Call a hook that we sent a message in a stream? (combine with flush hook?)
	}

	callResponseSent(ctx, s.hooks)
}

func (s *trackDataServiceServer) serveUploadTrackData(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveUploadTrackDataJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveUploadTrackDataProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *trackDataServiceServer) serveUploadTrackDataJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "UploadTrackData")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(TrackUpload)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request json")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *TrackServerId
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.UploadTrackData(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TrackServerId and nil error while calling UploadTrackData. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		err = wrapErr(err, "failed to marshal json response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	respBytes := buf.Bytes()
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *trackDataServiceServer) serveUploadTrackDataProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "UploadTrackData")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		err = wrapErr(err, "failed to read request body")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}
	reqContent := new(TrackUpload)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request proto")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *TrackServerId
	respFlusher, canFlush := resp.(http.Flusher)
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				if canFlush {
					respFlusher.Flush()
				}
				panic(r)
			}
		}()
		respContent, err = s.UploadTrackData(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil TrackUpload and nil error while calling UploadTrackData. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)
	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		err = wrapErr(err, "failed to marshal proto response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}

	callResponseSent(ctx, s.hooks)
}

func (s *trackDataServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor0, 0
}

func (s *trackDataServiceServer) ProtocGenTwirpVersion() string {
	return "v5.3.0"
}

type TrackChunkOrError struct {
	Msg *TrackChunk
	Err error
}

// =====
// Utils
// =====

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
//
// HTTPClient implementations should not follow redirects. Redirects are
// automatically disabled if *(net/http).Client is passed to client
// constructors. See the withoutRedirects function in this file for more
// details.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// TwirpServer is the interface generated server structs will support: they're
// HTTP handlers with additional methods for accessing metadata about the
// service. Those accessors are a low-level API for building reflection tools.
// Most people can think of TwirpServers as just http.Handlers.
type TwirpServer interface {
	http.Handler
	// ServiceDescriptor returns gzipped bytes describing the .proto file that
	// this service was generated from. Once unzipped, the bytes can be
	// unmarshalled as a
	// github.com/golang/protobuf/protoc-gen-go/descriptor.FileDescriptorProto.
	//
	// The returned integer is the index of this particular service within that
	// FileDescriptorProto's 'Service' slice of ServiceDescriptorProtos. This is a
	// low-level field, expected to be used for reflection.
	ServiceDescriptor() ([]byte, int)
	// ProtocGenTwirpVersion is the semantic version string of the version of
	// twirp used to generate this file.
	ProtocGenTwirpVersion() string
}

// WriteError writes an HTTP response with a valid Twirp error format.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func WriteError(resp http.ResponseWriter, err error) {
	writeError(context.Background(), resp, err, nil)
}

// writeError writes Twirp errors in the response and triggers hooks.
func writeError(ctx context.Context, resp http.ResponseWriter, err error, hooks *twirp.ServerHooks) {
	// Non-twirp errors are wrapped as Internal (default)
	twerr, ok := err.(twirp.Error)
	if !ok {
		twerr = twirp.InternalErrorWith(err)
	}

	statusCode := twirp.ServerHTTPStatusFromErrorCode(twerr.Code())
	ctx = ctxsetters.WithStatusCode(ctx, statusCode)
	ctx = callError(ctx, hooks, twerr)

	resp.Header().Set("Content-Type", "application/json") // Error responses are always JSON (instead of protobuf)
	resp.WriteHeader(statusCode)                          // HTTP response status code

	respBody := marshalErrorToJSON(twerr)
	_, writeErr := resp.Write(respBody)
	if writeErr != nil {
		// We have three options here. We could log the error, call the Error
		// hook, or just silently ignore the error.
		//
		// Logging is unacceptable because we don't have a user-controlled
		// logger; writing out to stderr without permission is too rude.
		//
		// Calling the Error hook would confuse users: it would mean the Error
		// hook got called twice for one request, which is likely to lead to
		// duplicated log messages and metrics, no matter how well we document
		// the behavior.
		//
		// Silently ignoring the error is our least-bad option. It's highly
		// likely that the connection is broken and the original 'err' says
		// so anyway.
		_ = writeErr
	}

	callResponseSent(ctx, hooks)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// getCustomHTTPReqHeaders retrieves a copy of any headers that are set in
// a context through the twirp.WithHTTPRequestHeaders function.
// If there are no headers set, or if they have the wrong type, nil is returned.
func getCustomHTTPReqHeaders(ctx context.Context) http.Header {
	header, ok := twirp.HTTPRequestHeaders(ctx)
	if !ok || header == nil {
		return nil
	}
	copied := make(http.Header)
	for k, vv := range header {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}
	return copied
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if customHeader := getCustomHTTPReqHeaders(ctx); customHeader != nil {
		req.Header = customHeader
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Twirp-Version", "v5.3.0")
	return req, nil
}

// JSON serialization for errors
type twerrJSON struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Meta map[string]string `json:"meta,omitempty"`
}

func (tj twerrJSON) toTwirpError() twirp.Error {
	errorCode := twirp.ErrorCode(tj.Code)
	if !twirp.IsValidErrorCode(errorCode) {
		msg := "invalid type returned from server error response: " + tj.Code
		return twirp.InternalError(msg)
	}

	twerr := twirp.NewError(errorCode, tj.Msg)
	for k, v := range tj.Meta {
		twerr = twerr.WithMeta(k, v)
	}
	return twerr
}

// marshalErrorToJSON returns JSON from a twirp.Error, that can be used as HTTP error response body.
// If serialization fails, it will use a descriptive Internal error instead.
func marshalErrorToJSON(twerr twirp.Error) []byte {
	// make sure that msg is not too large
	msg := twerr.Msg()
	if len(msg) > 1e6 {
		msg = msg[:1e6]
	}

	tj := twerrJSON{
		Code: string(twerr.Code()),
		Msg:  msg,
		Meta: twerr.MetaMap(),
	}

	buf, err := json.Marshal(&tj)
	if err != nil {
		buf = []byte("{\"type\": \"" + twirp.Internal + "\", \"msg\": \"There was an error but it could not be serialized into JSON\"}") // fallback
	}

	return buf
}

// errorFromResponse builds a twirp.Error from a non-200 HTTP response.
// If the response has a valid serialized Twirp error, then it's returned.
// If not, the response status code is used to generate a similar twirp
// error. See twirpErrorFromIntermediary for more info on intermediary errors.
func errorFromResponse(resp *http.Response) twirp.Error {
	statusCode := resp.StatusCode
	statusText := http.StatusText(statusCode)

	if isHTTPRedirect(statusCode) {
		// Unexpected redirect: it must be an error from an intermediary.
		// Twirp clients don't follow redirects automatically, Twirp only handles
		// POST requests, redirects should only happen on GET and HEAD requests.
		location := resp.Header.Get("Location")
		msg := fmt.Sprintf("unexpected HTTP status code %d %q received, Location=%q", statusCode, statusText, location)
		return twirpErrorFromIntermediary(statusCode, msg, location)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clientError("failed to read server error response body", err)
	}
	var tj twerrJSON
	if err := json.Unmarshal(respBodyBytes, &tj); err != nil {
		// Invalid JSON response; it must be an error from an intermediary.
		msg := fmt.Sprintf("Error from intermediary with HTTP status code %d %q", statusCode, statusText)
		return twirpErrorFromIntermediary(statusCode, msg, string(respBodyBytes))
	}

	return tj.toTwirpError()
}

// twirpErrorFromIntermediary maps HTTP errors from non-twirp sources to twirp errors.
// The mapping is similar to gRPC: https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md.
// Returned twirp Errors have some additional metadata for inspection.
func twirpErrorFromIntermediary(status int, msg string, bodyOrLocation string) twirp.Error {
	var code twirp.ErrorCode
	if isHTTPRedirect(status) { // 3xx
		code = twirp.Internal
	} else {
		switch status {
		case 400: // Bad Request
			code = twirp.Internal
		case 401: // Unauthorized
			code = twirp.Unauthenticated
		case 403: // Forbidden
			code = twirp.PermissionDenied
		case 404: // Not Found
			code = twirp.BadRoute
		case 429, 502, 503, 504: // Too Many Requests, Bad Gateway, Service Unavailable, Gateway Timeout
			code = twirp.Unavailable
		default: // All other codes
			code = twirp.Unknown
		}
	}

	twerr := twirp.NewError(code, msg)
	twerr = twerr.WithMeta("http_error_from_intermediary", "true") // to easily know if this error was from intermediary
	twerr = twerr.WithMeta("status_code", strconv.Itoa(status))
	if isHTTPRedirect(status) {
		twerr = twerr.WithMeta("location", bodyOrLocation)
	} else {
		twerr = twerr.WithMeta("body", bodyOrLocation)
	}
	return twerr
}
func isHTTPRedirect(status int) bool {
	return status >= 300 && status <= 399
}

// wrappedError implements the github.com/pkg/errors.Causer interface, allowing errors to be
// examined for their root cause.
type wrappedError struct {
	msg   string
	cause error
}

func wrapErr(err error, msg string) error { return &wrappedError{msg: msg, cause: err} }
func (e *wrappedError) Cause() error      { return e.cause }
func (e *wrappedError) Error() string     { return e.msg + ": " + e.cause.Error() }

// clientError adds consistency to errors generated in the client
func clientError(desc string, err error) twirp.Error {
	return twirp.InternalErrorWith(wrapErr(err, desc))
}

// badRouteError is used when the twirp server cannot route a request
func badRouteError(msg string, method, url string) twirp.Error {
	err := twirp.NewError(twirp.BadRoute, msg)
	err = err.WithMeta("twirp_invalid_route", method+" "+url)
	return err
}

// The standard library will, by default, redirect requests (including POSTs) if it gets a 302 or
// 303 response, and also 301s in go1.8. It redirects by making a second request, changing the
// method to GET and removing the body. This produces very confusing error messages, so instead we
// set a redirect policy that always errors. This stops Go from executing the redirect.
//
// We have to be a little careful in case the user-provided http.Client has its own CheckRedirect
// policy - if so, we'll run through that policy first.
//
// Because this requires modifying the http.Client, we make a new copy of the client and return it.
func withoutRedirects(in *http.Client) *http.Client {
	copy := *in
	copy.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if in.CheckRedirect != nil {
			// Run the input's redirect if it exists, in case it has side effects, but ignore any error it
			// returns, since we want to use ErrUseLastResponse.
			err := in.CheckRedirect(req, via)
			_ = err // Silly, but this makes sure generated code passes errcheck -blank, which some people use.
		}
		return http.ErrUseLastResponse
	}
	return &copy
}

// doProtobufRequest is common code to make a request to the remote twirp service.
func doProtobufRequest(ctx context.Context, client HTTPClient, url string, in, out proto.Message) (err error) {
	reqBodyBytes, err := proto.Marshal(in)
	if err != nil {
		return clientError("failed to marshal proto request", err)
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, url, reqBody, "application/protobuf")
	if err != nil {
		return clientError("could not build request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return clientError("failed to do request", err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = clientError("failed to close response body", cerr)
		}
	}()

	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clientError("failed to read response body", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if err = proto.Unmarshal(respBodyBytes, out); err != nil {
		return clientError("failed to unmarshal proto response", err)
	}
	return nil
}

// doJSONRequest is common code to make a request to the remote twirp service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out proto.Message) (err error) {
	reqBody := bytes.NewBuffer(nil)
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(reqBody, in); err != nil {
		return clientError("failed to marshal json request", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, url, reqBody, "application/json")
	if err != nil {
		return clientError("could not build request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return clientError("failed to do request", err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = clientError("failed to close response body", cerr)
		}
	}()

	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(resp.Body, out); err != nil {
		return clientError("failed to unmarshal json response", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}
	return nil
}

// Call twirp.ServerHooks.RequestReceived if the hook is available
func callRequestReceived(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestReceived == nil {
		return ctx, nil
	}
	return h.RequestReceived(ctx)
}

// Call twirp.ServerHooks.RequestRouted if the hook is available
func callRequestRouted(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestRouted == nil {
		return ctx, nil
	}
	return h.RequestRouted(ctx)
}

// Call twirp.ServerHooks.ResponsePrepared if the hook is available
func callResponsePrepared(ctx context.Context, h *twirp.ServerHooks) context.Context {
	if h == nil || h.ResponsePrepared == nil {
		return ctx
	}
	return h.ResponsePrepared(ctx)
}

// Call twirp.ServerHooks.ResponseSent if the hook is available
func callResponseSent(ctx context.Context, h *twirp.ServerHooks) {
	if h == nil || h.ResponseSent == nil {
		return
	}
	h.ResponseSent(ctx)
}

// Call twirp.ServerHooks.Error if the hook is available
func callError(ctx context.Context, h *twirp.ServerHooks, err twirp.Error) context.Context {
	if h == nil || h.Error == nil {
		return ctx
	}
	return h.Error(ctx, err)
}

type protoStreamReader struct {
	r       *bufio.Reader
	maxSize int
}

func (r protoStreamReader) Read(msg proto.Message) error {
	// Get next field tag.
	tag, err := binary.ReadUvarint(r.r)
	if err != nil {
		return err
	}

	const (
		msgTag     = (1 << 3) | 2
		trailerTag = (2 << 3) | 2
	)

	if tag == trailerTag { // Received a json twirp error or "EOF"
		// Read the length delimiter
		l, err := binary.ReadUvarint(r.r)
		if err != nil {
			return clientError("unable to read trailer's length delimiter", err)
		}
		sb := new(bytes.Buffer)
		sb.Grow(int(l))
		_, err = io.Copy(sb, r.r)
		if err != nil {
			return clientError("unable to read trailer", err)
		}
		if sb.String() == "EOF" {
			return io.EOF
		}
		var tj twerrJSON
		if err = json.Unmarshal([]byte(sb.String()), &tj); err != nil {
			return clientError("unable to decode trailer", err)
		}
		return tj.toTwirpError()
	}

	if tag != msgTag {
		return fmt.Errorf("invalid field tag: %v", tag)
	}

	// This is a real message. How long is it?
	l, err := binary.ReadUvarint(r.r)
	if err != nil {
		return err
	}
	if int(l) < 0 || int(l) > r.maxSize {
		return io.ErrShortBuffer
	}
	buf := make([]byte, int(l))

	// Go ahead and read a message.
	if _, err = io.ReadFull(r.r, buf); err != nil {
		return err
	}

	if err = proto.Unmarshal(buf, msg); err != nil {
		return err
	}
	return nil
}

type jsonStreamReader struct {
	dec               *json.Decoder
	unmarshaler       *jsonpb.Unmarshaler
	messageStreamDone bool
}

func newJSONStreamReader(r io.Reader) (*jsonStreamReader, error) {
	// stream should start with {"messages":[
	dec := json.NewDecoder(r)
	t, err := dec.Token()
	if err != nil {
		return nil, err
	}
	delim, ok := t.(json.Delim)
	if !ok || delim != '{' {
		return nil, fmt.Errorf("missing leading { in JSON stream, found %q", t)
	}

	t, err = dec.Token()
	if err != nil {
		return nil, err
	}
	key, ok := t.(string)
	if !ok || key != "messages" {
		return nil, fmt.Errorf("missing \"messages\" key in JSON stream, found %q", t)
	}

	t, err = dec.Token()
	if err != nil {
		return nil, err
	}
	delim, ok = t.(json.Delim)
	if !ok || delim != '[' {
		return nil, fmt.Errorf("missing [ to open messages array in JSON stream, found %q", t)
	}

	return &jsonStreamReader{
		dec:         dec,
		unmarshaler: &jsonpb.Unmarshaler{AllowUnknownFields: true},
	}, nil
}

func (r *jsonStreamReader) Read(msg proto.Message) error {
	if !r.messageStreamDone && r.dec.More() {
		return r.unmarshaler.UnmarshalNext(r.dec, msg)
	}

	// else, we hit the end of the message stream. finish up the array, and then read the trailer.
	r.messageStreamDone = true
	t, err := r.dec.Token()
	if err != nil {
		return err
	}
	delim, ok := t.(json.Delim)
	if !ok || delim != ']' {
		return fmt.Errorf("missing end of message array in JSON stream, found %q", t)
	}

	t, err = r.dec.Token()
	if err != nil {
		return err
	}
	key, ok := t.(string)
	if !ok || key != "trailer" {
		return fmt.Errorf("missing trailer after messages in JSON stream, found %q", t)
	}

	var tj twerrJSON
	err = r.dec.Decode(&tj)
	if err != nil {
		var eof string
		if _ = r.dec.Decode(&eof); eof == "EOF" {
			return io.EOF
		}
		return err
	}

	return tj.toTwirpError()
}

var twirpFileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4f, 0x32, 0x31,
	0x10, 0xce, 0xf2, 0xbe, 0x7c, 0xec, 0x20, 0xa2, 0x3d, 0x28, 0xe2, 0x05, 0xd7, 0x60, 0x38, 0xad,
	0x46, 0x0f, 0x1e, 0x4d, 0x50, 0x0f, 0x5c, 0x8c, 0x59, 0xe4, 0x42, 0x62, 0x48, 0xa1, 0x4d, 0xdc,
	0xe0, 0xb6, 0x4d, 0x3b, 0x4b, 0xc2, 0x8f, 0xf4, 0x3f, 0x99, 0xce, 0x22, 0x60, 0x82, 0xdc, 0x3a,
	0xf3, 0x7c, 0x6c, 0x9f, 0x67, 0x0b, 0xc7, 0xd6, 0xcc, 0xae, 0x9d, 0xb4, 0x8b, 0x74, 0x26, 0x63,
	0x63, 0x35, 0x6a, 0x76, 0x62, 0xa5, 0xd3, 0x8a, 0xa3, 0x8c, 0xb9, 0x49, 0x63, 0xb4, 0x7c, 0x36,
	0x17, 0x1c, 0x79, 0xf4, 0x00, 0xe1, 0xc8, 0x49, 0xfb, 0xe6, 0x17, 0xec, 0x14, 0xaa, 0xb9, 0x93,
	0x76, 0x92, 0x8a, 0x56, 0xd0, 0x09, 0x7a, 0x61, 0x52, 0xf1, 0xe3, 0x40, 0xb0, 0x33, 0xa8, 0x91,
	0xc4, 0x23, 0x25, 0x42, 0xaa, 0x34, 0x0f, 0x44, 0x74, 0x0f, 0x0d, 0x12, 0x0f, 0xa5, 0x5d, 0x10,
	0xf7, 0x0a, 0x9a, 0x05, 0xd7, 0xd1, 0x66, 0x63, 0xd6, 0xc0, 0x6d, 0x5e, 0x24, 0x00, 0x48, 0xf8,
	0xf8, 0x91, 0xab, 0x39, 0xeb, 0xc2, 0xa1, 0x43, 0x6e, 0x71, 0x62, 0xb4, 0x4b, 0x31, 0xd5, 0x8a,
	0x44, 0xe5, 0xa4, 0x41, 0xdb, 0xd7, 0xd5, 0x92, 0x9d, 0x43, 0xa8, 0xf2, 0x6c, 0x32, 0x5d, 0xa2,
	0x74, 0x74, 0x93, 0x72, 0x52, 0x53, 0x79, 0xd6, 0xf7, 0x33, 0x63, 0xf0, 0xdf, 0x67, 0x6a, 0xfd,
	0xeb, 0x04, 0xbd, 0x83, 0x84, 0xce, 0xd1, 0x0b, 0xd4, 0xe9, 0x2b, 0x23, 0xf3, 0xa9, 0xb9, 0xf0,
	0x14, 0xc5, 0x33, 0xb9, 0xba, 0x11, 0x9d, 0xb7, 0x53, 0x97, 0x7e, 0xa5, 0xde, 0xe5, 0x57, 0x85,
	0xf2, 0x73, 0x66, 0x70, 0x79, 0xfb, 0x15, 0xc0, 0x11, 0x39, 0x3f, 0x71, 0xe4, 0xc3, 0xa2, 0x6b,
	0x36, 0x86, 0xe6, 0x10, 0xad, 0xe4, 0xd9, 0x1a, 0x61, 0x17, 0xf1, 0xee, 0xe6, 0xe3, 0x75, 0xed,
	0xed, 0xe8, 0x2f, 0xca, 0xa6, 0x9f, 0x9b, 0x80, 0xbd, 0x43, 0xb3, 0x08, 0xb1, 0xf1, 0xbe, 0xdc,
	0x2b, 0x2c, 0xd8, 0xed, 0xee, 0x5e, 0xd2, 0xcf, 0xef, 0xe8, 0xd7, 0xc7, 0xe1, 0x1a, 0x9a, 0x56,
	0xe8, 0xd1, 0xdc, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xef, 0xc7, 0xf9, 0x49, 0x02, 0x00,
	0x00,
}

/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// request.go - the custom HTTP request for BCE

package http

import (
    "bufio"
    "bytes"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"

    "baidubce/util"
)

// BodyStream is used for internal request body stream implements the io.ReadCloser interface
// for adapting the `Body` field in the net/http.Request, as well as the `Len` method to set
// the content-length.
type BodyStream struct {
    Entity io.ReadCloser // abstract entity that can `Read` and `Close`
    Size   int64         // the body size that returned by `Len`
}

func (body *BodyStream) Read(p []byte) (int, error) { return body.Entity.Read(p) }

func (body *BodyStream) Close() error { return body.Entity.Close() }

func (body *BodyStream) Len() int64 { return body.Size }

// BufferedFileStream is an adapter of buffered file stream to `BodyStream`.
type bufferedFileStream struct {
    buffer *bufio.Reader
    closer io.Closer
}

func (buf *bufferedFileStream) Read(p []byte) (int, error) { return buf.buffer.Read(p) }

func (buf *bufferedFileStream) Close() error { return buf.closer.Close() }

func NewBodyStreamFromFile(fname string) (*BodyStream, error) {
    file, err := os.Open(fname)
    if err != nil {
        return nil, err
    }
    fileInfo, e := file.Stat()
    if e != nil {
        return nil, e
    }
    adapter := &bufferedFileStream{bufio.NewReader(file), file}
    return &BodyStream{adapter, fileInfo.Size()}, nil
}

// StringStream is an adapter of common strings to `BodyStream`.
type stringStream struct { buffer *bytes.Buffer }

func (str *stringStream) Read(p []byte) (int, error) { return str.buffer.Read(p) }

func (str *stringStream) Close() error { return nil }

func NewBodyStreamFromBytes(stream []byte) *BodyStream {
    buf := bytes.NewBuffer(stream)
    adapter := &stringStream{buf}
    return &BodyStream{adapter, int64(buf.Len())}
}

func NewBodyStreamFromString(stream string) *BodyStream {
    buf := bytes.NewBufferString(stream)
    adapter := &stringStream{buf}
    return &BodyStream{adapter, int64(buf.Len())}
}

// Reauest stands for the general http request structure to make request to the BCE services.
type Request struct {
    protocol    string
    host        string
    port        int
    method      string
    uri         string
    proxyUrl    string
    timeout     int
    headers     map[string]string
    params      map[string]string
    body        *BodyStream // Optional stream from which to read the payload
}

func (req *Request) Protocol() string {
    return req.protocol
}

func (req *Request) SetProtocol(protocol string) {
    req.protocol = protocol
}

func (req *Request) Endpoint() string {
    return req.protocol + "://" + req.host
}

func (req *Request) SetEndpoint(endpoint string) {
    pos := strings.Index(endpoint, "://")
    rest := endpoint
    if pos != -1 {
        req.protocol = endpoint[0:pos]
        rest = endpoint[pos + 3:]
    } else {
        req.protocol = "http"
    }

    req.SetHost(rest)
}

func (req *Request) Host() string {
    return req.host
}

func (req *Request) SetHost(host string) {
    req.host = host
    pos := strings.Index(host, ":")
    if pos != -1 {
        p, e := strconv.Atoi(host[pos + 1:])
        if e == nil {
            req.port = p
        }
    }

    if req.port == 0 {
        if req.protocol == "http" {
            req.port = 80
        } else if req.protocol == "https" {
            req.port = 443
        }
    }
}

func (req *Request) Port() int {
    return req.port
}

func (req *Request) SetPort(port int) {
    // Port can be set by the endpoint or host, this method is rarely used.
    req.port = port
}

func (req *Request) Headers() map[string]string {
    return req.headers
}

func (req *Request) SetHeaders(headers map[string]string) {
    req.headers = headers
}

func (req *Request) Header(key string) string {
    if v, ok := req.headers[key]; ok {
        return v
    }
    return ""
}

func (req *Request) SetHeader(key, value string) {
    if req.headers == nil {
        req.headers = make(map[string]string)
    }
    req.headers[key] = value
}

func (req *Request) Params() map[string]string {
    return req.params
}

func (req *Request) SetParams(params map[string]string) {
    req.params = params
}

func (req *Request) Param(key string) string {
    if v, ok := req.params[key]; ok {
        return v
    }
    return ""
}

func (req *Request) SetParam(key, value string) {
    if req.params == nil {
        req.params = make(map[string]string)
    }
    req.params[key] = value
}

func (req *Request) QueryString() string {
    buf := make([]string, 0, len(req.params))
    for k, v := range req.params {
        buf = append(buf, util.UriEncode(k, true) + "=" + util.UriEncode(v, true))
    }
    return strings.Join(buf, "&")
}

func (req *Request) Method() string {
    return req.method
}

func (req *Request) SetMethod(method string) {
    req.method = method
}

func (req *Request) Uri() string {
    return req.uri
}

func (req *Request) SetUri(uri string) {
    req.uri = uri
}

func (req *Request) ProxyUrl() string {
    return req.proxyUrl
}

func (req *Request) SetProxyUrl(url string) {
    req.proxyUrl = url
}

func (req *Request) Timeout() int {
    return req.timeout
}

func (req *Request) SetTimeout(timeout int) {
    req.timeout = timeout
}

func (req *Request) Body() *BodyStream {
    return req.body
}

func (req *Request) SetBody(stream *BodyStream) {
    req.body = stream
}

func (req *Request) GenerateUrl(addPort bool) string {
    if addPort {
        return fmt.Sprintf("%s://%s:%d%s?%s",
                           req.protocol, req.host, req.port, req.uri, req.QueryString())
    } else {
        return fmt.Sprintf("%s://%s%s?%s", req.protocol, req.host, req.uri, req.QueryString())
    }
}

func (req *Request) String() string {
    header := make([]string, 0, len(req.headers))
    for k, v := range req.headers {
        header = append(header, "\t" + k + "=" + v)
    }
    return fmt.Sprintf("\t%s %s\n%v",
                       req.method, req.GenerateUrl(false), strings.Join(header, "\n"))
}


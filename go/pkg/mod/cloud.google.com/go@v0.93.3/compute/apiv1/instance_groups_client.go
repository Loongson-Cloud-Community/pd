// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"sort"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newInstanceGroupsClientHook clientHook

// InstanceGroupsCallOptions contains the retry settings for each method of InstanceGroupsClient.
type InstanceGroupsCallOptions struct {
	AddInstances    []gax.CallOption
	AggregatedList  []gax.CallOption
	Delete          []gax.CallOption
	Get             []gax.CallOption
	Insert          []gax.CallOption
	List            []gax.CallOption
	ListInstances   []gax.CallOption
	RemoveInstances []gax.CallOption
	SetNamedPorts   []gax.CallOption
}

// internalInstanceGroupsClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalInstanceGroupsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddInstances(context.Context, *computepb.AddInstancesInstanceGroupRequest, ...gax.CallOption) (*Operation, error)
	AggregatedList(context.Context, *computepb.AggregatedListInstanceGroupsRequest, ...gax.CallOption) *InstanceGroupsScopedListPairIterator
	Delete(context.Context, *computepb.DeleteInstanceGroupRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetInstanceGroupRequest, ...gax.CallOption) (*computepb.InstanceGroup, error)
	Insert(context.Context, *computepb.InsertInstanceGroupRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListInstanceGroupsRequest, ...gax.CallOption) *InstanceGroupIterator
	ListInstances(context.Context, *computepb.ListInstancesInstanceGroupsRequest, ...gax.CallOption) *InstanceWithNamedPortsIterator
	RemoveInstances(context.Context, *computepb.RemoveInstancesInstanceGroupRequest, ...gax.CallOption) (*Operation, error)
	SetNamedPorts(context.Context, *computepb.SetNamedPortsInstanceGroupRequest, ...gax.CallOption) (*Operation, error)
}

// InstanceGroupsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The InstanceGroups API.
type InstanceGroupsClient struct {
	// The internal transport-dependent client.
	internalClient internalInstanceGroupsClient

	// The call options for this service.
	CallOptions *InstanceGroupsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *InstanceGroupsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *InstanceGroupsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *InstanceGroupsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddInstances adds a list of instances to the specified instance group. All of the instances in the instance group must be in the same network/subnetwork. Read  Adding instances for more information.
func (c *InstanceGroupsClient) AddInstances(ctx context.Context, req *computepb.AddInstancesInstanceGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.AddInstances(ctx, req, opts...)
}

// AggregatedList retrieves the list of instance groups and sorts them by zone.
func (c *InstanceGroupsClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListInstanceGroupsRequest, opts ...gax.CallOption) *InstanceGroupsScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified instance group. The instances in the group are not deleted. Note that instance group must not belong to a backend service. Read  Deleting an instance group for more information.
func (c *InstanceGroupsClient) Delete(ctx context.Context, req *computepb.DeleteInstanceGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified zonal instance group. Get a list of available zonal instance groups by making a list() request.
//
// For managed instance groups, use the instanceGroupManagers or regionInstanceGroupManagers methods instead.
func (c *InstanceGroupsClient) Get(ctx context.Context, req *computepb.GetInstanceGroupRequest, opts ...gax.CallOption) (*computepb.InstanceGroup, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates an instance group in the specified project using the parameters that are included in the request.
func (c *InstanceGroupsClient) Insert(ctx context.Context, req *computepb.InsertInstanceGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of zonal instance group resources contained within the specified zone.
//
// For managed instance groups, use the instanceGroupManagers or regionInstanceGroupManagers methods instead.
func (c *InstanceGroupsClient) List(ctx context.Context, req *computepb.ListInstanceGroupsRequest, opts ...gax.CallOption) *InstanceGroupIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// ListInstances lists the instances in the specified instance group. The orderBy query parameter is not supported.
func (c *InstanceGroupsClient) ListInstances(ctx context.Context, req *computepb.ListInstancesInstanceGroupsRequest, opts ...gax.CallOption) *InstanceWithNamedPortsIterator {
	return c.internalClient.ListInstances(ctx, req, opts...)
}

// RemoveInstances removes one or more instances from the specified instance group, but does not delete those instances.
//
// If the group is part of a backend service that has enabled connection draining, it can take up to 60 seconds after the connection draining duration before the VM instance is removed or deleted.
func (c *InstanceGroupsClient) RemoveInstances(ctx context.Context, req *computepb.RemoveInstancesInstanceGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.RemoveInstances(ctx, req, opts...)
}

// SetNamedPorts sets the named ports for the specified instance group.
func (c *InstanceGroupsClient) SetNamedPorts(ctx context.Context, req *computepb.SetNamedPortsInstanceGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.SetNamedPorts(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type instanceGroupsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewInstanceGroupsRESTClient creates a new instance groups rest client.
//
// The InstanceGroups API.
func NewInstanceGroupsRESTClient(ctx context.Context, opts ...option.ClientOption) (*InstanceGroupsClient, error) {
	clientOpts := append(defaultInstanceGroupsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &instanceGroupsRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &InstanceGroupsClient{internalClient: c, CallOptions: &InstanceGroupsCallOptions{}}, nil
}

func defaultInstanceGroupsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *instanceGroupsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *instanceGroupsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *instanceGroupsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AddInstances adds a list of instances to the specified instance group. All of the instances in the instance group must be in the same network/subnetwork. Read  Adding instances for more information.
func (c *instanceGroupsRESTClient) AddInstances(ctx context.Context, req *computepb.AddInstancesInstanceGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetInstanceGroupsAddInstancesRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/instanceGroups/%v/addInstances", req.GetProject(), req.GetZone(), req.GetInstanceGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// AggregatedList retrieves the list of instance groups and sorts them by zone.
func (c *instanceGroupsRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListInstanceGroupsRequest, opts ...gax.CallOption) *InstanceGroupsScopedListPairIterator {
	it := &InstanceGroupsScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListInstanceGroupsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]InstanceGroupsScopedListPair, string, error) {
		resp := &computepb.InstanceGroupAggregatedList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/aggregated/instanceGroups", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.IncludeAllScopes != nil {
			params.Add("includeAllScopes", fmt.Sprintf("%v", req.GetIncludeAllScopes()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if httpRsp.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf(httpRsp.Status)
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp

		elems := make([]InstanceGroupsScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, InstanceGroupsScopedListPair{k, v})
		}
		sort.Slice(elems, func(i, j int) bool { return elems[i].Key < elems[j].Key })

		return elems, resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Delete deletes the specified instance group. The instances in the group are not deleted. Note that instance group must not belong to a backend service. Read  Deleting an instance group for more information.
func (c *instanceGroupsRESTClient) Delete(ctx context.Context, req *computepb.DeleteInstanceGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/instanceGroups/%v", req.GetProject(), req.GetZone(), req.GetInstanceGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// Get returns the specified zonal instance group. Get a list of available zonal instance groups by making a list() request.
//
// For managed instance groups, use the instanceGroupManagers or regionInstanceGroupManagers methods instead.
func (c *instanceGroupsRESTClient) Get(ctx context.Context, req *computepb.GetInstanceGroupRequest, opts ...gax.CallOption) (*computepb.InstanceGroup, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/instanceGroups/%v", req.GetProject(), req.GetZone(), req.GetInstanceGroup())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.InstanceGroup{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates an instance group in the specified project using the parameters that are included in the request.
func (c *instanceGroupsRESTClient) Insert(ctx context.Context, req *computepb.InsertInstanceGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetInstanceGroupResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/instanceGroups", req.GetProject(), req.GetZone())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// List retrieves the list of zonal instance group resources contained within the specified zone.
//
// For managed instance groups, use the instanceGroupManagers or regionInstanceGroupManagers methods instead.
func (c *instanceGroupsRESTClient) List(ctx context.Context, req *computepb.ListInstanceGroupsRequest, opts ...gax.CallOption) *InstanceGroupIterator {
	it := &InstanceGroupIterator{}
	req = proto.Clone(req).(*computepb.ListInstanceGroupsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.InstanceGroup, string, error) {
		resp := &computepb.InstanceGroupList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/instanceGroups", req.GetProject(), req.GetZone())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if httpRsp.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf(httpRsp.Status)
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// ListInstances lists the instances in the specified instance group. The orderBy query parameter is not supported.
func (c *instanceGroupsRESTClient) ListInstances(ctx context.Context, req *computepb.ListInstancesInstanceGroupsRequest, opts ...gax.CallOption) *InstanceWithNamedPortsIterator {
	it := &InstanceWithNamedPortsIterator{}
	req = proto.Clone(req).(*computepb.ListInstancesInstanceGroupsRequest)
	m := protojson.MarshalOptions{AllowPartial: true, UseProtoNames: false}
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.InstanceWithNamedPorts, string, error) {
		resp := &computepb.InstanceGroupsListInstances{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		jsonReq, err := m.Marshal(req)
		if err != nil {
			return nil, "", err
		}

		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/instanceGroups/%v/listInstances", req.GetProject(), req.GetZone(), req.GetInstanceGroup())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if httpRsp.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf(httpRsp.Status)
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// RemoveInstances removes one or more instances from the specified instance group, but does not delete those instances.
//
// If the group is part of a backend service that has enabled connection draining, it can take up to 60 seconds after the connection draining duration before the VM instance is removed or deleted.
func (c *instanceGroupsRESTClient) RemoveInstances(ctx context.Context, req *computepb.RemoveInstancesInstanceGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetInstanceGroupsRemoveInstancesRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/instanceGroups/%v/removeInstances", req.GetProject(), req.GetZone(), req.GetInstanceGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// SetNamedPorts sets the named ports for the specified instance group.
func (c *instanceGroupsRESTClient) SetNamedPorts(ctx context.Context, req *computepb.SetNamedPortsInstanceGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetInstanceGroupsSetNamedPortsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/instanceGroups/%v/setNamedPorts", req.GetProject(), req.GetZone(), req.GetInstanceGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// InstanceGroupIterator manages a stream of *computepb.InstanceGroup.
type InstanceGroupIterator struct {
	items    []*computepb.InstanceGroup
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.InstanceGroup, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *InstanceGroupIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *InstanceGroupIterator) Next() (*computepb.InstanceGroup, error) {
	var item *computepb.InstanceGroup
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *InstanceGroupIterator) bufLen() int {
	return len(it.items)
}

func (it *InstanceGroupIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// InstanceGroupsScopedListPair is a holder type for string/*computepb.InstanceGroupsScopedList map entries
type InstanceGroupsScopedListPair struct {
	Key   string
	Value *computepb.InstanceGroupsScopedList
}

// InstanceGroupsScopedListPairIterator manages a stream of InstanceGroupsScopedListPair.
type InstanceGroupsScopedListPairIterator struct {
	items    []InstanceGroupsScopedListPair
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []InstanceGroupsScopedListPair, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *InstanceGroupsScopedListPairIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *InstanceGroupsScopedListPairIterator) Next() (InstanceGroupsScopedListPair, error) {
	var item InstanceGroupsScopedListPair
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *InstanceGroupsScopedListPairIterator) bufLen() int {
	return len(it.items)
}

func (it *InstanceGroupsScopedListPairIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// InstanceWithNamedPortsIterator manages a stream of *computepb.InstanceWithNamedPorts.
type InstanceWithNamedPortsIterator struct {
	items    []*computepb.InstanceWithNamedPorts
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.InstanceWithNamedPorts, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *InstanceWithNamedPortsIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *InstanceWithNamedPortsIterator) Next() (*computepb.InstanceWithNamedPorts, error) {
	var item *computepb.InstanceWithNamedPorts
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *InstanceWithNamedPortsIterator) bufLen() int {
	return len(it.items)
}

func (it *InstanceWithNamedPortsIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

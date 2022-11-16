// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package fcmdata provides access to the Firebase Cloud Messaging Data API.
//
// For product documentation, see: https://firebase.google.com/docs/cloud-messaging
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/fcmdata/v1beta1"
//   ...
//   ctx := context.Background()
//   fcmdataService, err := fcmdata.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   fcmdataService, err := fcmdata.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   fcmdataService, err := fcmdata.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package fcmdata // import "google.golang.org/api/fcmdata/v1beta1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "fcmdata:v1beta1"
const apiName = "fcmdata"
const apiVersion = "v1beta1"
const basePath = "https://fcmdata.googleapis.com/"
const mtlsBasePath = "https://fcmdata.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// See, edit, configure, and delete your Google Cloud data and see the
	// email address for your Google Account.
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/cloud-platform",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.AndroidApps = NewProjectsAndroidAppsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	AndroidApps *ProjectsAndroidAppsService
}

func NewProjectsAndroidAppsService(s *Service) *ProjectsAndroidAppsService {
	rs := &ProjectsAndroidAppsService{s: s}
	rs.DeliveryData = NewProjectsAndroidAppsDeliveryDataService(s)
	return rs
}

type ProjectsAndroidAppsService struct {
	s *Service

	DeliveryData *ProjectsAndroidAppsDeliveryDataService
}

func NewProjectsAndroidAppsDeliveryDataService(s *Service) *ProjectsAndroidAppsDeliveryDataService {
	rs := &ProjectsAndroidAppsDeliveryDataService{s: s}
	return rs
}

type ProjectsAndroidAppsDeliveryDataService struct {
	s *Service
}

// GoogleFirebaseFcmDataV1beta1AndroidDeliveryData: Message delivery
// data for a given date, app, and analytics label combination.
type GoogleFirebaseFcmDataV1beta1AndroidDeliveryData struct {
	// AnalyticsLabel: The analytics label associated with the messages
	// sent. All messages sent without an analytics label will be grouped
	// together in a single entry.
	AnalyticsLabel string `json:"analyticsLabel,omitempty"`

	// AppId: The app ID to which the messages were sent.
	AppId string `json:"appId,omitempty"`

	// Data: The data for the specified appId, date, and analyticsLabel.
	Data *GoogleFirebaseFcmDataV1beta1Data `json:"data,omitempty"`

	// Date: The date represented by this entry.
	Date *GoogleTypeDate `json:"date,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AnalyticsLabel") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnalyticsLabel") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirebaseFcmDataV1beta1AndroidDeliveryData) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirebaseFcmDataV1beta1AndroidDeliveryData
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirebaseFcmDataV1beta1Data: Data detailing messaging delivery
type GoogleFirebaseFcmDataV1beta1Data struct {
	// CountMessagesAccepted: Count of messages accepted by FCM intended to
	// Android devices. The targeted device must have opted in to the
	// collection of usage and diagnostic information.
	CountMessagesAccepted int64 `json:"countMessagesAccepted,omitempty,string"`

	// DeliveryPerformancePercents: Additional information about delivery
	// performance for messages that were successfully delivered.
	DeliveryPerformancePercents *GoogleFirebaseFcmDataV1beta1DeliveryPerformancePercents `json:"deliveryPerformancePercents,omitempty"`

	// MessageInsightPercents: Additional general insights about message
	// delivery.
	MessageInsightPercents *GoogleFirebaseFcmDataV1beta1MessageInsightPercents `json:"messageInsightPercents,omitempty"`

	// MessageOutcomePercents: Mutually exclusive breakdown of message
	// delivery outcomes.
	MessageOutcomePercents *GoogleFirebaseFcmDataV1beta1MessageOutcomePercents `json:"messageOutcomePercents,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "CountMessagesAccepted") to unconditionally include in API requests.
	// By default, fields with empty or default values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CountMessagesAccepted") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirebaseFcmDataV1beta1Data) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirebaseFcmDataV1beta1Data
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirebaseFcmDataV1beta1DeliveryPerformancePercents: Overview of
// delivery performance for messages that were successfully delivered.
// All percentages are calculated with countMessagesAccepted as the
// denominator. These categories are not mutually exclusive; a message
// can be delayed for multiple reasons.
type GoogleFirebaseFcmDataV1beta1DeliveryPerformancePercents struct {
	// DelayedDeviceDoze: The percentage of accepted messages that were
	// delayed because the device was in doze mode. Only normal priority
	// messages
	// (https://firebase.google.com/docs/cloud-messaging/concept-options#setting-the-priority-of-a-message)
	// should be delayed due to doze mode.
	DelayedDeviceDoze float64 `json:"delayedDeviceDoze,omitempty"`

	// DelayedDeviceOffline: The percentage of accepted messages that were
	// delayed because the target device was not connected at the time of
	// sending. These messages were eventually delivered when the device
	// reconnected.
	DelayedDeviceOffline float64 `json:"delayedDeviceOffline,omitempty"`

	// DelayedMessageThrottled: The percentage of accepted messages that
	// were delayed due to message throttling, such as collapsible message
	// throttling
	// (https://firebase.google.com/docs/cloud-messaging/concept-options#collapsible_throttling)
	// or maximum message rate throttling
	// (https://firebase.google.com/docs/cloud-messaging/concept-options#device_throttling).
	DelayedMessageThrottled float64 `json:"delayedMessageThrottled,omitempty"`

	// DelayedUserStopped: The percentage of accepted messages that were
	// delayed because the intended device user-profile was stopped
	// (https://firebase.google.com/docs/cloud-messaging/android/receive#handling_messages)
	// on the target device at the time of the send. The messages were
	// eventually delivered when the user-profile was started again.
	DelayedUserStopped float64 `json:"delayedUserStopped,omitempty"`

	// DeliveredNoDelay: The percentage of accepted messages that were
	// delivered to the device without delay from the FCM system.
	DeliveredNoDelay float64 `json:"deliveredNoDelay,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DelayedDeviceDoze")
	// to unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DelayedDeviceDoze") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirebaseFcmDataV1beta1DeliveryPerformancePercents) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirebaseFcmDataV1beta1DeliveryPerformancePercents
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleFirebaseFcmDataV1beta1DeliveryPerformancePercents) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleFirebaseFcmDataV1beta1DeliveryPerformancePercents
	var s1 struct {
		DelayedDeviceDoze       gensupport.JSONFloat64 `json:"delayedDeviceDoze"`
		DelayedDeviceOffline    gensupport.JSONFloat64 `json:"delayedDeviceOffline"`
		DelayedMessageThrottled gensupport.JSONFloat64 `json:"delayedMessageThrottled"`
		DelayedUserStopped      gensupport.JSONFloat64 `json:"delayedUserStopped"`
		DeliveredNoDelay        gensupport.JSONFloat64 `json:"deliveredNoDelay"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.DelayedDeviceDoze = float64(s1.DelayedDeviceDoze)
	s.DelayedDeviceOffline = float64(s1.DelayedDeviceOffline)
	s.DelayedMessageThrottled = float64(s1.DelayedMessageThrottled)
	s.DelayedUserStopped = float64(s1.DelayedUserStopped)
	s.DeliveredNoDelay = float64(s1.DeliveredNoDelay)
	return nil
}

// GoogleFirebaseFcmDataV1beta1ListAndroidDeliveryDataResponse: Response
// message for ListAndroidDeliveryData.
type GoogleFirebaseFcmDataV1beta1ListAndroidDeliveryDataResponse struct {
	// AndroidDeliveryData: The delivery data for the provided app. There
	// will be one entry per combination of app, date, and analytics label.
	AndroidDeliveryData []*GoogleFirebaseFcmDataV1beta1AndroidDeliveryData `json:"androidDeliveryData,omitempty"`

	// NextPageToken: A token, which can be sent as `page_token` to retrieve
	// the next page. If this field is omitted, there are no subsequent
	// pages.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AndroidDeliveryData")
	// to unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AndroidDeliveryData") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirebaseFcmDataV1beta1ListAndroidDeliveryDataResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirebaseFcmDataV1beta1ListAndroidDeliveryDataResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirebaseFcmDataV1beta1MessageInsightPercents: Additional
// information about message delivery. All percentages are calculated
// with countMessagesAccepted as the denominator.
type GoogleFirebaseFcmDataV1beta1MessageInsightPercents struct {
	// PriorityLowered: The percentage of accepted messages that had their
	// priority lowered from high to normal due to app standby buckets
	// (https://firebase.google.com/docs/cloud-messaging/concept-options#setting-the-priority-of-a-message).
	PriorityLowered float64 `json:"priorityLowered,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PriorityLowered") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PriorityLowered") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirebaseFcmDataV1beta1MessageInsightPercents) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirebaseFcmDataV1beta1MessageInsightPercents
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleFirebaseFcmDataV1beta1MessageInsightPercents) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleFirebaseFcmDataV1beta1MessageInsightPercents
	var s1 struct {
		PriorityLowered gensupport.JSONFloat64 `json:"priorityLowered"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.PriorityLowered = float64(s1.PriorityLowered)
	return nil
}

// GoogleFirebaseFcmDataV1beta1MessageOutcomePercents: Percentage
// breakdown of message delivery outcomes. These categories are mutually
// exclusive. All percentages are calculated with countMessagesAccepted
// as the denominator. These categories may not account for all message
// outcomes.
type GoogleFirebaseFcmDataV1beta1MessageOutcomePercents struct {
	// Delivered: The percentage of all accepted messages that were
	// successfully delivered to the device.
	Delivered float64 `json:"delivered,omitempty"`

	// DroppedAppForceStopped: The percentage of accepted messages that were
	// dropped because the application was force stopped on the device at
	// the time of delivery and retries were unsuccessful.
	DroppedAppForceStopped float64 `json:"droppedAppForceStopped,omitempty"`

	// DroppedDeviceInactive: The percentage of accepted messages that were
	// dropped because the target device is inactive. FCM will drop messages
	// if the target device is deemed inactive by our servers. If a device
	// does reconnect, we call OnDeletedMessages()
	// (https://firebase.google.com/docs/cloud-messaging/android/receive#override-ondeletedmessages)
	// in our SDK instead of delivering the messages.
	DroppedDeviceInactive float64 `json:"droppedDeviceInactive,omitempty"`

	// DroppedTooManyPendingMessages: The percentage of accepted messages
	// that were dropped due to too many undelivered non-collapsible
	// messages
	// (https://firebase.google.com/docs/cloud-messaging/concept-options#collapsible_and_non-collapsible_messages).
	// Specifically, each app instance can only have 100 pending messages
	// stored on our servers for a device which is disconnected. When that
	// device reconnects, those messages are delivered. When there are more
	// than the maximum pending messages, we call OnDeletedMessages()
	// (https://firebase.google.com/docs/cloud-messaging/android/receive#override-ondeletedmessages)
	// in our SDK instead of delivering the messages.
	DroppedTooManyPendingMessages float64 `json:"droppedTooManyPendingMessages,omitempty"`

	// Pending: The percentage of messages accepted on this day that were
	// not dropped and not delivered, due to the device being disconnected
	// (as of the end of the America/Los_Angeles day when the message was
	// sent to FCM). A portion of these messages will be delivered the next
	// day when the device connects but others may be destined to devices
	// that ultimately never reconnect.
	Pending float64 `json:"pending,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Delivered") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Delivered") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirebaseFcmDataV1beta1MessageOutcomePercents) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirebaseFcmDataV1beta1MessageOutcomePercents
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleFirebaseFcmDataV1beta1MessageOutcomePercents) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleFirebaseFcmDataV1beta1MessageOutcomePercents
	var s1 struct {
		Delivered                     gensupport.JSONFloat64 `json:"delivered"`
		DroppedAppForceStopped        gensupport.JSONFloat64 `json:"droppedAppForceStopped"`
		DroppedDeviceInactive         gensupport.JSONFloat64 `json:"droppedDeviceInactive"`
		DroppedTooManyPendingMessages gensupport.JSONFloat64 `json:"droppedTooManyPendingMessages"`
		Pending                       gensupport.JSONFloat64 `json:"pending"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Delivered = float64(s1.Delivered)
	s.DroppedAppForceStopped = float64(s1.DroppedAppForceStopped)
	s.DroppedDeviceInactive = float64(s1.DroppedDeviceInactive)
	s.DroppedTooManyPendingMessages = float64(s1.DroppedTooManyPendingMessages)
	s.Pending = float64(s1.Pending)
	return nil
}

// GoogleTypeDate: Represents a whole or partial calendar date, such as
// a birthday. The time of day and time zone are either specified
// elsewhere or are insignificant. The date is relative to the Gregorian
// Calendar. This can represent one of the following: * A full date,
// with non-zero year, month, and day values * A month and day value,
// with a zero year, such as an anniversary * A year on its own, with
// zero month and day values * A year and month value, with a zero day,
// such as a credit card expiration date Related types are
// google.type.TimeOfDay and `google.protobuf.Timestamp`.
type GoogleTypeDate struct {
	// Day: Day of a month. Must be from 1 to 31 and valid for the year and
	// month, or 0 to specify a year by itself or a year and month where the
	// day isn't significant.
	Day int64 `json:"day,omitempty"`

	// Month: Month of a year. Must be from 1 to 12, or 0 to specify a year
	// without a month and day.
	Month int64 `json:"month,omitempty"`

	// Year: Year of the date. Must be from 1 to 9999, or 0 to specify a
	// date without a year.
	Year int64 `json:"year,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Day") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Day") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleTypeDate) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleTypeDate
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "fcmdata.projects.androidApps.deliveryData.list":

type ProjectsAndroidAppsDeliveryDataListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List aggregate delivery data for the given Android application.
//
// - parent: The application for which to list delivery data. Format:
//   `projects/{project_id}/androidApps/{app_id}`.
func (r *ProjectsAndroidAppsDeliveryDataService) List(parent string) *ProjectsAndroidAppsDeliveryDataListCall {
	c := &ProjectsAndroidAppsDeliveryDataListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of entries to return. The service may return fewer than this value.
// If unspecified, at most 1,000 entries will be returned. The maximum
// value is 10,000; values above 10,000 will be capped to 10,000. This
// default may change over time.
func (c *ProjectsAndroidAppsDeliveryDataListCall) PageSize(pageSize int64) *ProjectsAndroidAppsDeliveryDataListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token,
// received from a previous `ListAndroidDeliveryDataRequest` call.
// Provide this to retrieve the subsequent page. When paginating, all
// other parameters provided to `ListAndroidDeliveryDataRequest` must
// match the call that provided the page token.
func (c *ProjectsAndroidAppsDeliveryDataListCall) PageToken(pageToken string) *ProjectsAndroidAppsDeliveryDataListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAndroidAppsDeliveryDataListCall) Fields(s ...googleapi.Field) *ProjectsAndroidAppsDeliveryDataListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAndroidAppsDeliveryDataListCall) IfNoneMatch(entityTag string) *ProjectsAndroidAppsDeliveryDataListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAndroidAppsDeliveryDataListCall) Context(ctx context.Context) *ProjectsAndroidAppsDeliveryDataListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAndroidAppsDeliveryDataListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAndroidAppsDeliveryDataListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210812")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/deliveryData")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "fcmdata.projects.androidApps.deliveryData.list" call.
// Exactly one of
// *GoogleFirebaseFcmDataV1beta1ListAndroidDeliveryDataResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleFirebaseFcmDataV1beta1ListAndroidDeliveryDataResponse.ServerRes
// ponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsAndroidAppsDeliveryDataListCall) Do(opts ...googleapi.CallOption) (*GoogleFirebaseFcmDataV1beta1ListAndroidDeliveryDataResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleFirebaseFcmDataV1beta1ListAndroidDeliveryDataResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List aggregate delivery data for the given Android application.",
	//   "flatPath": "v1beta1/projects/{projectsId}/androidApps/{androidAppsId}/deliveryData",
	//   "httpMethod": "GET",
	//   "id": "fcmdata.projects.androidApps.deliveryData.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of entries to return. The service may return fewer than this value. If unspecified, at most 1,000 entries will be returned. The maximum value is 10,000; values above 10,000 will be capped to 10,000. This default may change over time.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A page token, received from a previous `ListAndroidDeliveryDataRequest` call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to `ListAndroidDeliveryDataRequest` must match the call that provided the page token.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The application for which to list delivery data. Format: `projects/{project_id}/androidApps/{app_id}`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/androidApps/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/deliveryData",
	//   "response": {
	//     "$ref": "GoogleFirebaseFcmDataV1beta1ListAndroidDeliveryDataResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAndroidAppsDeliveryDataListCall) Pages(ctx context.Context, f func(*GoogleFirebaseFcmDataV1beta1ListAndroidDeliveryDataResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

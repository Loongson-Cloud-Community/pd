// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/maps/routes/v1/waypoint.proto

package routes

import (
	reflect "reflect"
	sync "sync"

	latlng "google.golang.org/genproto/googleapis/type/latlng"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Encapsulates a waypoint. Waypoints mark both the beginning and end of a
// route, and include intermediate stops along the route.
type Waypoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Different ways to represent a location.
	//
	// Types that are assignable to LocationType:
	//	*Waypoint_Location
	//	*Waypoint_PlaceId
	LocationType isWaypoint_LocationType `protobuf_oneof:"location_type"`
	// Marks this waypoint as a milestone, as opposed to a stopping point. For
	// each non-via waypoint in the request, the response appends an entry to the
	// `legs` array to provide the details for stopovers on that leg of the
	// trip. Set this value to true when you want the route to pass through this
	// waypoint without stopping over. Via waypoints don't cause an entry to be
	// added to the `legs` array, but they do route the journey through the
	// waypoint. You can only set this value on waypoints that are intermediates.
	// If you set this field on terminal waypoints, then the request fails.
	Via bool `protobuf:"varint,3,opt,name=via,proto3" json:"via,omitempty"`
	// Indicates that the waypoint is meant for vehicles to stop at, where the
	// intention is to either pickup or drop-off. When you set this value, the
	// calculated route won't include non-`via` waypoints on roads that are
	// unsuitable for pickup and drop-off. This option works only for `DRIVE` and
	// `TWO_WHEELER` travel modes, and when the `location_type` is `location`.
	VehicleStopover bool `protobuf:"varint,4,opt,name=vehicle_stopover,json=vehicleStopover,proto3" json:"vehicle_stopover,omitempty"`
	// Indicates that the location of this waypoint is meant to have a preference
	// for the vehicle to stop at a particular side of road. When you set this
	// value, the route will pass through the location so that the vehicle can
	// stop at the side of road that the location is biased towards from the
	// center of the road. This option works only for 'DRIVE' and 'TWO_WHEELER'
	// travel modes, and when the 'location_type' is set to 'location'.
	SideOfRoad bool `protobuf:"varint,5,opt,name=side_of_road,json=sideOfRoad,proto3" json:"side_of_road,omitempty"`
}

func (x *Waypoint) Reset() {
	*x = Waypoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routes_v1_waypoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Waypoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Waypoint) ProtoMessage() {}

func (x *Waypoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routes_v1_waypoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Waypoint.ProtoReflect.Descriptor instead.
func (*Waypoint) Descriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_waypoint_proto_rawDescGZIP(), []int{0}
}

func (m *Waypoint) GetLocationType() isWaypoint_LocationType {
	if m != nil {
		return m.LocationType
	}
	return nil
}

func (x *Waypoint) GetLocation() *Location {
	if x, ok := x.GetLocationType().(*Waypoint_Location); ok {
		return x.Location
	}
	return nil
}

func (x *Waypoint) GetPlaceId() string {
	if x, ok := x.GetLocationType().(*Waypoint_PlaceId); ok {
		return x.PlaceId
	}
	return ""
}

func (x *Waypoint) GetVia() bool {
	if x != nil {
		return x.Via
	}
	return false
}

func (x *Waypoint) GetVehicleStopover() bool {
	if x != nil {
		return x.VehicleStopover
	}
	return false
}

func (x *Waypoint) GetSideOfRoad() bool {
	if x != nil {
		return x.SideOfRoad
	}
	return false
}

type isWaypoint_LocationType interface {
	isWaypoint_LocationType()
}

type Waypoint_Location struct {
	// A point specified using geographic coordinates, including an optional
	// heading.
	Location *Location `protobuf:"bytes,1,opt,name=location,proto3,oneof"`
}

type Waypoint_PlaceId struct {
	// The POI Place ID associated with the waypoint.
	PlaceId string `protobuf:"bytes,2,opt,name=place_id,json=placeId,proto3,oneof"`
}

func (*Waypoint_Location) isWaypoint_LocationType() {}

func (*Waypoint_PlaceId) isWaypoint_LocationType() {}

// Encapsulates a location (a geographic point, and an optional heading).
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The waypoint's geographic coordinates.
	LatLng *latlng.LatLng `protobuf:"bytes,1,opt,name=lat_lng,json=latLng,proto3" json:"lat_lng,omitempty"`
	// The compass heading associated with the direction of the flow of traffic.
	// This value is used to specify the side of the road to use for pickup and
	// drop-off. Heading values can be from 0 to 360, where 0 specifies a heading
	// of due North, 90 specifies a heading of due East, etc. You can use this
	// field only for `DRIVE` and `TWO_WHEELER` travel modes.
	Heading *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=heading,proto3" json:"heading,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routes_v1_waypoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routes_v1_waypoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_waypoint_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetLatLng() *latlng.LatLng {
	if x != nil {
		return x.LatLng
	}
	return nil
}

func (x *Location) GetHeading() *wrapperspb.Int32Value {
	if x != nil {
		return x.Heading
	}
	return nil
}

var File_google_maps_routes_v1_waypoint_proto protoreflect.FileDescriptor

var file_google_maps_routes_v1_waypoint_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61, 0x74, 0x6c, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x08, 0x57, 0x61, 0x79, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x76,
	0x69, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x74,
	0x6f, 0x70, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x0c, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x72, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x69, 0x64, 0x65, 0x4f, 0x66, 0x52, 0x6f, 0x61, 0x64, 0x42,
	0x0f, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x6f, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x07,
	0x6c, 0x61, 0x74, 0x5f, 0x6c, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c,
	0x6e, 0x67, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x12, 0x35, 0x0a, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x42, 0xa3, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42,
	0x0d, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x04, 0x47, 0x4d, 0x52, 0x53, 0xaa, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routes_v1_waypoint_proto_rawDescOnce sync.Once
	file_google_maps_routes_v1_waypoint_proto_rawDescData = file_google_maps_routes_v1_waypoint_proto_rawDesc
)

func file_google_maps_routes_v1_waypoint_proto_rawDescGZIP() []byte {
	file_google_maps_routes_v1_waypoint_proto_rawDescOnce.Do(func() {
		file_google_maps_routes_v1_waypoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routes_v1_waypoint_proto_rawDescData)
	})
	return file_google_maps_routes_v1_waypoint_proto_rawDescData
}

var file_google_maps_routes_v1_waypoint_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_maps_routes_v1_waypoint_proto_goTypes = []interface{}{
	(*Waypoint)(nil),              // 0: google.maps.routes.v1.Waypoint
	(*Location)(nil),              // 1: google.maps.routes.v1.Location
	(*latlng.LatLng)(nil),         // 2: google.type.LatLng
	(*wrapperspb.Int32Value)(nil), // 3: google.protobuf.Int32Value
}
var file_google_maps_routes_v1_waypoint_proto_depIdxs = []int32{
	1, // 0: google.maps.routes.v1.Waypoint.location:type_name -> google.maps.routes.v1.Location
	2, // 1: google.maps.routes.v1.Location.lat_lng:type_name -> google.type.LatLng
	3, // 2: google.maps.routes.v1.Location.heading:type_name -> google.protobuf.Int32Value
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_maps_routes_v1_waypoint_proto_init() }
func file_google_maps_routes_v1_waypoint_proto_init() {
	if File_google_maps_routes_v1_waypoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_maps_routes_v1_waypoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Waypoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_maps_routes_v1_waypoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_maps_routes_v1_waypoint_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Waypoint_Location)(nil),
		(*Waypoint_PlaceId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routes_v1_waypoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routes_v1_waypoint_proto_goTypes,
		DependencyIndexes: file_google_maps_routes_v1_waypoint_proto_depIdxs,
		MessageInfos:      file_google_maps_routes_v1_waypoint_proto_msgTypes,
	}.Build()
	File_google_maps_routes_v1_waypoint_proto = out.File
	file_google_maps_routes_v1_waypoint_proto_rawDesc = nil
	file_google_maps_routes_v1_waypoint_proto_goTypes = nil
	file_google_maps_routes_v1_waypoint_proto_depIdxs = nil
}

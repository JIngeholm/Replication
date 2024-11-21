// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: proto/auction.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bidder int32 `protobuf:"varint,1,opt,name=bidder,proto3" json:"bidder,omitempty"`
	Amount int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{0}
}

func (x *BidRequest) GetBidder() int32 {
	if x != nil {
		return x.Bidder
	}
	return 0
}

func (x *BidRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BidResponse) Reset() {
	*x = BidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidResponse) ProtoMessage() {}

func (x *BidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidResponse.ProtoReflect.Descriptor instead.
func (*BidResponse) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{1}
}

func (x *BidResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResultRequest) Reset() {
	*x = ResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRequest) ProtoMessage() {}

func (x *ResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRequest.ProtoReflect.Descriptor instead.
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{2}
}

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{3}
}

func (x *ResultResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type BidderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *BidderRequest) Reset() {
	*x = BidderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidderRequest) ProtoMessage() {}

func (x *BidderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidderRequest.ProtoReflect.Descriptor instead.
func (*BidderRequest) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{4}
}

func (x *BidderRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type BidderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BidderResponse) Reset() {
	*x = BidderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidderResponse) ProtoMessage() {}

func (x *BidderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidderResponse.ProtoReflect.Descriptor instead.
func (*BidderResponse) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{5}
}

func (x *BidderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type SyncStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighestBidder int32 `protobuf:"varint,1,opt,name=highest_bidder,json=highestBidder,proto3" json:"highest_bidder,omitempty"`
	HighestBid    int32 `protobuf:"varint,2,opt,name=highest_bid,json=highestBid,proto3" json:"highest_bid,omitempty"`
}

func (x *SyncStateRequest) Reset() {
	*x = SyncStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncStateRequest) ProtoMessage() {}

func (x *SyncStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncStateRequest.ProtoReflect.Descriptor instead.
func (*SyncStateRequest) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{6}
}

func (x *SyncStateRequest) GetHighestBidder() int32 {
	if x != nil {
		return x.HighestBidder
	}
	return 0
}

func (x *SyncStateRequest) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

type SyncStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SyncStateResponse) Reset() {
	*x = SyncStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncStateResponse) ProtoMessage() {}

func (x *SyncStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncStateResponse.ProtoReflect.Descriptor instead.
func (*SyncStateResponse) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{7}
}

func (x *SyncStateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_auction_proto protoreflect.FileDescriptor

var file_proto_auction_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3c,
	0x0a, 0x0a, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x69,
	0x64, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x0b,
	0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x29,
	0x0a, 0x0d, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x28, 0x0a, 0x0e, 0x42, 0x69, 0x64,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x5a, 0x0a, 0x10, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x69, 0x67, 0x68, 0x65,
	0x73, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x22,
	0x2b, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x84, 0x02, 0x0a,
	0x0e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x12, 0x16,
	0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x42, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_auction_proto_rawDescOnce sync.Once
	file_proto_auction_proto_rawDescData = file_proto_auction_proto_rawDesc
)

func file_proto_auction_proto_rawDescGZIP() []byte {
	file_proto_auction_proto_rawDescOnce.Do(func() {
		file_proto_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auction_proto_rawDescData)
	})
	return file_proto_auction_proto_rawDescData
}

var file_proto_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_auction_proto_goTypes = []any{
	(*BidRequest)(nil),        // 0: auction.BidRequest
	(*BidResponse)(nil),       // 1: auction.BidResponse
	(*ResultRequest)(nil),     // 2: auction.ResultRequest
	(*ResultResponse)(nil),    // 3: auction.ResultResponse
	(*BidderRequest)(nil),     // 4: auction.BidderRequest
	(*BidderResponse)(nil),    // 5: auction.BidderResponse
	(*SyncStateRequest)(nil),  // 6: auction.SyncStateRequest
	(*SyncStateResponse)(nil), // 7: auction.SyncStateResponse
}
var file_proto_auction_proto_depIdxs = []int32{
	0, // 0: auction.AuctionService.Bid:input_type -> auction.BidRequest
	2, // 1: auction.AuctionService.Result:input_type -> auction.ResultRequest
	4, // 2: auction.AuctionService.RegisterBidder:input_type -> auction.BidderRequest
	6, // 3: auction.AuctionService.SyncState:input_type -> auction.SyncStateRequest
	1, // 4: auction.AuctionService.Bid:output_type -> auction.BidResponse
	3, // 5: auction.AuctionService.Result:output_type -> auction.ResultResponse
	5, // 6: auction.AuctionService.RegisterBidder:output_type -> auction.BidderResponse
	7, // 7: auction.AuctionService.SyncState:output_type -> auction.SyncStateResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_auction_proto_init() }
func file_proto_auction_proto_init() {
	if File_proto_auction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_auction_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BidRequest); i {
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
		file_proto_auction_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BidResponse); i {
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
		file_proto_auction_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ResultRequest); i {
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
		file_proto_auction_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ResultResponse); i {
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
		file_proto_auction_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*BidderRequest); i {
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
		file_proto_auction_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*BidderResponse); i {
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
		file_proto_auction_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SyncStateRequest); i {
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
		file_proto_auction_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SyncStateResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_auction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auction_proto_goTypes,
		DependencyIndexes: file_proto_auction_proto_depIdxs,
		MessageInfos:      file_proto_auction_proto_msgTypes,
	}.Build()
	File_proto_auction_proto = out.File
	file_proto_auction_proto_rawDesc = nil
	file_proto_auction_proto_goTypes = nil
	file_proto_auction_proto_depIdxs = nil
}

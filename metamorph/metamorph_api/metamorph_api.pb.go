// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: metamorph/metamorph_api/metamorph_api.proto

package metamorph_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_UNKNOWN              Status = 0
	Status_QUEUED               Status = 1
	Status_RECEIVED             Status = 2
	Status_STORED               Status = 3
	Status_ANNOUNCED_TO_NETWORK Status = 4
	Status_SENT_TO_NETWORK      Status = 5
	Status_SEEN_ON_NETWORK      Status = 6
	Status_MINED                Status = 7
	Status_CONFIRMED            Status = 108 // 100 confirmation blocks
	Status_REJECTED             Status = 109
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:   "UNKNOWN",
		1:   "QUEUED",
		2:   "RECEIVED",
		3:   "STORED",
		4:   "ANNOUNCED_TO_NETWORK",
		5:   "SENT_TO_NETWORK",
		6:   "SEEN_ON_NETWORK",
		7:   "MINED",
		108: "CONFIRMED",
		109: "REJECTED",
	}
	Status_value = map[string]int32{
		"UNKNOWN":              0,
		"QUEUED":               1,
		"RECEIVED":             2,
		"STORED":               3,
		"ANNOUNCED_TO_NETWORK": 4,
		"SENT_TO_NETWORK":      5,
		"SEEN_ON_NETWORK":      6,
		"MINED":                7,
		"CONFIRMED":            108,
		"REJECTED":             109,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_metamorph_metamorph_api_metamorph_api_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_metamorph_metamorph_api_metamorph_api_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_metamorph_metamorph_api_metamorph_api_proto_rawDescGZIP(), []int{0}
}

// swagger:model HealthResponse
type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok                bool                   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Details           string                 `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Timestamp         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Workers           int32                  `protobuf:"varint,4,opt,name=workers,proto3" json:"workers,omitempty"`
	Uptime            float32                `protobuf:"fixed32,5,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Queued            int32                  `protobuf:"varint,6,opt,name=queued,proto3" json:"queued,omitempty"`
	Processed         int32                  `protobuf:"varint,7,opt,name=processed,proto3" json:"processed,omitempty"`
	Waiting           int32                  `protobuf:"varint,8,opt,name=waiting,proto3" json:"waiting,omitempty"`
	Average           float32                `protobuf:"fixed32,9,opt,name=average,proto3" json:"average,omitempty"`
	MapSize           int32                  `protobuf:"varint,10,opt,name=mapSize,proto3" json:"mapSize,omitempty"`
	PeersConnected    string                 `protobuf:"bytes,11,opt,name=PeersConnected,proto3" json:"PeersConnected,omitempty"`
	PeersDisconnected string                 `protobuf:"bytes,12,opt,name=PeersDisconnected,proto3" json:"PeersDisconnected,omitempty"`
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metamorph_metamorph_api_metamorph_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metamorph_metamorph_api_metamorph_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_metamorph_metamorph_api_metamorph_api_proto_rawDescGZIP(), []int{0}
}

func (x *HealthResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *HealthResponse) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *HealthResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *HealthResponse) GetWorkers() int32 {
	if x != nil {
		return x.Workers
	}
	return 0
}

func (x *HealthResponse) GetUptime() float32 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *HealthResponse) GetQueued() int32 {
	if x != nil {
		return x.Queued
	}
	return 0
}

func (x *HealthResponse) GetProcessed() int32 {
	if x != nil {
		return x.Processed
	}
	return 0
}

func (x *HealthResponse) GetWaiting() int32 {
	if x != nil {
		return x.Waiting
	}
	return 0
}

func (x *HealthResponse) GetAverage() float32 {
	if x != nil {
		return x.Average
	}
	return 0
}

func (x *HealthResponse) GetMapSize() int32 {
	if x != nil {
		return x.MapSize
	}
	return 0
}

func (x *HealthResponse) GetPeersConnected() string {
	if x != nil {
		return x.PeersConnected
	}
	return ""
}

func (x *HealthResponse) GetPeersDisconnected() string {
	if x != nil {
		return x.PeersDisconnected
	}
	return ""
}

// swagger:model TransactionRequest
type TransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKeyId      int64  `protobuf:"varint,1,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	StandardFeeId int64  `protobuf:"varint,2,opt,name=standard_fee_id,json=standardFeeId,proto3" json:"standard_fee_id,omitempty"`
	DataFeeId     int64  `protobuf:"varint,3,opt,name=data_fee_id,json=dataFeeId,proto3" json:"data_fee_id,omitempty"`
	SourceIp      string `protobuf:"bytes,4,opt,name=source_ip,json=sourceIp,proto3" json:"source_ip,omitempty"`
	CallbackUrl   string `protobuf:"bytes,5,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url,omitempty"`
	CallbackToken string `protobuf:"bytes,6,opt,name=callback_token,json=callbackToken,proto3" json:"callback_token,omitempty"`
	MerkleProof   bool   `protobuf:"varint,7,opt,name=merkle_proof,json=merkleProof,proto3" json:"merkle_proof,omitempty"`
	RawTx         []byte `protobuf:"bytes,8,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
	WaitForStatus Status `protobuf:"varint,9,opt,name=wait_for_status,json=waitForStatus,proto3,enum=metamorph_api.Status" json:"wait_for_status,omitempty"`
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metamorph_metamorph_api_metamorph_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metamorph_metamorph_api_metamorph_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_metamorph_metamorph_api_metamorph_api_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionRequest) GetApiKeyId() int64 {
	if x != nil {
		return x.ApiKeyId
	}
	return 0
}

func (x *TransactionRequest) GetStandardFeeId() int64 {
	if x != nil {
		return x.StandardFeeId
	}
	return 0
}

func (x *TransactionRequest) GetDataFeeId() int64 {
	if x != nil {
		return x.DataFeeId
	}
	return 0
}

func (x *TransactionRequest) GetSourceIp() string {
	if x != nil {
		return x.SourceIp
	}
	return ""
}

func (x *TransactionRequest) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

func (x *TransactionRequest) GetCallbackToken() string {
	if x != nil {
		return x.CallbackToken
	}
	return ""
}

func (x *TransactionRequest) GetMerkleProof() bool {
	if x != nil {
		return x.MerkleProof
	}
	return false
}

func (x *TransactionRequest) GetRawTx() []byte {
	if x != nil {
		return x.RawTx
	}
	return nil
}

func (x *TransactionRequest) GetWaitForStatus() Status {
	if x != nil {
		return x.WaitForStatus
	}
	return Status_UNKNOWN
}

// swagger:model TransactionStatus
type TransactionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimedOut     bool                   `protobuf:"varint,1,opt,name=timed_out,json=timedOut,proto3" json:"timed_out,omitempty"`
	StoredAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=stored_at,json=storedAt,proto3" json:"stored_at,omitempty"`
	AnnouncedAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=announced_at,json=announcedAt,proto3" json:"announced_at,omitempty"`
	MinedAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=mined_at,json=minedAt,proto3" json:"mined_at,omitempty"`
	Txid         string                 `protobuf:"bytes,5,opt,name=txid,proto3" json:"txid,omitempty"`
	Status       Status                 `protobuf:"varint,6,opt,name=status,proto3,enum=metamorph_api.Status" json:"status,omitempty"`
	RejectReason string                 `protobuf:"bytes,7,opt,name=reject_reason,json=rejectReason,proto3" json:"reject_reason,omitempty"`
	BlockHeight  int32                  `protobuf:"varint,8,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	BlockHash    string                 `protobuf:"bytes,9,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (x *TransactionStatus) Reset() {
	*x = TransactionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metamorph_metamorph_api_metamorph_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionStatus) ProtoMessage() {}

func (x *TransactionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_metamorph_metamorph_api_metamorph_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionStatus.ProtoReflect.Descriptor instead.
func (*TransactionStatus) Descriptor() ([]byte, []int) {
	return file_metamorph_metamorph_api_metamorph_api_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionStatus) GetTimedOut() bool {
	if x != nil {
		return x.TimedOut
	}
	return false
}

func (x *TransactionStatus) GetStoredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StoredAt
	}
	return nil
}

func (x *TransactionStatus) GetAnnouncedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AnnouncedAt
	}
	return nil
}

func (x *TransactionStatus) GetMinedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.MinedAt
	}
	return nil
}

func (x *TransactionStatus) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *TransactionStatus) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

func (x *TransactionStatus) GetRejectReason() string {
	if x != nil {
		return x.RejectReason
	}
	return ""
}

func (x *TransactionStatus) GetBlockHeight() int32 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *TransactionStatus) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

// swagger:model TransactionRequest
type TransactionStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
}

func (x *TransactionStatusRequest) Reset() {
	*x = TransactionStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metamorph_metamorph_api_metamorph_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionStatusRequest) ProtoMessage() {}

func (x *TransactionStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metamorph_metamorph_api_metamorph_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionStatusRequest.ProtoReflect.Descriptor instead.
func (*TransactionStatusRequest) Descriptor() ([]byte, []int) {
	return file_metamorph_metamorph_api_metamorph_api_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionStatusRequest) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

var File_metamorph_metamorph_api_metamorph_api_proto protoreflect.FileDescriptor

var file_metamorph_metamorph_api_metamorph_api_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x6d, 0x6f, 0x72, 0x70, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x6f,
	0x72, 0x70, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d,
	0x65, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x03, 0x0a, 0x0e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x75, 0x70, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x61, 0x69,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x77, 0x61, 0x69, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x61, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6d, 0x61, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12,
	0x2c, 0x0a, 0x11, 0x50, 0x65, 0x65, 0x72, 0x73, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x50, 0x65, 0x65, 0x72,
	0x73, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0xda, 0x02,
	0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x66,
	0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x46, 0x65, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x46, 0x65, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x61, 0x77, 0x5f, 0x74, 0x78, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x61, 0x77, 0x54, 0x78, 0x12, 0x3d, 0x0a, 0x0f, 0x77,
	0x61, 0x69, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x70, 0x68,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x77, 0x61, 0x69,
	0x74, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x89, 0x03, 0x0a, 0x11, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x37, 0x0a,
	0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x78, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64,
	0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x22, 0x2e, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x2a, 0xa7, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x4f, 0x52,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x4e, 0x4e, 0x4f, 0x55, 0x4e, 0x43, 0x45,
	0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x04, 0x12, 0x13,
	0x0a, 0x0f, 0x53, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52,
	0x4b, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x45, 0x4e, 0x5f, 0x4f, 0x4e, 0x5f, 0x4e,
	0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44,
	0x10, 0x6c, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x6d,
	0x32, 0x8f, 0x02, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x4d, 0x6f, 0x72, 0x70, 0x68, 0x41, 0x50,
	0x49, 0x12, 0x41, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x72,
	0x70, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x6d, 0x6f, 0x72, 0x70, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x63, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x70,
	0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x70,
	0x68, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metamorph_metamorph_api_metamorph_api_proto_rawDescOnce sync.Once
	file_metamorph_metamorph_api_metamorph_api_proto_rawDescData = file_metamorph_metamorph_api_metamorph_api_proto_rawDesc
)

func file_metamorph_metamorph_api_metamorph_api_proto_rawDescGZIP() []byte {
	file_metamorph_metamorph_api_metamorph_api_proto_rawDescOnce.Do(func() {
		file_metamorph_metamorph_api_metamorph_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_metamorph_metamorph_api_metamorph_api_proto_rawDescData)
	})
	return file_metamorph_metamorph_api_metamorph_api_proto_rawDescData
}

var file_metamorph_metamorph_api_metamorph_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_metamorph_metamorph_api_metamorph_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_metamorph_metamorph_api_metamorph_api_proto_goTypes = []interface{}{
	(Status)(0),                      // 0: metamorph_api.Status
	(*HealthResponse)(nil),           // 1: metamorph_api.HealthResponse
	(*TransactionRequest)(nil),       // 2: metamorph_api.TransactionRequest
	(*TransactionStatus)(nil),        // 3: metamorph_api.TransactionStatus
	(*TransactionStatusRequest)(nil), // 4: metamorph_api.TransactionStatusRequest
	(*timestamppb.Timestamp)(nil),    // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),            // 6: google.protobuf.Empty
}
var file_metamorph_metamorph_api_metamorph_api_proto_depIdxs = []int32{
	5, // 0: metamorph_api.HealthResponse.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: metamorph_api.TransactionRequest.wait_for_status:type_name -> metamorph_api.Status
	5, // 2: metamorph_api.TransactionStatus.stored_at:type_name -> google.protobuf.Timestamp
	5, // 3: metamorph_api.TransactionStatus.announced_at:type_name -> google.protobuf.Timestamp
	5, // 4: metamorph_api.TransactionStatus.mined_at:type_name -> google.protobuf.Timestamp
	0, // 5: metamorph_api.TransactionStatus.status:type_name -> metamorph_api.Status
	6, // 6: metamorph_api.MetaMorphAPI.Health:input_type -> google.protobuf.Empty
	2, // 7: metamorph_api.MetaMorphAPI.PutTransaction:input_type -> metamorph_api.TransactionRequest
	4, // 8: metamorph_api.MetaMorphAPI.GetTransactionStatus:input_type -> metamorph_api.TransactionStatusRequest
	1, // 9: metamorph_api.MetaMorphAPI.Health:output_type -> metamorph_api.HealthResponse
	3, // 10: metamorph_api.MetaMorphAPI.PutTransaction:output_type -> metamorph_api.TransactionStatus
	3, // 11: metamorph_api.MetaMorphAPI.GetTransactionStatus:output_type -> metamorph_api.TransactionStatus
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_metamorph_metamorph_api_metamorph_api_proto_init() }
func file_metamorph_metamorph_api_metamorph_api_proto_init() {
	if File_metamorph_metamorph_api_metamorph_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metamorph_metamorph_api_metamorph_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
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
		file_metamorph_metamorph_api_metamorph_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRequest); i {
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
		file_metamorph_metamorph_api_metamorph_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionStatus); i {
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
		file_metamorph_metamorph_api_metamorph_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionStatusRequest); i {
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
			RawDescriptor: file_metamorph_metamorph_api_metamorph_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metamorph_metamorph_api_metamorph_api_proto_goTypes,
		DependencyIndexes: file_metamorph_metamorph_api_metamorph_api_proto_depIdxs,
		EnumInfos:         file_metamorph_metamorph_api_metamorph_api_proto_enumTypes,
		MessageInfos:      file_metamorph_metamorph_api_metamorph_api_proto_msgTypes,
	}.Build()
	File_metamorph_metamorph_api_metamorph_api_proto = out.File
	file_metamorph_metamorph_api_metamorph_api_proto_rawDesc = nil
	file_metamorph_metamorph_api_metamorph_api_proto_goTypes = nil
	file_metamorph_metamorph_api_metamorph_api_proto_depIdxs = nil
}

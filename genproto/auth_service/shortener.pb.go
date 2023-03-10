// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shortener.proto

package auth_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateShortUrlRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LongUrl              string   `protobuf:"bytes,2,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
	ExpireDate           string   `protobuf:"bytes,3,opt,name=expire_date,json=expireDate,proto3" json:"expire_date,omitempty"`
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	ShortUrl             string   `protobuf:"bytes,5,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	LimitClick           string   `protobuf:"bytes,6,opt,name=limit_click,json=limitClick,proto3" json:"limit_click,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateShortUrlRequest) Reset()         { *m = CreateShortUrlRequest{} }
func (m *CreateShortUrlRequest) String() string { return proto.CompactTextString(m) }
func (*CreateShortUrlRequest) ProtoMessage()    {}
func (*CreateShortUrlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a64040fb43d257f, []int{0}
}

func (m *CreateShortUrlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShortUrlRequest.Unmarshal(m, b)
}
func (m *CreateShortUrlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShortUrlRequest.Marshal(b, m, deterministic)
}
func (m *CreateShortUrlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShortUrlRequest.Merge(m, src)
}
func (m *CreateShortUrlRequest) XXX_Size() int {
	return xxx_messageInfo_CreateShortUrlRequest.Size(m)
}
func (m *CreateShortUrlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShortUrlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShortUrlRequest proto.InternalMessageInfo

func (m *CreateShortUrlRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateShortUrlRequest) GetLongUrl() string {
	if m != nil {
		return m.LongUrl
	}
	return ""
}

func (m *CreateShortUrlRequest) GetExpireDate() string {
	if m != nil {
		return m.ExpireDate
	}
	return ""
}

func (m *CreateShortUrlRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateShortUrlRequest) GetShortUrl() string {
	if m != nil {
		return m.ShortUrl
	}
	return ""
}

func (m *CreateShortUrlRequest) GetLimitClick() string {
	if m != nil {
		return m.LimitClick
	}
	return ""
}

type CreateShortUrlResponse struct {
	ShortUrl             string   `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ExpireDate           string   `protobuf:"bytes,3,opt,name=expire_date,json=expireDate,proto3" json:"expire_date,omitempty"`
	LongUrl              string   `protobuf:"bytes,5,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
	CraetedAt            string   `protobuf:"bytes,4,opt,name=craeted_at,json=craetedAt,proto3" json:"craeted_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Id                   string   `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateShortUrlResponse) Reset()         { *m = CreateShortUrlResponse{} }
func (m *CreateShortUrlResponse) String() string { return proto.CompactTextString(m) }
func (*CreateShortUrlResponse) ProtoMessage()    {}
func (*CreateShortUrlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a64040fb43d257f, []int{1}
}

func (m *CreateShortUrlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShortUrlResponse.Unmarshal(m, b)
}
func (m *CreateShortUrlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShortUrlResponse.Marshal(b, m, deterministic)
}
func (m *CreateShortUrlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShortUrlResponse.Merge(m, src)
}
func (m *CreateShortUrlResponse) XXX_Size() int {
	return xxx_messageInfo_CreateShortUrlResponse.Size(m)
}
func (m *CreateShortUrlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShortUrlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShortUrlResponse proto.InternalMessageInfo

func (m *CreateShortUrlResponse) GetShortUrl() string {
	if m != nil {
		return m.ShortUrl
	}
	return ""
}

func (m *CreateShortUrlResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateShortUrlResponse) GetExpireDate() string {
	if m != nil {
		return m.ExpireDate
	}
	return ""
}

func (m *CreateShortUrlResponse) GetLongUrl() string {
	if m != nil {
		return m.LongUrl
	}
	return ""
}

func (m *CreateShortUrlResponse) GetCraetedAt() string {
	if m != nil {
		return m.CraetedAt
	}
	return ""
}

func (m *CreateShortUrlResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *CreateShortUrlResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetShortUrlRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ShortUrl             string   `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShortUrlRequest) Reset()         { *m = GetShortUrlRequest{} }
func (m *GetShortUrlRequest) String() string { return proto.CompactTextString(m) }
func (*GetShortUrlRequest) ProtoMessage()    {}
func (*GetShortUrlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a64040fb43d257f, []int{2}
}

func (m *GetShortUrlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShortUrlRequest.Unmarshal(m, b)
}
func (m *GetShortUrlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShortUrlRequest.Marshal(b, m, deterministic)
}
func (m *GetShortUrlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShortUrlRequest.Merge(m, src)
}
func (m *GetShortUrlRequest) XXX_Size() int {
	return xxx_messageInfo_GetShortUrlRequest.Size(m)
}
func (m *GetShortUrlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShortUrlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShortUrlRequest proto.InternalMessageInfo

func (m *GetShortUrlRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetShortUrlRequest) GetShortUrl() string {
	if m != nil {
		return m.ShortUrl
	}
	return ""
}

type GetShortUrlResponse struct {
	ShortUrl             string   `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ExpireDate           string   `protobuf:"bytes,3,opt,name=expire_date,json=expireDate,proto3" json:"expire_date,omitempty"`
	LongUrl              string   `protobuf:"bytes,4,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ClickCount           string   `protobuf:"bytes,7,opt,name=click_count,json=clickCount,proto3" json:"click_count,omitempty"`
	LimitClick           string   `protobuf:"bytes,8,opt,name=limit_click,json=limitClick,proto3" json:"limit_click,omitempty"`
	UrlStatus            string   `protobuf:"bytes,9,opt,name=url_status,json=urlStatus,proto3" json:"url_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShortUrlResponse) Reset()         { *m = GetShortUrlResponse{} }
func (m *GetShortUrlResponse) String() string { return proto.CompactTextString(m) }
func (*GetShortUrlResponse) ProtoMessage()    {}
func (*GetShortUrlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a64040fb43d257f, []int{3}
}

func (m *GetShortUrlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShortUrlResponse.Unmarshal(m, b)
}
func (m *GetShortUrlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShortUrlResponse.Marshal(b, m, deterministic)
}
func (m *GetShortUrlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShortUrlResponse.Merge(m, src)
}
func (m *GetShortUrlResponse) XXX_Size() int {
	return xxx_messageInfo_GetShortUrlResponse.Size(m)
}
func (m *GetShortUrlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShortUrlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShortUrlResponse proto.InternalMessageInfo

func (m *GetShortUrlResponse) GetShortUrl() string {
	if m != nil {
		return m.ShortUrl
	}
	return ""
}

func (m *GetShortUrlResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetShortUrlResponse) GetExpireDate() string {
	if m != nil {
		return m.ExpireDate
	}
	return ""
}

func (m *GetShortUrlResponse) GetLongUrl() string {
	if m != nil {
		return m.LongUrl
	}
	return ""
}

func (m *GetShortUrlResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetShortUrlResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *GetShortUrlResponse) GetClickCount() string {
	if m != nil {
		return m.ClickCount
	}
	return ""
}

func (m *GetShortUrlResponse) GetLimitClick() string {
	if m != nil {
		return m.LimitClick
	}
	return ""
}

func (m *GetShortUrlResponse) GetUrlStatus() string {
	if m != nil {
		return m.UrlStatus
	}
	return ""
}

type IncClickCountRequest struct {
	ShortUrl             string   `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncClickCountRequest) Reset()         { *m = IncClickCountRequest{} }
func (m *IncClickCountRequest) String() string { return proto.CompactTextString(m) }
func (*IncClickCountRequest) ProtoMessage()    {}
func (*IncClickCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a64040fb43d257f, []int{4}
}

func (m *IncClickCountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncClickCountRequest.Unmarshal(m, b)
}
func (m *IncClickCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncClickCountRequest.Marshal(b, m, deterministic)
}
func (m *IncClickCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncClickCountRequest.Merge(m, src)
}
func (m *IncClickCountRequest) XXX_Size() int {
	return xxx_messageInfo_IncClickCountRequest.Size(m)
}
func (m *IncClickCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncClickCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncClickCountRequest proto.InternalMessageInfo

func (m *IncClickCountRequest) GetShortUrl() string {
	if m != nil {
		return m.ShortUrl
	}
	return ""
}

type IncClickCountResponse struct {
	ShortUrl             string   `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	ClickCount           int64    `protobuf:"varint,2,opt,name=click_count,json=clickCount,proto3" json:"click_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncClickCountResponse) Reset()         { *m = IncClickCountResponse{} }
func (m *IncClickCountResponse) String() string { return proto.CompactTextString(m) }
func (*IncClickCountResponse) ProtoMessage()    {}
func (*IncClickCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a64040fb43d257f, []int{5}
}

func (m *IncClickCountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncClickCountResponse.Unmarshal(m, b)
}
func (m *IncClickCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncClickCountResponse.Marshal(b, m, deterministic)
}
func (m *IncClickCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncClickCountResponse.Merge(m, src)
}
func (m *IncClickCountResponse) XXX_Size() int {
	return xxx_messageInfo_IncClickCountResponse.Size(m)
}
func (m *IncClickCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IncClickCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IncClickCountResponse proto.InternalMessageInfo

func (m *IncClickCountResponse) GetShortUrl() string {
	if m != nil {
		return m.ShortUrl
	}
	return ""
}

func (m *IncClickCountResponse) GetClickCount() int64 {
	if m != nil {
		return m.ClickCount
	}
	return 0
}

type HandleLongUrlRequest struct {
	ShortUrl             string   `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleLongUrlRequest) Reset()         { *m = HandleLongUrlRequest{} }
func (m *HandleLongUrlRequest) String() string { return proto.CompactTextString(m) }
func (*HandleLongUrlRequest) ProtoMessage()    {}
func (*HandleLongUrlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a64040fb43d257f, []int{6}
}

func (m *HandleLongUrlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleLongUrlRequest.Unmarshal(m, b)
}
func (m *HandleLongUrlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleLongUrlRequest.Marshal(b, m, deterministic)
}
func (m *HandleLongUrlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleLongUrlRequest.Merge(m, src)
}
func (m *HandleLongUrlRequest) XXX_Size() int {
	return xxx_messageInfo_HandleLongUrlRequest.Size(m)
}
func (m *HandleLongUrlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleLongUrlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleLongUrlRequest proto.InternalMessageInfo

func (m *HandleLongUrlRequest) GetShortUrl() string {
	if m != nil {
		return m.ShortUrl
	}
	return ""
}

type HandleLongUrlResponse struct {
	LongUrl              string   `protobuf:"bytes,1,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleLongUrlResponse) Reset()         { *m = HandleLongUrlResponse{} }
func (m *HandleLongUrlResponse) String() string { return proto.CompactTextString(m) }
func (*HandleLongUrlResponse) ProtoMessage()    {}
func (*HandleLongUrlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a64040fb43d257f, []int{7}
}

func (m *HandleLongUrlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleLongUrlResponse.Unmarshal(m, b)
}
func (m *HandleLongUrlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleLongUrlResponse.Marshal(b, m, deterministic)
}
func (m *HandleLongUrlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleLongUrlResponse.Merge(m, src)
}
func (m *HandleLongUrlResponse) XXX_Size() int {
	return xxx_messageInfo_HandleLongUrlResponse.Size(m)
}
func (m *HandleLongUrlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleLongUrlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HandleLongUrlResponse proto.InternalMessageInfo

func (m *HandleLongUrlResponse) GetLongUrl() string {
	if m != nil {
		return m.LongUrl
	}
	return ""
}

type GetAllUserUrlsRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllUserUrlsRequest) Reset()         { *m = GetAllUserUrlsRequest{} }
func (m *GetAllUserUrlsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllUserUrlsRequest) ProtoMessage()    {}
func (*GetAllUserUrlsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a64040fb43d257f, []int{8}
}

func (m *GetAllUserUrlsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllUserUrlsRequest.Unmarshal(m, b)
}
func (m *GetAllUserUrlsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllUserUrlsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllUserUrlsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllUserUrlsRequest.Merge(m, src)
}
func (m *GetAllUserUrlsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllUserUrlsRequest.Size(m)
}
func (m *GetAllUserUrlsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllUserUrlsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllUserUrlsRequest proto.InternalMessageInfo

func (m *GetAllUserUrlsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type UrlData struct {
	ShortUrl             string   `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ExpireDate           string   `protobuf:"bytes,3,opt,name=expire_date,json=expireDate,proto3" json:"expire_date,omitempty"`
	LongUrl              string   `protobuf:"bytes,4,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Id                   string   `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	ClickCount           string   `protobuf:"bytes,8,opt,name=click_count,json=clickCount,proto3" json:"click_count,omitempty"`
	LimitClick           string   `protobuf:"bytes,9,opt,name=limit_click,json=limitClick,proto3" json:"limit_click,omitempty"`
	UrlStatus            string   `protobuf:"bytes,10,opt,name=url_status,json=urlStatus,proto3" json:"url_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UrlData) Reset()         { *m = UrlData{} }
func (m *UrlData) String() string { return proto.CompactTextString(m) }
func (*UrlData) ProtoMessage()    {}
func (*UrlData) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a64040fb43d257f, []int{9}
}

func (m *UrlData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UrlData.Unmarshal(m, b)
}
func (m *UrlData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UrlData.Marshal(b, m, deterministic)
}
func (m *UrlData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UrlData.Merge(m, src)
}
func (m *UrlData) XXX_Size() int {
	return xxx_messageInfo_UrlData.Size(m)
}
func (m *UrlData) XXX_DiscardUnknown() {
	xxx_messageInfo_UrlData.DiscardUnknown(m)
}

var xxx_messageInfo_UrlData proto.InternalMessageInfo

func (m *UrlData) GetShortUrl() string {
	if m != nil {
		return m.ShortUrl
	}
	return ""
}

func (m *UrlData) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UrlData) GetExpireDate() string {
	if m != nil {
		return m.ExpireDate
	}
	return ""
}

func (m *UrlData) GetLongUrl() string {
	if m != nil {
		return m.LongUrl
	}
	return ""
}

func (m *UrlData) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UrlData) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *UrlData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UrlData) GetClickCount() string {
	if m != nil {
		return m.ClickCount
	}
	return ""
}

func (m *UrlData) GetLimitClick() string {
	if m != nil {
		return m.LimitClick
	}
	return ""
}

func (m *UrlData) GetUrlStatus() string {
	if m != nil {
		return m.UrlStatus
	}
	return ""
}

type GetAllUserUrlsResponse struct {
	Urls                 []*UrlData `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	TotalCount           int64      `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAllUserUrlsResponse) Reset()         { *m = GetAllUserUrlsResponse{} }
func (m *GetAllUserUrlsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllUserUrlsResponse) ProtoMessage()    {}
func (*GetAllUserUrlsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a64040fb43d257f, []int{10}
}

func (m *GetAllUserUrlsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllUserUrlsResponse.Unmarshal(m, b)
}
func (m *GetAllUserUrlsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllUserUrlsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllUserUrlsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllUserUrlsResponse.Merge(m, src)
}
func (m *GetAllUserUrlsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllUserUrlsResponse.Size(m)
}
func (m *GetAllUserUrlsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllUserUrlsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllUserUrlsResponse proto.InternalMessageInfo

func (m *GetAllUserUrlsResponse) GetUrls() []*UrlData {
	if m != nil {
		return m.Urls
	}
	return nil
}

func (m *GetAllUserUrlsResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateShortUrlRequest)(nil), "auth_service.CreateShortUrlRequest")
	proto.RegisterType((*CreateShortUrlResponse)(nil), "auth_service.CreateShortUrlResponse")
	proto.RegisterType((*GetShortUrlRequest)(nil), "auth_service.GetShortUrlRequest")
	proto.RegisterType((*GetShortUrlResponse)(nil), "auth_service.GetShortUrlResponse")
	proto.RegisterType((*IncClickCountRequest)(nil), "auth_service.IncClickCountRequest")
	proto.RegisterType((*IncClickCountResponse)(nil), "auth_service.IncClickCountResponse")
	proto.RegisterType((*HandleLongUrlRequest)(nil), "auth_service.HandleLongUrlRequest")
	proto.RegisterType((*HandleLongUrlResponse)(nil), "auth_service.HandleLongUrlResponse")
	proto.RegisterType((*GetAllUserUrlsRequest)(nil), "auth_service.GetAllUserUrlsRequest")
	proto.RegisterType((*UrlData)(nil), "auth_service.UrlData")
	proto.RegisterType((*GetAllUserUrlsResponse)(nil), "auth_service.GetAllUserUrlsResponse")
}

func init() {
	proto.RegisterFile("shortener.proto", fileDescriptor_6a64040fb43d257f)
}

var fileDescriptor_6a64040fb43d257f = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xcb, 0x6e, 0xd4, 0x40,
	0x10, 0x94, 0x9d, 0x7d, 0xb9, 0x17, 0x81, 0x64, 0x70, 0x62, 0x84, 0x50, 0x22, 0x9f, 0xc2, 0x65,
	0x41, 0xc9, 0x17, 0x2c, 0x1b, 0x29, 0x04, 0x71, 0xda, 0xc8, 0x17, 0x2e, 0xd6, 0x60, 0xb7, 0x12,
	0x8b, 0xc1, 0x5e, 0x7a, 0xda, 0x88, 0x2f, 0xe2, 0xc0, 0x27, 0xf0, 0x25, 0x7c, 0x0e, 0x9a, 0xde,
	0x59, 0x64, 0x8f, 0x10, 0xbb, 0x27, 0x24, 0x8e, 0xae, 0xb6, 0x6b, 0xaa, 0xaa, 0xcb, 0x03, 0x8f,
	0xcc, 0x7d, 0x4b, 0x8c, 0x0d, 0xd2, 0x62, 0x43, 0x2d, 0xb7, 0xf1, 0x03, 0xd5, 0xf1, 0x7d, 0x61,
	0x90, 0xbe, 0xd4, 0x25, 0x66, 0x3f, 0x02, 0x48, 0x56, 0x84, 0x8a, 0xf1, 0xd6, 0xbe, 0x97, 0x93,
	0x5e, 0xe3, 0xe7, 0x0e, 0x0d, 0xc7, 0x27, 0x30, 0xed, 0x0c, 0x52, 0x51, 0x57, 0x69, 0x70, 0x16,
	0x9c, 0x47, 0xeb, 0x89, 0x7d, 0xbc, 0xa9, 0xe2, 0xa7, 0x30, 0xd3, 0x6d, 0x73, 0x57, 0x74, 0xa4,
	0xd3, 0x50, 0x26, 0x53, 0xfb, 0x9c, 0x93, 0x8e, 0x4f, 0x61, 0x8e, 0x5f, 0x37, 0x35, 0x61, 0x51,
	0x29, 0xc6, 0xf4, 0x48, 0xa6, 0xb0, 0x85, 0xae, 0x14, 0x63, 0xfc, 0x10, 0xc2, 0xba, 0x4a, 0x47,
	0x82, 0x87, 0x75, 0x15, 0x3f, 0x83, 0x48, 0xf4, 0x09, 0xd9, 0x58, 0xe0, 0x99, 0x71, 0x42, 0x2c,
	0x9b, 0xae, 0x3f, 0xd5, 0x5c, 0x94, 0xba, 0x2e, 0x3f, 0xa6, 0x93, 0x2d, 0x9b, 0x40, 0x2b, 0x8b,
	0x64, 0x3f, 0x03, 0x38, 0xf6, 0xc5, 0x9b, 0x4d, 0xdb, 0x18, 0x1c, 0x12, 0x07, 0x1e, 0x71, 0xcf,
	0x5a, 0x38, 0xb0, 0xb6, 0x57, 0x7f, 0xdf, 0xfb, 0x78, 0xe8, 0xfd, 0x39, 0x40, 0x49, 0x0a, 0x19,
	0xab, 0x42, 0xb1, 0xb3, 0x18, 0x39, 0x64, 0xc9, 0x76, 0xdc, 0x6d, 0x2c, 0xab, 0x8c, 0xb7, 0x5e,
	0x22, 0x87, 0x2c, 0xd9, 0x05, 0x33, 0xdd, 0x05, 0x93, 0xbd, 0x85, 0xf8, 0x1a, 0xf9, 0xe0, 0x9d,
	0x0c, 0xec, 0x86, 0x43, 0xbb, 0xd9, 0xb7, 0x10, 0x1e, 0x0f, 0xc8, 0xfe, 0x5d, 0x46, 0xa3, 0x3f,
	0x64, 0x84, 0xbb, 0x10, 0xc6, 0xbb, 0x8c, 0x50, 0x1d, 0x94, 0xd1, 0x29, 0xcc, 0xa5, 0x09, 0x45,
	0xd9, 0x76, 0x0d, 0xbb, 0xb0, 0x40, 0xa0, 0x95, 0x45, 0xfc, 0xc2, 0xcc, 0xfc, 0xc2, 0xc8, 0x01,
	0xa4, 0x0b, 0xc3, 0x8a, 0x3b, 0x93, 0x46, 0xee, 0x00, 0xd2, 0xb7, 0x02, 0x64, 0x97, 0xf0, 0xe4,
	0xa6, 0x29, 0x57, 0xbf, 0x09, 0x77, 0xb1, 0xff, 0x2d, 0xa8, 0x2c, 0x87, 0xc4, 0xfb, 0xe8, 0x90,
	0x78, 0x3d, 0x2f, 0x36, 0xe2, 0xa3, 0xbe, 0x17, 0xab, 0xe5, 0x8d, 0x6a, 0x2a, 0x8d, 0xef, 0xb6,
	0xd9, 0x1d, 0xa4, 0xe5, 0x02, 0x12, 0xef, 0x23, 0xa7, 0xa5, 0xbf, 0x93, 0x60, 0xb0, 0x93, 0xec,
	0x15, 0x24, 0xd7, 0xc8, 0x4b, 0xad, 0x73, 0x83, 0x94, 0x93, 0x36, 0xfb, 0xca, 0x96, 0x7d, 0x0f,
	0x61, 0x9a, 0x93, 0xbe, 0x52, 0xac, 0xfe, 0xbb, 0x0e, 0x79, 0xff, 0x99, 0xbf, 0x87, 0xd9, 0xbe,
	0x4e, 0x45, 0x7b, 0x3a, 0x05, 0x7e, 0xa7, 0x2a, 0x38, 0xf6, 0xe3, 0x75, 0x3b, 0x79, 0x01, 0xa3,
	0x8e, 0xb4, 0x49, 0x83, 0xb3, 0xa3, 0xf3, 0xf9, 0x45, 0xb2, 0xe8, 0xdf, 0xcb, 0x0b, 0x97, 0xef,
	0x5a, 0x5e, 0xb1, 0x22, 0xb8, 0x65, 0xa5, 0x87, 0x6d, 0x11, 0x48, 0x54, 0xbe, 0x3e, 0x79, 0x9f,
	0xdc, 0x61, 0x23, 0x17, 0xfc, 0xcb, 0x3e, 0xcf, 0x87, 0x89, 0x60, 0x97, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x1f, 0x50, 0x8c, 0x34, 0x07, 0x06, 0x00, 0x00,
}

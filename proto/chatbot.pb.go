// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatbot.proto

package chatbotpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ChatAppType - chat app type
type ChatAppType int32

const (
	ChatAppType_CAT_TELEGRAM ChatAppType = 0
	ChatAppType_CAT_COOLQQ   ChatAppType = 1
)

var ChatAppType_name = map[int32]string{
	0: "CAT_TELEGRAM",
	1: "CAT_COOLQQ",
}
var ChatAppType_value = map[string]int32{
	"CAT_TELEGRAM": 0,
	"CAT_COOLQQ":   1,
}

func (x ChatAppType) String() string {
	return proto.EnumName(ChatAppType_name, int32(x))
}
func (ChatAppType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_ab43711db21b8f66, []int{0}
}

// UserAppInfo - user app info
type UserAppInfo struct {
	App                  ChatAppType `protobuf:"varint,1,opt,name=app,proto3,enum=chatbotpb.ChatAppType" json:"app,omitempty"`
	Appuid               string      `protobuf:"bytes,2,opt,name=appuid,proto3" json:"appuid,omitempty"`
	Appuname             string      `protobuf:"bytes,3,opt,name=appuname,proto3" json:"appuname,omitempty"`
	Chatnums             int32       `protobuf:"varint,4,opt,name=chatnums,proto3" json:"chatnums,omitempty"`
	UsernameAppServ      string      `protobuf:"bytes,5,opt,name=usernameAppServ,proto3" json:"usernameAppServ,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserAppInfo) Reset()         { *m = UserAppInfo{} }
func (m *UserAppInfo) String() string { return proto.CompactTextString(m) }
func (*UserAppInfo) ProtoMessage()    {}
func (*UserAppInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_ab43711db21b8f66, []int{0}
}
func (m *UserAppInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAppInfo.Unmarshal(m, b)
}
func (m *UserAppInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAppInfo.Marshal(b, m, deterministic)
}
func (dst *UserAppInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAppInfo.Merge(dst, src)
}
func (m *UserAppInfo) XXX_Size() int {
	return xxx_messageInfo_UserAppInfo.Size(m)
}
func (m *UserAppInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAppInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserAppInfo proto.InternalMessageInfo

func (m *UserAppInfo) GetApp() ChatAppType {
	if m != nil {
		return m.App
	}
	return ChatAppType_CAT_TELEGRAM
}

func (m *UserAppInfo) GetAppuid() string {
	if m != nil {
		return m.Appuid
	}
	return ""
}

func (m *UserAppInfo) GetAppuname() string {
	if m != nil {
		return m.Appuname
	}
	return ""
}

func (m *UserAppInfo) GetChatnums() int32 {
	if m != nil {
		return m.Chatnums
	}
	return 0
}

func (m *UserAppInfo) GetUsernameAppServ() string {
	if m != nil {
		return m.UsernameAppServ
	}
	return ""
}

// UserInfo - user info
type UserInfo struct {
	Uid                  int64          `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Apps                 []*UserAppInfo `protobuf:"bytes,3,rep,name=apps,proto3" json:"apps,omitempty"`
	Tags                 []string       `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	Money                int64          `protobuf:"varint,10,opt,name=money,proto3" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_ab43711db21b8f66, []int{1}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (dst *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(dst, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetApps() []*UserAppInfo {
	if m != nil {
		return m.Apps
	}
	return nil
}

func (m *UserInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *UserInfo) GetMoney() int64 {
	if m != nil {
		return m.Money
	}
	return 0
}

// AppGroupInfo - app group info
type AppGroupInfo struct {
	App                  ChatAppType `protobuf:"varint,1,opt,name=app,proto3,enum=chatbotpb.ChatAppType" json:"app,omitempty"`
	UsernameAppServ      string      `protobuf:"bytes,2,opt,name=usernameAppServ,proto3" json:"usernameAppServ,omitempty"`
	Groupid              string      `protobuf:"bytes,3,opt,name=groupid,proto3" json:"groupid,omitempty"`
	Groupname            string      `protobuf:"bytes,4,opt,name=groupname,proto3" json:"groupname,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AppGroupInfo) Reset()         { *m = AppGroupInfo{} }
func (m *AppGroupInfo) String() string { return proto.CompactTextString(m) }
func (*AppGroupInfo) ProtoMessage()    {}
func (*AppGroupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_ab43711db21b8f66, []int{2}
}
func (m *AppGroupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppGroupInfo.Unmarshal(m, b)
}
func (m *AppGroupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppGroupInfo.Marshal(b, m, deterministic)
}
func (dst *AppGroupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppGroupInfo.Merge(dst, src)
}
func (m *AppGroupInfo) XXX_Size() int {
	return xxx_messageInfo_AppGroupInfo.Size(m)
}
func (m *AppGroupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AppGroupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AppGroupInfo proto.InternalMessageInfo

func (m *AppGroupInfo) GetApp() ChatAppType {
	if m != nil {
		return m.App
	}
	return ChatAppType_CAT_TELEGRAM
}

func (m *AppGroupInfo) GetUsernameAppServ() string {
	if m != nil {
		return m.UsernameAppServ
	}
	return ""
}

func (m *AppGroupInfo) GetGroupid() string {
	if m != nil {
		return m.Groupid
	}
	return ""
}

func (m *AppGroupInfo) GetGroupname() string {
	if m != nil {
		return m.Groupname
	}
	return ""
}

// AppServInfo - app server info
type AppServInfo struct {
	Token                string      `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	AppType              ChatAppType `protobuf:"varint,2,opt,name=appType,proto3,enum=chatbotpb.ChatAppType" json:"appType,omitempty"`
	Username             string      `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Sessionid            string      `protobuf:"bytes,4,opt,name=sessionid,proto3" json:"sessionid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AppServInfo) Reset()         { *m = AppServInfo{} }
func (m *AppServInfo) String() string { return proto.CompactTextString(m) }
func (*AppServInfo) ProtoMessage()    {}
func (*AppServInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_ab43711db21b8f66, []int{3}
}
func (m *AppServInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppServInfo.Unmarshal(m, b)
}
func (m *AppServInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppServInfo.Marshal(b, m, deterministic)
}
func (dst *AppServInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppServInfo.Merge(dst, src)
}
func (m *AppServInfo) XXX_Size() int {
	return xxx_messageInfo_AppServInfo.Size(m)
}
func (m *AppServInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AppServInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AppServInfo proto.InternalMessageInfo

func (m *AppServInfo) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AppServInfo) GetAppType() ChatAppType {
	if m != nil {
		return m.AppType
	}
	return ChatAppType_CAT_TELEGRAM
}

func (m *AppServInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AppServInfo) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

// RegisterAppService - register app service
type RegisterAppService struct {
	AppServ              *AppServInfo `protobuf:"bytes,1,opt,name=appServ,proto3" json:"appServ,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RegisterAppService) Reset()         { *m = RegisterAppService{} }
func (m *RegisterAppService) String() string { return proto.CompactTextString(m) }
func (*RegisterAppService) ProtoMessage()    {}
func (*RegisterAppService) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_ab43711db21b8f66, []int{4}
}
func (m *RegisterAppService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterAppService.Unmarshal(m, b)
}
func (m *RegisterAppService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterAppService.Marshal(b, m, deterministic)
}
func (dst *RegisterAppService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterAppService.Merge(dst, src)
}
func (m *RegisterAppService) XXX_Size() int {
	return xxx_messageInfo_RegisterAppService.Size(m)
}
func (m *RegisterAppService) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterAppService.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterAppService proto.InternalMessageInfo

func (m *RegisterAppService) GetAppServ() *AppServInfo {
	if m != nil {
		return m.AppServ
	}
	return nil
}

// ReplyRegisterAppService - reply RegisterAppService
type ReplyRegisterAppService struct {
	AppType              ChatAppType `protobuf:"varint,1,opt,name=appType,proto3,enum=chatbotpb.ChatAppType" json:"appType,omitempty"`
	Error                string      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	SessionID            string      `protobuf:"bytes,3,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReplyRegisterAppService) Reset()         { *m = ReplyRegisterAppService{} }
func (m *ReplyRegisterAppService) String() string { return proto.CompactTextString(m) }
func (*ReplyRegisterAppService) ProtoMessage()    {}
func (*ReplyRegisterAppService) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_ab43711db21b8f66, []int{5}
}
func (m *ReplyRegisterAppService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyRegisterAppService.Unmarshal(m, b)
}
func (m *ReplyRegisterAppService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyRegisterAppService.Marshal(b, m, deterministic)
}
func (dst *ReplyRegisterAppService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyRegisterAppService.Merge(dst, src)
}
func (m *ReplyRegisterAppService) XXX_Size() int {
	return xxx_messageInfo_ReplyRegisterAppService.Size(m)
}
func (m *ReplyRegisterAppService) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyRegisterAppService.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyRegisterAppService proto.InternalMessageInfo

func (m *ReplyRegisterAppService) GetAppType() ChatAppType {
	if m != nil {
		return m.AppType
	}
	return ChatAppType_CAT_TELEGRAM
}

func (m *ReplyRegisterAppService) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ReplyRegisterAppService) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

// ChatMsg - chat message
type ChatMsg struct {
	Msg                  string        `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Uai                  *UserAppInfo  `protobuf:"bytes,2,opt,name=uai,proto3" json:"uai,omitempty"`
	Agi                  *AppGroupInfo `protobuf:"bytes,3,opt,name=agi,proto3" json:"agi,omitempty"`
	Filename             string        `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
	FileData             []byte        `protobuf:"bytes,5,opt,name=fileData,proto3" json:"fileData,omitempty"`
	Error                string        `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	Token                string        `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	SessionID            string        `protobuf:"bytes,8,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ChatMsg) Reset()         { *m = ChatMsg{} }
func (m *ChatMsg) String() string { return proto.CompactTextString(m) }
func (*ChatMsg) ProtoMessage()    {}
func (*ChatMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_ab43711db21b8f66, []int{6}
}
func (m *ChatMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMsg.Unmarshal(m, b)
}
func (m *ChatMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMsg.Marshal(b, m, deterministic)
}
func (dst *ChatMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMsg.Merge(dst, src)
}
func (m *ChatMsg) XXX_Size() int {
	return xxx_messageInfo_ChatMsg.Size(m)
}
func (m *ChatMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMsg proto.InternalMessageInfo

func (m *ChatMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ChatMsg) GetUai() *UserAppInfo {
	if m != nil {
		return m.Uai
	}
	return nil
}

func (m *ChatMsg) GetAgi() *AppGroupInfo {
	if m != nil {
		return m.Agi
	}
	return nil
}

func (m *ChatMsg) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *ChatMsg) GetFileData() []byte {
	if m != nil {
		return m.FileData
	}
	return nil
}

func (m *ChatMsg) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ChatMsg) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ChatMsg) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

// ChatMsgStream - chat message stream
type ChatMsgStream struct {
	// totalLength - If the message is too long, it will send data in multiple msg, this is the total length.
	TotalLength int32 `protobuf:"varint,1,opt,name=totalLength,proto3" json:"totalLength,omitempty"`
	// curStart - The starting point of the current data (in bytes).
	CurStart int32 `protobuf:"varint,2,opt,name=curStart,proto3" json:"curStart,omitempty"`
	// curLength - The length of the current data (in bytes).
	CurLength int32 `protobuf:"varint,3,opt,name=curLength,proto3" json:"curLength,omitempty"`
	// hashData - This is the hash of each paragraph.
	HashData string `protobuf:"bytes,4,opt,name=hashData,proto3" json:"hashData,omitempty"`
	// totalHashData - If multiple messages return data, this is the hash value of all data, only sent in the last message.
	TotalHashData string `protobuf:"bytes,5,opt,name=totalHashData,proto3" json:"totalHashData,omitempty"`
	// data
	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// error
	Error string `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
	// token
	Token string `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	// sessionID
	SessionID            string   `protobuf:"bytes,9,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	Chat                 *ChatMsg `protobuf:"bytes,10,opt,name=chat,proto3" json:"chat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMsgStream) Reset()         { *m = ChatMsgStream{} }
func (m *ChatMsgStream) String() string { return proto.CompactTextString(m) }
func (*ChatMsgStream) ProtoMessage()    {}
func (*ChatMsgStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_ab43711db21b8f66, []int{7}
}
func (m *ChatMsgStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMsgStream.Unmarshal(m, b)
}
func (m *ChatMsgStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMsgStream.Marshal(b, m, deterministic)
}
func (dst *ChatMsgStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMsgStream.Merge(dst, src)
}
func (m *ChatMsgStream) XXX_Size() int {
	return xxx_messageInfo_ChatMsgStream.Size(m)
}
func (m *ChatMsgStream) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMsgStream.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMsgStream proto.InternalMessageInfo

func (m *ChatMsgStream) GetTotalLength() int32 {
	if m != nil {
		return m.TotalLength
	}
	return 0
}

func (m *ChatMsgStream) GetCurStart() int32 {
	if m != nil {
		return m.CurStart
	}
	return 0
}

func (m *ChatMsgStream) GetCurLength() int32 {
	if m != nil {
		return m.CurLength
	}
	return 0
}

func (m *ChatMsgStream) GetHashData() string {
	if m != nil {
		return m.HashData
	}
	return ""
}

func (m *ChatMsgStream) GetTotalHashData() string {
	if m != nil {
		return m.TotalHashData
	}
	return ""
}

func (m *ChatMsgStream) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ChatMsgStream) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ChatMsgStream) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ChatMsgStream) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *ChatMsgStream) GetChat() *ChatMsg {
	if m != nil {
		return m.Chat
	}
	return nil
}

// RequestChatData - request chat data
type RequestChatData struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	SessionID            string   `protobuf:"bytes,2,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestChatData) Reset()         { *m = RequestChatData{} }
func (m *RequestChatData) String() string { return proto.CompactTextString(m) }
func (*RequestChatData) ProtoMessage()    {}
func (*RequestChatData) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatbot_ab43711db21b8f66, []int{8}
}
func (m *RequestChatData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestChatData.Unmarshal(m, b)
}
func (m *RequestChatData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestChatData.Marshal(b, m, deterministic)
}
func (dst *RequestChatData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestChatData.Merge(dst, src)
}
func (m *RequestChatData) XXX_Size() int {
	return xxx_messageInfo_RequestChatData.Size(m)
}
func (m *RequestChatData) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestChatData.DiscardUnknown(m)
}

var xxx_messageInfo_RequestChatData proto.InternalMessageInfo

func (m *RequestChatData) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RequestChatData) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func init() {
	proto.RegisterType((*UserAppInfo)(nil), "chatbotpb.UserAppInfo")
	proto.RegisterType((*UserInfo)(nil), "chatbotpb.UserInfo")
	proto.RegisterType((*AppGroupInfo)(nil), "chatbotpb.AppGroupInfo")
	proto.RegisterType((*AppServInfo)(nil), "chatbotpb.AppServInfo")
	proto.RegisterType((*RegisterAppService)(nil), "chatbotpb.RegisterAppService")
	proto.RegisterType((*ReplyRegisterAppService)(nil), "chatbotpb.ReplyRegisterAppService")
	proto.RegisterType((*ChatMsg)(nil), "chatbotpb.ChatMsg")
	proto.RegisterType((*ChatMsgStream)(nil), "chatbotpb.ChatMsgStream")
	proto.RegisterType((*RequestChatData)(nil), "chatbotpb.RequestChatData")
	proto.RegisterEnum("chatbotpb.ChatAppType", ChatAppType_name, ChatAppType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatBotServiceClient is the client API for ChatBotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatBotServiceClient interface {
	// registerAppService - register app service
	RegisterAppService(ctx context.Context, in *RegisterAppService, opts ...grpc.CallOption) (*ReplyRegisterAppService, error)
	// sendChat - send chat
	SendChat(ctx context.Context, opts ...grpc.CallOption) (ChatBotService_SendChatClient, error)
	// requestChat - request chat
	RequestChat(ctx context.Context, in *RequestChatData, opts ...grpc.CallOption) (ChatBotService_RequestChatClient, error)
}

type chatBotServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatBotServiceClient(cc *grpc.ClientConn) ChatBotServiceClient {
	return &chatBotServiceClient{cc}
}

func (c *chatBotServiceClient) RegisterAppService(ctx context.Context, in *RegisterAppService, opts ...grpc.CallOption) (*ReplyRegisterAppService, error) {
	out := new(ReplyRegisterAppService)
	err := c.cc.Invoke(ctx, "/chatbotpb.ChatBotService/registerAppService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatBotServiceClient) SendChat(ctx context.Context, opts ...grpc.CallOption) (ChatBotService_SendChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatBotService_serviceDesc.Streams[0], "/chatbotpb.ChatBotService/sendChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatBotServiceSendChatClient{stream}
	return x, nil
}

type ChatBotService_SendChatClient interface {
	Send(*ChatMsgStream) error
	Recv() (*ChatMsgStream, error)
	grpc.ClientStream
}

type chatBotServiceSendChatClient struct {
	grpc.ClientStream
}

func (x *chatBotServiceSendChatClient) Send(m *ChatMsgStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatBotServiceSendChatClient) Recv() (*ChatMsgStream, error) {
	m := new(ChatMsgStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatBotServiceClient) RequestChat(ctx context.Context, in *RequestChatData, opts ...grpc.CallOption) (ChatBotService_RequestChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatBotService_serviceDesc.Streams[1], "/chatbotpb.ChatBotService/requestChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatBotServiceRequestChatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatBotService_RequestChatClient interface {
	Recv() (*ChatMsgStream, error)
	grpc.ClientStream
}

type chatBotServiceRequestChatClient struct {
	grpc.ClientStream
}

func (x *chatBotServiceRequestChatClient) Recv() (*ChatMsgStream, error) {
	m := new(ChatMsgStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatBotServiceServer is the server API for ChatBotService service.
type ChatBotServiceServer interface {
	// registerAppService - register app service
	RegisterAppService(context.Context, *RegisterAppService) (*ReplyRegisterAppService, error)
	// sendChat - send chat
	SendChat(ChatBotService_SendChatServer) error
	// requestChat - request chat
	RequestChat(*RequestChatData, ChatBotService_RequestChatServer) error
}

func RegisterChatBotServiceServer(s *grpc.Server, srv ChatBotServiceServer) {
	s.RegisterService(&_ChatBotService_serviceDesc, srv)
}

func _ChatBotService_RegisterAppService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAppService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatBotServiceServer).RegisterAppService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatbotpb.ChatBotService/RegisterAppService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatBotServiceServer).RegisterAppService(ctx, req.(*RegisterAppService))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatBotService_SendChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatBotServiceServer).SendChat(&chatBotServiceSendChatServer{stream})
}

type ChatBotService_SendChatServer interface {
	Send(*ChatMsgStream) error
	Recv() (*ChatMsgStream, error)
	grpc.ServerStream
}

type chatBotServiceSendChatServer struct {
	grpc.ServerStream
}

func (x *chatBotServiceSendChatServer) Send(m *ChatMsgStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatBotServiceSendChatServer) Recv() (*ChatMsgStream, error) {
	m := new(ChatMsgStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatBotService_RequestChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestChatData)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatBotServiceServer).RequestChat(m, &chatBotServiceRequestChatServer{stream})
}

type ChatBotService_RequestChatServer interface {
	Send(*ChatMsgStream) error
	grpc.ServerStream
}

type chatBotServiceRequestChatServer struct {
	grpc.ServerStream
}

func (x *chatBotServiceRequestChatServer) Send(m *ChatMsgStream) error {
	return x.ServerStream.SendMsg(m)
}

var _ChatBotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chatbotpb.ChatBotService",
	HandlerType: (*ChatBotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerAppService",
			Handler:    _ChatBotService_RegisterAppService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sendChat",
			Handler:       _ChatBotService_SendChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "requestChat",
			Handler:       _ChatBotService_RequestChat_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chatbot.proto",
}

func init() { proto.RegisterFile("chatbot.proto", fileDescriptor_chatbot_ab43711db21b8f66) }

var fileDescriptor_chatbot_ab43711db21b8f66 = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xee, 0xc6, 0xf9, 0x1d, 0xf7, 0x27, 0x5a, 0x55, 0xad, 0x15, 0x81, 0x14, 0x59, 0x08, 0x85,
	0x1e, 0x4a, 0x15, 0x9e, 0x20, 0xb4, 0xa5, 0x54, 0x6a, 0x55, 0x75, 0x5b, 0x0e, 0x9c, 0xd0, 0xb6,
	0xd9, 0x3a, 0x16, 0x8d, 0xbd, 0x78, 0xd7, 0x48, 0x3d, 0x71, 0xe3, 0xca, 0x9d, 0xb7, 0x40, 0xbc,
	0x19, 0x47, 0x4e, 0x68, 0x26, 0xb6, 0x63, 0xa7, 0x09, 0x82, 0xdb, 0xce, 0xcf, 0xce, 0x7c, 0xdf,
	0xb7, 0x33, 0x36, 0x6c, 0xdc, 0x4e, 0xa4, 0xbd, 0x89, 0xed, 0xbe, 0x4e, 0x62, 0x1b, 0xf3, 0x4e,
	0x66, 0xea, 0x1b, 0xff, 0x07, 0x03, 0xf7, 0x9d, 0x51, 0xc9, 0x48, 0xeb, 0xd3, 0xe8, 0x2e, 0xe6,
	0x03, 0x70, 0xa4, 0xd6, 0x1e, 0xeb, 0xb3, 0xc1, 0xe6, 0x70, 0x67, 0xbf, 0x48, 0xdc, 0x3f, 0x9c,
	0x48, 0x3b, 0xd2, 0xfa, 0xfa, 0x41, 0x2b, 0x81, 0x29, 0x7c, 0x07, 0x9a, 0x52, 0xeb, 0x34, 0x1c,
	0x7b, 0xb5, 0x3e, 0x1b, 0x74, 0x44, 0x66, 0xf1, 0x1e, 0xb4, 0xf1, 0x14, 0xc9, 0xa9, 0xf2, 0x1c,
	0x8a, 0x14, 0x36, 0xc6, 0xb0, 0x62, 0x94, 0x4e, 0x8d, 0x57, 0xef, 0xb3, 0x41, 0x43, 0x14, 0x36,
	0x1f, 0xc0, 0x56, 0x6a, 0x54, 0x82, 0x79, 0x23, 0xad, 0xaf, 0x54, 0xf2, 0xd9, 0x6b, 0xd0, 0xf5,
	0x45, 0xb7, 0xff, 0x95, 0x41, 0x1b, 0x31, 0x13, 0xe0, 0x2e, 0x38, 0x88, 0x01, 0x01, 0x3b, 0x02,
	0x8f, 0x9c, 0x43, 0x9d, 0x9a, 0xcf, 0x60, 0xd1, 0x99, 0xef, 0x41, 0x5d, 0x6a, 0x6d, 0x3c, 0xa7,
	0xef, 0x0c, 0xdc, 0x0a, 0xaf, 0x12, 0x79, 0x41, 0x39, 0x78, 0xdf, 0xca, 0x00, 0x01, 0x3a, 0x78,
	0x1f, 0xcf, 0x7c, 0x1b, 0x1a, 0xd3, 0x38, 0x52, 0x0f, 0x1e, 0x50, 0x9f, 0x99, 0xe1, 0x7f, 0x67,
	0xb0, 0x3e, 0xd2, 0xfa, 0x24, 0x89, 0xd3, 0xff, 0x55, 0x6f, 0x09, 0xdb, 0xda, 0x52, 0xb6, 0xdc,
	0x83, 0x56, 0x80, 0x0d, 0xc2, 0x71, 0x26, 0x67, 0x6e, 0xf2, 0x27, 0xd0, 0xa1, 0x23, 0xb1, 0xad,
	0x53, 0x6c, 0xee, 0xf0, 0xbf, 0x31, 0x70, 0xb3, 0x1a, 0x84, 0x6d, 0x1b, 0x1a, 0x36, 0xfe, 0xa8,
	0x22, 0x42, 0xd7, 0x11, 0x33, 0x83, 0x1f, 0x40, 0x4b, 0xce, 0x70, 0x51, 0xff, 0xd5, 0xa8, 0xf3,
	0x34, 0x7c, 0xc3, 0x1c, 0x62, 0xfe, 0xbe, 0xb9, 0x8d, 0x88, 0x8c, 0x32, 0x26, 0x8c, 0xa3, 0x70,
	0x9c, 0x23, 0x2a, 0x1c, 0xfe, 0x1b, 0xe0, 0x42, 0x05, 0xa1, 0xb1, 0xa4, 0x38, 0x02, 0x0b, 0x6f,
	0x55, 0x86, 0x80, 0x14, 0x40, 0x64, 0xd5, 0xd7, 0x29, 0x11, 0x10, 0x79, 0x9a, 0xff, 0x05, 0x76,
	0x85, 0xd2, 0xf7, 0x0f, 0x2b, 0x8b, 0x11, 0x1d, 0xf6, 0x6f, 0x74, 0xb6, 0xa1, 0xa1, 0x92, 0x24,
	0x4e, 0x32, 0xf9, 0x67, 0x46, 0x89, 0xc8, 0xe9, 0x51, 0xc6, 0x72, 0xee, 0xf0, 0x7f, 0x31, 0x68,
	0x61, 0xb1, 0x73, 0x13, 0xe0, 0xfc, 0x4d, 0x4d, 0x90, 0x89, 0x8a, 0x47, 0x1c, 0x82, 0x54, 0x86,
	0x54, 0x6f, 0xf5, 0xa8, 0x61, 0x0a, 0x7f, 0x01, 0x8e, 0x0c, 0x42, 0xaa, 0xef, 0x0e, 0x77, 0xab,
	0xb4, 0x8b, 0xa1, 0x12, 0x98, 0x83, 0xaa, 0xdf, 0x85, 0xf7, 0xaa, 0xf4, 0xd4, 0x85, 0x9d, 0xc7,
	0x8e, 0xa4, 0x95, 0xb4, 0x32, 0xeb, 0xa2, 0xb0, 0xe7, 0xf4, 0x9a, 0x65, 0x7a, 0xc5, 0x2c, 0xb4,
	0xca, 0xb3, 0x50, 0x21, 0xdd, 0x5e, 0x24, 0xfd, 0xb3, 0x06, 0x1b, 0x19, 0xe9, 0x2b, 0x9b, 0x28,
	0x39, 0xe5, 0x7d, 0x70, 0x6d, 0x6c, 0xe5, 0xfd, 0x99, 0x8a, 0x02, 0x3b, 0x21, 0x09, 0x1a, 0xa2,
	0xec, 0xa2, 0x7d, 0x4f, 0x93, 0x2b, 0x2b, 0x13, 0x4b, 0x7a, 0xe0, 0xbe, 0x67, 0x36, 0x76, 0xbb,
	0x4d, 0x93, 0xec, 0xae, 0x43, 0xc1, 0xb9, 0x03, 0x6f, 0x4e, 0xa4, 0x99, 0x10, 0xa7, 0x8c, 0x6f,
	0x6e, 0xf3, 0x67, 0xb0, 0x41, 0x4d, 0xde, 0xe6, 0x09, 0xb3, 0xef, 0x44, 0xd5, 0x89, 0x6b, 0x3c,
	0xc6, 0x60, 0x93, 0x14, 0xa1, 0xf3, 0x5c, 0x8d, 0xd6, 0x52, 0x35, 0xda, 0x2b, 0xd5, 0xe8, 0x2c,
	0xa8, 0xc1, 0x9f, 0x43, 0x1d, 0x9f, 0x8b, 0xbe, 0x07, 0xee, 0x90, 0x2f, 0x4c, 0xd9, 0xb9, 0x09,
	0x04, 0xc5, 0xfd, 0x63, 0xd8, 0x12, 0xea, 0x53, 0xaa, 0x8c, 0x45, 0x7f, 0xfe, 0x24, 0x4b, 0x16,
	0xb1, 0xd2, 0xae, 0xb6, 0xd0, 0x6e, 0xef, 0x25, 0xb8, 0xa5, 0xe9, 0xe5, 0x5d, 0x58, 0x3f, 0x1c,
	0x5d, 0x7f, 0xb8, 0x3e, 0x3e, 0x3b, 0x3e, 0x11, 0xa3, 0xf3, 0xee, 0x1a, 0xdf, 0x04, 0x40, 0xcf,
	0xe1, 0xc5, 0xc5, 0xd9, 0xe5, 0x65, 0x97, 0x0d, 0x7f, 0x33, 0xd8, 0xc4, 0x1b, 0xaf, 0x63, 0x9b,
	0xef, 0xc6, 0x7b, 0xe0, 0xc9, 0xe3, 0x8d, 0x79, 0x5a, 0x82, 0xfe, 0x78, 0xa1, 0x7a, 0x7e, 0x25,
	0xbc, 0x74, 0xe9, 0xfc, 0x35, 0x7e, 0x04, 0x6d, 0xa3, 0xa2, 0x31, 0x36, 0xe4, 0xde, 0x63, 0x2d,
	0x66, 0xf3, 0xd2, 0x5b, 0x19, 0xf1, 0xd7, 0x06, 0xec, 0x80, 0xf1, 0x13, 0x70, 0x93, 0xb9, 0x56,
	0xbc, 0x57, 0x69, 0x5d, 0xd1, 0xf0, 0x6f, 0xa5, 0x0e, 0xd8, 0x4d, 0x93, 0x7e, 0x73, 0xaf, 0xfe,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xbc, 0x9a, 0xf2, 0xf7, 0x06, 0x00, 0x00,
}

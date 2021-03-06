// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Defines the applications that the user is associated with.
type ApplicationLink struct {
	ApplicationID   int64  `protobuf:"varint,1,opt,name=applicationID" json:"applicationID,omitempty"`
	ApplicationName string `protobuf:"bytes,2,opt,name=applicationName" json:"applicationName,omitempty"`
	IsAdmin         bool   `protobuf:"varint,3,opt,name=isAdmin" json:"isAdmin,omitempty"`
	CreatedAt       string `protobuf:"bytes,4,opt,name=createdAt" json:"createdAt,omitempty"`
	UpdatedAt       string `protobuf:"bytes,5,opt,name=updatedAt" json:"updatedAt,omitempty"`
}

func (m *ApplicationLink) Reset()                    { *m = ApplicationLink{} }
func (m *ApplicationLink) String() string            { return proto.CompactTextString(m) }
func (*ApplicationLink) ProtoMessage()               {}
func (*ApplicationLink) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *ApplicationLink) GetApplicationID() int64 {
	if m != nil {
		return m.ApplicationID
	}
	return 0
}

func (m *ApplicationLink) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *ApplicationLink) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *ApplicationLink) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ApplicationLink) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

// The user profile, indicating the capabilities of the user.
type UserProfile struct {
	Applications []*ApplicationLink `protobuf:"bytes,1,rep,name=applications" json:"applications,omitempty"`
}

func (m *UserProfile) Reset()                    { *m = UserProfile{} }
func (m *UserProfile) String() string            { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()               {}
func (*UserProfile) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *UserProfile) GetApplications() []*ApplicationLink {
	if m != nil {
		return m.Applications
	}
	return nil
}

// The request for profile requires not input as the profile is returned for
// the logged in user based on the JWT token passed in.
type ProfileRequest struct {
}

func (m *ProfileRequest) Reset()                    { *m = ProfileRequest{} }
func (m *ProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*ProfileRequest) ProtoMessage()               {}
func (*ProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

// The user's profile. This specifies the access the user has to the data
// and should be used limit what the UI presents as options (e.g., read-only
// access to an application should not allow an "edit" button on that
// application).
type ProfileResponse struct {
	User         *GetUserResponse   `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Applications []*ApplicationLink `protobuf:"bytes,2,rep,name=applications" json:"applications,omitempty"`
}

func (m *ProfileResponse) Reset()                    { *m = ProfileResponse{} }
func (m *ProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*ProfileResponse) ProtoMessage()               {}
func (*ProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *ProfileResponse) GetUser() *GetUserResponse {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ProfileResponse) GetApplications() []*ApplicationLink {
	if m != nil {
		return m.Applications
	}
	return nil
}

// The data for logging in to the system.  The connection is expected to be
// protected by SSL, allowing us to use the username and password in the clear.
type LoginRequest struct {
	// Username of the user.
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	// Password of the user.
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// The response to the login request upon success. The jwt token is to be
// placed in the header field named "Grpc-Metadata-Authorization" for all
// subsequent queries to the server.
type LoginResponse struct {
	// The JWT tag to be used to access lora-app-server interfaces.
	Jwt string `protobuf:"bytes,1,opt,name=jwt" json:"jwt,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *LoginResponse) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

// Request the users defined in the system.
type ListUserRequest struct {
	// Max number of user to return in the result-set.
	Limit int32 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int32 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	// When provided, the given string will be used to search on username.
	Search string `protobuf:"bytes,3,opt,name=search" json:"search,omitempty"`
}

func (m *ListUserRequest) Reset()                    { *m = ListUserRequest{} }
func (m *ListUserRequest) String() string            { return proto.CompactTextString(m) }
func (*ListUserRequest) ProtoMessage()               {}
func (*ListUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *ListUserRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListUserRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListUserRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

// Request the user information.
type UserRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *UserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type AddUserResponse struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *AddUserResponse) Reset()                    { *m = AddUserResponse{} }
func (m *AddUserResponse) String() string            { return proto.CompactTextString(m) }
func (*AddUserResponse) ProtoMessage()               {}
func (*AddUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *AddUserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// User data not including the data access profile.
type UserSettings struct {
	// ID of the user.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Username of the user.
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	// The session timeout, in minutes.
	SessionTTL int32 `protobuf:"varint,3,opt,name=sessionTTL" json:"sessionTTL,omitempty"`
	// If the user is a system admin, capable of creating other users.
	IsAdmin bool `protobuf:"varint,4,opt,name=isAdmin" json:"isAdmin,omitempty"`
	// If the user is active.
	IsActive bool `protobuf:"varint,5,opt,name=isActive" json:"isActive,omitempty"`
	// When the user was created.
	CreatedAt string `protobuf:"bytes,6,opt,name=createdAt" json:"createdAt,omitempty"`
	// When the user was last updated (excludes changes in application access).
	UpdatedAt string `protobuf:"bytes,7,opt,name=updatedAt" json:"updatedAt,omitempty"`
}

func (m *UserSettings) Reset()                    { *m = UserSettings{} }
func (m *UserSettings) String() string            { return proto.CompactTextString(m) }
func (*UserSettings) ProtoMessage()               {}
func (*UserSettings) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *UserSettings) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserSettings) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserSettings) GetSessionTTL() int32 {
	if m != nil {
		return m.SessionTTL
	}
	return 0
}

func (m *UserSettings) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *UserSettings) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *UserSettings) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserSettings) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

// UserInfo includes the general login settings as well as the user profile.
type UserInfo struct {
	UserSettings *UserSettings `protobuf:"bytes,1,opt,name=userSettings" json:"userSettings,omitempty"`
	UserProfile  *UserProfile  `protobuf:"bytes,2,opt,name=userProfile" json:"userProfile,omitempty"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

func (m *UserInfo) GetUserSettings() *UserSettings {
	if m != nil {
		return m.UserSettings
	}
	return nil
}

func (m *UserInfo) GetUserProfile() *UserProfile {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

type GetUserResponse struct {
	// ID of the user.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Username of the user.
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	// The session timeout, in minutes.
	SessionTTL int32 `protobuf:"varint,3,opt,name=sessionTTL" json:"sessionTTL,omitempty"`
	// If the user is a system admin, capable of creating other users.
	IsAdmin bool `protobuf:"varint,4,opt,name=isAdmin" json:"isAdmin,omitempty"`
	// If the user is active.
	IsActive bool `protobuf:"varint,5,opt,name=isActive" json:"isActive,omitempty"`
	// When the user was created.
	CreatedAt string `protobuf:"bytes,6,opt,name=createdAt" json:"createdAt,omitempty"`
	// When the user was last updated (excludes changes in application access).
	UpdatedAt string `protobuf:"bytes,7,opt,name=updatedAt" json:"updatedAt,omitempty"`
}

func (m *GetUserResponse) Reset()                    { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()               {}
func (*GetUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{11} }

func (m *GetUserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetUserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetUserResponse) GetSessionTTL() int32 {
	if m != nil {
		return m.SessionTTL
	}
	return 0
}

func (m *GetUserResponse) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *GetUserResponse) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *GetUserResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetUserResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

// Add a new user. Not quite the UserSettings data as it includes a password
// and excludes the ID and create/update dates.
type AddUserRequest struct {
	// Username of the user.
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	// Passowrd of the user.
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	// The session timeout, in minutes.
	SessionTTL int32 `protobuf:"varint,3,opt,name=sessionTTL" json:"sessionTTL,omitempty"`
	// If the user is a system-wide admin.
	IsAdmin bool `protobuf:"varint,4,opt,name=isAdmin" json:"isAdmin,omitempty"`
	// If the user is active.
	IsActive bool `protobuf:"varint,5,opt,name=isActive" json:"isActive,omitempty"`
}

func (m *AddUserRequest) Reset()                    { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string            { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()               {}
func (*AddUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{12} }

func (m *AddUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AddUserRequest) GetSessionTTL() int32 {
	if m != nil {
		return m.SessionTTL
	}
	return 0
}

func (m *AddUserRequest) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *AddUserRequest) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

// Not quite the AddUserRequest as no password.
type UpdateUserRequest struct {
	// The ID of the user to be updated.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The new username.
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	// The session timeout, in minutes.
	SessionTTL int32 `protobuf:"varint,3,opt,name=sessionTTL" json:"sessionTTL,omitempty"`
	// If the user is a system-wide admin.
	IsAdmin bool `protobuf:"varint,4,opt,name=isAdmin" json:"isAdmin,omitempty"`
	// If the user is active.
	IsActive bool `protobuf:"varint,5,opt,name=isActive" json:"isActive,omitempty"`
}

func (m *UpdateUserRequest) Reset()                    { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()               {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{13} }

func (m *UpdateUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpdateUserRequest) GetSessionTTL() int32 {
	if m != nil {
		return m.SessionTTL
	}
	return 0
}

func (m *UpdateUserRequest) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *UpdateUserRequest) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

type ListUserResponse struct {
	TotalCount int32              `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	Result     []*GetUserResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListUserResponse) Reset()                    { *m = ListUserResponse{} }
func (m *ListUserResponse) String() string            { return proto.CompactTextString(m) }
func (*ListUserResponse) ProtoMessage()               {}
func (*ListUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{14} }

func (m *ListUserResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListUserResponse) GetResult() []*GetUserResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

type UserEmptyResponse struct {
}

func (m *UserEmptyResponse) Reset()                    { *m = UserEmptyResponse{} }
func (m *UserEmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*UserEmptyResponse) ProtoMessage()               {}
func (*UserEmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{15} }

type UpdateUserPasswordRequest struct {
	// The ID of the user for which to update the password.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The new password to set.
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *UpdateUserPasswordRequest) Reset()                    { *m = UpdateUserPasswordRequest{} }
func (m *UpdateUserPasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserPasswordRequest) ProtoMessage()               {}
func (*UpdateUserPasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{16} }

func (m *UpdateUserPasswordRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*ApplicationLink)(nil), "api.ApplicationLink")
	proto.RegisterType((*UserProfile)(nil), "api.UserProfile")
	proto.RegisterType((*ProfileRequest)(nil), "api.ProfileRequest")
	proto.RegisterType((*ProfileResponse)(nil), "api.ProfileResponse")
	proto.RegisterType((*LoginRequest)(nil), "api.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "api.LoginResponse")
	proto.RegisterType((*ListUserRequest)(nil), "api.ListUserRequest")
	proto.RegisterType((*UserRequest)(nil), "api.UserRequest")
	proto.RegisterType((*AddUserResponse)(nil), "api.AddUserResponse")
	proto.RegisterType((*UserSettings)(nil), "api.UserSettings")
	proto.RegisterType((*UserInfo)(nil), "api.UserInfo")
	proto.RegisterType((*GetUserResponse)(nil), "api.GetUserResponse")
	proto.RegisterType((*AddUserRequest)(nil), "api.AddUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "api.UpdateUserRequest")
	proto.RegisterType((*ListUserResponse)(nil), "api.ListUserResponse")
	proto.RegisterType((*UserEmptyResponse)(nil), "api.UserEmptyResponse")
	proto.RegisterType((*UpdateUserPasswordRequest)(nil), "api.UpdateUserPasswordRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for User service

type UserClient interface {
	// Get user list.
	List(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
	// Get data for a particular user.
	Get(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// Create a new user.
	Create(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	// Update an existing user.
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserEmptyResponse, error)
	// Delete a user.
	Delete(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserEmptyResponse, error)
	// UpdatePassword updates a password.
	UpdatePassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*UserEmptyResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) List(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := grpc.Invoke(ctx, "/api.User/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Get(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := grpc.Invoke(ctx, "/api.User/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Create(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := grpc.Invoke(ctx, "/api.User/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserEmptyResponse, error) {
	out := new(UserEmptyResponse)
	err := grpc.Invoke(ctx, "/api.User/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Delete(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserEmptyResponse, error) {
	out := new(UserEmptyResponse)
	err := grpc.Invoke(ctx, "/api.User/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdatePassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*UserEmptyResponse, error) {
	out := new(UserEmptyResponse)
	err := grpc.Invoke(ctx, "/api.User/UpdatePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	// Get user list.
	List(context.Context, *ListUserRequest) (*ListUserResponse, error)
	// Get data for a particular user.
	Get(context.Context, *UserRequest) (*GetUserResponse, error)
	// Create a new user.
	Create(context.Context, *AddUserRequest) (*AddUserResponse, error)
	// Update an existing user.
	Update(context.Context, *UpdateUserRequest) (*UserEmptyResponse, error)
	// Delete a user.
	Delete(context.Context, *UserRequest) (*UserEmptyResponse, error)
	// UpdatePassword updates a password.
	UpdatePassword(context.Context, *UpdateUserPasswordRequest) (*UserEmptyResponse, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).List(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Get(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Create(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Delete(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdatePassword(ctx, req.(*UpdateUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _User_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _User_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _User_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _User_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _User_Delete_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _User_UpdatePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// Client API for Internal service

type InternalClient interface {
	// Log in a user
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Get the current user's profile
	Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
}

type internalClient struct {
	cc *grpc.ClientConn
}

func NewInternalClient(cc *grpc.ClientConn) InternalClient {
	return &internalClient{cc}
}

func (c *internalClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/api.Internal/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalClient) Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := grpc.Invoke(ctx, "/api.Internal/Profile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Internal service

type InternalServer interface {
	// Log in a user
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Get the current user's profile
	Profile(context.Context, *ProfileRequest) (*ProfileResponse, error)
}

func RegisterInternalServer(s *grpc.Server, srv InternalServer) {
	s.RegisterService(&_Internal_serviceDesc, srv)
}

func _Internal_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Internal/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internal_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Internal/Profile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).Profile(ctx, req.(*ProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Internal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Internal",
	HandlerType: (*InternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Internal_Login_Handler,
		},
		{
			MethodName: "Profile",
			Handler:    _Internal_Profile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 826 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x05, 0xf5, 0xb2, 0x7c, 0x25, 0x8b, 0xd2, 0xf8, 0xc5, 0x12, 0xb6, 0xa1, 0x4e, 0xbb, 0x10,
	0x8c, 0xc2, 0x02, 0x54, 0x14, 0x28, 0xbc, 0x13, 0xec, 0x5a, 0x70, 0x2b, 0x18, 0x06, 0x6d, 0xa3,
	0xdb, 0xb2, 0xe2, 0x48, 0x9d, 0x96, 0x22, 0x59, 0xce, 0xb0, 0x46, 0x51, 0x74, 0x93, 0x5f, 0x08,
	0xb2, 0xcd, 0x2f, 0x64, 0x99, 0x9f, 0x08, 0x90, 0x45, 0x7e, 0x21, 0xbf, 0x90, 0x7d, 0x30, 0xc3,
	0x21, 0x45, 0x8e, 0x1f, 0x08, 0x10, 0x24, 0x40, 0x76, 0x9a, 0x3b, 0x57, 0xe7, 0x9e, 0x7b, 0xe6,
	0xdc, 0xe1, 0x00, 0x24, 0x8c, 0xc4, 0x47, 0x51, 0x1c, 0xf2, 0x10, 0x55, 0xdd, 0x88, 0xda, 0x7b,
	0x8b, 0x30, 0x5c, 0xf8, 0x64, 0xe8, 0x46, 0x74, 0xe8, 0x06, 0x41, 0xc8, 0x5d, 0x4e, 0xc3, 0x80,
	0xa5, 0x29, 0xf8, 0xa5, 0x01, 0xe6, 0x38, 0x8a, 0x7c, 0x3a, 0x93, 0xe1, 0x29, 0x0d, 0xfe, 0x42,
	0xdf, 0xc2, 0x86, 0xbb, 0x0a, 0x9d, 0x9f, 0x5a, 0x46, 0xdf, 0x18, 0x54, 0x9d, 0x72, 0x10, 0x0d,
	0xc0, 0x2c, 0x04, 0x2e, 0xdc, 0x25, 0xb1, 0x2a, 0x7d, 0x63, 0xb0, 0xee, 0xe8, 0x61, 0x64, 0xc1,
	0x1a, 0x65, 0x63, 0x6f, 0x49, 0x03, 0xab, 0xda, 0x37, 0x06, 0x4d, 0x27, 0x5b, 0xa2, 0x3d, 0x58,
	0x9f, 0xc5, 0xc4, 0xe5, 0xc4, 0x1b, 0x73, 0xab, 0x26, 0xff, 0xbd, 0x0a, 0x88, 0xdd, 0x24, 0xf2,
	0xd4, 0x6e, 0x3d, 0xdd, 0xcd, 0x03, 0x78, 0x02, 0xad, 0x1b, 0x46, 0xe2, 0xcb, 0x38, 0x9c, 0x53,
	0x9f, 0xa0, 0x1f, 0xa1, 0x5d, 0xa8, 0xcb, 0x2c, 0xa3, 0x5f, 0x1d, 0xb4, 0x46, 0x5b, 0x47, 0x6e,
	0x44, 0x8f, 0xb4, 0x06, 0x9d, 0x52, 0x26, 0xee, 0x42, 0x47, 0x81, 0x38, 0xe4, 0xef, 0x84, 0x30,
	0x8e, 0x13, 0x30, 0xf3, 0x08, 0x8b, 0xc2, 0x80, 0x11, 0x34, 0x80, 0x9a, 0x10, 0x56, 0x4a, 0x91,
	0xc1, 0x4e, 0x08, 0x17, 0x0c, 0xb2, 0x1c, 0x47, 0x66, 0xdc, 0x21, 0x52, 0xf9, 0x60, 0x22, 0x67,
	0xd0, 0x9e, 0x86, 0x0b, 0x1a, 0x28, 0x1a, 0xc8, 0x86, 0xa6, 0x40, 0x0c, 0x84, 0xb4, 0x86, 0x6c,
	0x3f, 0x5f, 0x8b, 0xbd, 0xc8, 0x65, 0xec, 0x36, 0x8c, 0x3d, 0x25, 0x7b, 0xbe, 0xc6, 0x5f, 0xc3,
	0x86, 0xc2, 0x51, 0xe4, 0xbb, 0x50, 0xfd, 0xf3, 0x96, 0x2b, 0x0c, 0xf1, 0x13, 0xff, 0x0a, 0xe6,
	0x94, 0x32, 0x45, 0x3f, 0xad, 0xb6, 0x05, 0x75, 0x9f, 0x2e, 0x69, 0x9a, 0x56, 0x77, 0xd2, 0x05,
	0xda, 0x81, 0x46, 0x38, 0x9f, 0x33, 0xc2, 0x65, 0x95, 0xba, 0xa3, 0x56, 0x22, 0xce, 0x88, 0x1b,
	0xcf, 0xfe, 0x90, 0x47, 0xba, 0xee, 0xa8, 0x15, 0xde, 0x4f, 0x4f, 0x25, 0x03, 0xed, 0x40, 0x85,
	0x7a, 0xca, 0x3f, 0x15, 0x2a, 0xa8, 0x99, 0x63, 0xcf, 0x2b, 0xaa, 0x76, 0x27, 0xe5, 0x95, 0x01,
	0x6d, 0x91, 0x70, 0x45, 0x38, 0xa7, 0xc1, 0x82, 0xe9, 0x09, 0x25, 0x59, 0x2a, 0x9a, 0x2c, 0x07,
	0x00, 0x8c, 0x30, 0x46, 0xc3, 0xe0, 0xfa, 0x7a, 0x2a, 0xa9, 0xd5, 0x9d, 0x42, 0xa4, 0x68, 0xc5,
	0x5a, 0xd9, 0x8a, 0x36, 0x34, 0x29, 0x1b, 0xcf, 0x38, 0xfd, 0x87, 0x48, 0xaf, 0x35, 0x9d, 0x7c,
	0x5d, 0xb6, 0x69, 0xe3, 0x51, 0x9b, 0xae, 0xe9, 0x36, 0x4d, 0xa0, 0x29, 0xba, 0x39, 0x0f, 0xe6,
	0x21, 0xfa, 0x01, 0xda, 0x49, 0xa1, 0x33, 0x65, 0xa6, 0x9e, 0xb4, 0x46, 0xb1, 0x65, 0xa7, 0x94,
	0x86, 0x46, 0xd0, 0x4a, 0x56, 0x4e, 0x97, 0x3d, 0xb7, 0x46, 0xdd, 0xfc, 0x5f, 0x99, 0x55, 0x8b,
	0x49, 0xf8, 0xb5, 0x01, 0xa6, 0xe6, 0xcf, 0x2f, 0x5c, 0xc8, 0xe7, 0x06, 0x74, 0x72, 0xef, 0x7c,
	0xd4, 0x80, 0x7c, 0x9a, 0xe6, 0xf0, 0x33, 0x03, 0x7a, 0x37, 0x92, 0xee, 0x23, 0x13, 0xf0, 0xf9,
	0x45, 0xc7, 0xbf, 0x41, 0x77, 0x35, 0xeb, 0xca, 0x0a, 0x07, 0x00, 0x3c, 0xe4, 0xae, 0x7f, 0x12,
	0x26, 0x41, 0x36, 0xf1, 0x85, 0x08, 0xfa, 0x0e, 0x1a, 0x31, 0x61, 0x89, 0xcf, 0x4b, 0xd7, 0x97,
	0x7e, 0xe1, 0xa9, 0x1c, 0xbc, 0x09, 0x3d, 0x11, 0xff, 0x69, 0x19, 0xf1, 0x7f, 0xb3, 0x4d, 0x3c,
	0x81, 0xaf, 0x56, 0x6a, 0x5c, 0x2a, 0xe9, 0x1f, 0x51, 0xe5, 0xa1, 0xd3, 0x1a, 0xbd, 0xab, 0x42,
	0x4d, 0x60, 0xa0, 0x09, 0xd4, 0x44, 0x23, 0x28, 0x25, 0xa3, 0xdd, 0x5f, 0xf6, 0xb6, 0x16, 0x55,
	0x34, 0xd0, 0x93, 0x37, 0x6f, 0x9f, 0x56, 0xda, 0x08, 0xe4, 0x07, 0x50, 0x48, 0xcd, 0xd0, 0x19,
	0x54, 0x27, 0x84, 0xa3, 0xd5, 0x08, 0x65, 0x18, 0xf7, 0xb6, 0x89, 0x77, 0x25, 0x44, 0x0f, 0x99,
	0x2b, 0x88, 0xe1, 0x7f, 0xd4, 0xfb, 0x1f, 0xfd, 0x0c, 0x8d, 0x13, 0xe9, 0x5e, 0xb4, 0x99, 0x5e,
	0xef, 0x25, 0x7b, 0x2a, 0x34, 0xed, 0xbe, 0xc3, 0xdb, 0x12, 0xcd, 0xc4, 0x05, 0x42, 0xc7, 0xc6,
	0x21, 0xba, 0x86, 0x46, 0x2a, 0x17, 0xda, 0x49, 0x69, 0xe9, 0x4e, 0xb2, 0x77, 0x72, 0xba, 0x65,
	0xa1, 0x6d, 0x09, 0xb8, 0x65, 0xeb, 0xf4, 0x04, 0xea, 0x2f, 0xd0, 0x38, 0x25, 0x3e, 0xe1, 0xe4,
	0x9e, 0x66, 0x1f, 0xc2, 0x53, 0xed, 0x1e, 0xde, 0x69, 0x77, 0x09, 0x9d, 0x94, 0xd5, 0x65, 0x3e,
	0x48, 0x1a, 0x55, 0xed, 0x98, 0x1f, 0x2c, 0xf1, 0x8d, 0x2c, 0xb1, 0x6f, 0x5b, 0x5a, 0x89, 0x61,
	0x76, 0xe8, 0xc7, 0xc6, 0xe1, 0xe8, 0x85, 0x01, 0xcd, 0xf3, 0x80, 0x8b, 0xd9, 0xf0, 0xd1, 0x05,
	0xd4, 0xe5, 0x37, 0x0d, 0xa5, 0xb7, 0x65, 0xf1, 0x3b, 0x69, 0xa3, 0x62, 0x48, 0x55, 0x38, 0x90,
	0x15, 0x2c, 0xbc, 0x29, 0x2b, 0x50, 0x05, 0x33, 0xf4, 0x45, 0x92, 0x10, 0xe6, 0x0a, 0xd6, 0xb2,
	0x97, 0x43, 0x7a, 0x76, 0xe5, 0x27, 0x80, 0x3a, 0x3b, 0xed, 0x15, 0x80, 0xf7, 0x25, 0xea, 0x2e,
	0xda, 0x2e, 0xa3, 0x46, 0x69, 0xda, 0xef, 0x0d, 0xf9, 0xa6, 0xfa, 0xfe, 0x7d, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x93, 0x71, 0x71, 0x3b, 0x84, 0x09, 0x00, 0x00,
}

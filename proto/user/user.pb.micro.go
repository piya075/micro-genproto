// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*SuccessResponse, error)
	RegisterWithMobileNumber(ctx context.Context, in *RegisterWithMobileNumberRequest, opts ...client.CallOption) (*SuccessResponse, error)
	RegisterWithMobileNumberAis(ctx context.Context, in *RegisterWithMobileNumberAisRequest, opts ...client.CallOption) (*SuccessResponse, error)
	Profile(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*UserDetail, error)
	VerifyUser(ctx context.Context, in *VerifyUserRequest, opts ...client.CallOption) (*VerifyUserResponse, error)
	AutoRegisterBySocial(ctx context.Context, in *SocialRequest, opts ...client.CallOption) (*SocialResponse, error)
	UpdateSocialID(ctx context.Context, in *UpdateSocialIDRequest, opts ...client.CallOption) (*UpdateSocialIDRespose, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...client.CallOption) (*SuccessResponse, error)
	ResetPasswordToken(ctx context.Context, in *ResetPasswordTokenRequest, opts ...client.CallOption) (*ResetPasswordTokenResponse, error)
	UpdateResetPassword(ctx context.Context, in *UpdateResetPasswordRequest, opts ...client.CallOption) (*SuccessResponse, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...client.CallOption) (*SuccessResponse, error)
	UploadProfilePicture(ctx context.Context, in *UploadProfilePictureRequest, opts ...client.CallOption) (*SuccessResponse, error)
	ResendVerifyByOldToken(ctx context.Context, in *VerifyUserRequest, opts ...client.CallOption) (*VerifyUserResponse, error)
	Resendverify(ctx context.Context, in *ResetPasswordRequest, opts ...client.CallOption) (*VerifyUserResponse, error)
	CheckTerm(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CheckTermResponse, error)
	AcceptTerm(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*SuccessResponse, error)
	DeleteAccount(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*SuccessResponse, error)
	TriggerDeleteTemp(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*SuccessResponse, error)
	GetPdpa(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*PDPAResponse, error)
	AcceptPrivacyPolicy(ctx context.Context, in *PrivacyPolicyRequest, opts ...client.CallOption) (*SuccessResponse, error)
	CheckEmailHasAddress(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CheckEmailHasAddressResponse, error)
	EventCalendarLimitExceed(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EventCalendarLimitExceedResponse, error)
	RequestOtp(ctx context.Context, in *RequestOtpRequest, opts ...client.CallOption) (*RequestOtpResponse, error)
	VerifyOtp(ctx context.Context, in *VerifyOtpRequest, opts ...client.CallOption) (*VerifyOtpResponse, error)
	RequestOtpOperationTool(ctx context.Context, in *RequestOtpRequest, opts ...client.CallOption) (*RequestOtpOperationToolResponse, error)
	DeleteUserTemp(ctx context.Context, in *DeleteUserTempRequest, opts ...client.CallOption) (*DeleteUserTempResponse, error)
	Getbytel(ctx context.Context, in *GetByTelRequest, opts ...client.CallOption) (*UserDetailResponse, error)
	VerifyPartnerID(ctx context.Context, in *VerifyPartnerIDRequest, opts ...client.CallOption) (*VerifyPartnerIDResponse, error)
	ResendVerifyEmailMember(ctx context.Context, in *ResetPasswordRequest, opts ...client.CallOption) (*SuccessResponse, error)
	UpdateContactProfile(ctx context.Context, in *UpdateProfileRequest, opts ...client.CallOption) (*SuccessResponse, error)
	CreateUserInterest(ctx context.Context, in *UserInterestRequest, opts ...client.CallOption) (*UserInterestRequestResponse, error)
	CreateUserInterestList(ctx context.Context, in *UsersInterestReqList, opts ...client.CallOption) (*UsersInterestResList, error)
	UpdateUserInterest(ctx context.Context, in *UserInterestUpdateRequest, opts ...client.CallOption) (*UserInterestRequestResponse, error)
	UpdateUserInterestList(ctx context.Context, in *UsersInterestUpdateReqList, opts ...client.CallOption) (*UsersInterestResList, error)
	GetInterestItems(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*InterestItemList, error)
	TestPublishDeleteTel(ctx context.Context, in *UserDetail, opts ...client.CallOption) (*EmptyRequest, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) RegisterWithMobileNumber(ctx context.Context, in *RegisterWithMobileNumberRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.RegisterWithMobileNumber", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) RegisterWithMobileNumberAis(ctx context.Context, in *RegisterWithMobileNumberAisRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.RegisterWithMobileNumberAis", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Profile(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*UserDetail, error) {
	req := c.c.NewRequest(c.name, "User.Profile", in)
	out := new(UserDetail)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) VerifyUser(ctx context.Context, in *VerifyUserRequest, opts ...client.CallOption) (*VerifyUserResponse, error) {
	req := c.c.NewRequest(c.name, "User.VerifyUser", in)
	out := new(VerifyUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AutoRegisterBySocial(ctx context.Context, in *SocialRequest, opts ...client.CallOption) (*SocialResponse, error) {
	req := c.c.NewRequest(c.name, "User.AutoRegisterBySocial", in)
	out := new(SocialResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateSocialID(ctx context.Context, in *UpdateSocialIDRequest, opts ...client.CallOption) (*UpdateSocialIDRespose, error) {
	req := c.c.NewRequest(c.name, "User.UpdateSocialID", in)
	out := new(UpdateSocialIDRespose)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.ResetPassword", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ResetPasswordToken(ctx context.Context, in *ResetPasswordTokenRequest, opts ...client.CallOption) (*ResetPasswordTokenResponse, error) {
	req := c.c.NewRequest(c.name, "User.ResetPasswordToken", in)
	out := new(ResetPasswordTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateResetPassword(ctx context.Context, in *UpdateResetPasswordRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.UpdateResetPassword", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.UpdateProfile", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UploadProfilePicture(ctx context.Context, in *UploadProfilePictureRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.UploadProfilePicture", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ResendVerifyByOldToken(ctx context.Context, in *VerifyUserRequest, opts ...client.CallOption) (*VerifyUserResponse, error) {
	req := c.c.NewRequest(c.name, "User.ResendVerifyByOldToken", in)
	out := new(VerifyUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Resendverify(ctx context.Context, in *ResetPasswordRequest, opts ...client.CallOption) (*VerifyUserResponse, error) {
	req := c.c.NewRequest(c.name, "User.Resendverify", in)
	out := new(VerifyUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) CheckTerm(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CheckTermResponse, error) {
	req := c.c.NewRequest(c.name, "User.CheckTerm", in)
	out := new(CheckTermResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AcceptTerm(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.AcceptTerm", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteAccount(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.DeleteAccount", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) TriggerDeleteTemp(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.TriggerDeleteTemp", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetPdpa(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*PDPAResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetPdpa", in)
	out := new(PDPAResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AcceptPrivacyPolicy(ctx context.Context, in *PrivacyPolicyRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.AcceptPrivacyPolicy", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) CheckEmailHasAddress(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CheckEmailHasAddressResponse, error) {
	req := c.c.NewRequest(c.name, "User.CheckEmailHasAddress", in)
	out := new(CheckEmailHasAddressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) EventCalendarLimitExceed(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EventCalendarLimitExceedResponse, error) {
	req := c.c.NewRequest(c.name, "User.EventCalendarLimitExceed", in)
	out := new(EventCalendarLimitExceedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) RequestOtp(ctx context.Context, in *RequestOtpRequest, opts ...client.CallOption) (*RequestOtpResponse, error) {
	req := c.c.NewRequest(c.name, "User.RequestOtp", in)
	out := new(RequestOtpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) VerifyOtp(ctx context.Context, in *VerifyOtpRequest, opts ...client.CallOption) (*VerifyOtpResponse, error) {
	req := c.c.NewRequest(c.name, "User.VerifyOtp", in)
	out := new(VerifyOtpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) RequestOtpOperationTool(ctx context.Context, in *RequestOtpRequest, opts ...client.CallOption) (*RequestOtpOperationToolResponse, error) {
	req := c.c.NewRequest(c.name, "User.RequestOtpOperationTool", in)
	out := new(RequestOtpOperationToolResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteUserTemp(ctx context.Context, in *DeleteUserTempRequest, opts ...client.CallOption) (*DeleteUserTempResponse, error) {
	req := c.c.NewRequest(c.name, "User.DeleteUserTemp", in)
	out := new(DeleteUserTempResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Getbytel(ctx context.Context, in *GetByTelRequest, opts ...client.CallOption) (*UserDetailResponse, error) {
	req := c.c.NewRequest(c.name, "User.Getbytel", in)
	out := new(UserDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) VerifyPartnerID(ctx context.Context, in *VerifyPartnerIDRequest, opts ...client.CallOption) (*VerifyPartnerIDResponse, error) {
	req := c.c.NewRequest(c.name, "User.VerifyPartnerID", in)
	out := new(VerifyPartnerIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ResendVerifyEmailMember(ctx context.Context, in *ResetPasswordRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.ResendVerifyEmailMember", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateContactProfile(ctx context.Context, in *UpdateProfileRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "User.UpdateContactProfile", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) CreateUserInterest(ctx context.Context, in *UserInterestRequest, opts ...client.CallOption) (*UserInterestRequestResponse, error) {
	req := c.c.NewRequest(c.name, "User.CreateUserInterest", in)
	out := new(UserInterestRequestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) CreateUserInterestList(ctx context.Context, in *UsersInterestReqList, opts ...client.CallOption) (*UsersInterestResList, error) {
	req := c.c.NewRequest(c.name, "User.CreateUserInterestList", in)
	out := new(UsersInterestResList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUserInterest(ctx context.Context, in *UserInterestUpdateRequest, opts ...client.CallOption) (*UserInterestRequestResponse, error) {
	req := c.c.NewRequest(c.name, "User.UpdateUserInterest", in)
	out := new(UserInterestRequestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUserInterestList(ctx context.Context, in *UsersInterestUpdateReqList, opts ...client.CallOption) (*UsersInterestResList, error) {
	req := c.c.NewRequest(c.name, "User.UpdateUserInterestList", in)
	out := new(UsersInterestResList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetInterestItems(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*InterestItemList, error) {
	req := c.c.NewRequest(c.name, "User.GetInterestItems", in)
	out := new(InterestItemList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) TestPublishDeleteTel(ctx context.Context, in *UserDetail, opts ...client.CallOption) (*EmptyRequest, error) {
	req := c.c.NewRequest(c.name, "User.TestPublishDeleteTel", in)
	out := new(EmptyRequest)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Register(context.Context, *RegisterRequest, *SuccessResponse) error
	RegisterWithMobileNumber(context.Context, *RegisterWithMobileNumberRequest, *SuccessResponse) error
	RegisterWithMobileNumberAis(context.Context, *RegisterWithMobileNumberAisRequest, *SuccessResponse) error
	Profile(context.Context, *EmptyRequest, *UserDetail) error
	VerifyUser(context.Context, *VerifyUserRequest, *VerifyUserResponse) error
	AutoRegisterBySocial(context.Context, *SocialRequest, *SocialResponse) error
	UpdateSocialID(context.Context, *UpdateSocialIDRequest, *UpdateSocialIDRespose) error
	ResetPassword(context.Context, *ResetPasswordRequest, *SuccessResponse) error
	ResetPasswordToken(context.Context, *ResetPasswordTokenRequest, *ResetPasswordTokenResponse) error
	UpdateResetPassword(context.Context, *UpdateResetPasswordRequest, *SuccessResponse) error
	UpdateProfile(context.Context, *UpdateProfileRequest, *SuccessResponse) error
	UploadProfilePicture(context.Context, *UploadProfilePictureRequest, *SuccessResponse) error
	ResendVerifyByOldToken(context.Context, *VerifyUserRequest, *VerifyUserResponse) error
	Resendverify(context.Context, *ResetPasswordRequest, *VerifyUserResponse) error
	CheckTerm(context.Context, *EmptyRequest, *CheckTermResponse) error
	AcceptTerm(context.Context, *EmptyRequest, *SuccessResponse) error
	DeleteAccount(context.Context, *EmptyRequest, *SuccessResponse) error
	TriggerDeleteTemp(context.Context, *EmptyRequest, *SuccessResponse) error
	GetPdpa(context.Context, *EmptyRequest, *PDPAResponse) error
	AcceptPrivacyPolicy(context.Context, *PrivacyPolicyRequest, *SuccessResponse) error
	CheckEmailHasAddress(context.Context, *EmptyRequest, *CheckEmailHasAddressResponse) error
	EventCalendarLimitExceed(context.Context, *EmptyRequest, *EventCalendarLimitExceedResponse) error
	RequestOtp(context.Context, *RequestOtpRequest, *RequestOtpResponse) error
	VerifyOtp(context.Context, *VerifyOtpRequest, *VerifyOtpResponse) error
	RequestOtpOperationTool(context.Context, *RequestOtpRequest, *RequestOtpOperationToolResponse) error
	DeleteUserTemp(context.Context, *DeleteUserTempRequest, *DeleteUserTempResponse) error
	Getbytel(context.Context, *GetByTelRequest, *UserDetailResponse) error
	VerifyPartnerID(context.Context, *VerifyPartnerIDRequest, *VerifyPartnerIDResponse) error
	ResendVerifyEmailMember(context.Context, *ResetPasswordRequest, *SuccessResponse) error
	UpdateContactProfile(context.Context, *UpdateProfileRequest, *SuccessResponse) error
	CreateUserInterest(context.Context, *UserInterestRequest, *UserInterestRequestResponse) error
	CreateUserInterestList(context.Context, *UsersInterestReqList, *UsersInterestResList) error
	UpdateUserInterest(context.Context, *UserInterestUpdateRequest, *UserInterestRequestResponse) error
	UpdateUserInterestList(context.Context, *UsersInterestUpdateReqList, *UsersInterestResList) error
	GetInterestItems(context.Context, *EmptyRequest, *InterestItemList) error
	TestPublishDeleteTel(context.Context, *UserDetail, *EmptyRequest) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Register(ctx context.Context, in *RegisterRequest, out *SuccessResponse) error
		RegisterWithMobileNumber(ctx context.Context, in *RegisterWithMobileNumberRequest, out *SuccessResponse) error
		RegisterWithMobileNumberAis(ctx context.Context, in *RegisterWithMobileNumberAisRequest, out *SuccessResponse) error
		Profile(ctx context.Context, in *EmptyRequest, out *UserDetail) error
		VerifyUser(ctx context.Context, in *VerifyUserRequest, out *VerifyUserResponse) error
		AutoRegisterBySocial(ctx context.Context, in *SocialRequest, out *SocialResponse) error
		UpdateSocialID(ctx context.Context, in *UpdateSocialIDRequest, out *UpdateSocialIDRespose) error
		ResetPassword(ctx context.Context, in *ResetPasswordRequest, out *SuccessResponse) error
		ResetPasswordToken(ctx context.Context, in *ResetPasswordTokenRequest, out *ResetPasswordTokenResponse) error
		UpdateResetPassword(ctx context.Context, in *UpdateResetPasswordRequest, out *SuccessResponse) error
		UpdateProfile(ctx context.Context, in *UpdateProfileRequest, out *SuccessResponse) error
		UploadProfilePicture(ctx context.Context, in *UploadProfilePictureRequest, out *SuccessResponse) error
		ResendVerifyByOldToken(ctx context.Context, in *VerifyUserRequest, out *VerifyUserResponse) error
		Resendverify(ctx context.Context, in *ResetPasswordRequest, out *VerifyUserResponse) error
		CheckTerm(ctx context.Context, in *EmptyRequest, out *CheckTermResponse) error
		AcceptTerm(ctx context.Context, in *EmptyRequest, out *SuccessResponse) error
		DeleteAccount(ctx context.Context, in *EmptyRequest, out *SuccessResponse) error
		TriggerDeleteTemp(ctx context.Context, in *EmptyRequest, out *SuccessResponse) error
		GetPdpa(ctx context.Context, in *EmptyRequest, out *PDPAResponse) error
		AcceptPrivacyPolicy(ctx context.Context, in *PrivacyPolicyRequest, out *SuccessResponse) error
		CheckEmailHasAddress(ctx context.Context, in *EmptyRequest, out *CheckEmailHasAddressResponse) error
		EventCalendarLimitExceed(ctx context.Context, in *EmptyRequest, out *EventCalendarLimitExceedResponse) error
		RequestOtp(ctx context.Context, in *RequestOtpRequest, out *RequestOtpResponse) error
		VerifyOtp(ctx context.Context, in *VerifyOtpRequest, out *VerifyOtpResponse) error
		RequestOtpOperationTool(ctx context.Context, in *RequestOtpRequest, out *RequestOtpOperationToolResponse) error
		DeleteUserTemp(ctx context.Context, in *DeleteUserTempRequest, out *DeleteUserTempResponse) error
		Getbytel(ctx context.Context, in *GetByTelRequest, out *UserDetailResponse) error
		VerifyPartnerID(ctx context.Context, in *VerifyPartnerIDRequest, out *VerifyPartnerIDResponse) error
		ResendVerifyEmailMember(ctx context.Context, in *ResetPasswordRequest, out *SuccessResponse) error
		UpdateContactProfile(ctx context.Context, in *UpdateProfileRequest, out *SuccessResponse) error
		CreateUserInterest(ctx context.Context, in *UserInterestRequest, out *UserInterestRequestResponse) error
		CreateUserInterestList(ctx context.Context, in *UsersInterestReqList, out *UsersInterestResList) error
		UpdateUserInterest(ctx context.Context, in *UserInterestUpdateRequest, out *UserInterestRequestResponse) error
		UpdateUserInterestList(ctx context.Context, in *UsersInterestUpdateReqList, out *UsersInterestResList) error
		GetInterestItems(ctx context.Context, in *EmptyRequest, out *InterestItemList) error
		TestPublishDeleteTel(ctx context.Context, in *UserDetail, out *EmptyRequest) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Register(ctx context.Context, in *RegisterRequest, out *SuccessResponse) error {
	return h.UserHandler.Register(ctx, in, out)
}

func (h *userHandler) RegisterWithMobileNumber(ctx context.Context, in *RegisterWithMobileNumberRequest, out *SuccessResponse) error {
	return h.UserHandler.RegisterWithMobileNumber(ctx, in, out)
}

func (h *userHandler) RegisterWithMobileNumberAis(ctx context.Context, in *RegisterWithMobileNumberAisRequest, out *SuccessResponse) error {
	return h.UserHandler.RegisterWithMobileNumberAis(ctx, in, out)
}

func (h *userHandler) Profile(ctx context.Context, in *EmptyRequest, out *UserDetail) error {
	return h.UserHandler.Profile(ctx, in, out)
}

func (h *userHandler) VerifyUser(ctx context.Context, in *VerifyUserRequest, out *VerifyUserResponse) error {
	return h.UserHandler.VerifyUser(ctx, in, out)
}

func (h *userHandler) AutoRegisterBySocial(ctx context.Context, in *SocialRequest, out *SocialResponse) error {
	return h.UserHandler.AutoRegisterBySocial(ctx, in, out)
}

func (h *userHandler) UpdateSocialID(ctx context.Context, in *UpdateSocialIDRequest, out *UpdateSocialIDRespose) error {
	return h.UserHandler.UpdateSocialID(ctx, in, out)
}

func (h *userHandler) ResetPassword(ctx context.Context, in *ResetPasswordRequest, out *SuccessResponse) error {
	return h.UserHandler.ResetPassword(ctx, in, out)
}

func (h *userHandler) ResetPasswordToken(ctx context.Context, in *ResetPasswordTokenRequest, out *ResetPasswordTokenResponse) error {
	return h.UserHandler.ResetPasswordToken(ctx, in, out)
}

func (h *userHandler) UpdateResetPassword(ctx context.Context, in *UpdateResetPasswordRequest, out *SuccessResponse) error {
	return h.UserHandler.UpdateResetPassword(ctx, in, out)
}

func (h *userHandler) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, out *SuccessResponse) error {
	return h.UserHandler.UpdateProfile(ctx, in, out)
}

func (h *userHandler) UploadProfilePicture(ctx context.Context, in *UploadProfilePictureRequest, out *SuccessResponse) error {
	return h.UserHandler.UploadProfilePicture(ctx, in, out)
}

func (h *userHandler) ResendVerifyByOldToken(ctx context.Context, in *VerifyUserRequest, out *VerifyUserResponse) error {
	return h.UserHandler.ResendVerifyByOldToken(ctx, in, out)
}

func (h *userHandler) Resendverify(ctx context.Context, in *ResetPasswordRequest, out *VerifyUserResponse) error {
	return h.UserHandler.Resendverify(ctx, in, out)
}

func (h *userHandler) CheckTerm(ctx context.Context, in *EmptyRequest, out *CheckTermResponse) error {
	return h.UserHandler.CheckTerm(ctx, in, out)
}

func (h *userHandler) AcceptTerm(ctx context.Context, in *EmptyRequest, out *SuccessResponse) error {
	return h.UserHandler.AcceptTerm(ctx, in, out)
}

func (h *userHandler) DeleteAccount(ctx context.Context, in *EmptyRequest, out *SuccessResponse) error {
	return h.UserHandler.DeleteAccount(ctx, in, out)
}

func (h *userHandler) TriggerDeleteTemp(ctx context.Context, in *EmptyRequest, out *SuccessResponse) error {
	return h.UserHandler.TriggerDeleteTemp(ctx, in, out)
}

func (h *userHandler) GetPdpa(ctx context.Context, in *EmptyRequest, out *PDPAResponse) error {
	return h.UserHandler.GetPdpa(ctx, in, out)
}

func (h *userHandler) AcceptPrivacyPolicy(ctx context.Context, in *PrivacyPolicyRequest, out *SuccessResponse) error {
	return h.UserHandler.AcceptPrivacyPolicy(ctx, in, out)
}

func (h *userHandler) CheckEmailHasAddress(ctx context.Context, in *EmptyRequest, out *CheckEmailHasAddressResponse) error {
	return h.UserHandler.CheckEmailHasAddress(ctx, in, out)
}

func (h *userHandler) EventCalendarLimitExceed(ctx context.Context, in *EmptyRequest, out *EventCalendarLimitExceedResponse) error {
	return h.UserHandler.EventCalendarLimitExceed(ctx, in, out)
}

func (h *userHandler) RequestOtp(ctx context.Context, in *RequestOtpRequest, out *RequestOtpResponse) error {
	return h.UserHandler.RequestOtp(ctx, in, out)
}

func (h *userHandler) VerifyOtp(ctx context.Context, in *VerifyOtpRequest, out *VerifyOtpResponse) error {
	return h.UserHandler.VerifyOtp(ctx, in, out)
}

func (h *userHandler) RequestOtpOperationTool(ctx context.Context, in *RequestOtpRequest, out *RequestOtpOperationToolResponse) error {
	return h.UserHandler.RequestOtpOperationTool(ctx, in, out)
}

func (h *userHandler) DeleteUserTemp(ctx context.Context, in *DeleteUserTempRequest, out *DeleteUserTempResponse) error {
	return h.UserHandler.DeleteUserTemp(ctx, in, out)
}

func (h *userHandler) Getbytel(ctx context.Context, in *GetByTelRequest, out *UserDetailResponse) error {
	return h.UserHandler.Getbytel(ctx, in, out)
}

func (h *userHandler) VerifyPartnerID(ctx context.Context, in *VerifyPartnerIDRequest, out *VerifyPartnerIDResponse) error {
	return h.UserHandler.VerifyPartnerID(ctx, in, out)
}

func (h *userHandler) ResendVerifyEmailMember(ctx context.Context, in *ResetPasswordRequest, out *SuccessResponse) error {
	return h.UserHandler.ResendVerifyEmailMember(ctx, in, out)
}

func (h *userHandler) UpdateContactProfile(ctx context.Context, in *UpdateProfileRequest, out *SuccessResponse) error {
	return h.UserHandler.UpdateContactProfile(ctx, in, out)
}

func (h *userHandler) CreateUserInterest(ctx context.Context, in *UserInterestRequest, out *UserInterestRequestResponse) error {
	return h.UserHandler.CreateUserInterest(ctx, in, out)
}

func (h *userHandler) CreateUserInterestList(ctx context.Context, in *UsersInterestReqList, out *UsersInterestResList) error {
	return h.UserHandler.CreateUserInterestList(ctx, in, out)
}

func (h *userHandler) UpdateUserInterest(ctx context.Context, in *UserInterestUpdateRequest, out *UserInterestRequestResponse) error {
	return h.UserHandler.UpdateUserInterest(ctx, in, out)
}

func (h *userHandler) UpdateUserInterestList(ctx context.Context, in *UsersInterestUpdateReqList, out *UsersInterestResList) error {
	return h.UserHandler.UpdateUserInterestList(ctx, in, out)
}

func (h *userHandler) GetInterestItems(ctx context.Context, in *EmptyRequest, out *InterestItemList) error {
	return h.UserHandler.GetInterestItems(ctx, in, out)
}

func (h *userHandler) TestPublishDeleteTel(ctx context.Context, in *UserDetail, out *EmptyRequest) error {
	return h.UserHandler.TestPublishDeleteTel(ctx, in, out)
}

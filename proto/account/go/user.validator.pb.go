// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/account/user.proto

package _go

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *User) Validate() error {
	return nil
}

var _regex_RegisterUserRequest_Email = regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
var _regex_RegisterUserRequest_Password = regexp.MustCompile(`^[ -~]{8,32}$`)

func (this *RegisterUserRequest) Validate() error {
	if !_regex_RegisterUserRequest_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"`, this.Email))
	}
	if this.Email == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must not be an empty string`, this.Email))
	}
	if !_regex_RegisterUserRequest_Password.MatchString(this.Password) {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must be a string conforming to regex "^[ -~]{8,32}$"`, this.Password))
	}
	if this.Password == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if !(len(this.Name) < 255) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '255'`, this.Name))
	}
	if !(this.CompanyId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("CompanyId", fmt.Errorf(`value '%v' must be greater than '0'`, this.CompanyId))
	}
	return nil
}
func (this *RegisterUserResponse) Validate() error {
	return nil
}

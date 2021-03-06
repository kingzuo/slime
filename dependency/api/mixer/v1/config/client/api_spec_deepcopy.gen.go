// Code generated by protoc-gen-gogo. DO NOT EDIT.
// mixer/v1/config/client/api_spec.proto is a deprecated file.

package client

import (
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/api/mixer/v1"
	_ "istio.io/gogo-genproto/googleapis/google/api"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using HTTPAPISpec within kubernetes types, where deepcopy-gen is used.
func (in *HTTPAPISpec) DeepCopyInto(out *HTTPAPISpec) {
	p := proto.Clone(in).(*HTTPAPISpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPAPISpec. Required by controller-gen.
func (in *HTTPAPISpec) DeepCopy() *HTTPAPISpec {
	if in == nil {
		return nil
	}
	out := new(HTTPAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using HTTPAPISpecPattern within kubernetes types, where deepcopy-gen is used.
func (in *HTTPAPISpecPattern) DeepCopyInto(out *HTTPAPISpecPattern) {
	p := proto.Clone(in).(*HTTPAPISpecPattern)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPAPISpecPattern. Required by controller-gen.
func (in *HTTPAPISpecPattern) DeepCopy() *HTTPAPISpecPattern {
	if in == nil {
		return nil
	}
	out := new(HTTPAPISpecPattern)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using APIKey within kubernetes types, where deepcopy-gen is used.
func (in *APIKey) DeepCopyInto(out *APIKey) {
	p := proto.Clone(in).(*APIKey)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKey. Required by controller-gen.
func (in *APIKey) DeepCopy() *APIKey {
	if in == nil {
		return nil
	}
	out := new(APIKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using HTTPAPISpecReference within kubernetes types, where deepcopy-gen is used.
func (in *HTTPAPISpecReference) DeepCopyInto(out *HTTPAPISpecReference) {
	p := proto.Clone(in).(*HTTPAPISpecReference)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPAPISpecReference. Required by controller-gen.
func (in *HTTPAPISpecReference) DeepCopy() *HTTPAPISpecReference {
	if in == nil {
		return nil
	}
	out := new(HTTPAPISpecReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using HTTPAPISpecBinding within kubernetes types, where deepcopy-gen is used.
func (in *HTTPAPISpecBinding) DeepCopyInto(out *HTTPAPISpecBinding) {
	p := proto.Clone(in).(*HTTPAPISpecBinding)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPAPISpecBinding. Required by controller-gen.
func (in *HTTPAPISpecBinding) DeepCopy() *HTTPAPISpecBinding {
	if in == nil {
		return nil
	}
	out := new(HTTPAPISpecBinding)
	in.DeepCopyInto(out)
	return out
}

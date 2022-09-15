/*
Copyright 2022 The kubeall.com Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/kube-all/public-cloud/api/tencent/v1alpha1/generated.proto

package v1alpha1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *TencentCluster) Reset()      { *m = TencentCluster{} }
func (*TencentCluster) ProtoMessage() {}
func (*TencentCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_89190ef5971dcb84, []int{0}
}
func (m *TencentCluster) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TencentCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TencentCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TencentCluster.Merge(m, src)
}
func (m *TencentCluster) XXX_Size() int {
	return m.Size()
}
func (m *TencentCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_TencentCluster.DiscardUnknown(m)
}

var xxx_messageInfo_TencentCluster proto.InternalMessageInfo

func (m *TencentClusterList) Reset()      { *m = TencentClusterList{} }
func (*TencentClusterList) ProtoMessage() {}
func (*TencentClusterList) Descriptor() ([]byte, []int) {
	return fileDescriptor_89190ef5971dcb84, []int{1}
}
func (m *TencentClusterList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TencentClusterList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TencentClusterList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TencentClusterList.Merge(m, src)
}
func (m *TencentClusterList) XXX_Size() int {
	return m.Size()
}
func (m *TencentClusterList) XXX_DiscardUnknown() {
	xxx_messageInfo_TencentClusterList.DiscardUnknown(m)
}

var xxx_messageInfo_TencentClusterList proto.InternalMessageInfo

func (m *TencentClusterSpec) Reset()      { *m = TencentClusterSpec{} }
func (*TencentClusterSpec) ProtoMessage() {}
func (*TencentClusterSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_89190ef5971dcb84, []int{2}
}
func (m *TencentClusterSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TencentClusterSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TencentClusterSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TencentClusterSpec.Merge(m, src)
}
func (m *TencentClusterSpec) XXX_Size() int {
	return m.Size()
}
func (m *TencentClusterSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TencentClusterSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TencentClusterSpec proto.InternalMessageInfo

func (m *TencentClusterStatus) Reset()      { *m = TencentClusterStatus{} }
func (*TencentClusterStatus) ProtoMessage() {}
func (*TencentClusterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_89190ef5971dcb84, []int{3}
}
func (m *TencentClusterStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TencentClusterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TencentClusterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TencentClusterStatus.Merge(m, src)
}
func (m *TencentClusterStatus) XXX_Size() int {
	return m.Size()
}
func (m *TencentClusterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TencentClusterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TencentClusterStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TencentCluster)(nil), "github.com.kube_all.public_cloud.api.tencent.v1alpha1.TencentCluster")
	proto.RegisterType((*TencentClusterList)(nil), "github.com.kube_all.public_cloud.api.tencent.v1alpha1.TencentClusterList")
	proto.RegisterType((*TencentClusterSpec)(nil), "github.com.kube_all.public_cloud.api.tencent.v1alpha1.TencentClusterSpec")
	proto.RegisterType((*TencentClusterStatus)(nil), "github.com.kube_all.public_cloud.api.tencent.v1alpha1.TencentClusterStatus")
}

func init() {
	proto.RegisterFile("github.com/kube-all/public-cloud/api/tencent/v1alpha1/generated.proto", fileDescriptor_89190ef5971dcb84)
}

var fileDescriptor_89190ef5971dcb84 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x33, 0xfd, 0x89, 0x1a, 0xa7, 0x54, 0x60, 0x58, 0x44, 0x59, 0x4c, 0xab, 0x2c, 0x50,
	0x85, 0x14, 0xbb, 0xa9, 0x68, 0xc5, 0x7a, 0x4a, 0x16, 0x11, 0x45, 0x48, 0x43, 0x56, 0x80, 0x08,
	0x8e, 0xe3, 0x4e, 0xdc, 0xcc, 0x8f, 0x35, 0xf6, 0x44, 0x64, 0xc7, 0x23, 0xf0, 0x58, 0x59, 0x66,
	0xc1, 0xa2, 0xab, 0x88, 0x0c, 0x2f, 0x82, 0xec, 0x71, 0x9a, 0xb4, 0x05, 0xa9, 0xaa, 0xba, 0xcb,
	0xb9, 0xb9, 0xf7, 0x3b, 0xd7, 0xc7, 0xd6, 0x80, 0x76, 0xc0, 0xd5, 0x30, 0xeb, 0x23, 0x9a, 0x44,
	0x78, 0x94, 0xf5, 0x59, 0x93, 0x84, 0x21, 0x16, 0x59, 0x3f, 0xe4, 0xb4, 0x49, 0xc3, 0x24, 0x1b,
	0x60, 0x22, 0x38, 0x56, 0x2c, 0xa6, 0x2c, 0x56, 0x78, 0xdc, 0x22, 0xa1, 0x18, 0x92, 0x16, 0x0e,
	0x58, 0xcc, 0x52, 0xa2, 0xd8, 0x00, 0x89, 0x34, 0x51, 0x09, 0x3c, 0x59, 0x61, 0x90, 0xc6, 0xf4,
	0x48, 0x18, 0xa2, 0x02, 0xd3, 0x33, 0x18, 0x44, 0x04, 0x47, 0x16, 0x83, 0x96, 0x98, 0x7a, 0x73,
	0xcd, 0x3d, 0x48, 0x82, 0x04, 0x1b, 0x5a, 0x3f, 0xbb, 0x30, 0xca, 0x08, 0xf3, 0xab, 0x70, 0xa9,
	0xbf, 0x1e, 0xbd, 0x91, 0x88, 0x27, 0x7a, 0xa5, 0x88, 0xd0, 0x21, 0x8f, 0x59, 0x3a, 0xc1, 0x62,
	0x14, 0xe8, 0x82, 0xc4, 0x11, 0x53, 0x04, 0x8f, 0xef, 0xec, 0x56, 0x3f, 0xfd, 0xdf, 0x54, 0x9a,
	0xc5, 0x8a, 0x47, 0x0c, 0x4b, 0x3a, 0x64, 0x11, 0xb9, 0x3d, 0xd7, 0xf8, 0xb5, 0x01, 0xf6, 0xba,
	0xc5, 0xc6, 0x67, 0x61, 0x26, 0x15, 0x4b, 0xe1, 0x37, 0xb0, 0xa3, 0x5d, 0x06, 0x44, 0x91, 0x9a,
	0x73, 0xe0, 0x1c, 0x56, 0x8f, 0x8f, 0x50, 0x41, 0x47, 0xeb, 0x74, 0x24, 0x46, 0x81, 0x2e, 0x48,
	0xa4, 0xbb, 0xd1, 0xb8, 0x85, 0x3e, 0xf4, 0x2f, 0x19, 0x55, 0xef, 0x99, 0x22, 0x1e, 0x9c, 0xce,
	0xf7, 0x4b, 0xf9, 0x7c, 0x1f, 0xac, 0x6a, 0xfe, 0x35, 0x15, 0x8e, 0xc0, 0x96, 0x14, 0x8c, 0xd6,
	0x36, 0x0c, 0xbd, 0x83, 0x1e, 0x94, 0x2b, 0xba, 0xb9, 0xf6, 0x47, 0xc1, 0xa8, 0xb7, 0x6b, 0x6d,
	0xb7, 0xb4, 0xf2, 0x8d, 0x09, 0x94, 0xa0, 0x2c, 0x15, 0x51, 0x99, 0xac, 0x6d, 0x1a, 0xbb, 0x77,
	0x8f, 0x63, 0x67, 0x90, 0xde, 0x9e, 0x35, 0x2c, 0x17, 0xda, 0xb7, 0x56, 0x8d, 0xb9, 0x03, 0xe0,
	0xcd, 0x81, 0x73, 0x2e, 0x15, 0xfc, 0x72, 0x27, 0x5a, 0x74, 0xbf, 0x68, 0xf5, 0xb4, 0x09, 0xf6,
	0xa9, 0x35, 0xdc, 0x59, 0x56, 0xd6, 0x62, 0xbd, 0x04, 0xdb, 0x5c, 0xb1, 0x48, 0xd6, 0x36, 0x0e,
	0x36, 0x0f, 0xab, 0xc7, 0xed, 0x47, 0x39, 0xa8, 0xf7, 0xc4, 0x3a, 0x6e, 0x77, 0x34, 0xdb, 0x2f,
	0x2c, 0x1a, 0xe7, 0xb7, 0xcf, 0xa7, 0x13, 0x87, 0xa7, 0x60, 0x97, 0x16, 0xb2, 0xa7, 0x26, 0x82,
	0x99, 0x33, 0x56, 0xbc, 0xe7, 0x96, 0x50, 0xb5, 0xad, 0xdd, 0x89, 0x60, 0x7e, 0x95, 0xae, 0x44,
	0x63, 0xe6, 0x80, 0x17, 0xff, 0xca, 0x17, 0x1e, 0x01, 0xb0, 0x04, 0xf2, 0x81, 0xc5, 0x3d, 0xb3,
	0xb8, 0x8a, 0x6d, 0xed, 0xbc, 0xf5, 0x2b, 0xb6, 0xa9, 0x33, 0x80, 0x2f, 0x41, 0x99, 0x26, 0xf1,
	0x05, 0x0f, 0xcc, 0xeb, 0xaa, 0xac, 0x6e, 0xe8, 0xcc, 0x54, 0x7d, 0xfb, 0x2f, 0xfc, 0x0a, 0x00,
	0xfb, 0x2e, 0x78, 0x4a, 0x14, 0x4f, 0x62, 0xfb, 0x34, 0x5e, 0xdd, 0xef, 0x32, 0xba, 0x3c, 0x62,
	0xab, 0x17, 0xde, 0xbe, 0xa6, 0xf8, 0x6b, 0x44, 0xef, 0xf3, 0x74, 0xe1, 0x96, 0x66, 0x0b, 0xb7,
	0x74, 0xb5, 0x70, 0x4b, 0x3f, 0x72, 0xd7, 0x99, 0xe6, 0xae, 0x33, 0xcb, 0x5d, 0xe7, 0x2a, 0x77,
	0x9d, 0xdf, 0xb9, 0xeb, 0xfc, 0xfc, 0xe3, 0x96, 0x3e, 0x9d, 0x3c, 0xe8, 0xcb, 0xf4, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x73, 0x24, 0x5f, 0x06, 0xd1, 0x04, 0x00, 0x00,
}

func (m *TencentCluster) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TencentCluster) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TencentCluster) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TencentClusterList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TencentClusterList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TencentClusterList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TencentClusterSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TencentClusterSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TencentClusterSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.ClusterType)
	copy(dAtA[i:], m.ClusterType)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ClusterType)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TencentClusterStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TencentClusterStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TencentClusterStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Expiration.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	i -= len(m.Config)
	copy(dAtA[i:], m.Config)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Config)))
	i--
	dAtA[i] = 0x12
	i -= len(m.ClusterID)
	copy(dAtA[i:], m.ClusterID)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ClusterID)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TencentCluster) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *TencentClusterList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *TencentClusterSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterType)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *TencentClusterStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterID)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Config)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Expiration.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TencentCluster) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TencentCluster{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "TencentClusterSpec", "TencentClusterSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "TencentClusterStatus", "TencentClusterStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TencentClusterList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]TencentCluster{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "TencentCluster", "TencentCluster", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&TencentClusterList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *TencentClusterSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TencentClusterSpec{`,
		`ClusterType:` + fmt.Sprintf("%v", this.ClusterType) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TencentClusterStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TencentClusterStatus{`,
		`ClusterID:` + fmt.Sprintf("%v", this.ClusterID) + `,`,
		`Config:` + fmt.Sprintf("%v", this.Config) + `,`,
		`Expiration:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Expiration), "Time", "v1.Time", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TencentCluster) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TencentCluster: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TencentCluster: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TencentClusterList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TencentClusterList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TencentClusterList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, TencentCluster{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TencentClusterSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TencentClusterSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TencentClusterSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TencentClusterStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TencentClusterStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TencentClusterStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Config = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Expiration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
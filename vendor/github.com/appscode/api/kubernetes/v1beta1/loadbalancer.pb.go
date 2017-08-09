// Code generated by protoc-gen-go. DO NOT EDIT.
// source: loadbalancer.proto

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ListRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ListRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *ListRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ListResponse struct {
	LoadBalancers []*LoadBalancer `protobuf:"bytes,1,rep,name=load_balancers,json=loadBalancers" json:"load_balancers,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ListResponse) GetLoadBalancers() []*LoadBalancer {
	if m != nil {
		return m.LoadBalancers
	}
	return nil
}

type DescribeRequest struct {
	Kind      string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Cluster   string `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
	Raw       string `protobuf:"bytes,5,opt,name=raw" json:"raw,omitempty"`
}

func (m *DescribeRequest) Reset()                    { *m = DescribeRequest{} }
func (m *DescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRequest) ProtoMessage()               {}
func (*DescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *DescribeRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *DescribeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DescribeRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DescribeRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *DescribeRequest) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

type DescribeResponse struct {
	LoadBalancer *LoadBalancer `protobuf:"bytes,1,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
	Raw          *Raw          `protobuf:"bytes,2,opt,name=raw" json:"raw,omitempty"`
}

func (m *DescribeResponse) Reset()                    { *m = DescribeResponse{} }
func (m *DescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeResponse) ProtoMessage()               {}
func (*DescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *DescribeResponse) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

func (m *DescribeResponse) GetRaw() *Raw {
	if m != nil {
		return m.Raw
	}
	return nil
}

type CreateRequest struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Namespace    string        `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Cluster      string        `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
	LoadBalancer *LoadBalancer `protobuf:"bytes,4,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
	Raw          *Raw          `protobuf:"bytes,5,opt,name=raw" json:"raw,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *CreateRequest) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

func (m *CreateRequest) GetRaw() *Raw {
	if m != nil {
		return m.Raw
	}
	return nil
}

type UpdateRequest struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Cluster      string        `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	LoadBalancer *LoadBalancer `protobuf:"bytes,3,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
	Raw          *Raw          `protobuf:"bytes,5,opt,name=raw" json:"raw,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *UpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *UpdateRequest) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

func (m *UpdateRequest) GetRaw() *Raw {
	if m != nil {
		return m.Raw
	}
	return nil
}

type DeleteRequest struct {
	Kind      string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Cluster   string `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *DeleteRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

type Raw struct {
	Format string `protobuf:"bytes,1,opt,name=format" json:"format,omitempty"`
	Data   string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *Raw) Reset()                    { *m = Raw{} }
func (m *Raw) String() string            { return proto.CompactTextString(m) }
func (*Raw) ProtoMessage()               {}
func (*Raw) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *Raw) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *Raw) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type LoadBalancer struct {
	// 'kind' defines is it the regular kubernetes instance or the
	// appscode superset called Extended Ingress. This field will
	// strictly contains only those two values
	// 'ingress' - default kubernetes ingress object.
	// 'extendedIngress' - appscode superset of ingress.
	// when creating a Loadbalancer from UI this field will always
	// be only 'extendedIngress.' List, Describe, Update and Delete
	// will support both two modes.
	// Create will support only extendedIngress.
	// For Creating or Updating an regular ingress one must use the
	// kubectl or direct API calls directly to kubernetes.
	Kind              string            `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name              string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace         string            `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	CreationTimestamp int64             `protobuf:"varint,4,opt,name=creation_timestamp,json=creationTimestamp" json:"creation_timestamp,omitempty"`
	Options           map[string]string `protobuf:"bytes,5,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Spec              *Spec             `protobuf:"bytes,6,opt,name=spec" json:"spec,omitempty"`
	Status            *Status           `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
}

func (m *LoadBalancer) Reset()                    { *m = LoadBalancer{} }
func (m *LoadBalancer) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancer) ProtoMessage()               {}
func (*LoadBalancer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *LoadBalancer) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *LoadBalancer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoadBalancer) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *LoadBalancer) GetCreationTimestamp() int64 {
	if m != nil {
		return m.CreationTimestamp
	}
	return 0
}

func (m *LoadBalancer) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *LoadBalancer) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *LoadBalancer) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type Spec struct {
	Backend *HTTPLoadBalancerRule `protobuf:"bytes,1,opt,name=backend" json:"backend,omitempty"`
	Rules   []*LoadBalancerRule   `protobuf:"bytes,2,rep,name=rules" json:"rules,omitempty"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *Spec) GetBackend() *HTTPLoadBalancerRule {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *Spec) GetRules() []*LoadBalancerRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Status struct {
	Status []*LoadBalancerStatus `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *Status) GetStatus() []*LoadBalancerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type LoadBalancerStatus struct {
	IP   string `protobuf:"bytes,1,opt,name=IP" json:"IP,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *LoadBalancerStatus) Reset()                    { *m = LoadBalancerStatus{} }
func (m *LoadBalancerStatus) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerStatus) ProtoMessage()               {}
func (*LoadBalancerStatus) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *LoadBalancerStatus) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *LoadBalancerStatus) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type LoadBalancerBackend struct {
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	ServicePort string `protobuf:"bytes,2,opt,name=service_port,json=servicePort" json:"service_port,omitempty"`
}

func (m *LoadBalancerBackend) Reset()                    { *m = LoadBalancerBackend{} }
func (m *LoadBalancerBackend) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerBackend) ProtoMessage()               {}
func (*LoadBalancerBackend) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *LoadBalancerBackend) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *LoadBalancerBackend) GetServicePort() string {
	if m != nil {
		return m.ServicePort
	}
	return ""
}

type LoadBalancerRule struct {
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	// ssl secret name to enable https on the host.
	// ssl secret must contain data with the certs pem file.
	SSLSecretName string                  `protobuf:"bytes,5,opt,name=SSL_secret_name,json=SSLSecretName" json:"SSL_secret_name,omitempty"`
	Http          []*HTTPLoadBalancerRule `protobuf:"bytes,2,rep,name=http" json:"http,omitempty"`
	Tcp           []*TCPLoadBalancerRule  `protobuf:"bytes,3,rep,name=tcp" json:"tcp,omitempty"`
}

func (m *LoadBalancerRule) Reset()                    { *m = LoadBalancerRule{} }
func (m *LoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerRule) ProtoMessage()               {}
func (*LoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func (m *LoadBalancerRule) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *LoadBalancerRule) GetSSLSecretName() string {
	if m != nil {
		return m.SSLSecretName
	}
	return ""
}

func (m *LoadBalancerRule) GetHttp() []*HTTPLoadBalancerRule {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *LoadBalancerRule) GetTcp() []*TCPLoadBalancerRule {
	if m != nil {
		return m.Tcp
	}
	return nil
}

type HTTPLoadBalancerRule struct {
	Path         string               `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Backend      *LoadBalancerBackend `protobuf:"bytes,2,opt,name=backend" json:"backend,omitempty"`
	HeaderRules  []string             `protobuf:"bytes,3,rep,name=header_rules,json=headerRules" json:"header_rules,omitempty"`
	RewriteRules []string             `protobuf:"bytes,4,rep,name=rewrite_rules,json=rewriteRules" json:"rewrite_rules,omitempty"`
}

func (m *HTTPLoadBalancerRule) Reset()                    { *m = HTTPLoadBalancerRule{} }
func (m *HTTPLoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*HTTPLoadBalancerRule) ProtoMessage()               {}
func (*HTTPLoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

func (m *HTTPLoadBalancerRule) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HTTPLoadBalancerRule) GetBackend() *LoadBalancerBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *HTTPLoadBalancerRule) GetHeaderRules() []string {
	if m != nil {
		return m.HeaderRules
	}
	return nil
}

func (m *HTTPLoadBalancerRule) GetRewriteRules() []string {
	if m != nil {
		return m.RewriteRules
	}
	return nil
}

type TCPLoadBalancerRule struct {
	Port          string               `protobuf:"bytes,1,opt,name=port" json:"port,omitempty"`
	Backend       *LoadBalancerBackend `protobuf:"bytes,2,opt,name=backend" json:"backend,omitempty"`
	SSLSecretName string               `protobuf:"bytes,3,opt,name=SSL_secret_name,json=SSLSecretName" json:"SSL_secret_name,omitempty"`
	SecretPemName string               `protobuf:"bytes,4,opt,name=secret_pem_name,json=secretPemName" json:"secret_pem_name,omitempty"`
}

func (m *TCPLoadBalancerRule) Reset()                    { *m = TCPLoadBalancerRule{} }
func (m *TCPLoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*TCPLoadBalancerRule) ProtoMessage()               {}
func (*TCPLoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{15} }

func (m *TCPLoadBalancerRule) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *TCPLoadBalancerRule) GetBackend() *LoadBalancerBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *TCPLoadBalancerRule) GetSSLSecretName() string {
	if m != nil {
		return m.SSLSecretName
	}
	return ""
}

func (m *TCPLoadBalancerRule) GetSecretPemName() string {
	if m != nil {
		return m.SecretPemName
	}
	return ""
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "appscode.kubernetes.v1beta1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "appscode.kubernetes.v1beta1.ListResponse")
	proto.RegisterType((*DescribeRequest)(nil), "appscode.kubernetes.v1beta1.DescribeRequest")
	proto.RegisterType((*DescribeResponse)(nil), "appscode.kubernetes.v1beta1.DescribeResponse")
	proto.RegisterType((*CreateRequest)(nil), "appscode.kubernetes.v1beta1.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "appscode.kubernetes.v1beta1.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "appscode.kubernetes.v1beta1.DeleteRequest")
	proto.RegisterType((*Raw)(nil), "appscode.kubernetes.v1beta1.Raw")
	proto.RegisterType((*LoadBalancer)(nil), "appscode.kubernetes.v1beta1.LoadBalancer")
	proto.RegisterType((*Spec)(nil), "appscode.kubernetes.v1beta1.Spec")
	proto.RegisterType((*Status)(nil), "appscode.kubernetes.v1beta1.Status")
	proto.RegisterType((*LoadBalancerStatus)(nil), "appscode.kubernetes.v1beta1.LoadBalancerStatus")
	proto.RegisterType((*LoadBalancerBackend)(nil), "appscode.kubernetes.v1beta1.LoadBalancerBackend")
	proto.RegisterType((*LoadBalancerRule)(nil), "appscode.kubernetes.v1beta1.LoadBalancerRule")
	proto.RegisterType((*HTTPLoadBalancerRule)(nil), "appscode.kubernetes.v1beta1.HTTPLoadBalancerRule")
	proto.RegisterType((*TCPLoadBalancerRule)(nil), "appscode.kubernetes.v1beta1.TCPLoadBalancerRule")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LoadBalancers service

type LoadBalancersClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type loadBalancersClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancersClient(cc *grpc.ClientConn) LoadBalancersClient {
	return &loadBalancersClient{cc}
}

func (c *loadBalancersClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.LoadBalancers/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error) {
	out := new(DescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.LoadBalancers/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.LoadBalancers/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.LoadBalancers/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.LoadBalancers/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoadBalancers service

type LoadBalancersServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Describe(context.Context, *DescribeRequest) (*DescribeResponse, error)
	Create(context.Context, *CreateRequest) (*appscode_dtypes.VoidResponse, error)
	Update(context.Context, *UpdateRequest) (*appscode_dtypes.VoidResponse, error)
	Delete(context.Context, *DeleteRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterLoadBalancersServer(s *grpc.Server, srv LoadBalancersServer) {
	s.RegisterService(&_LoadBalancers_serviceDesc, srv)
}

func _LoadBalancers_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.LoadBalancers/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.LoadBalancers/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Describe(ctx, req.(*DescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.LoadBalancers/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.LoadBalancers/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.LoadBalancers/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadBalancers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.kubernetes.v1beta1.LoadBalancers",
	HandlerType: (*LoadBalancersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _LoadBalancers_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _LoadBalancers_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LoadBalancers_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LoadBalancers_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LoadBalancers_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loadbalancer.proto",
}

func init() { proto.RegisterFile("loadbalancer.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1055 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xd7, 0xec, 0x3a, 0x76, 0xf3, 0x62, 0x37, 0xe9, 0xb4, 0x42, 0x2b, 0x53, 0x20, 0xd9, 0x8a,
	0x2a, 0x8d, 0x88, 0xb7, 0x09, 0x02, 0x55, 0xe1, 0x96, 0x34, 0x2a, 0x01, 0x2b, 0x2c, 0xeb, 0x14,
	0x21, 0x40, 0x32, 0xe3, 0xf5, 0xd0, 0x2c, 0xb1, 0x77, 0x96, 0x9d, 0x71, 0xa2, 0xa8, 0xea, 0xa5,
	0x88, 0x4f, 0x80, 0x84, 0xf8, 0x06, 0x9c, 0x38, 0x80, 0xe0, 0x8a, 0xc4, 0xa1, 0x9f, 0x80, 0x13,
	0x27, 0x2e, 0x7c, 0x10, 0x34, 0x7f, 0x36, 0xde, 0x4d, 0xd2, 0xb5, 0x8b, 0xac, 0x5e, 0xac, 0xb7,
	0x6f, 0xde, 0x9f, 0xdf, 0xef, 0xbd, 0xe7, 0xd9, 0xb7, 0x80, 0x07, 0x8c, 0xf4, 0x7b, 0x64, 0x40,
	0xe2, 0x90, 0xa6, 0xad, 0x24, 0x65, 0x82, 0xe1, 0x57, 0x49, 0x92, 0xf0, 0x90, 0xf5, 0x69, 0xeb,
	0x68, 0xd4, 0xa3, 0x69, 0x4c, 0x05, 0xe5, 0xad, 0xe3, 0x8d, 0x1e, 0x15, 0x64, 0xa3, 0x79, 0xf3,
	0x11, 0x63, 0x8f, 0x06, 0xd4, 0x23, 0x49, 0xe4, 0x91, 0x38, 0x66, 0x82, 0x88, 0x88, 0xc5, 0x5c,
	0xbb, 0x36, 0x5f, 0xcf, 0x5c, 0x9f, 0x73, 0xfe, 0x46, 0xe1, 0xbc, 0x2f, 0x4e, 0x13, 0xca, 0x3d,
	0xf5, 0xab, 0x0d, 0xdc, 0x5d, 0x58, 0x68, 0x47, 0x5c, 0x04, 0xf4, 0x9b, 0x11, 0xe5, 0x02, 0x3b,
	0x50, 0x0b, 0x07, 0x23, 0x2e, 0x68, 0xea, 0xa0, 0x65, 0xb4, 0x3a, 0x1f, 0x64, 0x8f, 0xf8, 0x26,
	0xcc, 0xc7, 0x64, 0x48, 0x79, 0x42, 0x42, 0xea, 0x58, 0xea, 0x6c, 0xac, 0x70, 0xbf, 0x84, 0xba,
	0x0e, 0xc3, 0x13, 0x16, 0x73, 0x8a, 0x7d, 0xb8, 0x2a, 0x89, 0x76, 0x33, 0xa6, 0xdc, 0x41, 0xcb,
	0xf6, 0xea, 0xc2, 0xe6, 0x9d, 0x56, 0x09, 0xd7, 0x56, 0x9b, 0x91, 0xfe, 0xb6, 0xf1, 0x08, 0x1a,
	0x83, 0xdc, 0x13, 0x77, 0xbf, 0x45, 0xb0, 0x78, 0x9f, 0xf2, 0x30, 0x8d, 0x7a, 0x34, 0x43, 0x8b,
	0xa1, 0x72, 0x14, 0xc5, 0x7d, 0x03, 0x55, 0xc9, 0x52, 0x27, 0x61, 0x19, 0x88, 0x4a, 0x2e, 0x62,
	0xb7, 0xcf, 0x61, 0xcf, 0x73, 0xae, 0x14, 0x39, 0x2f, 0x81, 0x9d, 0x92, 0x13, 0x67, 0x4e, 0x69,
	0xa5, 0xe8, 0xfe, 0x80, 0x60, 0x69, 0x8c, 0xc2, 0x90, 0xdd, 0x87, 0x46, 0x81, 0xac, 0xc2, 0xf3,
	0x42, 0x5c, 0xeb, 0x79, 0xae, 0x78, 0x53, 0xa7, 0xb5, 0x54, 0x94, 0xe5, 0xd2, 0x28, 0x01, 0x39,
	0xd1, 0xc0, 0xfe, 0x41, 0xd0, 0xd8, 0x49, 0x29, 0x11, 0xf9, 0xe2, 0xa8, 0x42, 0xa0, 0xe7, 0x15,
	0xc2, 0x2a, 0x29, 0x84, 0x5d, 0x2c, 0xc4, 0x05, 0x86, 0x95, 0x99, 0x30, 0x9c, 0x7b, 0x11, 0x86,
	0x7f, 0x22, 0x68, 0x3c, 0x4c, 0xfa, 0x13, 0x18, 0xe6, 0x38, 0x58, 0x13, 0x38, 0xd8, 0x2f, 0x9f,
	0x03, 0x83, 0xc6, 0x7d, 0x3a, 0xa0, 0xe2, 0x65, 0x4d, 0xb0, 0xbb, 0x01, 0x76, 0x40, 0x4e, 0xf0,
	0x2b, 0x50, 0xfd, 0x8a, 0xa5, 0x43, 0x22, 0x4c, 0x22, 0xf3, 0x24, 0x53, 0xf5, 0x89, 0x20, 0x59,
	0x2a, 0x29, 0xbb, 0xdf, 0xd9, 0x50, 0xcf, 0xd3, 0x9e, 0x11, 0xc6, 0x75, 0xc0, 0xa1, 0x9c, 0xcf,
	0x88, 0xc5, 0x5d, 0x11, 0x0d, 0x29, 0x17, 0x64, 0x98, 0x28, 0xb8, 0x76, 0x70, 0x2d, 0x3b, 0x39,
	0xc8, 0x0e, 0xb0, 0x0f, 0x35, 0x96, 0xa8, 0x9b, 0xcc, 0x99, 0x53, 0x37, 0xc7, 0xbb, 0x53, 0xf7,
	0xa9, 0xf5, 0x91, 0x76, 0xdc, 0x8d, 0x45, 0x7a, 0x1a, 0x64, 0x61, 0xf0, 0x3b, 0x50, 0xe1, 0x09,
	0x0d, 0x9d, 0xaa, 0x6a, 0xd8, 0x4a, 0x69, 0xb8, 0x4e, 0x42, 0xc3, 0x40, 0x99, 0xe3, 0xf7, 0xa0,
	0xca, 0x05, 0x11, 0x23, 0xee, 0xd4, 0x94, 0xe3, 0xad, 0x72, 0x47, 0x65, 0x1a, 0x18, 0x97, 0xe6,
	0x16, 0xd4, 0xf3, 0x60, 0xe4, 0x85, 0x72, 0x44, 0x4f, 0x4d, 0x25, 0xa5, 0x88, 0x6f, 0xc0, 0xdc,
	0x31, 0x19, 0x8c, 0xb2, 0x4a, 0xea, 0x87, 0x2d, 0xeb, 0x1e, 0x72, 0x7f, 0x44, 0x50, 0x91, 0x38,
	0xf0, 0x87, 0x50, 0xeb, 0x91, 0xf0, 0x88, 0x9a, 0x16, 0x2c, 0x6c, 0x6e, 0x94, 0x42, 0x78, 0xff,
	0xe0, 0xc0, 0x2f, 0x8c, 0xed, 0x68, 0x40, 0x83, 0x2c, 0x02, 0xde, 0x81, 0xb9, 0x74, 0x34, 0xa0,
	0xdc, 0xb1, 0x54, 0x55, 0xd7, 0xa7, 0x9f, 0x7e, 0x19, 0x46, 0xfb, 0xba, 0x1f, 0x43, 0x55, 0x13,
	0xc5, 0x0f, 0xce, 0xaa, 0xa3, 0xef, 0x77, 0x6f, 0xea, 0x78, 0xc5, 0x4a, 0xb9, 0xf7, 0x00, 0x5f,
	0x3c, 0xc5, 0x57, 0xc1, 0xda, 0xf3, 0x4d, 0xb9, 0xac, 0x3d, 0x5f, 0x8e, 0xdd, 0x21, 0xe3, 0x22,
	0x1b, 0x3b, 0x29, 0xbb, 0x9f, 0xc3, 0xf5, 0xbc, 0xe7, 0xb6, 0x21, 0xba, 0x02, 0x75, 0x4e, 0xd3,
	0xe3, 0x28, 0xa4, 0xdd, 0xdc, 0x25, 0xb1, 0x60, 0x74, 0xfb, 0x72, 0x60, 0x73, 0x26, 0x09, 0x4b,
	0xb3, 0xa8, 0x99, 0x89, 0xcf, 0x52, 0xe1, 0xfe, 0x8d, 0x60, 0xe9, 0x7c, 0x15, 0xce, 0x50, 0xa0,
	0x31, 0x0a, 0x7c, 0x1b, 0x16, 0x3b, 0x9d, 0x76, 0x97, 0xd3, 0x30, 0xa5, 0x42, 0x67, 0xd4, 0xaf,
	0x8d, 0x46, 0xa7, 0xd3, 0xee, 0x28, 0xad, 0xca, 0xb9, 0x0b, 0x95, 0x43, 0x21, 0x12, 0x53, 0xfe,
	0xff, 0xd1, 0x49, 0xe5, 0x8e, 0xb7, 0xc1, 0x16, 0x61, 0xe2, 0xd8, 0x2a, 0xca, 0xdd, 0xd2, 0x28,
	0x07, 0x3b, 0x17, 0x83, 0x48, 0x67, 0xf7, 0x0f, 0x04, 0x37, 0x2e, 0x4b, 0x21, 0xf9, 0x25, 0x44,
	0x1c, 0x66, 0xfc, 0xa4, 0x8c, 0x3f, 0x18, 0x0f, 0xa1, 0x7e, 0x2f, 0xdd, 0x9d, 0xba, 0xd3, 0xa6,
	0x23, 0xe3, 0x19, 0x5c, 0x81, 0xfa, 0x21, 0x25, 0x7d, 0x9a, 0x76, 0xf5, 0x28, 0x4a, 0x16, 0xf3,
	0xc1, 0x82, 0xd6, 0x49, 0x04, 0x1c, 0xdf, 0x82, 0x46, 0x4a, 0x4f, 0xd2, 0x48, 0x50, 0x63, 0x53,
	0x51, 0x36, 0x75, 0xa3, 0x54, 0x46, 0xee, 0x33, 0x04, 0xd7, 0x2f, 0x61, 0xa7, 0xf0, 0xcb, 0x7e,
	0x66, 0xf8, 0x59, 0x2a, 0x66, 0x8a, 0xff, 0x92, 0x5e, 0xdb, 0x97, 0xf5, 0xfa, 0x36, 0x2c, 0x1a,
	0x9b, 0x84, 0x0e, 0xb5, 0x9d, 0xbe, 0x9e, 0x1b, 0x5a, 0xed, 0xd3, 0xa1, 0xb4, 0xdb, 0xfc, 0xa9,
	0x06, 0x8d, 0x7c, 0x42, 0x8e, 0x7f, 0x45, 0x50, 0x91, 0xfb, 0x14, 0x5e, 0x2d, 0x47, 0x39, 0xde,
	0xdc, 0x9a, 0x77, 0xa6, 0xb0, 0xd4, 0xfb, 0x8a, 0xfb, 0xf0, 0xe9, 0xef, 0x8e, 0x75, 0x05, 0x3d,
	0xfd, 0xeb, 0xdf, 0xef, 0xad, 0x3d, 0xfc, 0xc0, 0xeb, 0x16, 0x96, 0xc4, 0xb1, 0xb7, 0x67, 0xbc,
	0x3d, 0xf3, 0x36, 0xe1, 0xde, 0x63, 0x23, 0x3d, 0xf1, 0xf2, 0x6b, 0x2c, 0xf7, 0xbe, 0xe6, 0x2c,
	0xc6, 0xcf, 0x10, 0x5c, 0xc9, 0x76, 0x23, 0xfc, 0x56, 0x29, 0x9c, 0x73, 0x8b, 0x5c, 0x73, 0x7d,
	0x4a, 0x6b, 0x43, 0xe0, 0x8b, 0x1c, 0x01, 0x1f, 0xef, 0xcf, 0x80, 0xc0, 0x63, 0xd9, 0x9d, 0x27,
	0x9a, 0xc7, 0xcf, 0x08, 0xaa, 0x7a, 0x95, 0xc2, 0x6b, 0xa5, 0xb8, 0x0a, 0xfb, 0x56, 0xf3, 0xb5,
	0xb1, 0xad, 0xde, 0xb3, 0x5b, 0x9f, 0xb0, 0xa8, 0x7f, 0x86, 0xf9, 0xd3, 0x1c, 0xe6, 0xb6, 0x3b,
	0xab, 0xa2, 0x6f, 0xa1, 0x35, 0xfc, 0x1b, 0x82, 0xaa, 0x5e, 0x8c, 0x26, 0xe0, 0x2d, 0x6c, 0x4f,
	0x93, 0xf0, 0x76, 0x73, 0x78, 0x3b, 0xcd, 0x19, 0xd7, 0x58, 0xc2, 0xfe, 0x05, 0x41, 0x55, 0x2f,
	0x43, 0x13, 0x60, 0x17, 0x36, 0xa6, 0x49, 0xb0, 0x0b, 0xa3, 0xb1, 0x36, 0x63, 0xd8, 0xdb, 0xbb,
	0xf0, 0x66, 0xc8, 0x86, 0x63, 0x04, 0x24, 0x89, 0x5a, 0x85, 0x0f, 0x3a, 0x13, 0x75, 0xfb, 0x5a,
	0xfe, 0xff, 0xec, 0xcb, 0x2f, 0x2d, 0x1f, 0x7d, 0x56, 0x33, 0xa7, 0xbd, 0xaa, 0xfa, 0xf6, 0x7a,
	0xfb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xc6, 0xd8, 0x63, 0x0d, 0x0e, 0x00, 0x00,
}
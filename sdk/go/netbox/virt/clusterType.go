// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package virt

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#cluster-types):
//
// > A cluster type represents a technology or mechanism by which a cluster is formed. For example, you might create a cluster type named "VMware vSphere" for a locally hosted cluster or "DigitalOcean NYC3" for one hosted by a cloud provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/virt"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := virt.NewClusterType(ctx, "vmw_vsphere", &virt.ClusterTypeArgs{
//				Name: pulumi.String("VMware vSphere 6"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ClusterType struct {
	pulumi.CustomResourceState

	Name pulumi.StringOutput `pulumi:"name"`
	Slug pulumi.StringOutput `pulumi:"slug"`
}

// NewClusterType registers a new resource with the given unique name, arguments, and options.
func NewClusterType(ctx *pulumi.Context,
	name string, args *ClusterTypeArgs, opts ...pulumi.ResourceOption) (*ClusterType, error) {
	if args == nil {
		args = &ClusterTypeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterType
	err := ctx.RegisterResource("netbox:virt/clusterType:ClusterType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterType gets an existing ClusterType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterTypeState, opts ...pulumi.ResourceOption) (*ClusterType, error) {
	var resource ClusterType
	err := ctx.ReadResource("netbox:virt/clusterType:ClusterType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterType resources.
type clusterTypeState struct {
	Name *string `pulumi:"name"`
	Slug *string `pulumi:"slug"`
}

type ClusterTypeState struct {
	Name pulumi.StringPtrInput
	Slug pulumi.StringPtrInput
}

func (ClusterTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterTypeState)(nil)).Elem()
}

type clusterTypeArgs struct {
	Name *string `pulumi:"name"`
	Slug *string `pulumi:"slug"`
}

// The set of arguments for constructing a ClusterType resource.
type ClusterTypeArgs struct {
	Name pulumi.StringPtrInput
	Slug pulumi.StringPtrInput
}

func (ClusterTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterTypeArgs)(nil)).Elem()
}

type ClusterTypeInput interface {
	pulumi.Input

	ToClusterTypeOutput() ClusterTypeOutput
	ToClusterTypeOutputWithContext(ctx context.Context) ClusterTypeOutput
}

func (*ClusterType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterType)(nil)).Elem()
}

func (i *ClusterType) ToClusterTypeOutput() ClusterTypeOutput {
	return i.ToClusterTypeOutputWithContext(context.Background())
}

func (i *ClusterType) ToClusterTypeOutputWithContext(ctx context.Context) ClusterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTypeOutput)
}

// ClusterTypeArrayInput is an input type that accepts ClusterTypeArray and ClusterTypeArrayOutput values.
// You can construct a concrete instance of `ClusterTypeArrayInput` via:
//
//	ClusterTypeArray{ ClusterTypeArgs{...} }
type ClusterTypeArrayInput interface {
	pulumi.Input

	ToClusterTypeArrayOutput() ClusterTypeArrayOutput
	ToClusterTypeArrayOutputWithContext(context.Context) ClusterTypeArrayOutput
}

type ClusterTypeArray []ClusterTypeInput

func (ClusterTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterType)(nil)).Elem()
}

func (i ClusterTypeArray) ToClusterTypeArrayOutput() ClusterTypeArrayOutput {
	return i.ToClusterTypeArrayOutputWithContext(context.Background())
}

func (i ClusterTypeArray) ToClusterTypeArrayOutputWithContext(ctx context.Context) ClusterTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTypeArrayOutput)
}

// ClusterTypeMapInput is an input type that accepts ClusterTypeMap and ClusterTypeMapOutput values.
// You can construct a concrete instance of `ClusterTypeMapInput` via:
//
//	ClusterTypeMap{ "key": ClusterTypeArgs{...} }
type ClusterTypeMapInput interface {
	pulumi.Input

	ToClusterTypeMapOutput() ClusterTypeMapOutput
	ToClusterTypeMapOutputWithContext(context.Context) ClusterTypeMapOutput
}

type ClusterTypeMap map[string]ClusterTypeInput

func (ClusterTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterType)(nil)).Elem()
}

func (i ClusterTypeMap) ToClusterTypeMapOutput() ClusterTypeMapOutput {
	return i.ToClusterTypeMapOutputWithContext(context.Background())
}

func (i ClusterTypeMap) ToClusterTypeMapOutputWithContext(ctx context.Context) ClusterTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTypeMapOutput)
}

type ClusterTypeOutput struct{ *pulumi.OutputState }

func (ClusterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterType)(nil)).Elem()
}

func (o ClusterTypeOutput) ToClusterTypeOutput() ClusterTypeOutput {
	return o
}

func (o ClusterTypeOutput) ToClusterTypeOutputWithContext(ctx context.Context) ClusterTypeOutput {
	return o
}

func (o ClusterTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterTypeOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterType) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

type ClusterTypeArrayOutput struct{ *pulumi.OutputState }

func (ClusterTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterType)(nil)).Elem()
}

func (o ClusterTypeArrayOutput) ToClusterTypeArrayOutput() ClusterTypeArrayOutput {
	return o
}

func (o ClusterTypeArrayOutput) ToClusterTypeArrayOutputWithContext(ctx context.Context) ClusterTypeArrayOutput {
	return o
}

func (o ClusterTypeArrayOutput) Index(i pulumi.IntInput) ClusterTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterType {
		return vs[0].([]*ClusterType)[vs[1].(int)]
	}).(ClusterTypeOutput)
}

type ClusterTypeMapOutput struct{ *pulumi.OutputState }

func (ClusterTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterType)(nil)).Elem()
}

func (o ClusterTypeMapOutput) ToClusterTypeMapOutput() ClusterTypeMapOutput {
	return o
}

func (o ClusterTypeMapOutput) ToClusterTypeMapOutputWithContext(ctx context.Context) ClusterTypeMapOutput {
	return o
}

func (o ClusterTypeMapOutput) MapIndex(k pulumi.StringInput) ClusterTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterType {
		return vs[0].(map[string]*ClusterType)[vs[1].(string)]
	}).(ClusterTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTypeInput)(nil)).Elem(), &ClusterType{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTypeArrayInput)(nil)).Elem(), ClusterTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTypeMapInput)(nil)).Elem(), ClusterTypeMap{})
	pulumi.RegisterOutputType(ClusterTypeOutput{})
	pulumi.RegisterOutputType(ClusterTypeArrayOutput{})
	pulumi.RegisterOutputType(ClusterTypeMapOutput{})
}

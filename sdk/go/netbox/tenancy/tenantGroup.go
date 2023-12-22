// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tenancy

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/features/tenancy/#tenant-groups):
//
// > Tenants can be organized by custom groups. For instance, you might create one group called "Customers" and one called "Departments." The assignment of a tenant to a group is optional.
// >
// > Tenant groups may be nested recursively to achieve a multi-level hierarchy. For example, you might have a group called "Customers" containing subgroups of individual tenants grouped by product or account team.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/tenancy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tenancy.NewTenantGroup(ctx, "test", nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type TenantGroup struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	ParentId    pulumi.IntPtrOutput    `pulumi:"parentId"`
	Slug        pulumi.StringOutput    `pulumi:"slug"`
}

// NewTenantGroup registers a new resource with the given unique name, arguments, and options.
func NewTenantGroup(ctx *pulumi.Context,
	name string, args *TenantGroupArgs, opts ...pulumi.ResourceOption) (*TenantGroup, error) {
	if args == nil {
		args = &TenantGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TenantGroup
	err := ctx.RegisterResource("netbox:tenancy/tenantGroup:TenantGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenantGroup gets an existing TenantGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenantGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantGroupState, opts ...pulumi.ResourceOption) (*TenantGroup, error) {
	var resource TenantGroup
	err := ctx.ReadResource("netbox:tenancy/tenantGroup:TenantGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TenantGroup resources.
type tenantGroupState struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	ParentId    *int    `pulumi:"parentId"`
	Slug        *string `pulumi:"slug"`
}

type TenantGroupState struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	ParentId    pulumi.IntPtrInput
	Slug        pulumi.StringPtrInput
}

func (TenantGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantGroupState)(nil)).Elem()
}

type tenantGroupArgs struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	ParentId    *int    `pulumi:"parentId"`
	Slug        *string `pulumi:"slug"`
}

// The set of arguments for constructing a TenantGroup resource.
type TenantGroupArgs struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	ParentId    pulumi.IntPtrInput
	Slug        pulumi.StringPtrInput
}

func (TenantGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantGroupArgs)(nil)).Elem()
}

type TenantGroupInput interface {
	pulumi.Input

	ToTenantGroupOutput() TenantGroupOutput
	ToTenantGroupOutputWithContext(ctx context.Context) TenantGroupOutput
}

func (*TenantGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantGroup)(nil)).Elem()
}

func (i *TenantGroup) ToTenantGroupOutput() TenantGroupOutput {
	return i.ToTenantGroupOutputWithContext(context.Background())
}

func (i *TenantGroup) ToTenantGroupOutputWithContext(ctx context.Context) TenantGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantGroupOutput)
}

// TenantGroupArrayInput is an input type that accepts TenantGroupArray and TenantGroupArrayOutput values.
// You can construct a concrete instance of `TenantGroupArrayInput` via:
//
//	TenantGroupArray{ TenantGroupArgs{...} }
type TenantGroupArrayInput interface {
	pulumi.Input

	ToTenantGroupArrayOutput() TenantGroupArrayOutput
	ToTenantGroupArrayOutputWithContext(context.Context) TenantGroupArrayOutput
}

type TenantGroupArray []TenantGroupInput

func (TenantGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TenantGroup)(nil)).Elem()
}

func (i TenantGroupArray) ToTenantGroupArrayOutput() TenantGroupArrayOutput {
	return i.ToTenantGroupArrayOutputWithContext(context.Background())
}

func (i TenantGroupArray) ToTenantGroupArrayOutputWithContext(ctx context.Context) TenantGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantGroupArrayOutput)
}

// TenantGroupMapInput is an input type that accepts TenantGroupMap and TenantGroupMapOutput values.
// You can construct a concrete instance of `TenantGroupMapInput` via:
//
//	TenantGroupMap{ "key": TenantGroupArgs{...} }
type TenantGroupMapInput interface {
	pulumi.Input

	ToTenantGroupMapOutput() TenantGroupMapOutput
	ToTenantGroupMapOutputWithContext(context.Context) TenantGroupMapOutput
}

type TenantGroupMap map[string]TenantGroupInput

func (TenantGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TenantGroup)(nil)).Elem()
}

func (i TenantGroupMap) ToTenantGroupMapOutput() TenantGroupMapOutput {
	return i.ToTenantGroupMapOutputWithContext(context.Background())
}

func (i TenantGroupMap) ToTenantGroupMapOutputWithContext(ctx context.Context) TenantGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantGroupMapOutput)
}

type TenantGroupOutput struct{ *pulumi.OutputState }

func (TenantGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantGroup)(nil)).Elem()
}

func (o TenantGroupOutput) ToTenantGroupOutput() TenantGroupOutput {
	return o
}

func (o TenantGroupOutput) ToTenantGroupOutputWithContext(ctx context.Context) TenantGroupOutput {
	return o
}

func (o TenantGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TenantGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TenantGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TenantGroupOutput) ParentId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TenantGroup) pulumi.IntPtrOutput { return v.ParentId }).(pulumi.IntPtrOutput)
}

func (o TenantGroupOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantGroup) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

type TenantGroupArrayOutput struct{ *pulumi.OutputState }

func (TenantGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TenantGroup)(nil)).Elem()
}

func (o TenantGroupArrayOutput) ToTenantGroupArrayOutput() TenantGroupArrayOutput {
	return o
}

func (o TenantGroupArrayOutput) ToTenantGroupArrayOutputWithContext(ctx context.Context) TenantGroupArrayOutput {
	return o
}

func (o TenantGroupArrayOutput) Index(i pulumi.IntInput) TenantGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TenantGroup {
		return vs[0].([]*TenantGroup)[vs[1].(int)]
	}).(TenantGroupOutput)
}

type TenantGroupMapOutput struct{ *pulumi.OutputState }

func (TenantGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TenantGroup)(nil)).Elem()
}

func (o TenantGroupMapOutput) ToTenantGroupMapOutput() TenantGroupMapOutput {
	return o
}

func (o TenantGroupMapOutput) ToTenantGroupMapOutputWithContext(ctx context.Context) TenantGroupMapOutput {
	return o
}

func (o TenantGroupMapOutput) MapIndex(k pulumi.StringInput) TenantGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TenantGroup {
		return vs[0].(map[string]*TenantGroup)[vs[1].(string)]
	}).(TenantGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TenantGroupInput)(nil)).Elem(), &TenantGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*TenantGroupArrayInput)(nil)).Elem(), TenantGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TenantGroupMapInput)(nil)).Elem(), TenantGroupMap{})
	pulumi.RegisterOutputType(TenantGroupOutput{})
	pulumi.RegisterOutputType(TenantGroupArrayOutput{})
	pulumi.RegisterOutputType(TenantGroupMapOutput{})
}

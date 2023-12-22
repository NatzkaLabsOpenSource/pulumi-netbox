// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tenancy

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/features/tenancy/#tenants):
//
// > A tenant represents a discrete grouping of resources used for administrative purposes. Typically, tenants are used to represent individual customers or internal departments within an organization.
// >
// > Tenant assignment is used to signify the ownership of an object in NetBox. As such, each object may only be owned by a single tenant. For example, if you have a firewall dedicated to a particular customer, you would assign it to the tenant which represents that customer. However, if the firewall serves multiple customers, it doesn't belong to any particular customer, so tenant assignment would not be appropriate.
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
//			_, err := tenancy.NewTenant(ctx, "customerA", nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Tenant struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput   `pulumi:"description"`
	GroupId     pulumi.IntPtrOutput      `pulumi:"groupId"`
	Name        pulumi.StringOutput      `pulumi:"name"`
	Slug        pulumi.StringOutput      `pulumi:"slug"`
	Tags        pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewTenant registers a new resource with the given unique name, arguments, and options.
func NewTenant(ctx *pulumi.Context,
	name string, args *TenantArgs, opts ...pulumi.ResourceOption) (*Tenant, error) {
	if args == nil {
		args = &TenantArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Tenant
	err := ctx.RegisterResource("netbox:tenancy/tenant:Tenant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenant gets an existing Tenant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantState, opts ...pulumi.ResourceOption) (*Tenant, error) {
	var resource Tenant
	err := ctx.ReadResource("netbox:tenancy/tenant:Tenant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tenant resources.
type tenantState struct {
	Description *string  `pulumi:"description"`
	GroupId     *int     `pulumi:"groupId"`
	Name        *string  `pulumi:"name"`
	Slug        *string  `pulumi:"slug"`
	Tags        []string `pulumi:"tags"`
}

type TenantState struct {
	Description pulumi.StringPtrInput
	GroupId     pulumi.IntPtrInput
	Name        pulumi.StringPtrInput
	Slug        pulumi.StringPtrInput
	Tags        pulumi.StringArrayInput
}

func (TenantState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantState)(nil)).Elem()
}

type tenantArgs struct {
	Description *string  `pulumi:"description"`
	GroupId     *int     `pulumi:"groupId"`
	Name        *string  `pulumi:"name"`
	Slug        *string  `pulumi:"slug"`
	Tags        []string `pulumi:"tags"`
}

// The set of arguments for constructing a Tenant resource.
type TenantArgs struct {
	Description pulumi.StringPtrInput
	GroupId     pulumi.IntPtrInput
	Name        pulumi.StringPtrInput
	Slug        pulumi.StringPtrInput
	Tags        pulumi.StringArrayInput
}

func (TenantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantArgs)(nil)).Elem()
}

type TenantInput interface {
	pulumi.Input

	ToTenantOutput() TenantOutput
	ToTenantOutputWithContext(ctx context.Context) TenantOutput
}

func (*Tenant) ElementType() reflect.Type {
	return reflect.TypeOf((**Tenant)(nil)).Elem()
}

func (i *Tenant) ToTenantOutput() TenantOutput {
	return i.ToTenantOutputWithContext(context.Background())
}

func (i *Tenant) ToTenantOutputWithContext(ctx context.Context) TenantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantOutput)
}

// TenantArrayInput is an input type that accepts TenantArray and TenantArrayOutput values.
// You can construct a concrete instance of `TenantArrayInput` via:
//
//	TenantArray{ TenantArgs{...} }
type TenantArrayInput interface {
	pulumi.Input

	ToTenantArrayOutput() TenantArrayOutput
	ToTenantArrayOutputWithContext(context.Context) TenantArrayOutput
}

type TenantArray []TenantInput

func (TenantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tenant)(nil)).Elem()
}

func (i TenantArray) ToTenantArrayOutput() TenantArrayOutput {
	return i.ToTenantArrayOutputWithContext(context.Background())
}

func (i TenantArray) ToTenantArrayOutputWithContext(ctx context.Context) TenantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantArrayOutput)
}

// TenantMapInput is an input type that accepts TenantMap and TenantMapOutput values.
// You can construct a concrete instance of `TenantMapInput` via:
//
//	TenantMap{ "key": TenantArgs{...} }
type TenantMapInput interface {
	pulumi.Input

	ToTenantMapOutput() TenantMapOutput
	ToTenantMapOutputWithContext(context.Context) TenantMapOutput
}

type TenantMap map[string]TenantInput

func (TenantMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tenant)(nil)).Elem()
}

func (i TenantMap) ToTenantMapOutput() TenantMapOutput {
	return i.ToTenantMapOutputWithContext(context.Background())
}

func (i TenantMap) ToTenantMapOutputWithContext(ctx context.Context) TenantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantMapOutput)
}

type TenantOutput struct{ *pulumi.OutputState }

func (TenantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tenant)(nil)).Elem()
}

func (o TenantOutput) ToTenantOutput() TenantOutput {
	return o
}

func (o TenantOutput) ToTenantOutputWithContext(ctx context.Context) TenantOutput {
	return o
}

func (o TenantOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tenant) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TenantOutput) GroupId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Tenant) pulumi.IntPtrOutput { return v.GroupId }).(pulumi.IntPtrOutput)
}

func (o TenantOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Tenant) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TenantOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Tenant) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

func (o TenantOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Tenant) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

type TenantArrayOutput struct{ *pulumi.OutputState }

func (TenantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tenant)(nil)).Elem()
}

func (o TenantArrayOutput) ToTenantArrayOutput() TenantArrayOutput {
	return o
}

func (o TenantArrayOutput) ToTenantArrayOutputWithContext(ctx context.Context) TenantArrayOutput {
	return o
}

func (o TenantArrayOutput) Index(i pulumi.IntInput) TenantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Tenant {
		return vs[0].([]*Tenant)[vs[1].(int)]
	}).(TenantOutput)
}

type TenantMapOutput struct{ *pulumi.OutputState }

func (TenantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tenant)(nil)).Elem()
}

func (o TenantMapOutput) ToTenantMapOutput() TenantMapOutput {
	return o
}

func (o TenantMapOutput) ToTenantMapOutputWithContext(ctx context.Context) TenantMapOutput {
	return o
}

func (o TenantMapOutput) MapIndex(k pulumi.StringInput) TenantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Tenant {
		return vs[0].(map[string]*Tenant)[vs[1].(string)]
	}).(TenantOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TenantInput)(nil)).Elem(), &Tenant{})
	pulumi.RegisterInputType(reflect.TypeOf((*TenantArrayInput)(nil)).Elem(), TenantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TenantMapInput)(nil)).Elem(), TenantMap{})
	pulumi.RegisterOutputType(TenantOutput{})
	pulumi.RegisterOutputType(TenantArrayOutput{})
	pulumi.RegisterOutputType(TenantMapOutput{})
}

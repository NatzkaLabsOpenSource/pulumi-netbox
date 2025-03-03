// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tenancy

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/features/contacts/#contactroles):
//
// > A contact role defines the relationship of a contact to an assigned object. For example, you might define roles for administrative, operational, and emergency contacts
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
//			_, err := tenancy.NewContactRole(ctx, "test", &tenancy.ContactRoleArgs{
//				Name: pulumi.String("test"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ContactRole struct {
	pulumi.CustomResourceState

	Name pulumi.StringOutput `pulumi:"name"`
	Slug pulumi.StringOutput `pulumi:"slug"`
}

// NewContactRole registers a new resource with the given unique name, arguments, and options.
func NewContactRole(ctx *pulumi.Context,
	name string, args *ContactRoleArgs, opts ...pulumi.ResourceOption) (*ContactRole, error) {
	if args == nil {
		args = &ContactRoleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContactRole
	err := ctx.RegisterResource("netbox:tenancy/contactRole:ContactRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContactRole gets an existing ContactRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContactRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactRoleState, opts ...pulumi.ResourceOption) (*ContactRole, error) {
	var resource ContactRole
	err := ctx.ReadResource("netbox:tenancy/contactRole:ContactRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContactRole resources.
type contactRoleState struct {
	Name *string `pulumi:"name"`
	Slug *string `pulumi:"slug"`
}

type ContactRoleState struct {
	Name pulumi.StringPtrInput
	Slug pulumi.StringPtrInput
}

func (ContactRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactRoleState)(nil)).Elem()
}

type contactRoleArgs struct {
	Name *string `pulumi:"name"`
	Slug *string `pulumi:"slug"`
}

// The set of arguments for constructing a ContactRole resource.
type ContactRoleArgs struct {
	Name pulumi.StringPtrInput
	Slug pulumi.StringPtrInput
}

func (ContactRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactRoleArgs)(nil)).Elem()
}

type ContactRoleInput interface {
	pulumi.Input

	ToContactRoleOutput() ContactRoleOutput
	ToContactRoleOutputWithContext(ctx context.Context) ContactRoleOutput
}

func (*ContactRole) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactRole)(nil)).Elem()
}

func (i *ContactRole) ToContactRoleOutput() ContactRoleOutput {
	return i.ToContactRoleOutputWithContext(context.Background())
}

func (i *ContactRole) ToContactRoleOutputWithContext(ctx context.Context) ContactRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactRoleOutput)
}

// ContactRoleArrayInput is an input type that accepts ContactRoleArray and ContactRoleArrayOutput values.
// You can construct a concrete instance of `ContactRoleArrayInput` via:
//
//	ContactRoleArray{ ContactRoleArgs{...} }
type ContactRoleArrayInput interface {
	pulumi.Input

	ToContactRoleArrayOutput() ContactRoleArrayOutput
	ToContactRoleArrayOutputWithContext(context.Context) ContactRoleArrayOutput
}

type ContactRoleArray []ContactRoleInput

func (ContactRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContactRole)(nil)).Elem()
}

func (i ContactRoleArray) ToContactRoleArrayOutput() ContactRoleArrayOutput {
	return i.ToContactRoleArrayOutputWithContext(context.Background())
}

func (i ContactRoleArray) ToContactRoleArrayOutputWithContext(ctx context.Context) ContactRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactRoleArrayOutput)
}

// ContactRoleMapInput is an input type that accepts ContactRoleMap and ContactRoleMapOutput values.
// You can construct a concrete instance of `ContactRoleMapInput` via:
//
//	ContactRoleMap{ "key": ContactRoleArgs{...} }
type ContactRoleMapInput interface {
	pulumi.Input

	ToContactRoleMapOutput() ContactRoleMapOutput
	ToContactRoleMapOutputWithContext(context.Context) ContactRoleMapOutput
}

type ContactRoleMap map[string]ContactRoleInput

func (ContactRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContactRole)(nil)).Elem()
}

func (i ContactRoleMap) ToContactRoleMapOutput() ContactRoleMapOutput {
	return i.ToContactRoleMapOutputWithContext(context.Background())
}

func (i ContactRoleMap) ToContactRoleMapOutputWithContext(ctx context.Context) ContactRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactRoleMapOutput)
}

type ContactRoleOutput struct{ *pulumi.OutputState }

func (ContactRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactRole)(nil)).Elem()
}

func (o ContactRoleOutput) ToContactRoleOutput() ContactRoleOutput {
	return o
}

func (o ContactRoleOutput) ToContactRoleOutputWithContext(ctx context.Context) ContactRoleOutput {
	return o
}

func (o ContactRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ContactRoleOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactRole) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

type ContactRoleArrayOutput struct{ *pulumi.OutputState }

func (ContactRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContactRole)(nil)).Elem()
}

func (o ContactRoleArrayOutput) ToContactRoleArrayOutput() ContactRoleArrayOutput {
	return o
}

func (o ContactRoleArrayOutput) ToContactRoleArrayOutputWithContext(ctx context.Context) ContactRoleArrayOutput {
	return o
}

func (o ContactRoleArrayOutput) Index(i pulumi.IntInput) ContactRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContactRole {
		return vs[0].([]*ContactRole)[vs[1].(int)]
	}).(ContactRoleOutput)
}

type ContactRoleMapOutput struct{ *pulumi.OutputState }

func (ContactRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContactRole)(nil)).Elem()
}

func (o ContactRoleMapOutput) ToContactRoleMapOutput() ContactRoleMapOutput {
	return o
}

func (o ContactRoleMapOutput) ToContactRoleMapOutputWithContext(ctx context.Context) ContactRoleMapOutput {
	return o
}

func (o ContactRoleMapOutput) MapIndex(k pulumi.StringInput) ContactRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContactRole {
		return vs[0].(map[string]*ContactRole)[vs[1].(string)]
	}).(ContactRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContactRoleInput)(nil)).Elem(), &ContactRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactRoleArrayInput)(nil)).Elem(), ContactRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactRoleMapInput)(nil)).Elem(), ContactRoleMap{})
	pulumi.RegisterOutputType(ContactRoleOutput{})
	pulumi.RegisterOutputType(ContactRoleArrayOutput{})
	pulumi.RegisterOutputType(ContactRoleMapOutput{})
}

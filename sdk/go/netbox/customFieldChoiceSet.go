// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package netbox

import (
	"context"
	"reflect"

	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/models/extras/customfieldchoiceset/):
//
// Single- and multi-selection custom fields must define a set of valid choices from which the user may choose when defining the field value. These choices are defined as sets that may be reused among multiple custom fields.
//
// A choice set must define a base choice set and/or a set of arbitrary extra choices.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := netbox.NewCustomFieldChoiceSet(ctx, "test", &netbox.CustomFieldChoiceSetArgs{
//				Description: pulumi.String("Description"),
//				ExtraChoices: pulumi.StringArrayArray{
//					pulumi.StringArray("choice1"),
//					pulumi.StringArray("label1"),
//					pulumi.StringArray("choice2"),
//					pulumi.StringArray("choice2"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type CustomFieldChoiceSet struct {
	pulumi.CustomResourceState

	// Valid values are `IATA`, `ISO_3166` and `UN_LOCODE`. At least one of `baseChoices` or `extraChoices` must be given.
	BaseChoices  pulumi.StringPtrOutput `pulumi:"baseChoices"`
	CustomFields pulumi.StringMapOutput `pulumi:"customFields"`
	Description  pulumi.StringPtrOutput `pulumi:"description"`
	// This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice. At least one of `baseChoices` or `extraChoices` must be given.
	ExtraChoices pulumi.StringArrayArrayOutput `pulumi:"extraChoices"`
	Name         pulumi.StringOutput           `pulumi:"name"`
	// experimental. Defaults to `false`.
	OrderAlphabetically pulumi.BoolPtrOutput `pulumi:"orderAlphabetically"`
}

// NewCustomFieldChoiceSet registers a new resource with the given unique name, arguments, and options.
func NewCustomFieldChoiceSet(ctx *pulumi.Context,
	name string, args *CustomFieldChoiceSetArgs, opts ...pulumi.ResourceOption) (*CustomFieldChoiceSet, error) {
	if args == nil {
		args = &CustomFieldChoiceSetArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomFieldChoiceSet
	err := ctx.RegisterResource("netbox:index/customFieldChoiceSet:CustomFieldChoiceSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomFieldChoiceSet gets an existing CustomFieldChoiceSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomFieldChoiceSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomFieldChoiceSetState, opts ...pulumi.ResourceOption) (*CustomFieldChoiceSet, error) {
	var resource CustomFieldChoiceSet
	err := ctx.ReadResource("netbox:index/customFieldChoiceSet:CustomFieldChoiceSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomFieldChoiceSet resources.
type customFieldChoiceSetState struct {
	// Valid values are `IATA`, `ISO_3166` and `UN_LOCODE`. At least one of `baseChoices` or `extraChoices` must be given.
	BaseChoices  *string           `pulumi:"baseChoices"`
	CustomFields map[string]string `pulumi:"customFields"`
	Description  *string           `pulumi:"description"`
	// This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice. At least one of `baseChoices` or `extraChoices` must be given.
	ExtraChoices [][]string `pulumi:"extraChoices"`
	Name         *string    `pulumi:"name"`
	// experimental. Defaults to `false`.
	OrderAlphabetically *bool `pulumi:"orderAlphabetically"`
}

type CustomFieldChoiceSetState struct {
	// Valid values are `IATA`, `ISO_3166` and `UN_LOCODE`. At least one of `baseChoices` or `extraChoices` must be given.
	BaseChoices  pulumi.StringPtrInput
	CustomFields pulumi.StringMapInput
	Description  pulumi.StringPtrInput
	// This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice. At least one of `baseChoices` or `extraChoices` must be given.
	ExtraChoices pulumi.StringArrayArrayInput
	Name         pulumi.StringPtrInput
	// experimental. Defaults to `false`.
	OrderAlphabetically pulumi.BoolPtrInput
}

func (CustomFieldChoiceSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*customFieldChoiceSetState)(nil)).Elem()
}

type customFieldChoiceSetArgs struct {
	// Valid values are `IATA`, `ISO_3166` and `UN_LOCODE`. At least one of `baseChoices` or `extraChoices` must be given.
	BaseChoices  *string           `pulumi:"baseChoices"`
	CustomFields map[string]string `pulumi:"customFields"`
	Description  *string           `pulumi:"description"`
	// This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice. At least one of `baseChoices` or `extraChoices` must be given.
	ExtraChoices [][]string `pulumi:"extraChoices"`
	Name         *string    `pulumi:"name"`
	// experimental. Defaults to `false`.
	OrderAlphabetically *bool `pulumi:"orderAlphabetically"`
}

// The set of arguments for constructing a CustomFieldChoiceSet resource.
type CustomFieldChoiceSetArgs struct {
	// Valid values are `IATA`, `ISO_3166` and `UN_LOCODE`. At least one of `baseChoices` or `extraChoices` must be given.
	BaseChoices  pulumi.StringPtrInput
	CustomFields pulumi.StringMapInput
	Description  pulumi.StringPtrInput
	// This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice. At least one of `baseChoices` or `extraChoices` must be given.
	ExtraChoices pulumi.StringArrayArrayInput
	Name         pulumi.StringPtrInput
	// experimental. Defaults to `false`.
	OrderAlphabetically pulumi.BoolPtrInput
}

func (CustomFieldChoiceSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customFieldChoiceSetArgs)(nil)).Elem()
}

type CustomFieldChoiceSetInput interface {
	pulumi.Input

	ToCustomFieldChoiceSetOutput() CustomFieldChoiceSetOutput
	ToCustomFieldChoiceSetOutputWithContext(ctx context.Context) CustomFieldChoiceSetOutput
}

func (*CustomFieldChoiceSet) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomFieldChoiceSet)(nil)).Elem()
}

func (i *CustomFieldChoiceSet) ToCustomFieldChoiceSetOutput() CustomFieldChoiceSetOutput {
	return i.ToCustomFieldChoiceSetOutputWithContext(context.Background())
}

func (i *CustomFieldChoiceSet) ToCustomFieldChoiceSetOutputWithContext(ctx context.Context) CustomFieldChoiceSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomFieldChoiceSetOutput)
}

// CustomFieldChoiceSetArrayInput is an input type that accepts CustomFieldChoiceSetArray and CustomFieldChoiceSetArrayOutput values.
// You can construct a concrete instance of `CustomFieldChoiceSetArrayInput` via:
//
//	CustomFieldChoiceSetArray{ CustomFieldChoiceSetArgs{...} }
type CustomFieldChoiceSetArrayInput interface {
	pulumi.Input

	ToCustomFieldChoiceSetArrayOutput() CustomFieldChoiceSetArrayOutput
	ToCustomFieldChoiceSetArrayOutputWithContext(context.Context) CustomFieldChoiceSetArrayOutput
}

type CustomFieldChoiceSetArray []CustomFieldChoiceSetInput

func (CustomFieldChoiceSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomFieldChoiceSet)(nil)).Elem()
}

func (i CustomFieldChoiceSetArray) ToCustomFieldChoiceSetArrayOutput() CustomFieldChoiceSetArrayOutput {
	return i.ToCustomFieldChoiceSetArrayOutputWithContext(context.Background())
}

func (i CustomFieldChoiceSetArray) ToCustomFieldChoiceSetArrayOutputWithContext(ctx context.Context) CustomFieldChoiceSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomFieldChoiceSetArrayOutput)
}

// CustomFieldChoiceSetMapInput is an input type that accepts CustomFieldChoiceSetMap and CustomFieldChoiceSetMapOutput values.
// You can construct a concrete instance of `CustomFieldChoiceSetMapInput` via:
//
//	CustomFieldChoiceSetMap{ "key": CustomFieldChoiceSetArgs{...} }
type CustomFieldChoiceSetMapInput interface {
	pulumi.Input

	ToCustomFieldChoiceSetMapOutput() CustomFieldChoiceSetMapOutput
	ToCustomFieldChoiceSetMapOutputWithContext(context.Context) CustomFieldChoiceSetMapOutput
}

type CustomFieldChoiceSetMap map[string]CustomFieldChoiceSetInput

func (CustomFieldChoiceSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomFieldChoiceSet)(nil)).Elem()
}

func (i CustomFieldChoiceSetMap) ToCustomFieldChoiceSetMapOutput() CustomFieldChoiceSetMapOutput {
	return i.ToCustomFieldChoiceSetMapOutputWithContext(context.Background())
}

func (i CustomFieldChoiceSetMap) ToCustomFieldChoiceSetMapOutputWithContext(ctx context.Context) CustomFieldChoiceSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomFieldChoiceSetMapOutput)
}

type CustomFieldChoiceSetOutput struct{ *pulumi.OutputState }

func (CustomFieldChoiceSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomFieldChoiceSet)(nil)).Elem()
}

func (o CustomFieldChoiceSetOutput) ToCustomFieldChoiceSetOutput() CustomFieldChoiceSetOutput {
	return o
}

func (o CustomFieldChoiceSetOutput) ToCustomFieldChoiceSetOutputWithContext(ctx context.Context) CustomFieldChoiceSetOutput {
	return o
}

// Valid values are `IATA`, `ISO_3166` and `UN_LOCODE`. At least one of `baseChoices` or `extraChoices` must be given.
func (o CustomFieldChoiceSetOutput) BaseChoices() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomFieldChoiceSet) pulumi.StringPtrOutput { return v.BaseChoices }).(pulumi.StringPtrOutput)
}

func (o CustomFieldChoiceSetOutput) CustomFields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomFieldChoiceSet) pulumi.StringMapOutput { return v.CustomFields }).(pulumi.StringMapOutput)
}

func (o CustomFieldChoiceSetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomFieldChoiceSet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice. At least one of `baseChoices` or `extraChoices` must be given.
func (o CustomFieldChoiceSetOutput) ExtraChoices() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v *CustomFieldChoiceSet) pulumi.StringArrayArrayOutput { return v.ExtraChoices }).(pulumi.StringArrayArrayOutput)
}

func (o CustomFieldChoiceSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomFieldChoiceSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// experimental. Defaults to `false`.
func (o CustomFieldChoiceSetOutput) OrderAlphabetically() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomFieldChoiceSet) pulumi.BoolPtrOutput { return v.OrderAlphabetically }).(pulumi.BoolPtrOutput)
}

type CustomFieldChoiceSetArrayOutput struct{ *pulumi.OutputState }

func (CustomFieldChoiceSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomFieldChoiceSet)(nil)).Elem()
}

func (o CustomFieldChoiceSetArrayOutput) ToCustomFieldChoiceSetArrayOutput() CustomFieldChoiceSetArrayOutput {
	return o
}

func (o CustomFieldChoiceSetArrayOutput) ToCustomFieldChoiceSetArrayOutputWithContext(ctx context.Context) CustomFieldChoiceSetArrayOutput {
	return o
}

func (o CustomFieldChoiceSetArrayOutput) Index(i pulumi.IntInput) CustomFieldChoiceSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomFieldChoiceSet {
		return vs[0].([]*CustomFieldChoiceSet)[vs[1].(int)]
	}).(CustomFieldChoiceSetOutput)
}

type CustomFieldChoiceSetMapOutput struct{ *pulumi.OutputState }

func (CustomFieldChoiceSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomFieldChoiceSet)(nil)).Elem()
}

func (o CustomFieldChoiceSetMapOutput) ToCustomFieldChoiceSetMapOutput() CustomFieldChoiceSetMapOutput {
	return o
}

func (o CustomFieldChoiceSetMapOutput) ToCustomFieldChoiceSetMapOutputWithContext(ctx context.Context) CustomFieldChoiceSetMapOutput {
	return o
}

func (o CustomFieldChoiceSetMapOutput) MapIndex(k pulumi.StringInput) CustomFieldChoiceSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomFieldChoiceSet {
		return vs[0].(map[string]*CustomFieldChoiceSet)[vs[1].(string)]
	}).(CustomFieldChoiceSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomFieldChoiceSetInput)(nil)).Elem(), &CustomFieldChoiceSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomFieldChoiceSetArrayInput)(nil)).Elem(), CustomFieldChoiceSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomFieldChoiceSetMapInput)(nil)).Elem(), CustomFieldChoiceSetMap{})
	pulumi.RegisterOutputType(CustomFieldChoiceSetOutput{})
	pulumi.RegisterOutputType(CustomFieldChoiceSetArrayOutput{})
	pulumi.RegisterOutputType(CustomFieldChoiceSetMapOutput{})
}
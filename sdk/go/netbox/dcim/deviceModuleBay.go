// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcim

import (
	"context"
	"reflect"

	"errors"
	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/modulebay/):
//
// > Module bays represent a space or slot within a device in which a field-replaceable module may be installed. A common example is that of a chassis-based switch such as the Cisco Nexus 9000 or Juniper EX9200. Modules in turn hold additional components that become available to the parent device.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/dcim"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Note that some terraform code is not included in the example for brevity
//			test, err := dcim.NewDevice(ctx, "test", &dcim.DeviceArgs{
//				Name:         pulumi.String("%[1]s"),
//				DeviceTypeId: pulumi.Any(testNetboxDeviceType.Id),
//				RoleId:       pulumi.Any(testNetboxDeviceRole.Id),
//				SiteId:       pulumi.Any(testNetboxSite.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dcim.NewDeviceModuleBay(ctx, "test", &dcim.DeviceModuleBayArgs{
//				DeviceId: test.ID(),
//				Name:     pulumi.String("module bay 1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DeviceModuleBay struct {
	pulumi.CustomResourceState

	CustomFields pulumi.StringMapOutput   `pulumi:"customFields"`
	Description  pulumi.StringPtrOutput   `pulumi:"description"`
	DeviceId     pulumi.IntOutput         `pulumi:"deviceId"`
	Label        pulumi.StringPtrOutput   `pulumi:"label"`
	Name         pulumi.StringOutput      `pulumi:"name"`
	Position     pulumi.StringPtrOutput   `pulumi:"position"`
	Tags         pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewDeviceModuleBay registers a new resource with the given unique name, arguments, and options.
func NewDeviceModuleBay(ctx *pulumi.Context,
	name string, args *DeviceModuleBayArgs, opts ...pulumi.ResourceOption) (*DeviceModuleBay, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DeviceModuleBay
	err := ctx.RegisterResource("netbox:dcim/deviceModuleBay:DeviceModuleBay", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceModuleBay gets an existing DeviceModuleBay resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceModuleBay(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceModuleBayState, opts ...pulumi.ResourceOption) (*DeviceModuleBay, error) {
	var resource DeviceModuleBay
	err := ctx.ReadResource("netbox:dcim/deviceModuleBay:DeviceModuleBay", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceModuleBay resources.
type deviceModuleBayState struct {
	CustomFields map[string]string `pulumi:"customFields"`
	Description  *string           `pulumi:"description"`
	DeviceId     *int              `pulumi:"deviceId"`
	Label        *string           `pulumi:"label"`
	Name         *string           `pulumi:"name"`
	Position     *string           `pulumi:"position"`
	Tags         []string          `pulumi:"tags"`
}

type DeviceModuleBayState struct {
	CustomFields pulumi.StringMapInput
	Description  pulumi.StringPtrInput
	DeviceId     pulumi.IntPtrInput
	Label        pulumi.StringPtrInput
	Name         pulumi.StringPtrInput
	Position     pulumi.StringPtrInput
	Tags         pulumi.StringArrayInput
}

func (DeviceModuleBayState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceModuleBayState)(nil)).Elem()
}

type deviceModuleBayArgs struct {
	CustomFields map[string]string `pulumi:"customFields"`
	Description  *string           `pulumi:"description"`
	DeviceId     int               `pulumi:"deviceId"`
	Label        *string           `pulumi:"label"`
	Name         *string           `pulumi:"name"`
	Position     *string           `pulumi:"position"`
	Tags         []string          `pulumi:"tags"`
}

// The set of arguments for constructing a DeviceModuleBay resource.
type DeviceModuleBayArgs struct {
	CustomFields pulumi.StringMapInput
	Description  pulumi.StringPtrInput
	DeviceId     pulumi.IntInput
	Label        pulumi.StringPtrInput
	Name         pulumi.StringPtrInput
	Position     pulumi.StringPtrInput
	Tags         pulumi.StringArrayInput
}

func (DeviceModuleBayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceModuleBayArgs)(nil)).Elem()
}

type DeviceModuleBayInput interface {
	pulumi.Input

	ToDeviceModuleBayOutput() DeviceModuleBayOutput
	ToDeviceModuleBayOutputWithContext(ctx context.Context) DeviceModuleBayOutput
}

func (*DeviceModuleBay) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceModuleBay)(nil)).Elem()
}

func (i *DeviceModuleBay) ToDeviceModuleBayOutput() DeviceModuleBayOutput {
	return i.ToDeviceModuleBayOutputWithContext(context.Background())
}

func (i *DeviceModuleBay) ToDeviceModuleBayOutputWithContext(ctx context.Context) DeviceModuleBayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceModuleBayOutput)
}

// DeviceModuleBayArrayInput is an input type that accepts DeviceModuleBayArray and DeviceModuleBayArrayOutput values.
// You can construct a concrete instance of `DeviceModuleBayArrayInput` via:
//
//	DeviceModuleBayArray{ DeviceModuleBayArgs{...} }
type DeviceModuleBayArrayInput interface {
	pulumi.Input

	ToDeviceModuleBayArrayOutput() DeviceModuleBayArrayOutput
	ToDeviceModuleBayArrayOutputWithContext(context.Context) DeviceModuleBayArrayOutput
}

type DeviceModuleBayArray []DeviceModuleBayInput

func (DeviceModuleBayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceModuleBay)(nil)).Elem()
}

func (i DeviceModuleBayArray) ToDeviceModuleBayArrayOutput() DeviceModuleBayArrayOutput {
	return i.ToDeviceModuleBayArrayOutputWithContext(context.Background())
}

func (i DeviceModuleBayArray) ToDeviceModuleBayArrayOutputWithContext(ctx context.Context) DeviceModuleBayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceModuleBayArrayOutput)
}

// DeviceModuleBayMapInput is an input type that accepts DeviceModuleBayMap and DeviceModuleBayMapOutput values.
// You can construct a concrete instance of `DeviceModuleBayMapInput` via:
//
//	DeviceModuleBayMap{ "key": DeviceModuleBayArgs{...} }
type DeviceModuleBayMapInput interface {
	pulumi.Input

	ToDeviceModuleBayMapOutput() DeviceModuleBayMapOutput
	ToDeviceModuleBayMapOutputWithContext(context.Context) DeviceModuleBayMapOutput
}

type DeviceModuleBayMap map[string]DeviceModuleBayInput

func (DeviceModuleBayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceModuleBay)(nil)).Elem()
}

func (i DeviceModuleBayMap) ToDeviceModuleBayMapOutput() DeviceModuleBayMapOutput {
	return i.ToDeviceModuleBayMapOutputWithContext(context.Background())
}

func (i DeviceModuleBayMap) ToDeviceModuleBayMapOutputWithContext(ctx context.Context) DeviceModuleBayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceModuleBayMapOutput)
}

type DeviceModuleBayOutput struct{ *pulumi.OutputState }

func (DeviceModuleBayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceModuleBay)(nil)).Elem()
}

func (o DeviceModuleBayOutput) ToDeviceModuleBayOutput() DeviceModuleBayOutput {
	return o
}

func (o DeviceModuleBayOutput) ToDeviceModuleBayOutputWithContext(ctx context.Context) DeviceModuleBayOutput {
	return o
}

func (o DeviceModuleBayOutput) CustomFields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeviceModuleBay) pulumi.StringMapOutput { return v.CustomFields }).(pulumi.StringMapOutput)
}

func (o DeviceModuleBayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceModuleBay) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DeviceModuleBayOutput) DeviceId() pulumi.IntOutput {
	return o.ApplyT(func(v *DeviceModuleBay) pulumi.IntOutput { return v.DeviceId }).(pulumi.IntOutput)
}

func (o DeviceModuleBayOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceModuleBay) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

func (o DeviceModuleBayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceModuleBay) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeviceModuleBayOutput) Position() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceModuleBay) pulumi.StringPtrOutput { return v.Position }).(pulumi.StringPtrOutput)
}

func (o DeviceModuleBayOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DeviceModuleBay) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

type DeviceModuleBayArrayOutput struct{ *pulumi.OutputState }

func (DeviceModuleBayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceModuleBay)(nil)).Elem()
}

func (o DeviceModuleBayArrayOutput) ToDeviceModuleBayArrayOutput() DeviceModuleBayArrayOutput {
	return o
}

func (o DeviceModuleBayArrayOutput) ToDeviceModuleBayArrayOutputWithContext(ctx context.Context) DeviceModuleBayArrayOutput {
	return o
}

func (o DeviceModuleBayArrayOutput) Index(i pulumi.IntInput) DeviceModuleBayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeviceModuleBay {
		return vs[0].([]*DeviceModuleBay)[vs[1].(int)]
	}).(DeviceModuleBayOutput)
}

type DeviceModuleBayMapOutput struct{ *pulumi.OutputState }

func (DeviceModuleBayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceModuleBay)(nil)).Elem()
}

func (o DeviceModuleBayMapOutput) ToDeviceModuleBayMapOutput() DeviceModuleBayMapOutput {
	return o
}

func (o DeviceModuleBayMapOutput) ToDeviceModuleBayMapOutputWithContext(ctx context.Context) DeviceModuleBayMapOutput {
	return o
}

func (o DeviceModuleBayMapOutput) MapIndex(k pulumi.StringInput) DeviceModuleBayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeviceModuleBay {
		return vs[0].(map[string]*DeviceModuleBay)[vs[1].(string)]
	}).(DeviceModuleBayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceModuleBayInput)(nil)).Elem(), &DeviceModuleBay{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceModuleBayArrayInput)(nil)).Elem(), DeviceModuleBayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceModuleBayMapInput)(nil)).Elem(), DeviceModuleBayMap{})
	pulumi.RegisterOutputType(DeviceModuleBayOutput{})
	pulumi.RegisterOutputType(DeviceModuleBayArrayOutput{})
	pulumi.RegisterOutputType(DeviceModuleBayMapOutput{})
}

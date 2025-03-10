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

// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/consoleport/):
//
// > A console port provides connectivity to the physical console of a device. These are typically used for temporary access by someone who is physically near the device, or for remote out-of-band access provided via a networked console server.
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
//			_, err = dcim.NewDeviceConsolePort(ctx, "test", &dcim.DeviceConsolePortArgs{
//				DeviceId:      test.ID(),
//				Name:          pulumi.String("console port"),
//				Type:          pulumi.String("de-9"),
//				Speed:         pulumi.Int(1200),
//				MarkConnected: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DeviceConsolePort struct {
	pulumi.CustomResourceState

	CustomFields pulumi.StringMapOutput `pulumi:"customFields"`
	Description  pulumi.StringPtrOutput `pulumi:"description"`
	DeviceId     pulumi.IntOutput       `pulumi:"deviceId"`
	Label        pulumi.StringPtrOutput `pulumi:"label"`
	// Defaults to `false`.
	MarkConnected pulumi.BoolPtrOutput `pulumi:"markConnected"`
	ModuleId      pulumi.IntPtrOutput  `pulumi:"moduleId"`
	Name          pulumi.StringOutput  `pulumi:"name"`
	// One of [1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200].
	Speed pulumi.IntPtrOutput      `pulumi:"speed"`
	Tags  pulumi.StringArrayOutput `pulumi:"tags"`
	// One of [de-9, db-25, rj-11, rj-12, rj-45, mini-din-8, usb-a, usb-b, usb-c, usb-mini-a, usb-mini-b, usb-micro-a, usb-micro-b, usb-micro-ab, other].
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewDeviceConsolePort registers a new resource with the given unique name, arguments, and options.
func NewDeviceConsolePort(ctx *pulumi.Context,
	name string, args *DeviceConsolePortArgs, opts ...pulumi.ResourceOption) (*DeviceConsolePort, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DeviceConsolePort
	err := ctx.RegisterResource("netbox:dcim/deviceConsolePort:DeviceConsolePort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceConsolePort gets an existing DeviceConsolePort resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceConsolePort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceConsolePortState, opts ...pulumi.ResourceOption) (*DeviceConsolePort, error) {
	var resource DeviceConsolePort
	err := ctx.ReadResource("netbox:dcim/deviceConsolePort:DeviceConsolePort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceConsolePort resources.
type deviceConsolePortState struct {
	CustomFields map[string]string `pulumi:"customFields"`
	Description  *string           `pulumi:"description"`
	DeviceId     *int              `pulumi:"deviceId"`
	Label        *string           `pulumi:"label"`
	// Defaults to `false`.
	MarkConnected *bool   `pulumi:"markConnected"`
	ModuleId      *int    `pulumi:"moduleId"`
	Name          *string `pulumi:"name"`
	// One of [1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200].
	Speed *int     `pulumi:"speed"`
	Tags  []string `pulumi:"tags"`
	// One of [de-9, db-25, rj-11, rj-12, rj-45, mini-din-8, usb-a, usb-b, usb-c, usb-mini-a, usb-mini-b, usb-micro-a, usb-micro-b, usb-micro-ab, other].
	Type *string `pulumi:"type"`
}

type DeviceConsolePortState struct {
	CustomFields pulumi.StringMapInput
	Description  pulumi.StringPtrInput
	DeviceId     pulumi.IntPtrInput
	Label        pulumi.StringPtrInput
	// Defaults to `false`.
	MarkConnected pulumi.BoolPtrInput
	ModuleId      pulumi.IntPtrInput
	Name          pulumi.StringPtrInput
	// One of [1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200].
	Speed pulumi.IntPtrInput
	Tags  pulumi.StringArrayInput
	// One of [de-9, db-25, rj-11, rj-12, rj-45, mini-din-8, usb-a, usb-b, usb-c, usb-mini-a, usb-mini-b, usb-micro-a, usb-micro-b, usb-micro-ab, other].
	Type pulumi.StringPtrInput
}

func (DeviceConsolePortState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceConsolePortState)(nil)).Elem()
}

type deviceConsolePortArgs struct {
	CustomFields map[string]string `pulumi:"customFields"`
	Description  *string           `pulumi:"description"`
	DeviceId     int               `pulumi:"deviceId"`
	Label        *string           `pulumi:"label"`
	// Defaults to `false`.
	MarkConnected *bool   `pulumi:"markConnected"`
	ModuleId      *int    `pulumi:"moduleId"`
	Name          *string `pulumi:"name"`
	// One of [1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200].
	Speed *int     `pulumi:"speed"`
	Tags  []string `pulumi:"tags"`
	// One of [de-9, db-25, rj-11, rj-12, rj-45, mini-din-8, usb-a, usb-b, usb-c, usb-mini-a, usb-mini-b, usb-micro-a, usb-micro-b, usb-micro-ab, other].
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a DeviceConsolePort resource.
type DeviceConsolePortArgs struct {
	CustomFields pulumi.StringMapInput
	Description  pulumi.StringPtrInput
	DeviceId     pulumi.IntInput
	Label        pulumi.StringPtrInput
	// Defaults to `false`.
	MarkConnected pulumi.BoolPtrInput
	ModuleId      pulumi.IntPtrInput
	Name          pulumi.StringPtrInput
	// One of [1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200].
	Speed pulumi.IntPtrInput
	Tags  pulumi.StringArrayInput
	// One of [de-9, db-25, rj-11, rj-12, rj-45, mini-din-8, usb-a, usb-b, usb-c, usb-mini-a, usb-mini-b, usb-micro-a, usb-micro-b, usb-micro-ab, other].
	Type pulumi.StringPtrInput
}

func (DeviceConsolePortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceConsolePortArgs)(nil)).Elem()
}

type DeviceConsolePortInput interface {
	pulumi.Input

	ToDeviceConsolePortOutput() DeviceConsolePortOutput
	ToDeviceConsolePortOutputWithContext(ctx context.Context) DeviceConsolePortOutput
}

func (*DeviceConsolePort) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceConsolePort)(nil)).Elem()
}

func (i *DeviceConsolePort) ToDeviceConsolePortOutput() DeviceConsolePortOutput {
	return i.ToDeviceConsolePortOutputWithContext(context.Background())
}

func (i *DeviceConsolePort) ToDeviceConsolePortOutputWithContext(ctx context.Context) DeviceConsolePortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceConsolePortOutput)
}

// DeviceConsolePortArrayInput is an input type that accepts DeviceConsolePortArray and DeviceConsolePortArrayOutput values.
// You can construct a concrete instance of `DeviceConsolePortArrayInput` via:
//
//	DeviceConsolePortArray{ DeviceConsolePortArgs{...} }
type DeviceConsolePortArrayInput interface {
	pulumi.Input

	ToDeviceConsolePortArrayOutput() DeviceConsolePortArrayOutput
	ToDeviceConsolePortArrayOutputWithContext(context.Context) DeviceConsolePortArrayOutput
}

type DeviceConsolePortArray []DeviceConsolePortInput

func (DeviceConsolePortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceConsolePort)(nil)).Elem()
}

func (i DeviceConsolePortArray) ToDeviceConsolePortArrayOutput() DeviceConsolePortArrayOutput {
	return i.ToDeviceConsolePortArrayOutputWithContext(context.Background())
}

func (i DeviceConsolePortArray) ToDeviceConsolePortArrayOutputWithContext(ctx context.Context) DeviceConsolePortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceConsolePortArrayOutput)
}

// DeviceConsolePortMapInput is an input type that accepts DeviceConsolePortMap and DeviceConsolePortMapOutput values.
// You can construct a concrete instance of `DeviceConsolePortMapInput` via:
//
//	DeviceConsolePortMap{ "key": DeviceConsolePortArgs{...} }
type DeviceConsolePortMapInput interface {
	pulumi.Input

	ToDeviceConsolePortMapOutput() DeviceConsolePortMapOutput
	ToDeviceConsolePortMapOutputWithContext(context.Context) DeviceConsolePortMapOutput
}

type DeviceConsolePortMap map[string]DeviceConsolePortInput

func (DeviceConsolePortMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceConsolePort)(nil)).Elem()
}

func (i DeviceConsolePortMap) ToDeviceConsolePortMapOutput() DeviceConsolePortMapOutput {
	return i.ToDeviceConsolePortMapOutputWithContext(context.Background())
}

func (i DeviceConsolePortMap) ToDeviceConsolePortMapOutputWithContext(ctx context.Context) DeviceConsolePortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceConsolePortMapOutput)
}

type DeviceConsolePortOutput struct{ *pulumi.OutputState }

func (DeviceConsolePortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceConsolePort)(nil)).Elem()
}

func (o DeviceConsolePortOutput) ToDeviceConsolePortOutput() DeviceConsolePortOutput {
	return o
}

func (o DeviceConsolePortOutput) ToDeviceConsolePortOutputWithContext(ctx context.Context) DeviceConsolePortOutput {
	return o
}

func (o DeviceConsolePortOutput) CustomFields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeviceConsolePort) pulumi.StringMapOutput { return v.CustomFields }).(pulumi.StringMapOutput)
}

func (o DeviceConsolePortOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceConsolePort) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DeviceConsolePortOutput) DeviceId() pulumi.IntOutput {
	return o.ApplyT(func(v *DeviceConsolePort) pulumi.IntOutput { return v.DeviceId }).(pulumi.IntOutput)
}

func (o DeviceConsolePortOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceConsolePort) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

// Defaults to `false`.
func (o DeviceConsolePortOutput) MarkConnected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DeviceConsolePort) pulumi.BoolPtrOutput { return v.MarkConnected }).(pulumi.BoolPtrOutput)
}

func (o DeviceConsolePortOutput) ModuleId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeviceConsolePort) pulumi.IntPtrOutput { return v.ModuleId }).(pulumi.IntPtrOutput)
}

func (o DeviceConsolePortOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceConsolePort) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// One of [1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200].
func (o DeviceConsolePortOutput) Speed() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeviceConsolePort) pulumi.IntPtrOutput { return v.Speed }).(pulumi.IntPtrOutput)
}

func (o DeviceConsolePortOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DeviceConsolePort) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// One of [de-9, db-25, rj-11, rj-12, rj-45, mini-din-8, usb-a, usb-b, usb-c, usb-mini-a, usb-mini-b, usb-micro-a, usb-micro-b, usb-micro-ab, other].
func (o DeviceConsolePortOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceConsolePort) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type DeviceConsolePortArrayOutput struct{ *pulumi.OutputState }

func (DeviceConsolePortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceConsolePort)(nil)).Elem()
}

func (o DeviceConsolePortArrayOutput) ToDeviceConsolePortArrayOutput() DeviceConsolePortArrayOutput {
	return o
}

func (o DeviceConsolePortArrayOutput) ToDeviceConsolePortArrayOutputWithContext(ctx context.Context) DeviceConsolePortArrayOutput {
	return o
}

func (o DeviceConsolePortArrayOutput) Index(i pulumi.IntInput) DeviceConsolePortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeviceConsolePort {
		return vs[0].([]*DeviceConsolePort)[vs[1].(int)]
	}).(DeviceConsolePortOutput)
}

type DeviceConsolePortMapOutput struct{ *pulumi.OutputState }

func (DeviceConsolePortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceConsolePort)(nil)).Elem()
}

func (o DeviceConsolePortMapOutput) ToDeviceConsolePortMapOutput() DeviceConsolePortMapOutput {
	return o
}

func (o DeviceConsolePortMapOutput) ToDeviceConsolePortMapOutputWithContext(ctx context.Context) DeviceConsolePortMapOutput {
	return o
}

func (o DeviceConsolePortMapOutput) MapIndex(k pulumi.StringInput) DeviceConsolePortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeviceConsolePort {
		return vs[0].(map[string]*DeviceConsolePort)[vs[1].(string)]
	}).(DeviceConsolePortOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceConsolePortInput)(nil)).Elem(), &DeviceConsolePort{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceConsolePortArrayInput)(nil)).Elem(), DeviceConsolePortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceConsolePortMapInput)(nil)).Elem(), DeviceConsolePortMap{})
	pulumi.RegisterOutputType(DeviceConsolePortOutput{})
	pulumi.RegisterOutputType(DeviceConsolePortArrayOutput{})
	pulumi.RegisterOutputType(DeviceConsolePortMapOutput{})
}

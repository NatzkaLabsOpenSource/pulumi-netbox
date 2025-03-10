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

// From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/poweroutlet/):
//
// > Power outlets represent the outlets on a power distribution unit (PDU) or other device that supplies power to dependent devices. Each power port may be assigned a physical type, and may be associated with a specific feed leg (where three-phase power is used) and/or a specific upstream power port. This association can be used to model the distribution of power within a device.
//
// For example, imagine a PDU with one power port which draws from a three-phase feed and 48 power outlets arranged into three banks of 16 outlets each. Outlets 1-16 would be associated with leg A on the port, and outlets 17-32 and 33-48 would be associated with legs B and C, respectively.
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
//			_, err = dcim.NewDevicePowerOutlet(ctx, "test", &dcim.DevicePowerOutletArgs{
//				DeviceId: test.ID(),
//				Name:     pulumi.String("power outlet"),
//				Type:     pulumi.String("iec-60320-c5"),
//				FeedLeg:  pulumi.String("A"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DevicePowerOutlet struct {
	pulumi.CustomResourceState

	CustomFields pulumi.StringMapOutput `pulumi:"customFields"`
	Description  pulumi.StringPtrOutput `pulumi:"description"`
	DeviceId     pulumi.IntOutput       `pulumi:"deviceId"`
	// One of [A, B, C].
	FeedLeg pulumi.StringPtrOutput `pulumi:"feedLeg"`
	Label   pulumi.StringPtrOutput `pulumi:"label"`
	// Defaults to `false`.
	MarkConnected pulumi.BoolPtrOutput     `pulumi:"markConnected"`
	ModuleId      pulumi.IntPtrOutput      `pulumi:"moduleId"`
	Name          pulumi.StringOutput      `pulumi:"name"`
	PowerPortId   pulumi.IntPtrOutput      `pulumi:"powerPortId"`
	Tags          pulumi.StringArrayOutput `pulumi:"tags"`
	// One of [iec-60320-c5, iec-60320-c7, iec-60320-c13, iec-60320-c15, iec-60320-c19, iec-60320-c21, iec-60309-p-n-e-4h, iec-60309-p-n-e-6h, iec-60309-p-n-e-9h, iec-60309-2p-e-4h, iec-60309-2p-e-6h, iec-60309-2p-e-9h, iec-60309-3p-e-4h, iec-60309-3p-e-6h, iec-60309-3p-e-9h, iec-60309-3p-n-e-4h, iec-60309-3p-n-e-6h, iec-60309-3p-n-e-9h, nema-1-15r, nema-5-15r, nema-5-20r, nema-5-30r, nema-5-50r, nema-6-15r, nema-6-20r, nema-6-30r, nema-6-50r, nema-10-30r, nema-10-50r, nema-14-20r, nema-14-30r, nema-14-50r, nema-14-60r, nema-15-15r, nema-15-20r, nema-15-30r, nema-15-50r, nema-15-60r, nema-l1-15r, nema-l5-15r, nema-l5-20r, nema-l5-30r, nema-l5-50r, nema-l6-15r, nema-l6-20r, nema-l6-30r, nema-l6-50r, nema-l10-30r, nema-l14-20r, nema-l14-30r, nema-l14-50r, nema-l14-60r, nema-l15-20r, nema-l15-30r, nema-l15-50r, nema-l15-60r, nema-l21-20r, nema-l21-30r, nema-l22-30r, CS6360C, CS6364C, CS8164C, CS8264C, CS8364C, CS8464C, ita-e, ita-f, ita-g, ita-h, ita-i, ita-j, ita-k, ita-l, ita-m, ita-n, ita-o, ita-multistandard, usb-a, usb-micro-b, usb-c, dc-terminal, hdot-cx, saf-d-grid, neutrik-powercon-20a, neutrik-powercon-32a, neutrik-powercon-true1, neutrik-powercon-true1-top, ubiquiti-smartpower, hardwired, other].
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewDevicePowerOutlet registers a new resource with the given unique name, arguments, and options.
func NewDevicePowerOutlet(ctx *pulumi.Context,
	name string, args *DevicePowerOutletArgs, opts ...pulumi.ResourceOption) (*DevicePowerOutlet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DevicePowerOutlet
	err := ctx.RegisterResource("netbox:dcim/devicePowerOutlet:DevicePowerOutlet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevicePowerOutlet gets an existing DevicePowerOutlet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevicePowerOutlet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevicePowerOutletState, opts ...pulumi.ResourceOption) (*DevicePowerOutlet, error) {
	var resource DevicePowerOutlet
	err := ctx.ReadResource("netbox:dcim/devicePowerOutlet:DevicePowerOutlet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DevicePowerOutlet resources.
type devicePowerOutletState struct {
	CustomFields map[string]string `pulumi:"customFields"`
	Description  *string           `pulumi:"description"`
	DeviceId     *int              `pulumi:"deviceId"`
	// One of [A, B, C].
	FeedLeg *string `pulumi:"feedLeg"`
	Label   *string `pulumi:"label"`
	// Defaults to `false`.
	MarkConnected *bool    `pulumi:"markConnected"`
	ModuleId      *int     `pulumi:"moduleId"`
	Name          *string  `pulumi:"name"`
	PowerPortId   *int     `pulumi:"powerPortId"`
	Tags          []string `pulumi:"tags"`
	// One of [iec-60320-c5, iec-60320-c7, iec-60320-c13, iec-60320-c15, iec-60320-c19, iec-60320-c21, iec-60309-p-n-e-4h, iec-60309-p-n-e-6h, iec-60309-p-n-e-9h, iec-60309-2p-e-4h, iec-60309-2p-e-6h, iec-60309-2p-e-9h, iec-60309-3p-e-4h, iec-60309-3p-e-6h, iec-60309-3p-e-9h, iec-60309-3p-n-e-4h, iec-60309-3p-n-e-6h, iec-60309-3p-n-e-9h, nema-1-15r, nema-5-15r, nema-5-20r, nema-5-30r, nema-5-50r, nema-6-15r, nema-6-20r, nema-6-30r, nema-6-50r, nema-10-30r, nema-10-50r, nema-14-20r, nema-14-30r, nema-14-50r, nema-14-60r, nema-15-15r, nema-15-20r, nema-15-30r, nema-15-50r, nema-15-60r, nema-l1-15r, nema-l5-15r, nema-l5-20r, nema-l5-30r, nema-l5-50r, nema-l6-15r, nema-l6-20r, nema-l6-30r, nema-l6-50r, nema-l10-30r, nema-l14-20r, nema-l14-30r, nema-l14-50r, nema-l14-60r, nema-l15-20r, nema-l15-30r, nema-l15-50r, nema-l15-60r, nema-l21-20r, nema-l21-30r, nema-l22-30r, CS6360C, CS6364C, CS8164C, CS8264C, CS8364C, CS8464C, ita-e, ita-f, ita-g, ita-h, ita-i, ita-j, ita-k, ita-l, ita-m, ita-n, ita-o, ita-multistandard, usb-a, usb-micro-b, usb-c, dc-terminal, hdot-cx, saf-d-grid, neutrik-powercon-20a, neutrik-powercon-32a, neutrik-powercon-true1, neutrik-powercon-true1-top, ubiquiti-smartpower, hardwired, other].
	Type *string `pulumi:"type"`
}

type DevicePowerOutletState struct {
	CustomFields pulumi.StringMapInput
	Description  pulumi.StringPtrInput
	DeviceId     pulumi.IntPtrInput
	// One of [A, B, C].
	FeedLeg pulumi.StringPtrInput
	Label   pulumi.StringPtrInput
	// Defaults to `false`.
	MarkConnected pulumi.BoolPtrInput
	ModuleId      pulumi.IntPtrInput
	Name          pulumi.StringPtrInput
	PowerPortId   pulumi.IntPtrInput
	Tags          pulumi.StringArrayInput
	// One of [iec-60320-c5, iec-60320-c7, iec-60320-c13, iec-60320-c15, iec-60320-c19, iec-60320-c21, iec-60309-p-n-e-4h, iec-60309-p-n-e-6h, iec-60309-p-n-e-9h, iec-60309-2p-e-4h, iec-60309-2p-e-6h, iec-60309-2p-e-9h, iec-60309-3p-e-4h, iec-60309-3p-e-6h, iec-60309-3p-e-9h, iec-60309-3p-n-e-4h, iec-60309-3p-n-e-6h, iec-60309-3p-n-e-9h, nema-1-15r, nema-5-15r, nema-5-20r, nema-5-30r, nema-5-50r, nema-6-15r, nema-6-20r, nema-6-30r, nema-6-50r, nema-10-30r, nema-10-50r, nema-14-20r, nema-14-30r, nema-14-50r, nema-14-60r, nema-15-15r, nema-15-20r, nema-15-30r, nema-15-50r, nema-15-60r, nema-l1-15r, nema-l5-15r, nema-l5-20r, nema-l5-30r, nema-l5-50r, nema-l6-15r, nema-l6-20r, nema-l6-30r, nema-l6-50r, nema-l10-30r, nema-l14-20r, nema-l14-30r, nema-l14-50r, nema-l14-60r, nema-l15-20r, nema-l15-30r, nema-l15-50r, nema-l15-60r, nema-l21-20r, nema-l21-30r, nema-l22-30r, CS6360C, CS6364C, CS8164C, CS8264C, CS8364C, CS8464C, ita-e, ita-f, ita-g, ita-h, ita-i, ita-j, ita-k, ita-l, ita-m, ita-n, ita-o, ita-multistandard, usb-a, usb-micro-b, usb-c, dc-terminal, hdot-cx, saf-d-grid, neutrik-powercon-20a, neutrik-powercon-32a, neutrik-powercon-true1, neutrik-powercon-true1-top, ubiquiti-smartpower, hardwired, other].
	Type pulumi.StringPtrInput
}

func (DevicePowerOutletState) ElementType() reflect.Type {
	return reflect.TypeOf((*devicePowerOutletState)(nil)).Elem()
}

type devicePowerOutletArgs struct {
	CustomFields map[string]string `pulumi:"customFields"`
	Description  *string           `pulumi:"description"`
	DeviceId     int               `pulumi:"deviceId"`
	// One of [A, B, C].
	FeedLeg *string `pulumi:"feedLeg"`
	Label   *string `pulumi:"label"`
	// Defaults to `false`.
	MarkConnected *bool    `pulumi:"markConnected"`
	ModuleId      *int     `pulumi:"moduleId"`
	Name          *string  `pulumi:"name"`
	PowerPortId   *int     `pulumi:"powerPortId"`
	Tags          []string `pulumi:"tags"`
	// One of [iec-60320-c5, iec-60320-c7, iec-60320-c13, iec-60320-c15, iec-60320-c19, iec-60320-c21, iec-60309-p-n-e-4h, iec-60309-p-n-e-6h, iec-60309-p-n-e-9h, iec-60309-2p-e-4h, iec-60309-2p-e-6h, iec-60309-2p-e-9h, iec-60309-3p-e-4h, iec-60309-3p-e-6h, iec-60309-3p-e-9h, iec-60309-3p-n-e-4h, iec-60309-3p-n-e-6h, iec-60309-3p-n-e-9h, nema-1-15r, nema-5-15r, nema-5-20r, nema-5-30r, nema-5-50r, nema-6-15r, nema-6-20r, nema-6-30r, nema-6-50r, nema-10-30r, nema-10-50r, nema-14-20r, nema-14-30r, nema-14-50r, nema-14-60r, nema-15-15r, nema-15-20r, nema-15-30r, nema-15-50r, nema-15-60r, nema-l1-15r, nema-l5-15r, nema-l5-20r, nema-l5-30r, nema-l5-50r, nema-l6-15r, nema-l6-20r, nema-l6-30r, nema-l6-50r, nema-l10-30r, nema-l14-20r, nema-l14-30r, nema-l14-50r, nema-l14-60r, nema-l15-20r, nema-l15-30r, nema-l15-50r, nema-l15-60r, nema-l21-20r, nema-l21-30r, nema-l22-30r, CS6360C, CS6364C, CS8164C, CS8264C, CS8364C, CS8464C, ita-e, ita-f, ita-g, ita-h, ita-i, ita-j, ita-k, ita-l, ita-m, ita-n, ita-o, ita-multistandard, usb-a, usb-micro-b, usb-c, dc-terminal, hdot-cx, saf-d-grid, neutrik-powercon-20a, neutrik-powercon-32a, neutrik-powercon-true1, neutrik-powercon-true1-top, ubiquiti-smartpower, hardwired, other].
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a DevicePowerOutlet resource.
type DevicePowerOutletArgs struct {
	CustomFields pulumi.StringMapInput
	Description  pulumi.StringPtrInput
	DeviceId     pulumi.IntInput
	// One of [A, B, C].
	FeedLeg pulumi.StringPtrInput
	Label   pulumi.StringPtrInput
	// Defaults to `false`.
	MarkConnected pulumi.BoolPtrInput
	ModuleId      pulumi.IntPtrInput
	Name          pulumi.StringPtrInput
	PowerPortId   pulumi.IntPtrInput
	Tags          pulumi.StringArrayInput
	// One of [iec-60320-c5, iec-60320-c7, iec-60320-c13, iec-60320-c15, iec-60320-c19, iec-60320-c21, iec-60309-p-n-e-4h, iec-60309-p-n-e-6h, iec-60309-p-n-e-9h, iec-60309-2p-e-4h, iec-60309-2p-e-6h, iec-60309-2p-e-9h, iec-60309-3p-e-4h, iec-60309-3p-e-6h, iec-60309-3p-e-9h, iec-60309-3p-n-e-4h, iec-60309-3p-n-e-6h, iec-60309-3p-n-e-9h, nema-1-15r, nema-5-15r, nema-5-20r, nema-5-30r, nema-5-50r, nema-6-15r, nema-6-20r, nema-6-30r, nema-6-50r, nema-10-30r, nema-10-50r, nema-14-20r, nema-14-30r, nema-14-50r, nema-14-60r, nema-15-15r, nema-15-20r, nema-15-30r, nema-15-50r, nema-15-60r, nema-l1-15r, nema-l5-15r, nema-l5-20r, nema-l5-30r, nema-l5-50r, nema-l6-15r, nema-l6-20r, nema-l6-30r, nema-l6-50r, nema-l10-30r, nema-l14-20r, nema-l14-30r, nema-l14-50r, nema-l14-60r, nema-l15-20r, nema-l15-30r, nema-l15-50r, nema-l15-60r, nema-l21-20r, nema-l21-30r, nema-l22-30r, CS6360C, CS6364C, CS8164C, CS8264C, CS8364C, CS8464C, ita-e, ita-f, ita-g, ita-h, ita-i, ita-j, ita-k, ita-l, ita-m, ita-n, ita-o, ita-multistandard, usb-a, usb-micro-b, usb-c, dc-terminal, hdot-cx, saf-d-grid, neutrik-powercon-20a, neutrik-powercon-32a, neutrik-powercon-true1, neutrik-powercon-true1-top, ubiquiti-smartpower, hardwired, other].
	Type pulumi.StringPtrInput
}

func (DevicePowerOutletArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devicePowerOutletArgs)(nil)).Elem()
}

type DevicePowerOutletInput interface {
	pulumi.Input

	ToDevicePowerOutletOutput() DevicePowerOutletOutput
	ToDevicePowerOutletOutputWithContext(ctx context.Context) DevicePowerOutletOutput
}

func (*DevicePowerOutlet) ElementType() reflect.Type {
	return reflect.TypeOf((**DevicePowerOutlet)(nil)).Elem()
}

func (i *DevicePowerOutlet) ToDevicePowerOutletOutput() DevicePowerOutletOutput {
	return i.ToDevicePowerOutletOutputWithContext(context.Background())
}

func (i *DevicePowerOutlet) ToDevicePowerOutletOutputWithContext(ctx context.Context) DevicePowerOutletOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicePowerOutletOutput)
}

// DevicePowerOutletArrayInput is an input type that accepts DevicePowerOutletArray and DevicePowerOutletArrayOutput values.
// You can construct a concrete instance of `DevicePowerOutletArrayInput` via:
//
//	DevicePowerOutletArray{ DevicePowerOutletArgs{...} }
type DevicePowerOutletArrayInput interface {
	pulumi.Input

	ToDevicePowerOutletArrayOutput() DevicePowerOutletArrayOutput
	ToDevicePowerOutletArrayOutputWithContext(context.Context) DevicePowerOutletArrayOutput
}

type DevicePowerOutletArray []DevicePowerOutletInput

func (DevicePowerOutletArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DevicePowerOutlet)(nil)).Elem()
}

func (i DevicePowerOutletArray) ToDevicePowerOutletArrayOutput() DevicePowerOutletArrayOutput {
	return i.ToDevicePowerOutletArrayOutputWithContext(context.Background())
}

func (i DevicePowerOutletArray) ToDevicePowerOutletArrayOutputWithContext(ctx context.Context) DevicePowerOutletArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicePowerOutletArrayOutput)
}

// DevicePowerOutletMapInput is an input type that accepts DevicePowerOutletMap and DevicePowerOutletMapOutput values.
// You can construct a concrete instance of `DevicePowerOutletMapInput` via:
//
//	DevicePowerOutletMap{ "key": DevicePowerOutletArgs{...} }
type DevicePowerOutletMapInput interface {
	pulumi.Input

	ToDevicePowerOutletMapOutput() DevicePowerOutletMapOutput
	ToDevicePowerOutletMapOutputWithContext(context.Context) DevicePowerOutletMapOutput
}

type DevicePowerOutletMap map[string]DevicePowerOutletInput

func (DevicePowerOutletMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DevicePowerOutlet)(nil)).Elem()
}

func (i DevicePowerOutletMap) ToDevicePowerOutletMapOutput() DevicePowerOutletMapOutput {
	return i.ToDevicePowerOutletMapOutputWithContext(context.Background())
}

func (i DevicePowerOutletMap) ToDevicePowerOutletMapOutputWithContext(ctx context.Context) DevicePowerOutletMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicePowerOutletMapOutput)
}

type DevicePowerOutletOutput struct{ *pulumi.OutputState }

func (DevicePowerOutletOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevicePowerOutlet)(nil)).Elem()
}

func (o DevicePowerOutletOutput) ToDevicePowerOutletOutput() DevicePowerOutletOutput {
	return o
}

func (o DevicePowerOutletOutput) ToDevicePowerOutletOutputWithContext(ctx context.Context) DevicePowerOutletOutput {
	return o
}

func (o DevicePowerOutletOutput) CustomFields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DevicePowerOutlet) pulumi.StringMapOutput { return v.CustomFields }).(pulumi.StringMapOutput)
}

func (o DevicePowerOutletOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevicePowerOutlet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DevicePowerOutletOutput) DeviceId() pulumi.IntOutput {
	return o.ApplyT(func(v *DevicePowerOutlet) pulumi.IntOutput { return v.DeviceId }).(pulumi.IntOutput)
}

// One of [A, B, C].
func (o DevicePowerOutletOutput) FeedLeg() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevicePowerOutlet) pulumi.StringPtrOutput { return v.FeedLeg }).(pulumi.StringPtrOutput)
}

func (o DevicePowerOutletOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevicePowerOutlet) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

// Defaults to `false`.
func (o DevicePowerOutletOutput) MarkConnected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DevicePowerOutlet) pulumi.BoolPtrOutput { return v.MarkConnected }).(pulumi.BoolPtrOutput)
}

func (o DevicePowerOutletOutput) ModuleId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DevicePowerOutlet) pulumi.IntPtrOutput { return v.ModuleId }).(pulumi.IntPtrOutput)
}

func (o DevicePowerOutletOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DevicePowerOutlet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DevicePowerOutletOutput) PowerPortId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DevicePowerOutlet) pulumi.IntPtrOutput { return v.PowerPortId }).(pulumi.IntPtrOutput)
}

func (o DevicePowerOutletOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DevicePowerOutlet) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// One of [iec-60320-c5, iec-60320-c7, iec-60320-c13, iec-60320-c15, iec-60320-c19, iec-60320-c21, iec-60309-p-n-e-4h, iec-60309-p-n-e-6h, iec-60309-p-n-e-9h, iec-60309-2p-e-4h, iec-60309-2p-e-6h, iec-60309-2p-e-9h, iec-60309-3p-e-4h, iec-60309-3p-e-6h, iec-60309-3p-e-9h, iec-60309-3p-n-e-4h, iec-60309-3p-n-e-6h, iec-60309-3p-n-e-9h, nema-1-15r, nema-5-15r, nema-5-20r, nema-5-30r, nema-5-50r, nema-6-15r, nema-6-20r, nema-6-30r, nema-6-50r, nema-10-30r, nema-10-50r, nema-14-20r, nema-14-30r, nema-14-50r, nema-14-60r, nema-15-15r, nema-15-20r, nema-15-30r, nema-15-50r, nema-15-60r, nema-l1-15r, nema-l5-15r, nema-l5-20r, nema-l5-30r, nema-l5-50r, nema-l6-15r, nema-l6-20r, nema-l6-30r, nema-l6-50r, nema-l10-30r, nema-l14-20r, nema-l14-30r, nema-l14-50r, nema-l14-60r, nema-l15-20r, nema-l15-30r, nema-l15-50r, nema-l15-60r, nema-l21-20r, nema-l21-30r, nema-l22-30r, CS6360C, CS6364C, CS8164C, CS8264C, CS8364C, CS8464C, ita-e, ita-f, ita-g, ita-h, ita-i, ita-j, ita-k, ita-l, ita-m, ita-n, ita-o, ita-multistandard, usb-a, usb-micro-b, usb-c, dc-terminal, hdot-cx, saf-d-grid, neutrik-powercon-20a, neutrik-powercon-32a, neutrik-powercon-true1, neutrik-powercon-true1-top, ubiquiti-smartpower, hardwired, other].
func (o DevicePowerOutletOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevicePowerOutlet) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type DevicePowerOutletArrayOutput struct{ *pulumi.OutputState }

func (DevicePowerOutletArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DevicePowerOutlet)(nil)).Elem()
}

func (o DevicePowerOutletArrayOutput) ToDevicePowerOutletArrayOutput() DevicePowerOutletArrayOutput {
	return o
}

func (o DevicePowerOutletArrayOutput) ToDevicePowerOutletArrayOutputWithContext(ctx context.Context) DevicePowerOutletArrayOutput {
	return o
}

func (o DevicePowerOutletArrayOutput) Index(i pulumi.IntInput) DevicePowerOutletOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DevicePowerOutlet {
		return vs[0].([]*DevicePowerOutlet)[vs[1].(int)]
	}).(DevicePowerOutletOutput)
}

type DevicePowerOutletMapOutput struct{ *pulumi.OutputState }

func (DevicePowerOutletMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DevicePowerOutlet)(nil)).Elem()
}

func (o DevicePowerOutletMapOutput) ToDevicePowerOutletMapOutput() DevicePowerOutletMapOutput {
	return o
}

func (o DevicePowerOutletMapOutput) ToDevicePowerOutletMapOutputWithContext(ctx context.Context) DevicePowerOutletMapOutput {
	return o
}

func (o DevicePowerOutletMapOutput) MapIndex(k pulumi.StringInput) DevicePowerOutletOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DevicePowerOutlet {
		return vs[0].(map[string]*DevicePowerOutlet)[vs[1].(string)]
	}).(DevicePowerOutletOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DevicePowerOutletInput)(nil)).Elem(), &DevicePowerOutlet{})
	pulumi.RegisterInputType(reflect.TypeOf((*DevicePowerOutletArrayInput)(nil)).Elem(), DevicePowerOutletArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DevicePowerOutletMapInput)(nil)).Elem(), DevicePowerOutletMap{})
	pulumi.RegisterOutputType(DevicePowerOutletOutput{})
	pulumi.RegisterOutputType(DevicePowerOutletArrayOutput{})
	pulumi.RegisterOutputType(DevicePowerOutletMapOutput{})
}

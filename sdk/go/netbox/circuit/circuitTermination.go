// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package circuit

import (
	"context"
	"reflect"

	"errors"
	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/features/circuits/#circuit-terminations):
//
// > The association of a circuit with a particular site and/or device is modeled separately as a circuit termination. A circuit may have up to two terminations, labeled A and Z. A single-termination circuit can be used when you don't know (or care) about the far end of a circuit (for example, an Internet access circuit which connects to a transit provider). A dual-termination circuit is useful for tracking circuits which connect two sites.
// >
// > Each circuit termination is attached to either a site or to a provider network. Site terminations may optionally be connected via a cable to a specific device interface or port within that site. Each termination must be assigned a port speed, and can optionally be assigned an upstream speed if it differs from the downstream speed (a common scenario with e.g. DOCSIS cable modems). Fields are also available to track cross-connect and patch panel details.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/circuit"
//	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/dcim"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testSite, err := dcim.NewSite(ctx, "testSite", &dcim.SiteArgs{
//				Status: pulumi.String("active"),
//			})
//			if err != nil {
//				return err
//			}
//			testCircuitProvider, err := circuit.NewCircuitProvider(ctx, "testCircuitProvider", nil)
//			if err != nil {
//				return err
//			}
//			testCircuitType, err := circuit.NewCircuitType(ctx, "testCircuitType", nil)
//			if err != nil {
//				return err
//			}
//			testCircuit, err := circuit.NewCircuit(ctx, "testCircuit", &circuit.CircuitArgs{
//				Cid:        pulumi.String("%[1]s"),
//				Status:     pulumi.String("active"),
//				ProviderId: testCircuitProvider.ID(),
//				TypeId:     testCircuitType.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = circuit.NewCircuitTermination(ctx, "testCircuitTermination", &circuit.CircuitTerminationArgs{
//				CircuitId:     testCircuit.ID(),
//				TermSide:      pulumi.String("A"),
//				SiteId:        testSite.ID(),
//				PortSpeed:     pulumi.Int(100000),
//				UpstreamSpeed: pulumi.Int(50000),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type CircuitTermination struct {
	pulumi.CustomResourceState

	CircuitId    pulumi.IntOutput         `pulumi:"circuitId"`
	CustomFields pulumi.StringMapOutput   `pulumi:"customFields"`
	PortSpeed    pulumi.IntPtrOutput      `pulumi:"portSpeed"`
	SiteId       pulumi.IntOutput         `pulumi:"siteId"`
	Tags         pulumi.StringArrayOutput `pulumi:"tags"`
	// Valid values are `A` and `Z`.
	TermSide      pulumi.StringOutput `pulumi:"termSide"`
	UpstreamSpeed pulumi.IntPtrOutput `pulumi:"upstreamSpeed"`
}

// NewCircuitTermination registers a new resource with the given unique name, arguments, and options.
func NewCircuitTermination(ctx *pulumi.Context,
	name string, args *CircuitTerminationArgs, opts ...pulumi.ResourceOption) (*CircuitTermination, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CircuitId == nil {
		return nil, errors.New("invalid value for required argument 'CircuitId'")
	}
	if args.SiteId == nil {
		return nil, errors.New("invalid value for required argument 'SiteId'")
	}
	if args.TermSide == nil {
		return nil, errors.New("invalid value for required argument 'TermSide'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CircuitTermination
	err := ctx.RegisterResource("netbox:circuit/circuitTermination:CircuitTermination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCircuitTermination gets an existing CircuitTermination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCircuitTermination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CircuitTerminationState, opts ...pulumi.ResourceOption) (*CircuitTermination, error) {
	var resource CircuitTermination
	err := ctx.ReadResource("netbox:circuit/circuitTermination:CircuitTermination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CircuitTermination resources.
type circuitTerminationState struct {
	CircuitId    *int              `pulumi:"circuitId"`
	CustomFields map[string]string `pulumi:"customFields"`
	PortSpeed    *int              `pulumi:"portSpeed"`
	SiteId       *int              `pulumi:"siteId"`
	Tags         []string          `pulumi:"tags"`
	// Valid values are `A` and `Z`.
	TermSide      *string `pulumi:"termSide"`
	UpstreamSpeed *int    `pulumi:"upstreamSpeed"`
}

type CircuitTerminationState struct {
	CircuitId    pulumi.IntPtrInput
	CustomFields pulumi.StringMapInput
	PortSpeed    pulumi.IntPtrInput
	SiteId       pulumi.IntPtrInput
	Tags         pulumi.StringArrayInput
	// Valid values are `A` and `Z`.
	TermSide      pulumi.StringPtrInput
	UpstreamSpeed pulumi.IntPtrInput
}

func (CircuitTerminationState) ElementType() reflect.Type {
	return reflect.TypeOf((*circuitTerminationState)(nil)).Elem()
}

type circuitTerminationArgs struct {
	CircuitId    int               `pulumi:"circuitId"`
	CustomFields map[string]string `pulumi:"customFields"`
	PortSpeed    *int              `pulumi:"portSpeed"`
	SiteId       int               `pulumi:"siteId"`
	Tags         []string          `pulumi:"tags"`
	// Valid values are `A` and `Z`.
	TermSide      string `pulumi:"termSide"`
	UpstreamSpeed *int   `pulumi:"upstreamSpeed"`
}

// The set of arguments for constructing a CircuitTermination resource.
type CircuitTerminationArgs struct {
	CircuitId    pulumi.IntInput
	CustomFields pulumi.StringMapInput
	PortSpeed    pulumi.IntPtrInput
	SiteId       pulumi.IntInput
	Tags         pulumi.StringArrayInput
	// Valid values are `A` and `Z`.
	TermSide      pulumi.StringInput
	UpstreamSpeed pulumi.IntPtrInput
}

func (CircuitTerminationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*circuitTerminationArgs)(nil)).Elem()
}

type CircuitTerminationInput interface {
	pulumi.Input

	ToCircuitTerminationOutput() CircuitTerminationOutput
	ToCircuitTerminationOutputWithContext(ctx context.Context) CircuitTerminationOutput
}

func (*CircuitTermination) ElementType() reflect.Type {
	return reflect.TypeOf((**CircuitTermination)(nil)).Elem()
}

func (i *CircuitTermination) ToCircuitTerminationOutput() CircuitTerminationOutput {
	return i.ToCircuitTerminationOutputWithContext(context.Background())
}

func (i *CircuitTermination) ToCircuitTerminationOutputWithContext(ctx context.Context) CircuitTerminationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CircuitTerminationOutput)
}

// CircuitTerminationArrayInput is an input type that accepts CircuitTerminationArray and CircuitTerminationArrayOutput values.
// You can construct a concrete instance of `CircuitTerminationArrayInput` via:
//
//	CircuitTerminationArray{ CircuitTerminationArgs{...} }
type CircuitTerminationArrayInput interface {
	pulumi.Input

	ToCircuitTerminationArrayOutput() CircuitTerminationArrayOutput
	ToCircuitTerminationArrayOutputWithContext(context.Context) CircuitTerminationArrayOutput
}

type CircuitTerminationArray []CircuitTerminationInput

func (CircuitTerminationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CircuitTermination)(nil)).Elem()
}

func (i CircuitTerminationArray) ToCircuitTerminationArrayOutput() CircuitTerminationArrayOutput {
	return i.ToCircuitTerminationArrayOutputWithContext(context.Background())
}

func (i CircuitTerminationArray) ToCircuitTerminationArrayOutputWithContext(ctx context.Context) CircuitTerminationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CircuitTerminationArrayOutput)
}

// CircuitTerminationMapInput is an input type that accepts CircuitTerminationMap and CircuitTerminationMapOutput values.
// You can construct a concrete instance of `CircuitTerminationMapInput` via:
//
//	CircuitTerminationMap{ "key": CircuitTerminationArgs{...} }
type CircuitTerminationMapInput interface {
	pulumi.Input

	ToCircuitTerminationMapOutput() CircuitTerminationMapOutput
	ToCircuitTerminationMapOutputWithContext(context.Context) CircuitTerminationMapOutput
}

type CircuitTerminationMap map[string]CircuitTerminationInput

func (CircuitTerminationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CircuitTermination)(nil)).Elem()
}

func (i CircuitTerminationMap) ToCircuitTerminationMapOutput() CircuitTerminationMapOutput {
	return i.ToCircuitTerminationMapOutputWithContext(context.Background())
}

func (i CircuitTerminationMap) ToCircuitTerminationMapOutputWithContext(ctx context.Context) CircuitTerminationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CircuitTerminationMapOutput)
}

type CircuitTerminationOutput struct{ *pulumi.OutputState }

func (CircuitTerminationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CircuitTermination)(nil)).Elem()
}

func (o CircuitTerminationOutput) ToCircuitTerminationOutput() CircuitTerminationOutput {
	return o
}

func (o CircuitTerminationOutput) ToCircuitTerminationOutputWithContext(ctx context.Context) CircuitTerminationOutput {
	return o
}

func (o CircuitTerminationOutput) CircuitId() pulumi.IntOutput {
	return o.ApplyT(func(v *CircuitTermination) pulumi.IntOutput { return v.CircuitId }).(pulumi.IntOutput)
}

func (o CircuitTerminationOutput) CustomFields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CircuitTermination) pulumi.StringMapOutput { return v.CustomFields }).(pulumi.StringMapOutput)
}

func (o CircuitTerminationOutput) PortSpeed() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CircuitTermination) pulumi.IntPtrOutput { return v.PortSpeed }).(pulumi.IntPtrOutput)
}

func (o CircuitTerminationOutput) SiteId() pulumi.IntOutput {
	return o.ApplyT(func(v *CircuitTermination) pulumi.IntOutput { return v.SiteId }).(pulumi.IntOutput)
}

func (o CircuitTerminationOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CircuitTermination) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Valid values are `A` and `Z`.
func (o CircuitTerminationOutput) TermSide() pulumi.StringOutput {
	return o.ApplyT(func(v *CircuitTermination) pulumi.StringOutput { return v.TermSide }).(pulumi.StringOutput)
}

func (o CircuitTerminationOutput) UpstreamSpeed() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CircuitTermination) pulumi.IntPtrOutput { return v.UpstreamSpeed }).(pulumi.IntPtrOutput)
}

type CircuitTerminationArrayOutput struct{ *pulumi.OutputState }

func (CircuitTerminationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CircuitTermination)(nil)).Elem()
}

func (o CircuitTerminationArrayOutput) ToCircuitTerminationArrayOutput() CircuitTerminationArrayOutput {
	return o
}

func (o CircuitTerminationArrayOutput) ToCircuitTerminationArrayOutputWithContext(ctx context.Context) CircuitTerminationArrayOutput {
	return o
}

func (o CircuitTerminationArrayOutput) Index(i pulumi.IntInput) CircuitTerminationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CircuitTermination {
		return vs[0].([]*CircuitTermination)[vs[1].(int)]
	}).(CircuitTerminationOutput)
}

type CircuitTerminationMapOutput struct{ *pulumi.OutputState }

func (CircuitTerminationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CircuitTermination)(nil)).Elem()
}

func (o CircuitTerminationMapOutput) ToCircuitTerminationMapOutput() CircuitTerminationMapOutput {
	return o
}

func (o CircuitTerminationMapOutput) ToCircuitTerminationMapOutputWithContext(ctx context.Context) CircuitTerminationMapOutput {
	return o
}

func (o CircuitTerminationMapOutput) MapIndex(k pulumi.StringInput) CircuitTerminationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CircuitTermination {
		return vs[0].(map[string]*CircuitTermination)[vs[1].(string)]
	}).(CircuitTerminationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CircuitTerminationInput)(nil)).Elem(), &CircuitTermination{})
	pulumi.RegisterInputType(reflect.TypeOf((*CircuitTerminationArrayInput)(nil)).Elem(), CircuitTerminationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CircuitTerminationMapInput)(nil)).Elem(), CircuitTerminationMap{})
	pulumi.RegisterOutputType(CircuitTerminationOutput{})
	pulumi.RegisterOutputType(CircuitTerminationArrayOutput{})
	pulumi.RegisterOutputType(CircuitTerminationMapOutput{})
}
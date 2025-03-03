// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipam

import (
	"context"
	"reflect"

	"errors"
	"github.com/NatzkaLabsOpenSource/pulumi-netbox/sdk/go/netbox/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// From the [official documentation](https://docs.netbox.dev/en/stable/features/services/#services):
//
// > A service represents a layer four TCP or UDP service available on a device or virtual machine. For example, you might want to document that an HTTP service is running on a device. Each service includes a name, protocol, and port number; for example, "SSH (TCP/22)" or "DNS (UDP/53)."
// >
// > A service may optionally be bound to one or more specific IP addresses belonging to its parent device or VM. (If no IP addresses are bound, the service is assumed to be reachable via any assigned IP address.
type Service struct {
	pulumi.CustomResourceState

	CustomFields pulumi.StringMapOutput `pulumi:"customFields"`
	Description  pulumi.StringPtrOutput `pulumi:"description"`
	// Exactly one of `virtualMachineId` or `deviceId` must be given.
	DeviceId pulumi.IntPtrOutput `pulumi:"deviceId"`
	Name     pulumi.StringOutput `pulumi:"name"`
	// Exactly one of `port` or `ports` must be given.
	//
	// Deprecated: This field is deprecated. Please use the new "ports" attribute instead.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// Exactly one of `port` or `ports` must be given.
	Ports pulumi.IntArrayOutput `pulumi:"ports"`
	// Valid values are `tcp`, `udp` and `sctp`.
	Protocol pulumi.StringOutput      `pulumi:"protocol"`
	Tags     pulumi.StringArrayOutput `pulumi:"tags"`
	// Exactly one of `virtualMachineId` or `deviceId` must be given.
	VirtualMachineId pulumi.IntPtrOutput `pulumi:"virtualMachineId"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("netbox:ipam/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("netbox:ipam/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	CustomFields map[string]string `pulumi:"customFields"`
	Description  *string           `pulumi:"description"`
	// Exactly one of `virtualMachineId` or `deviceId` must be given.
	DeviceId *int    `pulumi:"deviceId"`
	Name     *string `pulumi:"name"`
	// Exactly one of `port` or `ports` must be given.
	//
	// Deprecated: This field is deprecated. Please use the new "ports" attribute instead.
	Port *int `pulumi:"port"`
	// Exactly one of `port` or `ports` must be given.
	Ports []int `pulumi:"ports"`
	// Valid values are `tcp`, `udp` and `sctp`.
	Protocol *string  `pulumi:"protocol"`
	Tags     []string `pulumi:"tags"`
	// Exactly one of `virtualMachineId` or `deviceId` must be given.
	VirtualMachineId *int `pulumi:"virtualMachineId"`
}

type ServiceState struct {
	CustomFields pulumi.StringMapInput
	Description  pulumi.StringPtrInput
	// Exactly one of `virtualMachineId` or `deviceId` must be given.
	DeviceId pulumi.IntPtrInput
	Name     pulumi.StringPtrInput
	// Exactly one of `port` or `ports` must be given.
	//
	// Deprecated: This field is deprecated. Please use the new "ports" attribute instead.
	Port pulumi.IntPtrInput
	// Exactly one of `port` or `ports` must be given.
	Ports pulumi.IntArrayInput
	// Valid values are `tcp`, `udp` and `sctp`.
	Protocol pulumi.StringPtrInput
	Tags     pulumi.StringArrayInput
	// Exactly one of `virtualMachineId` or `deviceId` must be given.
	VirtualMachineId pulumi.IntPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	CustomFields map[string]string `pulumi:"customFields"`
	Description  *string           `pulumi:"description"`
	// Exactly one of `virtualMachineId` or `deviceId` must be given.
	DeviceId *int    `pulumi:"deviceId"`
	Name     *string `pulumi:"name"`
	// Exactly one of `port` or `ports` must be given.
	//
	// Deprecated: This field is deprecated. Please use the new "ports" attribute instead.
	Port *int `pulumi:"port"`
	// Exactly one of `port` or `ports` must be given.
	Ports []int `pulumi:"ports"`
	// Valid values are `tcp`, `udp` and `sctp`.
	Protocol string   `pulumi:"protocol"`
	Tags     []string `pulumi:"tags"`
	// Exactly one of `virtualMachineId` or `deviceId` must be given.
	VirtualMachineId *int `pulumi:"virtualMachineId"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	CustomFields pulumi.StringMapInput
	Description  pulumi.StringPtrInput
	// Exactly one of `virtualMachineId` or `deviceId` must be given.
	DeviceId pulumi.IntPtrInput
	Name     pulumi.StringPtrInput
	// Exactly one of `port` or `ports` must be given.
	//
	// Deprecated: This field is deprecated. Please use the new "ports" attribute instead.
	Port pulumi.IntPtrInput
	// Exactly one of `port` or `ports` must be given.
	Ports pulumi.IntArrayInput
	// Valid values are `tcp`, `udp` and `sctp`.
	Protocol pulumi.StringInput
	Tags     pulumi.StringArrayInput
	// Exactly one of `virtualMachineId` or `deviceId` must be given.
	VirtualMachineId pulumi.IntPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//	ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//	ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func (o ServiceOutput) CustomFields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.CustomFields }).(pulumi.StringMapOutput)
}

func (o ServiceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Exactly one of `virtualMachineId` or `deviceId` must be given.
func (o ServiceOutput) DeviceId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.IntPtrOutput { return v.DeviceId }).(pulumi.IntPtrOutput)
}

func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Exactly one of `port` or `ports` must be given.
//
// Deprecated: This field is deprecated. Please use the new "ports" attribute instead.
func (o ServiceOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// Exactly one of `port` or `ports` must be given.
func (o ServiceOutput) Ports() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Service) pulumi.IntArrayOutput { return v.Ports }).(pulumi.IntArrayOutput)
}

// Valid values are `tcp`, `udp` and `sctp`.
func (o ServiceOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o ServiceOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Service) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Exactly one of `virtualMachineId` or `deviceId` must be given.
func (o ServiceOutput) VirtualMachineId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.IntPtrOutput { return v.VirtualMachineId }).(pulumi.IntPtrOutput)
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Service {
		return vs[0].([]*Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Service {
		return vs[0].(map[string]*Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceArrayInput)(nil)).Elem(), ServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMapInput)(nil)).Elem(), ServiceMap{})
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}

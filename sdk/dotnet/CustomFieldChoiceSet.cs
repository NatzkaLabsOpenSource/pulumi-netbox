// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Netbox
{
    /// <summary>
    /// From the [official documentation](https://docs.netbox.dev/en/stable/models/extras/customfieldchoiceset/):
    /// 
    /// Single- and multi-selection custom fields must define a set of valid choices from which the user may choose when defining the field value. These choices are defined as sets that may be reused among multiple custom fields.
    /// 
    /// A choice set must define a base choice set and/or a set of arbitrary extra choices.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Netbox = Pulumi.Netbox;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Netbox.CustomFieldChoiceSet("test", new()
    ///     {
    ///         Name = "my-custom-field-set",
    ///         Description = "Description",
    ///         ExtraChoices = new[]
    ///         {
    ///             new[]
    ///             {
    ///                 "choice1",
    ///                 "label1",
    ///             },
    ///             new[]
    ///             {
    ///                 "choice2",
    ///                 "choice2",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [NetboxResourceType("netbox:index/customFieldChoiceSet:CustomFieldChoiceSet")]
    public partial class CustomFieldChoiceSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Valid values are `IATA`, `ISO_3166` and `UN_LOCODE`. At least one of `base_choices` or `extra_choices` must be given.
        /// </summary>
        [Output("baseChoices")]
        public Output<string?> BaseChoices { get; private set; } = null!;

        [Output("customFields")]
        public Output<ImmutableDictionary<string, string>?> CustomFields { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice. At least one of `base_choices` or `extra_choices` must be given.
        /// </summary>
        [Output("extraChoices")]
        public Output<ImmutableArray<ImmutableArray<string>>> ExtraChoices { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// experimental. Defaults to `false`.
        /// </summary>
        [Output("orderAlphabetically")]
        public Output<bool?> OrderAlphabetically { get; private set; } = null!;


        /// <summary>
        /// Create a CustomFieldChoiceSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomFieldChoiceSet(string name, CustomFieldChoiceSetArgs? args = null, CustomResourceOptions? options = null)
            : base("netbox:index/customFieldChoiceSet:CustomFieldChoiceSet", name, args ?? new CustomFieldChoiceSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomFieldChoiceSet(string name, Input<string> id, CustomFieldChoiceSetState? state = null, CustomResourceOptions? options = null)
            : base("netbox:index/customFieldChoiceSet:CustomFieldChoiceSet", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/NatzkaLabsOpenSource/pulumi-netbox",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CustomFieldChoiceSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomFieldChoiceSet Get(string name, Input<string> id, CustomFieldChoiceSetState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomFieldChoiceSet(name, id, state, options);
        }
    }

    public sealed class CustomFieldChoiceSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Valid values are `IATA`, `ISO_3166` and `UN_LOCODE`. At least one of `base_choices` or `extra_choices` must be given.
        /// </summary>
        [Input("baseChoices")]
        public Input<string>? BaseChoices { get; set; }

        [Input("customFields")]
        private InputMap<string>? _customFields;
        public InputMap<string> CustomFields
        {
            get => _customFields ?? (_customFields = new InputMap<string>());
            set => _customFields = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("extraChoices")]
        private InputList<ImmutableArray<string>>? _extraChoices;

        /// <summary>
        /// This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice. At least one of `base_choices` or `extra_choices` must be given.
        /// </summary>
        public InputList<ImmutableArray<string>> ExtraChoices
        {
            get => _extraChoices ?? (_extraChoices = new InputList<ImmutableArray<string>>());
            set => _extraChoices = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// experimental. Defaults to `false`.
        /// </summary>
        [Input("orderAlphabetically")]
        public Input<bool>? OrderAlphabetically { get; set; }

        public CustomFieldChoiceSetArgs()
        {
        }
        public static new CustomFieldChoiceSetArgs Empty => new CustomFieldChoiceSetArgs();
    }

    public sealed class CustomFieldChoiceSetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Valid values are `IATA`, `ISO_3166` and `UN_LOCODE`. At least one of `base_choices` or `extra_choices` must be given.
        /// </summary>
        [Input("baseChoices")]
        public Input<string>? BaseChoices { get; set; }

        [Input("customFields")]
        private InputMap<string>? _customFields;
        public InputMap<string> CustomFields
        {
            get => _customFields ?? (_customFields = new InputMap<string>());
            set => _customFields = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("extraChoices")]
        private InputList<ImmutableArray<string>>? _extraChoices;

        /// <summary>
        /// This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice. At least one of `base_choices` or `extra_choices` must be given.
        /// </summary>
        public InputList<ImmutableArray<string>> ExtraChoices
        {
            get => _extraChoices ?? (_extraChoices = new InputList<ImmutableArray<string>>());
            set => _extraChoices = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// experimental. Defaults to `false`.
        /// </summary>
        [Input("orderAlphabetically")]
        public Input<bool>? OrderAlphabetically { get; set; }

        public CustomFieldChoiceSetState()
        {
        }
        public static new CustomFieldChoiceSetState Empty => new CustomFieldChoiceSetState();
    }
}

// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Gitpod
{
    [GitpodResourceType("pulumi:providers:gitpod")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// Your Gitpod access token
        /// </summary>
        [Output("accessToken")]
        public Output<string?> AccessToken { get; private set; } = null!;

        [Output("organizationId")]
        public Output<string?> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Id of owner account
        /// </summary>
        [Output("ownerId")]
        public Output<string?> OwnerId { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("gitpod", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pierskarsenbarg/pulumi-gitpod",
                AdditionalSecretOutputs =
                {
                    "accessToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken")]
        private Input<string>? _accessToken;

        /// <summary>
        /// Your Gitpod access token
        /// </summary>
        public Input<string>? AccessToken
        {
            get => _accessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// Id of owner account
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        public ProviderArgs()
        {
            AccessToken = Utilities.GetEnv("GITPOD_ACCESSTOKEN") ?? "";
            OrganizationId = Utilities.GetEnv("GITPOD_ORGANISATIONID") ?? "";
            OwnerId = Utilities.GetEnv("GITPOD_OWNERID") ?? "";
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}

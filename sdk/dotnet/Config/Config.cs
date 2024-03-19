// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace PiersKarsenbarg.Gitpod
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("gitpod");

        private static readonly __Value<string?> _accessToken = new __Value<string?>(() => __config.Get("accessToken") ?? Utilities.GetEnv("GITPOD_ACCESSTOKEN") ?? "");
        /// <summary>
        /// Your Gitpod access token. You can create these in your user settings: https://gitpod.io/user/tokens
        /// </summary>
        public static string? AccessToken
        {
            get => _accessToken.Get();
            set => _accessToken.Set(value);
        }

        private static readonly __Value<string?> _organizationId = new __Value<string?>(() => __config.Get("organizationId") ?? Utilities.GetEnv("GITPOD_ORGANISATIONID") ?? "");
        /// <summary>
        /// Id of the organisation who is going to own the workspaces. You can find this on the organisation settings page of the currently selected organisation: https://gitpod.io/settings
        /// </summary>
        public static string? OrganizationId
        {
            get => _organizationId.Get();
            set => _organizationId.Set(value);
        }

        private static readonly __Value<string?> _ownerId = new __Value<string?>(() => __config.Get("ownerId") ?? Utilities.GetEnv("GITPOD_OWNERID") ?? "");
        /// <summary>
        /// Id of owner account. This can be found on your user account page: https://gitpod.io/user/account
        /// </summary>
        public static string? OwnerId
        {
            get => _ownerId.Get();
            set => _ownerId.Set(value);
        }

    }
}

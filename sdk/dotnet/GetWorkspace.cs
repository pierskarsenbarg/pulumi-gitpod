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
    public static class GetWorkspace
    {
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("gitpod:index:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithDefaults());

        public static Output<GetWorkspaceResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkspaceResult>("gitpod:index:getWorkspace", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetWorkspaceArgs : global::Pulumi.InvokeArgs
    {
        public GetWorkspaceArgs()
        {
        }
        public static new GetWorkspaceArgs Empty => new GetWorkspaceArgs();
    }


    [OutputType]
    public sealed class GetWorkspaceResult
    {
        /// <summary>
        /// The information from git used to build the workspace
        /// </summary>
        public readonly Outputs.WorkspaceStateGitInfo GitInfo;
        /// <summary>
        /// The metadata of the workspace
        /// </summary>
        public readonly Outputs.WorkspaceStateMetaData Metadata;
        public readonly Outputs.WorkspaceStateSpec Spec;
        /// <summary>
        /// The ID of the workspace
        /// </summary>
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetWorkspaceResult(
            Outputs.WorkspaceStateGitInfo gitInfo,

            Outputs.WorkspaceStateMetaData metadata,

            Outputs.WorkspaceStateSpec spec,

            string workspaceId)
        {
            GitInfo = gitInfo;
            Metadata = metadata;
            Spec = spec;
            WorkspaceId = workspaceId;
        }
    }
}

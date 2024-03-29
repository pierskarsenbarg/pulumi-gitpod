// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Gitpod.Inputs
{

    public sealed class WorkspaceEditorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the editor that you'd like to use in your Gitpod workspace. Defaults to VS Code desktop.
        /// </summary>
        [Input("name")]
        public Input<PiersKarsenbarg.Gitpod.EditorName>? Name { get; set; }

        /// <summary>
        /// Version of editor to use. Options are `latest` or `stable`. Defaults to `stable`
        /// </summary>
        [Input("version")]
        public Input<PiersKarsenbarg.Gitpod.EditorVersion>? Version { get; set; }

        public WorkspaceEditorArgs()
        {
            Name = PiersKarsenbarg.Gitpod.EditorName.Code;
            Version = PiersKarsenbarg.Gitpod.EditorVersion.Stable;
        }
        public static new WorkspaceEditorArgs Empty => new WorkspaceEditorArgs();
    }
}

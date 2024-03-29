// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitpod

type EditorName string

const (
	// Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
	EditorNameEditorNameCode = EditorName("code")
	// Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
	EditorName_EditorName_Code_Desktop = EditorName("code-desktop")
	// Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
	EditorNameEditorNameIntellij = EditorName("intellij")
	// Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
	EditorNameEditorNameGoland = EditorName("goland")
	// Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
	EditorNameEditorNamePhpstorm = EditorName("phpstorm")
	// Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
	EditorNameEditorNamePycharm = EditorName("pycharm")
	// Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
	EditorNameEditorNameRubymine = EditorName("rubymine")
	// Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
	EditorNameEditorNameWebstorm = EditorName("webstorm")
	// Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
	EditorNameEditorNameRider = EditorName("rider")
	// Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
	EditorNameEditorNameClion = EditorName("clion")
)

type EditorVersion string

const (
	// Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
	EditorVersionEditorVersionLatest = EditorVersion("latest")
	// Standard workspace class - up to 4 cores, up to 8GB RAM, 30GB storage
	EditorVersionEditorVersionStable = EditorVersion("stable")
)

type WorkspaceClass string

const (
	// Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
	WorkspaceClass_WorkspaceClass_G1_Large = WorkspaceClass("g1-large")
	// Standard workspace class - up to 4 cores, up to 8GB RAM, 30GB storage
	WorkspaceClass_WorkspaceClass_G1_Standard = WorkspaceClass("g1-standard")
)

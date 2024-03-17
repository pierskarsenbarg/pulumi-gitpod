package pkg

import "github.com/pulumi/pulumi-go-provider/infer"

type WorkspaceClass string
type EditorName string
type EditorVersion string

var _ = (infer.Enum[WorkspaceClass])((*WorkspaceClass)(nil))
var _ = (infer.Enum[EditorName])((*EditorName)(nil))
var _ = (infer.Enum[EditorVersion])((*EditorVersion)(nil))

const (
	G1Standard WorkspaceClass = "g1-standard"
	G1Large    WorkspaceClass = "g1-large"
)

const (
	Code        EditorName = "code"
	CodeDesktop EditorName = "code-desktop"
	IntelliJ    EditorName = "intellij"
	GoLand      EditorName = "goland"
	PhpStorm    EditorName = "phpstorm"
	PyCharm     EditorName = "pycharm"
	RubyMine    EditorName = "rubymine"
	WebStorm    EditorName = "webstorm"
	Rider       EditorName = "rider"
	CLion       EditorName = "clion"
)

const (
	Stable EditorVersion = "stable"
	Latest EditorVersion = "latest"
)

func (*WorkspaceClass) Values() []infer.EnumValue[WorkspaceClass] {
	return []infer.EnumValue[WorkspaceClass]{
		{Value: G1Large, Description: "Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage."},
		{Value: G1Standard, Description: "Standard workspace class - up to 4 cores, up to 8GB RAM, 30GB storage"},
	}
}

func (*EditorName) Values() []infer.EnumValue[EditorName] {
	return []infer.EnumValue[EditorName]{
		{Value: Code, Description: "Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage."},
		{Value: CodeDesktop, Description: "Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage."},
		{Value: IntelliJ, Description: "Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage."},
		{Value: GoLand, Description: "Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage."},
		{Value: PhpStorm, Description: "Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage."},
		{Value: PyCharm, Description: "Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage."},
		{Value: RubyMine, Description: "Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage."},
		{Value: WebStorm, Description: "Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage."},
		{Value: Rider, Description: "Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage."},
		{Value: CLion, Description: "Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage."},
	}
}

func (*EditorVersion) Values() []infer.EnumValue[EditorVersion] {
	return []infer.EnumValue[EditorVersion]{
		{Value: Latest, Description: "Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage."},
		{Value: Stable, Description: "Standard workspace class - up to 4 cores, up to 8GB RAM, 30GB storage"},
	}
}

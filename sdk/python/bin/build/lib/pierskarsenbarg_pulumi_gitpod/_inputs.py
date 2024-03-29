# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._enums import *

__all__ = [
    'WorkspaceEditorArgs',
]

@pulumi.input_type
class WorkspaceEditorArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input['EditorName']] = None,
                 version: Optional[pulumi.Input['EditorVersion']] = None):
        """
        :param pulumi.Input['EditorName'] name: Name of the editor that you'd like to use in your Gitpod workspace. Defaults to VS Code desktop.
        :param pulumi.Input['EditorVersion'] version: Version of editor to use. Options are `latest` or `stable`. Defaults to `stable`
        """
        if name is None:
            name = 'code'
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is None:
            version = 'stable'
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input['EditorName']]:
        """
        Name of the editor that you'd like to use in your Gitpod workspace. Defaults to VS Code desktop.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input['EditorName']]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input['EditorVersion']]:
        """
        Version of editor to use. Options are `latest` or `stable`. Defaults to `stable`
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input['EditorVersion']]):
        pulumi.set(self, "version", value)



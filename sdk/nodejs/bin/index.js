"use strict";
// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __exportStar = (this && this.__exportStar) || function(m, exports) {
    for (var p in m) if (p !== "default" && !Object.prototype.hasOwnProperty.call(exports, p)) __createBinding(exports, m, p);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.types = exports.config = exports.Workspace = exports.Provider = exports.Organization = exports.getOrganizationOutput = exports.getOrganization = void 0;
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
exports.getOrganization = null;
exports.getOrganizationOutput = null;
utilities.lazyLoad(exports, ["getOrganization", "getOrganizationOutput"], () => require("./getOrganization"));
exports.Organization = null;
utilities.lazyLoad(exports, ["Organization"], () => require("./organization"));
exports.Provider = null;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));
exports.Workspace = null;
utilities.lazyLoad(exports, ["Workspace"], () => require("./workspace"));
// Export enums:
__exportStar(require("./types/enums"), exports);
// Export sub-modules:
const config = require("./config");
exports.config = config;
const types = require("./types");
exports.types = types;
const _module = {
    version: utilities.getVersion(),
    construct: (name, type, urn) => {
        switch (type) {
            case "gitpod:index:Organization":
                return new exports.Organization(name, undefined, { urn });
            case "gitpod:index:Workspace":
                return new exports.Workspace(name, undefined, { urn });
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gitpod", "index", _module);
pulumi.runtime.registerResourcePackage("gitpod", {
    version: utilities.getVersion(),
    constructProvider: (name, type, urn) => {
        if (type !== "pulumi:providers:gitpod") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new exports.Provider(name, undefined, { urn });
    },
});
//# sourceMappingURL=index.js.map
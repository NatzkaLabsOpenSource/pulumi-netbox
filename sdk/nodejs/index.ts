// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { ConfigContextArgs, ConfigContextState } from "./configContext";
export type ConfigContext = import("./configContext").ConfigContext;
export const ConfigContext: typeof import("./configContext").ConfigContext = null as any;
utilities.lazyLoad(exports, ["ConfigContext"], () => require("./configContext"));

export { ConfigTemplateArgs, ConfigTemplateState } from "./configTemplate";
export type ConfigTemplate = import("./configTemplate").ConfigTemplate;
export const ConfigTemplate: typeof import("./configTemplate").ConfigTemplate = null as any;
utilities.lazyLoad(exports, ["ConfigTemplate"], () => require("./configTemplate"));

export { CustomFieldArgs, CustomFieldState } from "./customField";
export type CustomField = import("./customField").CustomField;
export const CustomField: typeof import("./customField").CustomField = null as any;
utilities.lazyLoad(exports, ["CustomField"], () => require("./customField"));

export { CustomFieldChoiceSetArgs, CustomFieldChoiceSetState } from "./customFieldChoiceSet";
export type CustomFieldChoiceSet = import("./customFieldChoiceSet").CustomFieldChoiceSet;
export const CustomFieldChoiceSet: typeof import("./customFieldChoiceSet").CustomFieldChoiceSet = null as any;
utilities.lazyLoad(exports, ["CustomFieldChoiceSet"], () => require("./customFieldChoiceSet"));

export { EventRuleArgs, EventRuleState } from "./eventRule";
export type EventRule = import("./eventRule").EventRule;
export const EventRule: typeof import("./eventRule").EventRule = null as any;
utilities.lazyLoad(exports, ["EventRule"], () => require("./eventRule"));

export { GetConfigContextArgs, GetConfigContextResult, GetConfigContextOutputArgs } from "./getConfigContext";
export const getConfigContext: typeof import("./getConfigContext").getConfigContext = null as any;
export const getConfigContextOutput: typeof import("./getConfigContext").getConfigContextOutput = null as any;
utilities.lazyLoad(exports, ["getConfigContext","getConfigContextOutput"], () => require("./getConfigContext"));

export { GetTagArgs, GetTagResult, GetTagOutputArgs } from "./getTag";
export const getTag: typeof import("./getTag").getTag = null as any;
export const getTagOutput: typeof import("./getTag").getTagOutput = null as any;
utilities.lazyLoad(exports, ["getTag","getTagOutput"], () => require("./getTag"));

export { GetTagsArgs, GetTagsResult, GetTagsOutputArgs } from "./getTags";
export const getTags: typeof import("./getTags").getTags = null as any;
export const getTagsOutput: typeof import("./getTags").getTagsOutput = null as any;
utilities.lazyLoad(exports, ["getTags","getTagsOutput"], () => require("./getTags"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { TagArgs, TagState } from "./tag";
export type Tag = import("./tag").Tag;
export const Tag: typeof import("./tag").Tag = null as any;
utilities.lazyLoad(exports, ["Tag"], () => require("./tag"));

export { WebhookArgs, WebhookState } from "./webhook";
export type Webhook = import("./webhook").Webhook;
export const Webhook: typeof import("./webhook").Webhook = null as any;
utilities.lazyLoad(exports, ["Webhook"], () => require("./webhook"));


// Export sub-modules:
import * as auth from "./auth";
import * as circuit from "./circuit";
import * as config from "./config";
import * as dcim from "./dcim";
import * as ipam from "./ipam";
import * as tenancy from "./tenancy";
import * as types from "./types";
import * as virt from "./virt";
import * as vpn from "./vpn";

export {
    auth,
    circuit,
    config,
    dcim,
    ipam,
    tenancy,
    types,
    virt,
    vpn,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "netbox:index/configContext:ConfigContext":
                return new ConfigContext(name, <any>undefined, { urn })
            case "netbox:index/configTemplate:ConfigTemplate":
                return new ConfigTemplate(name, <any>undefined, { urn })
            case "netbox:index/customField:CustomField":
                return new CustomField(name, <any>undefined, { urn })
            case "netbox:index/customFieldChoiceSet:CustomFieldChoiceSet":
                return new CustomFieldChoiceSet(name, <any>undefined, { urn })
            case "netbox:index/eventRule:EventRule":
                return new EventRule(name, <any>undefined, { urn })
            case "netbox:index/tag:Tag":
                return new Tag(name, <any>undefined, { urn })
            case "netbox:index/webhook:Webhook":
                return new Webhook(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("netbox", "index/configContext", _module)
pulumi.runtime.registerResourceModule("netbox", "index/configTemplate", _module)
pulumi.runtime.registerResourceModule("netbox", "index/customField", _module)
pulumi.runtime.registerResourceModule("netbox", "index/customFieldChoiceSet", _module)
pulumi.runtime.registerResourceModule("netbox", "index/eventRule", _module)
pulumi.runtime.registerResourceModule("netbox", "index/tag", _module)
pulumi.runtime.registerResourceModule("netbox", "index/webhook", _module)
pulumi.runtime.registerResourcePackage("netbox", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:netbox") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});

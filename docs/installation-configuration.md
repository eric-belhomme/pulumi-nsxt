---
title: Nsxt Installation & Configuration
meta_desc: Information on how to install the Nsxt provider.
layout: installation
---

## Installation

The Pulumi Nsxt provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@SCC-Hyperscale-fr/nsxt`](https://www.npmjs.com/package/@SCC-Hyperscale-fr/nsxt)
* Python: [`hyperscale_pulumi_nsxt`](https://pypi.org/project/hyperscale_pulumi_nsxt/)
* Go: [`github.com/SCC-Hyperscale-fr/pulumi-nsxt/sdk/go/nsxt`](https://pkg.go.dev/github.com/SCC-Hyperscale-fr/pulumi-nsxt/sdk/go/nsxt)
* .NET: [`SCC-Hyperscale-fr.Nsxt`](https://www.nuget.org/packages/SCC-Hyperscale-fr.Nsxt)


## Configuration

> Note:  
> Replace the following **sample content**, with the configuration options
> of the wrapped Terraform provider and remove this note.

The following configuration points are available for the `nsxt` provider:

- `nsxt:apiKey` (environment: `nsxt_API_KEY`) - the API key for `nsxt`
- `nsxt:region` (environment: `nsxt_REGION`) - the region in which to deploy resources

### Provider Binary

The Nsxt provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource nsxt <version>
```

Replace the version string `<version>` with your desired version.

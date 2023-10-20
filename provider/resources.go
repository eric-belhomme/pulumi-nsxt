// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nsxt

import (
	
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ettle/strcase"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/vmware/terraform-provider-nsxt/nsxt"
	"github.com/SCC-Hyperscale-fr/pulumi-nsxt/provider/pkg/version"
)



// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	mainMod = "index" // the nsxt module
)

func convertName(name string) string {
	idx := strings.Index(name, "_")
	contract.Assertf(idx > 0 && idx < len(name)-1, "Invalid snake case name %s", name)
	name = name[idx+1:]
	contract.Assertf(len(name) > 0, "Invalid snake case name %s", name)
	return strcase.ToPascal(name)
}

func makeDataSource(mod string, name string) tokens.ModuleMember {
	name = convertName(name)
	return tfbridge.MakeDataSource("nsxt", mod, "get"+name)
}

func makeResource(mod string, res string) tokens.Type {
	return tfbridge.MakeResource("nsxt", mod, convertName(res))
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(nsxt.Provider())
			// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "nsxt",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Nsxt",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "SCC-Hyperscale-fr",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://raw.githubusercontent.com/SCC-Hyperscale-fr/pulumi-nsxt/main/docs/nsxt.png",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/SCC-Hyperscale-fr/pulumi-nsxt",
		Description:       "A Pulumi package for creating and managing Nsxt resources",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{
			"pulumi",
			"nsxt",
			"category/network",
		},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/SCC-Hyperscale-fr/pulumi-nsxt",
		Repository: "https://github.com/SCC-Hyperscale-fr/pulumi-nsxt",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		Version:   version.Version,
		GitHubOrg: "vmware",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources:            map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: makeResource(mainMod(mainMod, "aws_iam_role")}
			//
			// "aws_acm_certificate": {
			// 	Tok: Tok: makeResource(mainMod(mainMod, "aws_acm_certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType("nsxt", "Tags")},
			// 	},
			// },
			"nsxt_dhcp_relay_profile": {Tok: makeResource(mainMod, "nsxt_dhcp_relay_profile")},
			"nsxt_dhcp_relay_service": {Tok: makeResource(mainMod, "nsxt_dhcp_relay_service")},
			"nsxt_dhcp_server_profile": {Tok: makeResource(mainMod, "nsxt_dhcp_server_profile")},
			"nsxt_logical_dhcp_server": {Tok: makeResource(mainMod, "nsxt_logical_dhcp_server")},
			"nsxt_dhcp_server_ip_pool": {Tok: makeResource(mainMod, "nsxt_dhcp_server_ip_pool")},
			"nsxt_logical_switch": {Tok: makeResource(mainMod, "nsxt_logical_switch")},
			"nsxt_vlan_logical_switch": {Tok: makeResource(mainMod, "nsxt_vlan_logical_switch")},
			"nsxt_logical_dhcp_port": {Tok: makeResource(mainMod, "nsxt_logical_dhcp_port")},
			"nsxt_logical_port": {Tok: makeResource(mainMod, "nsxt_logical_port")},
			"nsxt_logical_tier0_router": {Tok: makeResource(mainMod, "nsxt_logical_tier0_router")},
			"nsxt_logical_tier1_router": {Tok: makeResource(mainMod, "nsxt_logical_tier1_router")},
			"nsxt_logical_router_centralized_service_port": {Tok: makeResource(mainMod, "nsxt_logical_router_centralized_service_port")},
			"nsxt_logical_router_downlink_port": {Tok: makeResource(mainMod, "nsxt_logical_router_downlink_port")},
			"nsxt_logical_router_link_port_on_tier0": {Tok: makeResource(mainMod, "nsxt_logical_router_link_port_on_tier0")},
			"nsxt_logical_router_link_port_on_tier1": {Tok: makeResource(mainMod, "nsxt_logical_router_link_port_on_tier1")},
			"nsxt_ip_discovery_switching_profile": {Tok: makeResource(mainMod, "nsxt_ip_discovery_switching_profile")},
			"nsxt_mac_management_switching_profile": {Tok: makeResource(mainMod, "nsxt_mac_management_switching_profile")},
			"nsxt_qos_switching_profile": {Tok: makeResource(mainMod, "nsxt_qos_switching_profile")},
			"nsxt_spoofguard_switching_profile": {Tok: makeResource(mainMod, "nsxt_spoofguard_switching_profile")},
			"nsxt_switch_security_switching_profile": {Tok: makeResource(mainMod, "nsxt_switch_security_switching_profile")},
			"nsxt_l4_port_set_ns_service": {Tok: makeResource(mainMod, "nsxt_l4_port_set_ns_service")},
			"nsxt_algorithm_type_ns_service": {Tok: makeResource(mainMod, "nsxt_algorithm_type_ns_service")},
			"nsxt_icmp_type_ns_service": {Tok: makeResource(mainMod, "nsxt_icmp_type_ns_service")},
			"nsxt_igmp_type_ns_service": {Tok: makeResource(mainMod, "nsxt_igmp_type_ns_service")},
			"nsxt_ether_type_ns_service": {Tok: makeResource(mainMod, "nsxt_ether_type_ns_service")},
			"nsxt_ip_protocol_ns_service": {Tok: makeResource(mainMod, "nsxt_ip_protocol_ns_service")},
			"nsxt_ns_service_group": {Tok: makeResource(mainMod, "nsxt_ns_service_group")},
			"nsxt_ns_group": {Tok: makeResource(mainMod, "nsxt_ns_group")},
			"nsxt_firewall_section": {Tok: makeResource(mainMod, "nsxt_firewall_section")},
			"nsxt_nat_rule": {Tok: makeResource(mainMod, "nsxt_nat_rule")},
			"nsxt_ip_block": {Tok: makeResource(mainMod, "nsxt_ip_block")},
			"nsxt_ip_block_subnet": {Tok: makeResource(mainMod, "nsxt_ip_block_subnet")},
			"nsxt_ip_pool": {Tok: makeResource(mainMod, "nsxt_ip_pool")},
			"nsxt_ip_pool_allocation_ip_address": {Tok: makeResource(mainMod, "nsxt_ip_pool_allocation_ip_address")},
			"nsxt_ip_set": {Tok: makeResource(mainMod, "nsxt_ip_set")},
			"nsxt_static_route": {Tok: makeResource(mainMod, "nsxt_static_route")},
			"nsxt_vm_tags": {Tok: makeResource(mainMod, "nsxt_vm_tags")},
			"nsxt_lb_icmp_monitor": {Tok: makeResource(mainMod, "nsxt_lb_icmp_monitor")},
			"nsxt_lb_tcp_monitor": {Tok: makeResource(mainMod, "nsxt_lb_tcp_monitor")},
			"nsxt_lb_udp_monitor": {Tok: makeResource(mainMod, "nsxt_lb_udp_monitor")},
			"nsxt_lb_http_monitor": {Tok: makeResource(mainMod, "nsxt_lb_http_monitor")},
			"nsxt_lb_https_monitor": {Tok: makeResource(mainMod, "nsxt_lb_https_monitor")},
			"nsxt_lb_passive_monitor": {Tok: makeResource(mainMod, "nsxt_lb_passive_monitor")},
			"nsxt_lb_pool": {Tok: makeResource(mainMod, "nsxt_lb_pool")},
			"nsxt_lb_tcp_virtual_server": {Tok: makeResource(mainMod, "nsxt_lb_tcp_virtual_server")},
			"nsxt_lb_udp_virtual_server": {Tok: makeResource(mainMod, "nsxt_lb_udp_virtual_server")},
			"nsxt_lb_http_virtual_server": {Tok: makeResource(mainMod, "nsxt_lb_http_virtual_server")},
			"nsxt_lb_http_forwarding_rule": {Tok: makeResource(mainMod, "nsxt_lb_http_forwarding_rule")},
			"nsxt_lb_http_request_rewrite_rule": {Tok: makeResource(mainMod, "nsxt_lb_http_request_rewrite_rule")},
			"nsxt_lb_http_response_rewrite_rule": {Tok: makeResource(mainMod, "nsxt_lb_http_response_rewrite_rule")},
			"nsxt_lb_cookie_persistence_profile": {Tok: makeResource(mainMod, "nsxt_lb_cookie_persistence_profile")},
			"nsxt_lb_source_ip_persistence_profile": {Tok: makeResource(mainMod, "nsxt_lb_source_ip_persistence_profile")},
			"nsxt_lb_client_ssl_profile": {Tok: makeResource(mainMod, "nsxt_lb_client_ssl_profile")},
			"nsxt_lb_server_ssl_profile": {Tok: makeResource(mainMod, "nsxt_lb_server_ssl_profile")},
			"nsxt_lb_service": {Tok: makeResource(mainMod, "nsxt_lb_service")},
			"nsxt_lb_fast_tcp_application_profile": {Tok: makeResource(mainMod, "nsxt_lb_fast_tcp_application_profile")},
			"nsxt_lb_fast_udp_application_profile": {Tok: makeResource(mainMod, "nsxt_lb_fast_udp_application_profile")},
			"nsxt_lb_http_application_profile": {Tok: makeResource(mainMod, "nsxt_lb_http_application_profile")},
			"nsxt_policy_tier1_gateway": {Tok: makeResource(mainMod, "nsxt_policy_tier1_gateway")},
			"nsxt_policy_tier1_gateway_interface": {Tok: makeResource(mainMod, "nsxt_policy_tier1_gateway_interface")},
			"nsxt_policy_tier0_gateway": {Tok: makeResource(mainMod, "nsxt_policy_tier0_gateway")},
			"nsxt_policy_tier0_gateway_interface": {Tok: makeResource(mainMod, "nsxt_policy_tier0_gateway_interface")},
			"nsxt_policy_tier0_gateway_ha_vip_config": {Tok: makeResource(mainMod, "nsxt_policy_tier0_gateway_ha_vip_config")},
			"nsxt_policy_group": {Tok: makeResource(mainMod, "nsxt_policy_group")},
			"nsxt_policy_domain": {Tok: makeResource(mainMod, "nsxt_policy_domain")},
			"nsxt_policy_security_policy": {Tok: makeResource(mainMod, "nsxt_policy_security_policy")},
			"nsxt_policy_service": {Tok: makeResource(mainMod, "nsxt_policy_service")},
			"nsxt_policy_gateway_policy": {Tok: makeResource(mainMod, "nsxt_policy_gateway_policy")},
			"nsxt_policy_predefined_gateway_policy": {Tok: makeResource(mainMod, "nsxt_policy_predefined_gateway_policy")},
			"nsxt_policy_predefined_security_policy": {Tok: makeResource(mainMod, "nsxt_policy_predefined_security_policy")},
			"nsxt_policy_segment": {Tok: makeResource(mainMod, "nsxt_policy_segment")},
			"nsxt_policy_vlan_segment": {Tok: makeResource(mainMod, "nsxt_policy_vlan_segment")},
			"nsxt_policy_fixed_segment": {Tok: makeResource(mainMod, "nsxt_policy_fixed_segment")},
			"nsxt_policy_static_route": {Tok: makeResource(mainMod, "nsxt_policy_static_route")},
			"nsxt_policy_gateway_prefix_list": {Tok: makeResource(mainMod, "nsxt_policy_gateway_prefix_list")},
			"nsxt_policy_vm_tags": {Tok: makeResource(mainMod, "nsxt_policy_vm_tags")},
			"nsxt_policy_nat_rule": {Tok: makeResource(mainMod, "nsxt_policy_nat_rule")},
			"nsxt_policy_ip_block": {Tok: makeResource(mainMod, "nsxt_policy_ip_block")},
			"nsxt_policy_lb_pool": {Tok: makeResource(mainMod, "nsxt_policy_lb_pool")},
			"nsxt_policy_ip_pool": {Tok: makeResource(mainMod, "nsxt_policy_ip_pool")},
			"nsxt_policy_ip_pool_block_subnet": {Tok: makeResource(mainMod, "nsxt_policy_ip_pool_block_subnet")},
			"nsxt_policy_ip_pool_static_subnet": {Tok: makeResource(mainMod, "nsxt_policy_ip_pool_static_subnet")},
			"nsxt_policy_lb_service": {Tok: makeResource(mainMod, "nsxt_policy_lb_service")},
			"nsxt_policy_lb_virtual_server": {Tok: makeResource(mainMod, "nsxt_policy_lb_virtual_server")},
			"nsxt_policy_ip_address_allocation": {Tok: makeResource(mainMod, "nsxt_policy_ip_address_allocation")},
			"nsxt_policy_bgp_neighbor": {Tok: makeResource(mainMod, "nsxt_policy_bgp_neighbor")},
			"nsxt_policy_bgp_config": {Tok: makeResource(mainMod, "nsxt_policy_bgp_config")},
			"nsxt_policy_dhcp_relay": {Tok: makeResource(mainMod, "nsxt_policy_dhcp_relay")},
			"nsxt_policy_dhcp_server": {Tok: makeResource(mainMod, "nsxt_policy_dhcp_server")},
			"nsxt_policy_context_profile": {Tok: makeResource(mainMod, "nsxt_policy_context_profile")},
			"nsxt_policy_dhcp_v4_static_binding": {Tok: makeResource(mainMod, "nsxt_policy_dhcp_v4_static_binding")},
			"nsxt_policy_dhcp_v6_static_binding": {Tok: makeResource(mainMod, "nsxt_policy_dhcp_v6_static_binding")},
			"nsxt_policy_dns_forwarder_zone": {Tok: makeResource(mainMod, "nsxt_policy_dns_forwarder_zone")},
			"nsxt_policy_gateway_dns_forwarder": {Tok: makeResource(mainMod, "nsxt_policy_gateway_dns_forwarder")},
			"nsxt_policy_gateway_community_list": {Tok: makeResource(mainMod, "nsxt_policy_gateway_community_list")},
			"nsxt_policy_gateway_route_map": {Tok: makeResource(mainMod, "nsxt_policy_gateway_route_map")},
			"nsxt_policy_intrusion_service_policy": {Tok: makeResource(mainMod, "nsxt_policy_intrusion_service_policy")},
			"nsxt_policy_static_route_bfd_peer": {Tok: makeResource(mainMod, "nsxt_policy_static_route_bfd_peer")},
			"nsxt_policy_intrusion_service_profile": {Tok: makeResource(mainMod, "nsxt_policy_intrusion_service_profile")},
			"nsxt_policy_evpn_tenant": {Tok: makeResource(mainMod, "nsxt_policy_evpn_tenant")},
			"nsxt_policy_evpn_config": {Tok: makeResource(mainMod, "nsxt_policy_evpn_config")},
			"nsxt_policy_evpn_tunnel_endpoint": {Tok: makeResource(mainMod, "nsxt_policy_evpn_tunnel_endpoint")},
			"nsxt_policy_vni_pool": {Tok: makeResource(mainMod, "nsxt_policy_vni_pool")},
			"nsxt_policy_qos_profile": {Tok: makeResource(mainMod, "nsxt_policy_qos_profile")},
			"nsxt_policy_ospf_config": {Tok: makeResource(mainMod, "nsxt_policy_ospf_config")},
			"nsxt_policy_ospf_area": {Tok: makeResource(mainMod, "nsxt_policy_ospf_area")},
			"nsxt_policy_gateway_redistribution_config": {Tok: makeResource(mainMod, "nsxt_policy_gateway_redistribution_config")},
			"nsxt_policy_mac_discovery_profile": {Tok: makeResource(mainMod, "nsxt_policy_mac_discovery_profile")},
			"nsxt_policy_ipsec_vpn_ike_profile": {Tok: makeResource(mainMod, "nsxt_policy_ipsec_vpn_ike_profile")},
			"nsxt_policy_ipsec_vpn_tunnel_profile": {Tok: makeResource(mainMod, "nsxt_policy_ipsec_vpn_tunnel_profile")},
			"nsxt_policy_ipsec_vpn_dpd_profile": {Tok: makeResource(mainMod, "nsxt_policy_ipsec_vpn_dpd_profile")},
			"nsxt_policy_ipsec_vpn_session": {Tok: makeResource(mainMod, "nsxt_policy_ipsec_vpn_session")},
			"nsxt_policy_l2_vpn_session": {Tok: makeResource(mainMod, "nsxt_policy_l2_vpn_session")},
			"nsxt_policy_ipsec_vpn_service": {Tok: makeResource(mainMod, "nsxt_policy_ipsec_vpn_service")},
			"nsxt_policy_l2_vpn_service": {Tok: makeResource(mainMod, "nsxt_policy_l2_vpn_service")},
			"nsxt_policy_ipsec_vpn_local_endpoint": {Tok: makeResource(mainMod, "nsxt_policy_ipsec_vpn_local_endpoint")},
			"nsxt_policy_ip_discovery_profile": {Tok: makeResource(mainMod, "nsxt_policy_ip_discovery_profile")},
			"nsxt_policy_context_profile_custom_attribute": {Tok: makeResource(mainMod, "nsxt_policy_context_profile_custom_attribute")},
			"nsxt_policy_segment_security_profile": {Tok: makeResource(mainMod, "nsxt_policy_segment_security_profile")},
			"nsxt_policy_spoof_guard_profile": {Tok: makeResource(mainMod, "nsxt_policy_spoof_guard_profile")},
			"nsxt_policy_gateway_qos_profile": {Tok: makeResource(mainMod, "nsxt_policy_gateway_qos_profile")},
			"nsxt_policy_project": {Tok: makeResource(mainMod, "nsxt_policy_project")},
			"nsxt_policy_transport_zone": {Tok: makeResource(mainMod, "nsxt_policy_transport_zone")},
			"nsxt_edge_cluster": {Tok: makeResource(mainMod, "nsxt_edge_cluster")},
			"nsxt_compute_manager": {Tok: makeResource(mainMod, "nsxt_compute_manager")},
			"nsxt_manager_cluster": {Tok: makeResource(mainMod, "nsxt_manager_cluster")},
			"nsxt_uplink_host_switch_profile": {Tok: makeResource(mainMod, "nsxt_uplink_host_switch_profile")},
			"nsxt_transport_node": {Tok: makeResource(mainMod, "nsxt_transport_node")},
			"nsxt_failure_domain": {Tok: makeResource(mainMod, "nsxt_failure_domain")},
			"nsxt_cluster_virtual_ip": {Tok: makeResource(mainMod, "nsxt_cluster_virtual_ip")},
			"nsxt_policy_host_transport_node_profile": {Tok: makeResource(mainMod, "nsxt_policy_host_transport_node_profile")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: makeDataSource(mainMod, "aws_ami")},
			"nsxt_provider_info": {Tok: makeDataSource(mainMod, "nsxt_provider_info")},
			"nsxt_transport_zone": {Tok: makeDataSource(mainMod, "nsxt_transport_zone")},
			"nsxt_switching_profile": {Tok: makeDataSource(mainMod, "nsxt_switching_profile")},
			"nsxt_logical_tier0_router": {Tok: makeDataSource(mainMod, "nsxt_logical_tier0_router")},
			"nsxt_logical_tier1_router": {Tok: makeDataSource(mainMod, "nsxt_logical_tier1_router")},
			"nsxt_mac_pool": {Tok: makeDataSource(mainMod, "nsxt_mac_pool")},
			"nsxt_ns_group": {Tok: makeDataSource(mainMod, "nsxt_ns_group")},
			"nsxt_ns_groups": {Tok: makeDataSource(mainMod, "nsxt_ns_groups")},
			"nsxt_ns_service": {Tok: makeDataSource(mainMod, "nsxt_ns_service")},
			"nsxt_ns_services": {Tok: makeDataSource(mainMod, "nsxt_ns_services")},
			"nsxt_edge_cluster": {Tok: makeDataSource(mainMod, "nsxt_edge_cluster")},
			"nsxt_certificate": {Tok: makeDataSource(mainMod, "nsxt_certificate")},
			"nsxt_ip_pool": {Tok: makeDataSource(mainMod, "nsxt_ip_pool")},
			"nsxt_firewall_section": {Tok: makeDataSource(mainMod, "nsxt_firewall_section")},
			"nsxt_management_cluster": {Tok: makeDataSource(mainMod, "nsxt_management_cluster")},
			"nsxt_policy_edge_cluster": {Tok: makeDataSource(mainMod, "nsxt_policy_edge_cluster")},
			"nsxt_policy_edge_node": {Tok: makeDataSource(mainMod, "nsxt_policy_edge_node")},
			"nsxt_policy_tier0_gateway": {Tok: makeDataSource(mainMod, "nsxt_policy_tier0_gateway")},
			"nsxt_policy_tier1_gateway": {Tok: makeDataSource(mainMod, "nsxt_policy_tier1_gateway")},
			"nsxt_policy_service": {Tok: makeDataSource(mainMod, "nsxt_policy_service")},
			"nsxt_policy_realization_info": {Tok: makeDataSource(mainMod, "nsxt_policy_realization_info")},
			"nsxt_policy_segment_realization": {Tok: makeDataSource(mainMod, "nsxt_policy_segment_realization")},
			"nsxt_policy_transport_zone": {Tok: makeDataSource(mainMod, "nsxt_policy_transport_zone")},
			"nsxt_policy_ip_discovery_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_ip_discovery_profile")},
			"nsxt_policy_spoofguard_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_spoofguard_profile")},
			"nsxt_policy_qos_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_qos_profile")},
			"nsxt_policy_ipv6_ndra_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_ipv6_ndra_profile")},
			"nsxt_policy_ipv6_dad_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_ipv6_dad_profile")},
			"nsxt_policy_gateway_qos_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_gateway_qos_profile")},
			"nsxt_policy_segment_security_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_segment_security_profile")},
			"nsxt_policy_mac_discovery_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_mac_discovery_profile")},
			"nsxt_policy_vm": {Tok: makeDataSource(mainMod, "nsxt_policy_vm")},
			"nsxt_policy_vms": {Tok: makeDataSource(mainMod, "nsxt_policy_vms")},
			"nsxt_policy_lb_app_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_lb_app_profile")},
			"nsxt_policy_lb_client_ssl_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_lb_client_ssl_profile")},
			"nsxt_policy_lb_server_ssl_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_lb_server_ssl_profile")},
			"nsxt_policy_lb_monitor": {Tok: makeDataSource(mainMod, "nsxt_policy_lb_monitor")},
			"nsxt_policy_certificate": {Tok: makeDataSource(mainMod, "nsxt_policy_certificate")},
			"nsxt_policy_lb_persistence_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_lb_persistence_profile")},
			"nsxt_policy_vni_pool": {Tok: makeDataSource(mainMod, "nsxt_policy_vni_pool")},
			"nsxt_policy_ip_block": {Tok: makeDataSource(mainMod, "nsxt_policy_ip_block")},
			"nsxt_policy_ip_pool": {Tok: makeDataSource(mainMod, "nsxt_policy_ip_pool")},
			"nsxt_policy_site": {Tok: makeDataSource(mainMod, "nsxt_policy_site")},
			"nsxt_policy_gateway_policy": {Tok: makeDataSource(mainMod, "nsxt_policy_gateway_policy")},
			"nsxt_policy_security_policy": {Tok: makeDataSource(mainMod, "nsxt_policy_security_policy")},
			"nsxt_policy_group": {Tok: makeDataSource(mainMod, "nsxt_policy_group")},
			"nsxt_policy_context_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_context_profile")},
			"nsxt_policy_dhcp_server": {Tok: makeDataSource(mainMod, "nsxt_policy_dhcp_server")},
			"nsxt_policy_bfd_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_bfd_profile")},
			"nsxt_policy_intrusion_service_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_intrusion_service_profile")},
			"nsxt_policy_lb_service": {Tok: makeDataSource(mainMod, "nsxt_policy_lb_service")},
			"nsxt_policy_gateway_locale_service": {Tok: makeDataSource(mainMod, "nsxt_policy_gateway_locale_service")},
			"nsxt_policy_bridge_profile": {Tok: makeDataSource(mainMod, "nsxt_policy_bridge_profile")},
			"nsxt_policy_ipsec_vpn_local_endpoint": {Tok: makeDataSource(mainMod, "nsxt_policy_ipsec_vpn_local_endpoint")},
			"nsxt_policy_ipsec_vpn_service": {Tok: makeDataSource(mainMod, "nsxt_policy_ipsec_vpn_service")},
			"nsxt_policy_l2_vpn_service": {Tok: makeDataSource(mainMod, "nsxt_policy_l2_vpn_service")},
			"nsxt_policy_segment": {Tok: makeDataSource(mainMod, "nsxt_policy_segment")},
			"nsxt_policy_project": {Tok: makeDataSource(mainMod, "nsxt_policy_project")},
			"nsxt_policy_gateway_prefix_list": {Tok: makeDataSource(mainMod, "nsxt_policy_gateway_prefix_list")},
			"nsxt_policy_gateway_route_map": {Tok: makeDataSource(mainMod, "nsxt_policy_gateway_route_map")},
			"nsxt_uplink_host_switch_profile": {Tok: makeDataSource(mainMod, "nsxt_uplink_host_switch_profile")},
			"nsxt_compute_manager": {Tok: makeDataSource(mainMod, "nsxt_compute_manager")},
			"nsxt_transport_node_realization": {Tok: makeDataSource(mainMod, "nsxt_transport_node_realization")},
			"nsxt_failure_domain": {Tok: makeDataSource(mainMod, "nsxt_failure_domain")},
			"nsxt_compute_collection": {Tok: makeDataSource(mainMod, "nsxt_compute_collection")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@SCC-Hyperscale-fr/nsxt",

			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "hyperscale_pulumi_nsxt",

			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/SCC-Hyperscale-fr/pulumi-%[1]s/sdk/", "nsxt"),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				"nsxt",
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "SCC-Hyperscale-fr",

			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.hyperscale",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}

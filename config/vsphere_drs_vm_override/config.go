package vSphereDrsVmOverride

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_drs_vm_override", func(r *config.Resource) {
		r.ShortGroup = "Host_and_Cluster_Management"
		r.Kind = "VSphereDrsVmOverride"
		r.Version = "v1alpha1"
	})
}
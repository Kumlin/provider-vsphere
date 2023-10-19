package vspherevmstoragepolicy

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_vm_storage_policy", func(r *config.Resource) {
		r.ShortGroup = "Storage"
		r.Kind = "VSphereVmStoragePolicy"
		r.Version = "v1alpha1"
	})
}

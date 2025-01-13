/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"

	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha12 "github.com/kirillinda/provider-vsphere/apis/hostandclustermanagement/v1alpha1"
	v1alpha1 "github.com/kirillinda/provider-vsphere/apis/inventory/v1alpha1"
	v1alpha11 "github.com/kirillinda/provider-vsphere/apis/storage/v1alpha1"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this VSphereVirtualDisk.
func (mg *VSphereVirtualDisk) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Datacenter),
		Extract:      resource.ExtractParamPath("Name", true),
		Reference:    mg.Spec.ForProvider.DatacenterRef,
		Selector:     mg.Spec.ForProvider.DatacenterSelector,
		To: reference.To{
			List:    &v1alpha1.VSphereDatacenterList{},
			Managed: &v1alpha1.VSphereDatacenter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Datacenter")
	}
	mg.Spec.ForProvider.Datacenter = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatacenterRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Datastore),
		Extract:      resource.ExtractParamPath("Name", true),
		Reference:    mg.Spec.ForProvider.DatastoreRef,
		Selector:     mg.Spec.ForProvider.DatastoreSelector,
		To: reference.To{
			List:    &v1alpha11.VSphereVmfsDatastoreList{},
			Managed: &v1alpha11.VSphereVmfsDatastore{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Datastore")
	}
	mg.Spec.ForProvider.Datastore = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatastoreRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VSphereVirtualMachine.
func (mg *VSphereVirtualMachine) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatastoreID),
		Extract:      resource.ExtractParamPath("ID", true),
		Reference:    mg.Spec.ForProvider.DatastoreIDRef,
		Selector:     mg.Spec.ForProvider.DatastoreIDSelector,
		To: reference.To{
			List:    &v1alpha11.VSphereVmfsDatastoreList{},
			Managed: &v1alpha11.VSphereVmfsDatastore{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatastoreID")
	}
	mg.Spec.ForProvider.DatastoreID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatastoreIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourcePoolID),
		Extract:      resource.ExtractParamPath("ID", true),
		Reference:    mg.Spec.ForProvider.ResourcePoolIDRef,
		Selector:     mg.Spec.ForProvider.ResourcePoolIDSelector,
		To: reference.To{
			List:    &v1alpha12.VSphereComputeClusterList{},
			Managed: &v1alpha12.VSphereComputeCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourcePoolID")
	}
	mg.Spec.ForProvider.ResourcePoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourcePoolIDRef = rsp.ResolvedReference

	return nil
}

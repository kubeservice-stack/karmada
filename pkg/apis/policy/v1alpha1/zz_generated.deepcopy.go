//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAffinity) DeepCopyInto(out *ClusterAffinity) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.FieldSelector != nil {
		in, out := &in.FieldSelector, &out.FieldSelector
		*out = new(FieldSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterNames != nil {
		in, out := &in.ClusterNames, &out.ClusterNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeClusters != nil {
		in, out := &in.ExcludeClusters, &out.ExcludeClusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAffinity.
func (in *ClusterAffinity) DeepCopy() *ClusterAffinity {
	if in == nil {
		return nil
	}
	out := new(ClusterAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOverridePolicy) DeepCopyInto(out *ClusterOverridePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOverridePolicy.
func (in *ClusterOverridePolicy) DeepCopy() *ClusterOverridePolicy {
	if in == nil {
		return nil
	}
	out := new(ClusterOverridePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterOverridePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOverridePolicyList) DeepCopyInto(out *ClusterOverridePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterOverridePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOverridePolicyList.
func (in *ClusterOverridePolicyList) DeepCopy() *ClusterOverridePolicyList {
	if in == nil {
		return nil
	}
	out := new(ClusterOverridePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterOverridePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPreferences) DeepCopyInto(out *ClusterPreferences) {
	*out = *in
	if in.StaticWeightList != nil {
		in, out := &in.StaticWeightList, &out.StaticWeightList
		*out = make([]StaticClusterWeight, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPreferences.
func (in *ClusterPreferences) DeepCopy() *ClusterPreferences {
	if in == nil {
		return nil
	}
	out := new(ClusterPreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPropagationPolicy) DeepCopyInto(out *ClusterPropagationPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPropagationPolicy.
func (in *ClusterPropagationPolicy) DeepCopy() *ClusterPropagationPolicy {
	if in == nil {
		return nil
	}
	out := new(ClusterPropagationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterPropagationPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPropagationPolicyList) DeepCopyInto(out *ClusterPropagationPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterPropagationPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPropagationPolicyList.
func (in *ClusterPropagationPolicyList) DeepCopy() *ClusterPropagationPolicyList {
	if in == nil {
		return nil
	}
	out := new(ClusterPropagationPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterPropagationPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandArgsOverrider) DeepCopyInto(out *CommandArgsOverrider) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandArgsOverrider.
func (in *CommandArgsOverrider) DeepCopy() *CommandArgsOverrider {
	if in == nil {
		return nil
	}
	out := new(CommandArgsOverrider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldSelector) DeepCopyInto(out *FieldSelector) {
	*out = *in
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]corev1.NodeSelectorRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldSelector.
func (in *FieldSelector) DeepCopy() *FieldSelector {
	if in == nil {
		return nil
	}
	out := new(FieldSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageOverrider) DeepCopyInto(out *ImageOverrider) {
	*out = *in
	if in.Predicate != nil {
		in, out := &in.Predicate, &out.Predicate
		*out = new(ImagePredicate)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageOverrider.
func (in *ImageOverrider) DeepCopy() *ImageOverrider {
	if in == nil {
		return nil
	}
	out := new(ImageOverrider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePredicate) DeepCopyInto(out *ImagePredicate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePredicate.
func (in *ImagePredicate) DeepCopy() *ImagePredicate {
	if in == nil {
		return nil
	}
	out := new(ImagePredicate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OverridePolicy) DeepCopyInto(out *OverridePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverridePolicy.
func (in *OverridePolicy) DeepCopy() *OverridePolicy {
	if in == nil {
		return nil
	}
	out := new(OverridePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OverridePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OverridePolicyList) DeepCopyInto(out *OverridePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OverridePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverridePolicyList.
func (in *OverridePolicyList) DeepCopy() *OverridePolicyList {
	if in == nil {
		return nil
	}
	out := new(OverridePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OverridePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OverrideSpec) DeepCopyInto(out *OverrideSpec) {
	*out = *in
	if in.ResourceSelectors != nil {
		in, out := &in.ResourceSelectors, &out.ResourceSelectors
		*out = make([]ResourceSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OverrideRules != nil {
		in, out := &in.OverrideRules, &out.OverrideRules
		*out = make([]RuleWithCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TargetCluster != nil {
		in, out := &in.TargetCluster, &out.TargetCluster
		*out = new(ClusterAffinity)
		(*in).DeepCopyInto(*out)
	}
	in.Overriders.DeepCopyInto(&out.Overriders)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverrideSpec.
func (in *OverrideSpec) DeepCopy() *OverrideSpec {
	if in == nil {
		return nil
	}
	out := new(OverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Overriders) DeepCopyInto(out *Overriders) {
	*out = *in
	if in.Plaintext != nil {
		in, out := &in.Plaintext, &out.Plaintext
		*out = make([]PlaintextOverrider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageOverrider != nil {
		in, out := &in.ImageOverrider, &out.ImageOverrider
		*out = make([]ImageOverrider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CommandOverrider != nil {
		in, out := &in.CommandOverrider, &out.CommandOverrider
		*out = make([]CommandArgsOverrider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ArgsOverrider != nil {
		in, out := &in.ArgsOverrider, &out.ArgsOverrider
		*out = make([]CommandArgsOverrider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Overriders.
func (in *Overriders) DeepCopy() *Overriders {
	if in == nil {
		return nil
	}
	out := new(Overriders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Placement) DeepCopyInto(out *Placement) {
	*out = *in
	if in.ClusterAffinity != nil {
		in, out := &in.ClusterAffinity, &out.ClusterAffinity
		*out = new(ClusterAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterTolerations != nil {
		in, out := &in.ClusterTolerations, &out.ClusterTolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SpreadConstraints != nil {
		in, out := &in.SpreadConstraints, &out.SpreadConstraints
		*out = make([]SpreadConstraint, len(*in))
		copy(*out, *in)
	}
	if in.ReplicaScheduling != nil {
		in, out := &in.ReplicaScheduling, &out.ReplicaScheduling
		*out = new(ReplicaSchedulingStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Placement.
func (in *Placement) DeepCopy() *Placement {
	if in == nil {
		return nil
	}
	out := new(Placement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlaintextOverrider) DeepCopyInto(out *PlaintextOverrider) {
	*out = *in
	in.Value.DeepCopyInto(&out.Value)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlaintextOverrider.
func (in *PlaintextOverrider) DeepCopy() *PlaintextOverrider {
	if in == nil {
		return nil
	}
	out := new(PlaintextOverrider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropagationPolicy) DeepCopyInto(out *PropagationPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropagationPolicy.
func (in *PropagationPolicy) DeepCopy() *PropagationPolicy {
	if in == nil {
		return nil
	}
	out := new(PropagationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PropagationPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropagationPolicyList) DeepCopyInto(out *PropagationPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PropagationPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropagationPolicyList.
func (in *PropagationPolicyList) DeepCopy() *PropagationPolicyList {
	if in == nil {
		return nil
	}
	out := new(PropagationPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PropagationPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropagationSpec) DeepCopyInto(out *PropagationSpec) {
	*out = *in
	if in.ResourceSelectors != nil {
		in, out := &in.ResourceSelectors, &out.ResourceSelectors
		*out = make([]ResourceSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Placement.DeepCopyInto(&out.Placement)
	if in.DependentOverrides != nil {
		in, out := &in.DependentOverrides, &out.DependentOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropagationSpec.
func (in *PropagationSpec) DeepCopy() *PropagationSpec {
	if in == nil {
		return nil
	}
	out := new(PropagationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSchedulingPolicy) DeepCopyInto(out *ReplicaSchedulingPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSchedulingPolicy.
func (in *ReplicaSchedulingPolicy) DeepCopy() *ReplicaSchedulingPolicy {
	if in == nil {
		return nil
	}
	out := new(ReplicaSchedulingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicaSchedulingPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSchedulingPolicyList) DeepCopyInto(out *ReplicaSchedulingPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicaSchedulingPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSchedulingPolicyList.
func (in *ReplicaSchedulingPolicyList) DeepCopy() *ReplicaSchedulingPolicyList {
	if in == nil {
		return nil
	}
	out := new(ReplicaSchedulingPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicaSchedulingPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSchedulingSpec) DeepCopyInto(out *ReplicaSchedulingSpec) {
	*out = *in
	if in.ResourceSelectors != nil {
		in, out := &in.ResourceSelectors, &out.ResourceSelectors
		*out = make([]ResourceSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Preferences.DeepCopyInto(&out.Preferences)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSchedulingSpec.
func (in *ReplicaSchedulingSpec) DeepCopy() *ReplicaSchedulingSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicaSchedulingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSchedulingStrategy) DeepCopyInto(out *ReplicaSchedulingStrategy) {
	*out = *in
	if in.WeightPreference != nil {
		in, out := &in.WeightPreference, &out.WeightPreference
		*out = new(ClusterPreferences)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSchedulingStrategy.
func (in *ReplicaSchedulingStrategy) DeepCopy() *ReplicaSchedulingStrategy {
	if in == nil {
		return nil
	}
	out := new(ReplicaSchedulingStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSelector) DeepCopyInto(out *ResourceSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSelector.
func (in *ResourceSelector) DeepCopy() *ResourceSelector {
	if in == nil {
		return nil
	}
	out := new(ResourceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleWithCluster) DeepCopyInto(out *RuleWithCluster) {
	*out = *in
	if in.TargetCluster != nil {
		in, out := &in.TargetCluster, &out.TargetCluster
		*out = new(ClusterAffinity)
		(*in).DeepCopyInto(*out)
	}
	in.Overriders.DeepCopyInto(&out.Overriders)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleWithCluster.
func (in *RuleWithCluster) DeepCopy() *RuleWithCluster {
	if in == nil {
		return nil
	}
	out := new(RuleWithCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpreadConstraint) DeepCopyInto(out *SpreadConstraint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpreadConstraint.
func (in *SpreadConstraint) DeepCopy() *SpreadConstraint {
	if in == nil {
		return nil
	}
	out := new(SpreadConstraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticClusterWeight) DeepCopyInto(out *StaticClusterWeight) {
	*out = *in
	in.TargetCluster.DeepCopyInto(&out.TargetCluster)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticClusterWeight.
func (in *StaticClusterWeight) DeepCopy() *StaticClusterWeight {
	if in == nil {
		return nil
	}
	out := new(StaticClusterWeight)
	in.DeepCopyInto(out)
	return out
}

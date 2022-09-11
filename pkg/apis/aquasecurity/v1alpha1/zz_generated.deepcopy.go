//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/aquasecurity/trivy-db/pkg/types"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Artifact) DeepCopyInto(out *Artifact) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Artifact.
func (in *Artifact) DeepCopy() *Artifact {
	if in == nil {
		return nil
	}
	out := new(Artifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Check) DeepCopyInto(out *Check) {
	*out = *in
	if in.Messages != nil {
		in, out := &in.Messages, &out.Messages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(CheckScope)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Check.
func (in *Check) DeepCopy() *Check {
	if in == nil {
		return nil
	}
	out := new(Check)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheckScope) DeepCopyInto(out *CheckScope) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckScope.
func (in *CheckScope) DeepCopy() *CheckScope {
	if in == nil {
		return nil
	}
	out := new(CheckScope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComplianceDetailReport) DeepCopyInto(out *ClusterComplianceDetailReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComplianceDetailReport.
func (in *ClusterComplianceDetailReport) DeepCopy() *ClusterComplianceDetailReport {
	if in == nil {
		return nil
	}
	out := new(ClusterComplianceDetailReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterComplianceDetailReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComplianceDetailReportData) DeepCopyInto(out *ClusterComplianceDetailReportData) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
	out.Type = in.Type
	out.Summary = in.Summary
	if in.ControlChecks != nil {
		in, out := &in.ControlChecks, &out.ControlChecks
		*out = make([]ControlCheckDetails, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComplianceDetailReportData.
func (in *ClusterComplianceDetailReportData) DeepCopy() *ClusterComplianceDetailReportData {
	if in == nil {
		return nil
	}
	out := new(ClusterComplianceDetailReportData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComplianceDetailReportList) DeepCopyInto(out *ClusterComplianceDetailReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterComplianceReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComplianceDetailReportList.
func (in *ClusterComplianceDetailReportList) DeepCopy() *ClusterComplianceDetailReportList {
	if in == nil {
		return nil
	}
	out := new(ClusterComplianceDetailReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterComplianceDetailReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComplianceReport) DeepCopyInto(out *ClusterComplianceReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComplianceReport.
func (in *ClusterComplianceReport) DeepCopy() *ClusterComplianceReport {
	if in == nil {
		return nil
	}
	out := new(ClusterComplianceReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterComplianceReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComplianceReportList) DeepCopyInto(out *ClusterComplianceReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterComplianceReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComplianceReportList.
func (in *ClusterComplianceReportList) DeepCopy() *ClusterComplianceReportList {
	if in == nil {
		return nil
	}
	out := new(ClusterComplianceReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterComplianceReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComplianceSummary) DeepCopyInto(out *ClusterComplianceSummary) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComplianceSummary.
func (in *ClusterComplianceSummary) DeepCopy() *ClusterComplianceSummary {
	if in == nil {
		return nil
	}
	out := new(ClusterComplianceSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigAuditReport) DeepCopyInto(out *ClusterConfigAuditReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigAuditReport.
func (in *ClusterConfigAuditReport) DeepCopy() *ClusterConfigAuditReport {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigAuditReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfigAuditReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigAuditReportList) DeepCopyInto(out *ClusterConfigAuditReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterConfigAuditReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigAuditReportList.
func (in *ClusterConfigAuditReportList) DeepCopy() *ClusterConfigAuditReportList {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigAuditReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfigAuditReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRbacAssessmentReport) DeepCopyInto(out *ClusterRbacAssessmentReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRbacAssessmentReport.
func (in *ClusterRbacAssessmentReport) DeepCopy() *ClusterRbacAssessmentReport {
	if in == nil {
		return nil
	}
	out := new(ClusterRbacAssessmentReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterRbacAssessmentReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRbacAssessmentReportList) DeepCopyInto(out *ClusterRbacAssessmentReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterRbacAssessmentReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRbacAssessmentReportList.
func (in *ClusterRbacAssessmentReportList) DeepCopy() *ClusterRbacAssessmentReportList {
	if in == nil {
		return nil
	}
	out := new(ClusterRbacAssessmentReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterRbacAssessmentReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Compliance) DeepCopyInto(out *Compliance) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Compliance.
func (in *Compliance) DeepCopy() *Compliance {
	if in == nil {
		return nil
	}
	out := new(Compliance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigAuditReport) DeepCopyInto(out *ConfigAuditReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigAuditReport.
func (in *ConfigAuditReport) DeepCopy() *ConfigAuditReport {
	if in == nil {
		return nil
	}
	out := new(ConfigAuditReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigAuditReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigAuditReportData) DeepCopyInto(out *ConfigAuditReportData) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
	out.Scanner = in.Scanner
	out.Summary = in.Summary
	if in.Checks != nil {
		in, out := &in.Checks, &out.Checks
		*out = make([]Check, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigAuditReportData.
func (in *ConfigAuditReportData) DeepCopy() *ConfigAuditReportData {
	if in == nil {
		return nil
	}
	out := new(ConfigAuditReportData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigAuditReportList) DeepCopyInto(out *ConfigAuditReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigAuditReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigAuditReportList.
func (in *ConfigAuditReportList) DeepCopy() *ConfigAuditReportList {
	if in == nil {
		return nil
	}
	out := new(ConfigAuditReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigAuditReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigAuditSummary) DeepCopyInto(out *ConfigAuditSummary) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigAuditSummary.
func (in *ConfigAuditSummary) DeepCopy() *ConfigAuditSummary {
	if in == nil {
		return nil
	}
	out := new(ConfigAuditSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Control) DeepCopyInto(out *Control) {
	*out = *in
	if in.Kinds != nil {
		in, out := &in.Kinds, &out.Kinds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Mapping.DeepCopyInto(&out.Mapping)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Control.
func (in *Control) DeepCopy() *Control {
	if in == nil {
		return nil
	}
	out := new(Control)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlCheck) DeepCopyInto(out *ControlCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlCheck.
func (in *ControlCheck) DeepCopy() *ControlCheck {
	if in == nil {
		return nil
	}
	out := new(ControlCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlCheckDetails) DeepCopyInto(out *ControlCheckDetails) {
	*out = *in
	if in.ScannerCheckResult != nil {
		in, out := &in.ScannerCheckResult, &out.ScannerCheckResult
		*out = make([]ScannerCheckResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlCheckDetails.
func (in *ControlCheckDetails) DeepCopy() *ControlCheckDetails {
	if in == nil {
		return nil
	}
	out := new(ControlCheckDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposedSecret) DeepCopyInto(out *ExposedSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposedSecret.
func (in *ExposedSecret) DeepCopy() *ExposedSecret {
	if in == nil {
		return nil
	}
	out := new(ExposedSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposedSecretReport) DeepCopyInto(out *ExposedSecretReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposedSecretReport.
func (in *ExposedSecretReport) DeepCopy() *ExposedSecretReport {
	if in == nil {
		return nil
	}
	out := new(ExposedSecretReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExposedSecretReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposedSecretReportData) DeepCopyInto(out *ExposedSecretReportData) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
	out.Scanner = in.Scanner
	out.Registry = in.Registry
	out.Artifact = in.Artifact
	out.Summary = in.Summary
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]ExposedSecret, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposedSecretReportData.
func (in *ExposedSecretReportData) DeepCopy() *ExposedSecretReportData {
	if in == nil {
		return nil
	}
	out := new(ExposedSecretReportData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposedSecretReportList) DeepCopyInto(out *ExposedSecretReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExposedSecretReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposedSecretReportList.
func (in *ExposedSecretReportList) DeepCopy() *ExposedSecretReportList {
	if in == nil {
		return nil
	}
	out := new(ExposedSecretReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExposedSecretReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposedSecretSummary) DeepCopyInto(out *ExposedSecretSummary) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposedSecretSummary.
func (in *ExposedSecretSummary) DeepCopy() *ExposedSecretSummary {
	if in == nil {
		return nil
	}
	out := new(ExposedSecretSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mapping) DeepCopyInto(out *Mapping) {
	*out = *in
	if in.Checks != nil {
		in, out := &in.Checks, &out.Checks
		*out = make([]SpecCheck, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mapping.
func (in *Mapping) DeepCopy() *Mapping {
	if in == nil {
		return nil
	}
	out := new(Mapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbacAssessmentReport) DeepCopyInto(out *RbacAssessmentReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbacAssessmentReport.
func (in *RbacAssessmentReport) DeepCopy() *RbacAssessmentReport {
	if in == nil {
		return nil
	}
	out := new(RbacAssessmentReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RbacAssessmentReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbacAssessmentReportData) DeepCopyInto(out *RbacAssessmentReportData) {
	*out = *in
	out.Scanner = in.Scanner
	out.Summary = in.Summary
	if in.Checks != nil {
		in, out := &in.Checks, &out.Checks
		*out = make([]Check, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbacAssessmentReportData.
func (in *RbacAssessmentReportData) DeepCopy() *RbacAssessmentReportData {
	if in == nil {
		return nil
	}
	out := new(RbacAssessmentReportData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbacAssessmentReportList) DeepCopyInto(out *RbacAssessmentReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RbacAssessmentReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbacAssessmentReportList.
func (in *RbacAssessmentReportList) DeepCopy() *RbacAssessmentReportList {
	if in == nil {
		return nil
	}
	out := new(RbacAssessmentReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RbacAssessmentReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbacAssessmentSummary) DeepCopyInto(out *RbacAssessmentSummary) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbacAssessmentSummary.
func (in *RbacAssessmentSummary) DeepCopy() *RbacAssessmentSummary {
	if in == nil {
		return nil
	}
	out := new(RbacAssessmentSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportSpec) DeepCopyInto(out *ReportSpec) {
	*out = *in
	if in.Controls != nil {
		in, out := &in.Controls, &out.Controls
		*out = make([]Control, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportSpec.
func (in *ReportSpec) DeepCopy() *ReportSpec {
	if in == nil {
		return nil
	}
	out := new(ReportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportStatus) DeepCopyInto(out *ReportStatus) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
	out.Summary = in.Summary
	if in.ControlChecks != nil {
		in, out := &in.ControlChecks, &out.ControlChecks
		*out = make([]ControlCheck, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportStatus.
func (in *ReportStatus) DeepCopy() *ReportStatus {
	if in == nil {
		return nil
	}
	out := new(ReportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultDetails) DeepCopyInto(out *ResultDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultDetails.
func (in *ResultDetails) DeepCopy() *ResultDetails {
	if in == nil {
		return nil
	}
	out := new(ResultDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scanner) DeepCopyInto(out *Scanner) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scanner.
func (in *Scanner) DeepCopy() *Scanner {
	if in == nil {
		return nil
	}
	out := new(Scanner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScannerCheckResult) DeepCopyInto(out *ScannerCheckResult) {
	*out = *in
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]ResultDetails, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScannerCheckResult.
func (in *ScannerCheckResult) DeepCopy() *ScannerCheckResult {
	if in == nil {
		return nil
	}
	out := new(ScannerCheckResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecCheck) DeepCopyInto(out *SpecCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecCheck.
func (in *SpecCheck) DeepCopy() *SpecCheck {
	if in == nil {
		return nil
	}
	out := new(SpecCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vulnerability) DeepCopyInto(out *Vulnerability) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Score != nil {
		in, out := &in.Score, &out.Score
		*out = new(float64)
		**out = **in
	}
	if in.CVSS != nil {
		in, out := &in.CVSS, &out.CVSS
		*out = make(types.VendorCVSS, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vulnerability.
func (in *Vulnerability) DeepCopy() *Vulnerability {
	if in == nil {
		return nil
	}
	out := new(Vulnerability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityReport) DeepCopyInto(out *VulnerabilityReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityReport.
func (in *VulnerabilityReport) DeepCopy() *VulnerabilityReport {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VulnerabilityReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityReportData) DeepCopyInto(out *VulnerabilityReportData) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
	out.Scanner = in.Scanner
	out.Registry = in.Registry
	out.Artifact = in.Artifact
	out.Summary = in.Summary
	if in.Vulnerabilities != nil {
		in, out := &in.Vulnerabilities, &out.Vulnerabilities
		*out = make([]Vulnerability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityReportData.
func (in *VulnerabilityReportData) DeepCopy() *VulnerabilityReportData {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityReportData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityReportList) DeepCopyInto(out *VulnerabilityReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VulnerabilityReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityReportList.
func (in *VulnerabilityReportList) DeepCopy() *VulnerabilityReportList {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VulnerabilityReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilitySummary) DeepCopyInto(out *VulnerabilitySummary) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilitySummary.
func (in *VulnerabilitySummary) DeepCopy() *VulnerabilitySummary {
	if in == nil {
		return nil
	}
	out := new(VulnerabilitySummary)
	in.DeepCopyInto(out)
	return out
}

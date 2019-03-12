package tomcat

import (
	"fmt"

	"k8s.io/apimachinery/pkg/labels"

	tomcatv1alpha1 "github.com/kube-incubator/tomcat-operator/pkg/apis/tomcat/v1alpha1"
)

// Tomcat embeds tomcatv1alpha1.Tomcat and adds utility functions
type Tomcat struct {
	*tomcatv1alpha1.Tomcat
}

type component struct {
	name       string
	objNameFmt string
	objName    string
}

var (
	// TomcatDeployment component
	TomcatDeployment = component{name: "tomcat-server", objNameFmt: "%s"}
	// TomcatService component
	TomcatService = component{name: "tomcat-server", objNameFmt: "%s"}
)

// New wraps a tomcatv1alpha1.Tomcat into a Tomcat object
func New(obj *tomcatv1alpha1.Tomcat) *Tomcat {
	return &Tomcat{obj}
}

// Unwrap returns the wrapped tomcatv1alpha1.Tomcat object
func (o *Tomcat) Unwrap() *tomcatv1alpha1.Tomcat {
	return o.Tomcat
}

// Labels returns default label set for tomcatv1alpha1.Tomcat
func (o *Tomcat) Labels() labels.Set {
	partOf := "tomcat"
	if o.ObjectMeta.Labels != nil && len(o.ObjectMeta.Labels["app.kubernetes.io/part-of"]) > 0 {
		partOf = o.ObjectMeta.Labels["app.kubernetes.io/part-of"]
	}

	labels := labels.Set{
		"app.kubernetes.io/name":     "tomcat",
		"app.kubernetes.io/part-of":  partOf,
		"app.kubernetes.io/instance": o.ObjectMeta.Name,
	}

	return labels
}

// ComponentLabels returns labels for a label set for a tomcatv1alpha1.Tomcat component
func (o *Tomcat) ComponentLabels(component component) labels.Set {
	l := o.Labels()
	l["app.kubernetes.io/component"] = component.name
	return l
}

// ComponentName returns the object name for a component
func (o *Tomcat) ComponentName(component component) string {
	name := component.objName
	if len(component.objNameFmt) > 0 {
		name = fmt.Sprintf(component.objNameFmt, o.ObjectMeta.Name)
	}

	return name
}

// TomcatServerPodLabels return labels to apply to tomcat server pods
func (o *Tomcat) TomcatServerPodLabels() labels.Set {
	l := o.Labels()
	l["app.kubernetes.io/component"] = "tomcat-server"
	return l
}

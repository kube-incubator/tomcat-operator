package tomcat

var (
	defaultServicePort int32 = 80
)

const (
	defaultImage           string = "tomcat:latest"
	defaultWebArchiveImage string = "ananwaresystems/webarchive:1.0"
	defaultDeployDirectory string = "/usr/local/tomcat/webapps"
)

// SetDefaults sets Tomcat field defaults
func (o *Tomcat) SetDefaults() {

	if len(o.Spec.Image) == 0 {
		o.Spec.Image = defaultImage
	}

	if o.Spec.ServicePort == nil {
		o.Spec.ServicePort = &defaultServicePort
	}

	if len(o.Spec.WebArchiveImage) == 0 {
		o.Spec.WebArchiveImage = defaultWebArchiveImage
	}

	if len(o.Spec.DeployDirectory) == 0 {
		o.Spec.DeployDirectory = defaultDeployDirectory
	}
}

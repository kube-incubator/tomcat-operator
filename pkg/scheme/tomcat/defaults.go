package tomcat

var (
	defaultServicePort int32 = 80
)

const (
	defaultImage string = "tomcat:latest"
)

// SetDefaults sets Tomcat field defaults
func (o *Tomcat) SetDefaults() {

	if len(o.Spec.Image) == 0 {
		o.Spec.Image = defaultImage
	}

	if o.Spec.ServicePort == nil {
		o.Spec.ServicePort = &defaultServicePort
	}
}

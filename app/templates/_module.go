package <%=modelName%>

import (
    "go.viam.com/rdk/components/<%=apiName%>"
	"go.viam.com/rdk/resource"
	<%- moreImports %>
)

var Model = resource.NewModel("<%=nameSpace%>", "<%=moduleName%>", " <%=modelName%>")



func init() {
	resource.RegisterComponent(<%=apiName%>.API, Model,
		resource.Registration[<%=apiName%>.<%=apiNameUppercase%>,  *Config]{
			Constructor: new<%=moduleName%><%=modelName%>,
		},
	)
}

type Config struct {
	// Put config attributes here
	// if you dont have config add nonativeconfgif

	// trivallyvalidate if you want

}

func (cfg *Config) Validate(path string) ([]string, error) {
	// put config validation here
	 return nil, nil
}

// add services and more than one module



// make this private
// register with no native config
type <%=moduleName%><%=modelName%> struct {
	name   resource.Name

	// if docommand do named
	logger logging.Logger
	cfg    *Config

	cancelCtx  context.Context
	cancelFunc func()
	// if you dont want a close or reconfigure
	// TrivallyCloseable
	// if you dont have any goroutines you can do trivally closeable
	// TrivallyReconfigurable



}


func new<%=moduleName%><%=modelName%>(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (<%=apiName%>.<%=apiNameUppercase%>, error) {
	conf, err := resource.NativeConfig[*Config](rawConf)
	if err != nil {
		return nil, err
	}

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	s := &<%=moduleName%><%=modelName%>{
		name:       rawConf.ResourceName(),
		logger:     logger,
		cfg:        conf,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
	}
	return s, nil
}


func (s *<%=moduleName%><%=modelName%>) Name() resource.Name {
	return s.name
}

// Reconfigures the model. Most models can be reconfigured in place without needing to rebuild. If you need to instead create a new instance of the sensor, throw a NewMustBuildError.
func (s *<%=moduleName%><%=modelName%>) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {
	// put reconfigure code here
	return nil
}

<%=funcs%>

func (s *<%=moduleName%><%=modelName%>) Close() error {
	// put close code here
	return nil
}

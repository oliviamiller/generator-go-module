package <%=modelName%>

import (
    "go.viam.com/rdk/<%=resourceTypeLower%>s/<%=apiName%>"
	<%- moreImports %>
)

var Model = resource.NewModel("<%=nameSpace%>", "<%=moduleName%>", "<%=modelName%>")

func init() {
	resource.Register<%=resourceType%>(<%=apiName%>.API, Model,
		resource.Registration[<%=apiName%>.<%=objName%>,  *Config]{
			Constructor: new<%=moduleName%><%=modelName%>,
		},
	)
}

type Config struct {
	// Put config attributes here

	/* if your model  does not need a config,
	   replace *Config on line 13 with resource.NoNativeConfig */

	/* Uncomment this if your model does not need to be validated
	    and has no implicit dependecies. */
	// resource.TriviallyValidateConfig

}

func (cfg *Config) Validate(path string) ([]string, error) {
	// Add config validation code here
	 return nil, nil
}

type <%=moduleName%><%=modelName%> struct {
	name   resource.Name

	logger logging.Logger
	cfg    *Config

	cancelCtx  context.Context
	cancelFunc func()


	/* Uncomment this if your model does not need to reconfigure. */
	// resource.TriviallyReconfigurable

	// Uncomment this if the model does not have any goroutines that
	// need to be shut down while closing.
	// resource.TriviallyCloseable

}


func new<%=moduleName%><%=modelName%>(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (<%=apiName%>.<%=objName%>, error) {
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

func (s *<%=moduleName%><%=modelName%>) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {
	// Put reconfigure code here
	return nil
}

<%=funcs%>

func (s *<%=moduleName%><%=modelName%>) Close(context.Context) error {
	return nil
}

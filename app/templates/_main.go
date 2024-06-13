package main

import (
	"context"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
	<% let i = 0 %>
	<% for (api of apis) { %>
		<%=models[i]%> "<%=moduleName%>/<%=models[i]%>"
		"go.viam.com/rdk/<%=resourceTypeLower%>s/<%=api%>"
		<% i += 1 %>
   <% } %>
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("<%=moduleName%>"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) error {
	<%= moduleName%>, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}
	<% let j = 0 %>
	<% for (model of models) { %>
		if err = <%= moduleName%>.AddModelFromRegistry(ctx, <%=apis[j]%>.API, <%=model%>.Model); err != nil {
			return err
		}
		<% j += 1 %>
   <% } %>


	err = <%=moduleName%>.Start(ctx)
	defer <%=moduleName%>.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

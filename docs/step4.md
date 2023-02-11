# Microservice tracing

## Status

Adding datadog tracing

## Content

It is using datadog tracing middleware:
  - https://docs.datadoghq.com/tracing/trace_collection/compatibility/go/
  - https://docs.datadoghq.com/tracing/trace_collection/custom_instrumentation/go/

## Code explanation

if you checkout step4 branch, you need to understand the code:

### datadog middleware

By including

```
package config

import (
  ...
	negronitrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/urfave/negroni"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
  ...
)

// SetupGlobalMiddleware setup the global middleware
func SetupGlobalMiddleware(handler http.Handler) http.Handler {

  ...
	if Config.DatadogTraceEnabled {
		tracer.Start()
		n.Use(negronitrace.Middleware(negronitrace.WithServiceName("appconfigr")))
	}
  ...
```

It will add a Datadog middleware that will send back to Datadog a trace each time there is a call to this REST API service.

You can enhance the existing `span` with something like

```
	if config.Config.DatadogTraceEnabled {
		ctx := params.HTTPRequest.Context()
		if span, ok := tracer.SpanFromContext(ctx); ok {
			span.SetTag("update.todoid", params.TodoID)
			span.SetTag("update.title", params.Body.Title)
			span.SetTag("update.body", params.Body.Body)
		}
```

you can also create `sub span` with a code like 

```
	if config.Config.DatadogEnabled {
		subspan, _ := tracer.StartSpanFromContext(ctx, "group")
		subspan.SetTag("graphql.operation.resolver", "groupResolver.Channels")
		if obj != nil {
			subspan.SetTag("graphql.group", obj.Name)
		}
		defer subspan.Finish()
	}
```

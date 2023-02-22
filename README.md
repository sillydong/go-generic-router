# go-generic-router

For so many years, we are creating different web frameworks defining different types of `Request Context` to try to carry data within request life time.
First solution is use `context.WithValue` to wrap data and every time you use the value, you need to force type it. Second solution is like gin framework, use its own stucture to store values.
Now with go generic type, we can define our own Context and pass it through the whole request life time.

**!!NOTE!!: THIS IS ONLY A DEMO NOW, FAR FROM A PRODUCTION READY FRAMEWORK, HOPE YOU CAN GET SOME INSPIRATION FROM IT**

The main idea is define an interface named `Context` and create a base structure for context named `ReqContext`, it will include things like `Logger`, `*http.Request`, `http.ResponseWriter`, etc. All user define context should implement from this structure. See the example code for detail [main.go](./example/main.go).


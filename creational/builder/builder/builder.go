package builder

/*
	A Builder always has a dependency class as a property
	Once the builder is instantiated, it initializes a dependency class instance with some default values
	Then, each property of that dependency is set with its own method, "building" the obj
	Once its done, the method "Build" returns the dependency object complete

	PROS:
	- Readable object construction - Allows you to set up complex structs step by step in a fluent, easy-to-read way.
	- Optional fields handled cleanly - Avoids huge constructors with many optional params.
	- Encapsulation of contstruction logic - You can validate or compute derived values inside Build().
	- Immutability support - You can build immutable objects (return a copy and avoid setters later).
	- Fluent interface - Easy chaining like .SetName().SetEmail().
	CONS:
	- Extra boilerplate - Requires defining and extra builder struct and multiple methods.
	- Overkill for simple structs - In Go, a simple struct literal is often enough.
	- Not idiomatic in Go - Functional options are preferred for optimal configuration.
	- Can hide simple logic - If overused, it makes creation look complex when it isn't

	WHEN TO USE:
	- When your struct has many optional fields or complex validation.
	- When you need a fluent configuration API (e.g., bulding an HTTP client, request, or config).
	- When you want immutability (the final object isn't modified after build).
*/

type Dependency struct {
	property1 string
	property2 int
	property3 bool
}

type Builder struct {
	Dependency Dependency
}

func NewBuilder() *Builder {
	return &Builder{
		Dependency: Dependency{
			property1: "default",
			property2: 1,
			property3: true,
		},
	}
}

func (b *Builder) Property1(value string) *Builder {
	b.Dependency.property1 = value
	return b
}

func (b *Builder) Property2(value int) *Builder {
	b.Dependency.property2 = value
	return b
}

func (b *Builder) Property3(value bool) *Builder {
	b.Dependency.property3 = value
	return b
}

func (b *Builder) Build() *Dependency {
	return &b.Dependency
}

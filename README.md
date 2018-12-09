# Swagger Schema Gen

Generate [Swagger data models (schemas)](https://swagger.io/docs/specification/data-models/) from Go type specifications.


```
Usage of schemagen:
	schemagen [flags] [directory]
Flags:
  -initialisms string
    	comma-separated list of initalism to lowercase in schema object property names
  -output string
    	output file name; if not file is given output is written to stdout
  -tags string
    	comma-separated list of build tags to apply
  -trimprefix prefix
    	trim the prefix from the generated schema object names
```

## Example

Given a package with the following Go types defined:

```
package foo

type Foo struct {
	FooProperty string `json:"fooProperty"`
	Name        string `json:"name"`
}

func (f Foo) GetGreeting(msg string) string {
	return msg + f.Name
}

type Bar struct {
	BarProperty string `json:"barProperty"`
}

type BarExtended struct {
	Bar
	AnotherProperty    string `json:"anotherProperty"`
	YetAnotherProperty int
}

type FooWrapper struct {
	Foo Foo
}

type BarContainer struct {
	Bars []Bar
}

type FooBar struct {
	Foo
	Bar
}

```

Swagger Schema Gen will generate the following Swagger schema:

```
Bar:
  type: object
  properties:
    barProperty:
      type: string
BarContainer:
  type: object
  properties:
    bars:
      type: array
      items:
        $ref: '#/components/schemas/Bar'
BarExtended:
  allOf:
    - $ref: '#/components/schemas/Bar'
    - type: object
      properties:
        anotherProperty:
          type: string
        yetAnotherProperty:
          type: integer
Foo:
  type: object
  properties:
    fooProperty:
      type: string
    name:
      type: string
FooBar:
  allOf:
    - $ref: '#/components/schemas/Foo'
    - $ref: '#/components/schemas/Bar'
FooWrapper:
  type: object
  properties:
    foo:
      $ref: '#/components/schemas/Foo'
```
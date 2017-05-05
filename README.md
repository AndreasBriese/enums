## **enums**: a Go enum generator ##

**enums** uses go generate functionality to produce text/template encoded go files within the same folder, that hold the enum like constant constructs.

in contrast to other approaches **enums** parses a special comment section to compose the enum type definition - instead of AST parsing to get the specs. **enums** will create a go file in the same folder sharing the package name and including all code lines to implement the enum. therefore the the enum's name, type and const definition should not be given in the code go file that calls the generator. 

#### load & install ####

```golang
  $ go get -u github.com/AndreasBriese/enums
```

#### usage ####

for the enum(s) to be generated include a line `//go:generate enums $GOFILE` in your go file and in the same file give the enum(s) specifics in a commented out block starting with `// [@enums` and ending with `// @]` like the below:

```golang

//go:generate enums $GOFILE

// [@enums
// type colors int
// -Grau colr = iota + 1
// -Weiss
// -Gelb
// -Lila = Weiss
// @]

```

#### call #### 

```golang
  $ go generate
```

in the same folder will result in a file "colors_enums.go" that holds this code:

```golang

type colors int

const (
	Grau colors = (iota + 1) * 5
	Weiss
	Gelb
	Lila = Weiss
)

var colors_Constants = [7]colors{
	Grau,
	Weiss,
	Gelb,
	Lila,
}

func (t colors) String() string {
	switch t {
	case Grau:
		return "Grau"
	case Weiss:
		return "Weiss"
	case Gelb:
		return "Gelb"
	}
	return ""
}
```

colors has got a method String() satisfying the interface `fmt.Stringer` to print out the color names and an array `colors_Constants` that can be used to loop over the enum i.e. to search for a color constant.   

note that eventually all information on the type will be public if given a Uppercase type name in the encoding section. 

#### License ####

MIT-License


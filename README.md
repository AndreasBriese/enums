## **enums**: a Go enum generator ##

**enums** uses go generate functionality to produce text/template encoded go files within the same folder, that hold the enum like constant constructs.

in contrast to other approaches **enums** parses a special comment section to compose the enum type definition - instead of AST parsing to get the specs. **enums** will create a go file in the same folder sharing the package name and including all code lines to implement the enum. therefore the the enum's name, type and const definition should not be given in the code of the go file that calls the generator - otherwise a build error because of duplicate definition might occur. 

**enums** deals with double values by commenting out lines that hinder building the output file. try out the example to learn more. 

#### load & install ####

```sh
  $ go get -u github.com/AndreasBriese/enums
```

#### usage ####

for the enum(s) to be generated include a line `//go:generate enums $GOFILE` in your go file and in the same file give the enum(s) specifics in a commented out block starting with `// [@enums` and ending with `// @]` like the below:

```golang

//go:generate enums $GOFILE

// [@enums
// type colors int
// -Grau colors = iota + 1
// -Weiss
// -Gelb
// -Lila = Weiss
// @]

```

and call enums in the folder containing the go file by 

```sh
  $ go generate <YOURFile.go>
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
//	case Lila:
//		return "Lila"
	}
	return ""
}
```

colors has got a method String() satisfying the interface `fmt.Stringer` to print out the color names and an array `colors_Constants` that can be used to loop over the enum i.e. to search for a color constant.   

note that eventually all information on the type will be public if given a Uppercase type name in the encoding section. 

#### try ####

change into example folder and run 

```sh
$ go generate 
generated enum type 'colors' in colors_enums.go
generated enum type 'values' in values_enums.go
$ go run *.go
List of colors constants:
colors_Constants[0] = 1 translating to "Grau"
colors_Constants[1] = 2 translating to "Weiss"
colors_Constants[2] = 3 translating to "Gelb"
colors_Constants[3] = 4 translating to "Grün"
colors_Constants[4] = 5 translating to "Rot"
colors_Constants[5] = 6 translating to "Blau"
colors_Constants[6] = 2 translating to "Weiss"
List of values constants:
values_Constants[0] = 3.141593 translating to "pi"
values_Constants[1] = 6.283185 translating to "pi2"
values_Constants[2] = 9.424778 translating to "pi3"
values_Constants[3] = 12.566371 translating to "pi4"
values_Constants[4] = 3.141593 translating to "pi"
values_Constants[5] = 31.415927 translating to "pi10"
```

#### License ####

MIT-License


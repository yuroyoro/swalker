# swalker

extract value from map/struct by dot notated syntax like "Foo.Bar[0].Baz"

## Usage

```
package main

import (
	"fmt"
	"github.com/yuroyoro/swalker"
)

type A struct {
	Foo *B
}
type B struct {
	Bar []*C
}
type C struct {
	Hoge string
}

func main() {
	obj := A{Foo: &B{Bar: []*C{&C{Hoge: "aaa"}, &C{Hoge: "bbb"}}}}

	v, err := swalker.Read("Foo", obj)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T : %v\n", v, v) // -> *B

	v, err = swalker.Read("Foo.Bar", obj)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T : %v\n", v, v) // -> []*C

	v, err = swalker.Read("Foo.Bar[0]", obj)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T : %v\n", v, v) // -> *C

	v, err = swalker.Read("Foo.Bar[0].Hoge", obj)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T : %v\n", v, v) // -> "aaaa"
}
```

## License

MIT

## Author

Tomothio Ozaki (@yuroyoro)

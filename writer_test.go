package swalker

import (
	"reflect"
	"testing"
)

func TestWriteHash(t *testing.T) {

	obj := makeHash()
	hash := obj.(map[string]interface{})
	fooval := hash["foo"]
	foo0 := fooval.([]interface{})[0]
	bar1 := foo0.(map[string]interface{})["bar1"]
	hbar1 := bar1.(map[string]interface{})

	// Write(obj, "foo.bar1.hoge", "xxxx") string value
	err := Write("foo[0].bar1.hoge", obj, "xxxx")
	if err != nil {
		t.Fatalf("unexpected error : [%s]", err)
	}
	if hbar1["hoge"] != "xxxx" {
		t.Fatalf("could not write : value %v", hbar1["hoge"])
	}
	// Write(obj, "foo.bar1.hoge", "xxxx") new key
	err = Write("foo[0].bar1.fuga", obj, "wwww")
	if err != nil {
		t.Fatalf("unexpected error : [%s]", err)
	}
	if hbar1["fuga"] != "wwww" {
		t.Fatalf("could not write : value %v", hbar1["hoge"])
	}

	// Write(obj, "foo.bar1", "xxxx") map value
	newhoge := map[string]interface{}{"hoge": "yyyy"}
	err = Write("foo[0].bar1", obj, newhoge)
	if err != nil {
		t.Fatalf("unexpected error : [%s]", err)
	}
	bar1 = foo0.(map[string]interface{})["bar1"]
	if reflect.DeepEqual(bar1, newhoge) == false {
		t.Fatalf("could not write : value %v", bar1)
	}

	// Write(obj, "noooo[1]", 99) replace slice value
	err = Write("noooo[1]", obj, 99)
	if err != nil {
		t.Fatalf("unexpected error : [%s]", err)
	}
	noooo := hash["noooo"].([]interface{})
	if reflect.DeepEqual(99, noooo[1]) == false {
		t.Fatalf("unexpected value : [%v]", noooo[1])
	}

}

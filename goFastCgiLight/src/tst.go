package main

import "fmt"

type Fruit struct {
    Name string
}

func (f Fruit) Rename(name string) {
    f.Name = name
}

type Candy struct {
    Name string
}

func (c *Candy) Rename(name string) {
    c.Name = name
}

type Renamable interface {
    Rename(string)
}

func Rename(v Renamable, name string) {
    v.Rename(name)
    // at this point, we don't know if v is a pointer type or not.
}

func main() {
    c := Candy{Name: "Snickers"}
    f := Fruit{Name: "Apple"}
    fmt.Println(f)
    fmt.Println(c)
    Rename(f, "Zemo Fruit")
    Rename(&c, "Zemo Bar")
    fmt.Println(f)
    fmt.Println(c)
}


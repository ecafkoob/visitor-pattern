package main

import "fmt"

type Visitor interface {
    Visit(VisitorFunc) error
}
type VisitorFunc func() error

type VisitorList []Visitor

func (l VisitorList) Visit(f VisitorFunc) error {
    for i := range l {
        if err := l[i].Visit(func() error {
            fmt.Println("In VisitorList before fn")
            f()
            fmt.Println("In VisitorList after fn")
            return  nil
        }); err != nil {return err}
    }
    return nil
}

type Visitor1 struct {
}

func (l Visitor1) Visit(f VisitorFunc) error {
    fmt.Println("In Visitor1 before fn")
    f()
    fmt.Println("In Visitor1 after fn")
    return  nil
}

type Visitor2 struct {
    visitor Visitor
}

func (v Visitor2) Visit(f VisitorFunc) error {
   v.visitor.Visit(func() error {
       fmt.Println("In Visitor2 before fn")
       f()
       fmt.Println("In Visitor2 after fn")
       return  nil
   })
    return nil
}

type Visitor3 struct {
    visitor Visitor
}

func (v Visitor3) Visit(f VisitorFunc) error {
    v.visitor.Visit(func() error {
        fmt.Println("In Visitor3 before fn")
        f()
        fmt.Println("In Visitor3 after fn")
        return  nil
    })
    return nil
}


func main() {
   var visitor Visitor
   var visitors []Visitor
   visitor = Visitor1{}
   visitors = append(visitors,visitor)
   visitor = Visitor2{VisitorList(visitors)}
   visitor = Visitor3{visitor}
   visitor.Visit(func() error {
      fmt.Println("In VisitorFunc")
       return nil
   })
}
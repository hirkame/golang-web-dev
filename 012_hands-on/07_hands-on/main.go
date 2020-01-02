package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Descrip string
	Price         float64
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := menu{
		meal{
			Meal: "Brakfast",
			Item: []item{
				item{
					Name:    "aa",
					Descrip: "This is aa.",
					Price:   10000,
				},
				item{
					Name:    "bb",
					Descrip: "This is bb.",
					Price:   20000,
				},
			},
		},
		meal{
			Meal: "Lunch",
			Item: []item{
				item{
					Name:    "aa",
					Descrip: "This is aa.",
					Price:   10000,
				},
				item{
					Name:    "bb",
					Descrip: "This is bb.",
					Price:   20000,
				},
			},
		},
	}
	r := restaurants{
		restaurant{
			Name: "aiueo",
			Menu: m,
		},
		restaurant{
			Name: "aiueo2",
			Menu: m,
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}

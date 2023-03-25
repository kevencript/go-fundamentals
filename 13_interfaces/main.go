package main

import "fmt"

// Interfaces define which methods a struct needs to be consider as an implemented interace
// obs: we can use interfaces it as "types"
type Languages interface {
	Greeting(string) string
	Goodbye(string) string
	BePartOfLanguage() 
}

// Struct for Portuguese fields
type Portugues struct {
	nome string
	continente string
	falantes int
}

func (p Portugues) Greeting(name string) string {
	return "Olá "+name+"! Falamos "+p.nome+" e o nosso continente é aqui: "+ p.continente  
}	

func (p Portugues) Goodbye(name string) string {
	return "Tchau "+name+" :)"
}	



func (p *Portugues) BePartOfLanguage() {
	oldSpeakerValue := p.falantes
	p.falantes++
	fmt.Printf("Novo falante de %v! Total antes>depois: %d>%d", p.nome, oldSpeakerValue, p.falantes )
}

// Struct for English fields
type English struct {
	name string
	continent string
	speakers int
}

func (e English) Greeting(name string) string {
	return "Hi "+name+"! We speak "+e.name+" our continent is here: "+ e.continent 
}	

func (e English) Goodbye(name string) string {
	return "Bye "+name+" :)"
}	

func (e *English) BePartOfLanguage() {
	oldSpeakerValue := e.speakers
	e.speakers++
	fmt.Printf("New %v speaker! Total before>afer: %d>%d", e.name, oldSpeakerValue, e.speakers )
}


func main() {
	EnglishLanguage := English{name:"English", continent:"North America", speakers:1823}
	PortugueseLanguage := Portugues{nome:"Português", continente: "Amérida do Sul", falantes: 7362}

	// Here we define that every array list item must be Languages interface type 
	// OBS: If we remove one of the methods described on Interface, Go will complain
	// and say that the struct is not implemented by Languages interface
	languagesList := []Languages {&EnglishLanguage, &PortugueseLanguage}

	// Here we are iterating the array and using the methods from each struct
	for _,language := range languagesList {
		fmt.Printf("\nGreetings: %v\n Goodbye: %v\n", language.Greeting("Lucas"), language.Goodbye("Lucas"))
		language.BePartOfLanguage()
	}
}
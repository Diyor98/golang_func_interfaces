package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, movement, sound string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.movement)
}

func (c Cow) Speak() {
	fmt.Println(c.sound)
}

type Bird struct {
	food, movement, sound string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.movement)
}

func (b Bird) Speak() {
	fmt.Println(b.sound)
}

type Snake struct {
	food, movement, sound string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.movement)
}

func (s Snake) Speak() {
	fmt.Println(s.sound)
}

func AnimalAction(a Animal, action string) {
	switch action {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Incorrect inputer parameters. Please, try again.")
	}

}

func main() {
	animals := make(map[string]Animal)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		words := strings.Fields(text)
		if len(words) < 3 {
			fmt.Println("Incorrect inputer parameters. Please, try again.")
			continue
		}

		name := strings.ToLower(words[1])

		switch strings.ToLower(words[0]) {
		case "newanimal":
			switch strings.ToLower(words[2]) {
			case "cow":
				animals[name] = Cow{"grass", "walk", "moo"}
			case "bird":
				animals[name] = Bird{"worms", "fly", "peep"}
			case "snake":
				animals[name] = Snake{"mice", "slither", "hsss"}
			default:
				fmt.Println("Incorrect inputer parameters. Please, try again.")
				continue
			}
			fmt.Println("Created it!")

		case "query":
			animal, flag := animals[name]
			if flag == true {
				AnimalAction(animal, words[2])
			} else {
				fmt.Println("Animal with such a name does not exist")
			}
		}
	}
}

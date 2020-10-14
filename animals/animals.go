package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, movement, sound string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.movement)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		words := strings.Fields(text)
		if words[0] == "cow" {
			if words[1] == "eat" {
				cow.Eat()
			} else if words[1] == "move" {
				cow.Move()
			} else if words[1] == "speak" {
				cow.Speak()
			}
		} else if words[0] == "bird" {
			if words[1] == "eat" {
				bird.Eat()
			} else if words[1] == "move" {
				bird.Move()
			} else if words[1] == "speak" {
				bird.Speak()
			}
		} else if words[0] == "snake" {
			if words[1] == "eat" {
				snake.Eat()
			} else if words[1] == "move" {
				snake.Move()
			} else if words[1] == "speak" {
				snake.Speak()
			}
		}

	}
}

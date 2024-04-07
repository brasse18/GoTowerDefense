package main

import (
	"fmt"
	"os"
    "os/exec"
    "runtime"
	"bufio"
	"strings"
	"time"
	"GoTowerDefense/objekts"
)

var clear map[string]func() //create a map for storing clear funcs

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

func setup() {
	fmt.Println("Seting up for game loop")
	clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func getInput(reader *bufio.Reader) (string) {
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func getKey(reader *bufio.Reader) rune {
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Kunde inte läsa tangenten:", err)
		return 'q'
	}
	return char
}

func inputLoop(ticker *time.Ticker, player *objekts.Player) {
	reader := bufio.NewReader(os.Stdin)
	for {
		//input := getInput(reader)
		keyPress := getKey(reader)
		switch keyPress {
		case 'q':
			ticker.Stop()
			return
		case 'w':
			player.AMove(-(256/8),0)
		case 's':
			player.AMove((256/8),0)
		case 'd':
			player.AMove(0,(256/8))
		case 'a':
			player.AMove(0,-(256/8))
		}
	
	}
}

func loop(ticker *time.Ticker, gWorld *GameWorld) {
	setup()
	
	for range ticker.C{
		//CallClear()
		for _, entatyPtr := range gWorld.Entatys {
            if entatyPtr != nil {
                //fmt.Println(*entatyPtr) // Använd *entatyPtr för att skriva ut det pekaren pekar på
            }
        }
		//fmt.Println("Tick ")
		//if player.ReDraw {
		//	CallClear()
		//	player.ReDraw = false
		//	fmt.Println(player.Name, ": ", player.PlayerEnataty.Posison)
		//}
	}
}
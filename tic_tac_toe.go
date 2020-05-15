package main

import (
	"bufio"
	"os"
	"strings"
)

type workspace struct {
	ttt_buff [9]string
	cursor_pos int
	turn int
	game_fin bool
	text string
	reader *bufio.Reader
}


func main() {
	work := workspace{
		cursor_pos: 0,
		turn:       0,
		game_fin:	false,
		reader:		nil,
	}

	work.reader = bufio.NewReader(os.Stdin)

	Init( &work )
	Build_Output( work )

	for work.game_fin == false {
		Read_Cmd( &work )
		Proc_cmd( &work )

		print(work.text)

		Build_Output( work )
	}
}


func Init( work *workspace ) {
	// Init
	for ii:= 1; ii < 9; ii++ {
		work.ttt_buff[ii] = " "
	}
}

func Build_Output( work workspace ){
	kk := 0

	for ii:= 0; ii < 3; ii++ {
		for jj:= 0; jj < 3; jj++ {

			// Print cursor
			if work.cursor_pos == kk {
				print("$")
			}

			// Print marker
			print(work.ttt_buff[kk])

			// Print dividing line
			if jj < 2 {
				print(" | ")
			}
			kk++
		}
		if ii < 2 {
			print("\n----------\n")
		} else{
			print("\n\n")
		}
	}
}

func Read_Cmd(work *workspace){

	work.text, _ = work.reader.ReadString('\n')
}




func Proc_cmd( work *workspace) {
	if strings.Compare(work.text, "d\r\n") == 0 && ((work.cursor_pos + 1) % 3) != 0 {
		work.cursor_pos++
	} else if strings.Compare(work.text, "a\r\n") == 0 && (work.cursor_pos != 0 && work.cursor_pos != 3 && work.cursor_pos != 6) {
		work.cursor_pos--
	} else if strings.Compare(work.text, "s\r\n") == 0 && (work.cursor_pos != 6 && work.cursor_pos != 7 && work.cursor_pos != 8) {
		work.cursor_pos = work.cursor_pos + 3;
	} else if strings.Compare(work.text, "w\r\n") == 0 && (work.cursor_pos != 0 && work.cursor_pos != 1 && work.cursor_pos != 2) {
		work.cursor_pos = work.cursor_pos - 3;
	}

	if strings.Compare(work.text, "x\r\n") == 0 && strings.Compare(work.ttt_buff[work.cursor_pos], " ") == 0 {
		work.ttt_buff[work.cursor_pos] = "x"
	}

	if strings.Compare(work.text, "o\r\n") == 0 && strings.Compare(work.ttt_buff[work.cursor_pos], " ") == 0 {
		work.ttt_buff[work.cursor_pos] = "o"
	}
}


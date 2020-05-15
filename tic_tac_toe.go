package main

import(
	"bufio"
	"os"
	"strings"
)
func Init( ttt_buff []string ) int{
	// Init
	for ii:= 1; ii < 9; ii++ {
		ttt_buff[ii] = " "
	}

	return 0;
}

func Build_Output( ttt_buff []string, cursor_pos int ){
	kk := 0

	for ii:= 0; ii < 3; ii++ {
		for jj:= 0; jj < 3; jj++ {

			// Print cursor
			if cursor_pos == kk {
				print("$")
			}

			// Print marker
			print(ttt_buff[kk])

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

func Read_Cmd(reader *bufio.Reader) string {

	text, _ := reader.ReadString('\n')

	return text
}


func main() {

	game_fin := false
	ttt_buff := [9]string{" "}

	reader := bufio.NewReader(os.Stdin)

	cursor_pos := Init( ttt_buff[:] )
	Build_Output( ttt_buff[:], cursor_pos )

	for game_fin == false {
		text := Read_Cmd( reader )
		cursor_pos = Proc_cmd( text, ttt_buff[:], cursor_pos )
		print(text)

		if strings.Contains(text, "x") {

		}
		Build_Output( ttt_buff[:], cursor_pos )
	}
}

func Proc_cmd(text string, ttt_buff []string, cursor_pos int) int {
	if strings.Compare(text, "d\r\n") == 0 && ((cursor_pos + 1) % 3) != 0 {
		cursor_pos++
	} else if strings.Compare(text, "a\r\n") == 0 && (cursor_pos != 0 && cursor_pos != 3 && cursor_pos != 6) {
		cursor_pos--
	} else if strings.Compare(text, "s\r\n") == 0 && (cursor_pos != 6 && cursor_pos != 7 && cursor_pos != 8) {
		cursor_pos = cursor_pos + 3;
	} else if strings.Compare(text, "w\r\n") == 0 && (cursor_pos != 0 && cursor_pos != 1 && cursor_pos != 2) {
		cursor_pos = cursor_pos - 3;
	}
	return( cursor_pos )
}


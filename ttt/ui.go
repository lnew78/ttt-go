package ttt

import(
  "fmt"
  "bufio"
  "os"
)

type UI struct{}

func (ui UI) PrintMsg(message string) {
  fmt.Println(message)
}

func (ui UI) GetInput() string {
  in := bufio.NewReader(os.Stdin);
  input, err := in.ReadString('\n');
  if err != nil {
    fmt.Println(err)
  }
  return input
}

func (ui UI) PrintBoard(tttRules TTTRules) {
  row1, row2, row3 := tttRules.Rows()
  rows := [][]string{row1, row2, row3}
  for index1 := range rows {
    for index2 := range rows {
      fmt.Print(rows[index1][index2])
      fmt.Print(" | ")
      if index2 == len(rows) - 1 {
        fmt.Println("\n")
      }
    }
  }
}

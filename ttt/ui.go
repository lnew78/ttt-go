package ttt

import(
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)

type UI struct{}

func (ui UI) PrintMsg(message string) {
  fmt.Println(message)
}

func (ui UI) PrintMsgWithData(message, data string) {
  fmt.Println(message, data)
}

func (ui UI) GetInput() string {
  in := bufio.NewReader(os.Stdin);
  input, err := in.ReadString('\n');
  if err != nil {
    fmt.Println(err)
  }
  return input
}

func (ui UI) GetNumericInput(message string) int {
  ui.PrintMsg(message)
  input := strings.TrimSuffix(ui.GetInput(), "\n")

  number, err := strconv.Atoi(input)
  if err != nil {
    return ui.GetNumericInput(message)
  }
  return number
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

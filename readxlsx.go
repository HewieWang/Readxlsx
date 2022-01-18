package main

import(
    "fmt"
    "github.com/tealeg/xlsx"
    "os"
    "strconv"
)

var (
    inFile = os.Args[1]
)

func main(){
    hang, _ := strconv.Atoi(os.Args[2])
    lie, _ := strconv.Atoi(os.Args[3])
    xlFile, err := xlsx.OpenFile(inFile)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    for _, sheet := range xlFile.Sheets {
        //遍历行读取
        for l, row := range sheet.Rows {
            // 遍历每行的列读取
            for c, cell := range row.Cells {
                text := cell.String()
                if l==hang-1 && c==lie-1 {
                  fmt.Println(text)
                }

            }
        }
    }
}

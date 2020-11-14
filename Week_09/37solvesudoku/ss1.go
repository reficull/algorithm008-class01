package main

import(
    "fmt"
    "strconv"
)
var rowMap [9]int
var colMap [9]int
var boxMap [3][3]int

func main(){
    sdk := [][]byte{
        {'5','3','.','.','7','.','.','.','.'},  
        {'6','.','.','1','9','5','.','.','.'},
        {'.','9','8','.','.','.','.','6','.'},
        {'8','.','.','.','6','.','.','.','3'},
        {'4','.','.','8','.','3','.','.','1'},
        {'7','.','.','.','2','.','.','.','6'},
        {'.','6','.','.','.','.','2','8','.'},
        {'.','.','.','4','1','9','.','.','5'},
        {'.','.','.','.','8','.','.','7','9'},
    }

    //rowMap = [9]int{0,0,0,0,0,0,0,0,0}
    //colMap := [9]int{0,0,0,0,0,0,0,0,0}
    //boxMap := [3][3]int{{0,0,0},{0,0,0},{0,0,0}}
    solveSudoku(sdk)
    printMatrix(sdk)
}

func solveSudoku(board [][]byte)  {
    backtrace(board, 0, 0)

}

func backtrace(board [][]byte, row int, col int) bool {
    if row == 9 { //遍历完成，找到答案
        return true
    }
    if col == 9 { //换行
        return backtrace(board, row+1, 0)
    }
    if board[row][col] != '.' { //跳过已填有数字的格子
        return backtrace(board, row, col+1)
    }
    for i := '1'; i <= '9'; i++ { //遍历1--9
        if isOK(board, row, col, byte(i)) { //判断行列单元格是否存在当前数字
            board[row][col] = byte(i) //放置数字
            if backtrace(board, row, col+1){ //放置下一个格子
                return true
            }
            board[row][col] = '.' //若当前放置的格子无法使下一个格子放置成功，则清除当前操作
        }
    }
    return false

}

func checkOK(row,col int, c byte) bool{
    bit := 1
    if ((rowMap[row]&bit)!=bit && bit!=(colMap[col]&bit) && bit!=(boxMap[row/3][col/3]&bit)){ //判断行列单元格是否存在当前数字


    }
    return true
}

func isOK(board [][]byte, row int, col int, c byte) bool {
    for i := 0; i < 9; i++ {
        // 判断行是否存在重复
        if board[row][i] == c {
            return false

        }
        // 判断列是否存在重复
        if board[i][col] == c {
            return false

        }
        // 判断单元格是否存在重复
        if board[(row/3)*3+i/3][(col/3)*3+i%3] == c {
            return false

        }

    }
    return true

}

func printMatrix(matrix [][]byte){
    for _,col1 := range matrix{
        for _,c := range col1{
            num,_ := strconv.Atoi(string(c))
            fmt.Printf("%d,",num)

        }
        fmt.Println()
    }
}

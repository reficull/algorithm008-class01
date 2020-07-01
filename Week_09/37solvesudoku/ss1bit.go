package main

import(
    "fmt"
    "strconv"
)
var    rowMap = [9]int{0,0,0,0,0,0,0,0,0}
var    colMap = [9]int{0,0,0,0,0,0,0,0,0}
var    boxMap = [3][3]int{{0,0,0},{0,0,0},{0,0,0}}
func main(){
    /*
    {
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
    */
    
    sdk := [][]byte{
        {'.','.','9','7','4','8','.','.','.'},
        {'7','.','.','.','.','.','.','.','.'},
        {'.','2','.','1','.','9','.','.','.'},
        {'.','.','7','.','.','.','2','4','.'},
        {'.','6','4','.','1','.','5','9','.'},
        {'.','9','8','.','.','.','3','.','.'},
        {'.','.','.','8','.','3','.','2','.'},
        {'.','.','.','.','.','.','.','.','6'},
        {'.','.','.','2','7','5','9','.','.'},
    }
    solveSudoku(sdk)
    printMatrix(sdk)
}

func solveSudoku(board [][]byte)  {
    for i:=0;i<len(board);i++{
        for j:=0;j<len(board[i]);j++{
            if board[i][j] == '.' {continue}
            num,_ := strconv.Atoi(string(board[i][j]))
            bit := 1 << (num -1)
            rowMap[i]|=bit;colMap[j]|=bit;boxMap[i/3][j/3]|=bit;
        }
    }
    backtrace(board,0,0 )

}

func backtrace(board [][]byte, row,col int) bool{
    if row == 9 { //遍历完成，找到答案
       // fmt.Println("return result")
      //  printMatrix(board)
        return true
    }
    if col == 9 { //换行
        return backtrace(board, row+1, 0)
    }
    if board[row][col] != '.' { //跳过已填有数字的格子
        return backtrace(board, row, col+1)
    }
    for i := '1'; i <= '9'; i++ { //遍历1--9
        num,_ := strconv.Atoi(string(i))
        bit := 1 << (num -1)
        if (((rowMap[row]&bit)!=bit) && (bit!=(colMap[col]&bit)) && (bit!=(boxMap[row/3][col/3]&bit))){ //判断行列单元格是否存在当前数字
            board[row][col] = byte(i) //放置数字
            rowMap[row]|=bit;colMap[col]|=bit;boxMap[row/3][col/3]|=bit;
    //        fmt.Printf("r : %d map : %#b,col:%d map:%#b,box:%#b\n",row,rowMap[row],col,colMap[col],boxMap[row/3][col/3])
            //fmt.Printf("i:%d,j:%d,num:%d,set row:%d,box:%d\n",row,col,i,rowMap[row],boxMap[row/3][col/3])
        //    fmt.Printf("fill i:%d,j:%d,num:%d,bit:%d\n\n",row,col,i,bit)
        //    printMatrix(board)
            if backtrace(board, row, col+1){ //放置下一个格子
                return true
            }
         //   fmt.Printf("unfill i:%d,j:%d,num:%d\n\n",row,col,i)
            board[row][col] = '.' //若当前放置的格子无法使下一个格子放置成功，则清除当前操作
            rowMap[row]&=^bit;colMap[col]&=^bit;boxMap[row/3][col/3]&=^bit;
        }
    }
    return false

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

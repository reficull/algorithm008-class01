package main

import(
    "fmt"
    "strconv"
)
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

    solveSudoku(sdk)

}

func solveSudoku(board [][]byte)  {
    //dfs(board,0,0,9)
    col := [9][10]bool{}
    row := [9][10]bool{}
    box := [9][10]bool{}

    for r:=0;r<9;r++{
        for c:=0;c<9;c++{
            if board[r][c] != '.'{
                num,_ := strconv.Atoi(string(board[r][c]))  
                boxI :=  r/3 * 3 + c/3
                col[c][num] ,row[r][num] ,box[boxI][num] = true,true,true
            }
        }
    }
    fill(board,0,row,col,box)
    printMatrix(board)
}

func fill(board [][]byte,n int,row,col,box [9][10]bool) bool{
    if n >= 81 {
        fmt.Println("=================================")
        printMatrix(board)
        return true
    } 
//    fmt.Println("fill num:",n)
    rk := n / 9
    ck := n % 9
    if board[rk][ck] != '.'{
        fill(board,n+1,row,col,box)
    }
    boxK := rk / 3 * 3 + ck / 3
    for i := 1;i <= 9;i++{
        if (!row[rk][i]) && (!col[ck][i]) && (!box[boxK][i]){
            board[rk][ck] = byte( i)
            fmt.Printf("fill i:%d,j:%d,num:%d\n",rk,ck,i)
            row[rk][i],col[ck][i],box[boxK][i]=true,true,true
            if fill(board,n+1,row,col,box){
               return  true
            }
            fmt.Printf("========unfill i:%d,j:%d\n",rk,ck)
            row[rk][i],col[ck][i],box[boxK][i]=false,false,false
        } 
    }
    printLine(board[rk])
    board[rk][ck] = '.'
    return false
}
func printLine(l []byte){
    for _,i:= range(l){
        num,_ := strconv.Atoi(string(i)) 
        fmt.Printf("%d,",num)
    }
    fmt.Println()
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

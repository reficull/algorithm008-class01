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
    col := [9]int{}
    row := [9]int{}
    box := [9]int{}

    for r:=0;r<9;r++{
        for c:=0;c<9;c++{
            if board[r][c] != '.'{
                num := board[r][c] - '1'
                boxI :=  r/3 * 3 + c/3
                bit := 1 << num 
                col[c] |= bit
                row[r] |= bit
                box[boxI] |= bit
            }
        }
    }
    fill(board,0,row,col,box)
}

func fill(board [][]byte,n int,row,col,box [9]int) bool{
    if n >= 81 {
        printMatrix(board)
        return true
    } 
    fmt.Println("fill num:",n)
    rk := n/9
    ck := n%9
    if board[rk][ck] != '.'{
        fill(board,n+1,row,col,box)
    }
    boxK := rk/3*3 + ck/3
    for i:=0;i<9;i++{
        bit := 1 << i
        if (row[rk]&bit==0) && (col[ck]&bit==0) && (box[boxK]&bit==0){
            board[rk][ck] = byte('1' + i)
        //    fmt.Printf("fill i:%d,j:%d,num:%d\n",rk,ck,i)
            row[rk]|=bit;col[ck]|=bit;box[boxK]|=bit
            if fill(board,n+1,row,col,box){
                printMatrix(board)
                return true
            }

            row[rk]&= (^bit);col[ck]&=(^bit);box[boxK]&=(^bit)
        } 
    }
    board[rk][ck] = '.'
    return false
}

func dfs(board [][]byte,si,sj,n int) bool{
    fmt.Printf("i:%d,j:%d\n",si,sj)
    if si >= n{
        return true
    }
    for i:=si;i<n;i++{
        for j:=sj;j<n;j++{
            num,_ := strconv.Atoi(string(board[i][j]))
            if num > 0{ 
                continue
            }else{
                for t := 1;t<=n; t++{
                    board[i][j] = byte('1' + t)
                    if isValidSudoku(board){
                        fmt.Printf("drill down:i:%d,j:%d,%d\n",i,j,t)
                         if  dfs(board,i,j+1,n){
                             return true
                         }else{
                             
                         }
                        
                    }
                }
            }
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
func isValidSudoku(board [][]byte) bool {
    rowMap := []int{0,0,0,0,0,0,0,0,0,0}
    colMap := []int{0,0,0,0,0,0,0,0,0,0}
    block := []int{0,0,0,0,0,0,0,0,0,0}
    for i:=0;i<len(board);i++{
        for j:=0;j<len(board[i]);j++{
            num,_ := strconv.Atoi(string(board[i][j]))
            if num > 0{
                //fmt.Printf("i:%d,j:%d\n",i,j)
                bitI := (1 <<  num)
                if (rowMap[i] &  bitI) ==0{ 
                    rowMap[i] |= bitI
                 //   fmt.Printf("set row bit:%d,%d,biTI:%d\n",i,rowMap[i],bitI)
                }else{ 
                  //  fmt.Printf("rowMap got %d , %d = %d\n",i,j,num)
                    return false
                }
                if colMap[j] & bitI==0{ 
                    colMap[j] |= bitI
                }else{ 
                   // fmt.Printf("colMap got %d , %d = %d\n",i,j,num)
                    return false
                }
                bI := i/3 * 3 + j/3       
                //fmt.Printf("block:%d,i:%d,j:%d\n",bI,i,j)
                if block[bI] & bitI != 0{
                 //   fmt.Printf("dup:%d,%d,num:%d\n",i,j,num)
                    return false
                }else{
                    block[bI] |= bitI 
                }

            }
        }
    }
    return true
}

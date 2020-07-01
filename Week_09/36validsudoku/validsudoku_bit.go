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
/*    
        {'.','.','4','.','.','.','6','3','.'},
        {'.','.','.','.','.','.','.','.','.'},
        {'5','.','.','.','.','.','.','9','.'},
        {'.','.','.','5','6','.','.','.','.'},
        {'4','.','3','.','.','.','.','.','1'},
        {'.','.','.','7','.','.','.','.','.'},
        {'.','.','.','5','.','.','.','.','.'},
        {'.','.','.','.','.','.','.','.','.'},
        {'.','.','.','.','.','.','.','.','.'},
        */
    }
    ret := isValidSudoku(sdk)
    fmt.Println("ret:",ret)
}

func isValidSudoku(board [][]byte) bool {
    rowMap := []int{0,0,0,0,0,0,0,0,0,0}
    colMap := []int{0,0,0,0,0,0,0,0,0,0}
    block := []int{0,0,0,0,0,0,0,0,0,0}
    for i:=0;i<len(board);i++{
        for j:=0;j<len(board[i]);j++{
            num,_ := strconv.Atoi(string(board[i][j]))
            if num > 0{
                fmt.Printf("i:%d,j:%d\n",i,j)
                bitI := (1 <<  num)
                if (rowMap[i] &  bitI) ==0{ 
                    rowMap[i] |= bitI
                    fmt.Printf("set row bit:%d,%d,biTI:%d\n",i,rowMap[i],bitI)
                }else{ 
                    fmt.Printf("rowMap got %d , %d = %d\n",i,j,num)
                    return false
                }
                if colMap[j] & bitI==0{ 
                    colMap[j] |= bitI
                }else{ 
                    fmt.Printf("colMap got %d , %d = %d\n",i,j,num)
                    return false
                }
                bI := i/3 * 3 + j/3       
                fmt.Printf("block:%d,i:%d,j:%d\n",bI,i,j)
                if block[bI] & bitI != 0{
                    fmt.Printf("dup:%d,%d,num:%d\n",i,j,num)
                    return false
                }else{
                    block[bI] |= bitI 
                }

            }
        }
        fmt.Println()
    }
    return true
}

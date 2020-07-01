package main

import(
    "fmt"
)

func main(){
    beginWord := "hit"
    endWord := "cog"
    wordList := []string{"hot","dot","dog","lot","log","cog"}
    ret := ladderLength(beginWord,endWord,wordList)
    fmt.Println("ret:",ret) 
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    memo := make(map[string]bool)
    ret := dfs(beginWord,endWord,0,[]string{beginWord},wordList,memo)

    fmt.Println("ret:",ret)
    return len(ret)-1
}

func dfs(beginWord,endWord string,index int, tmp,wordList []string,memo map[string]bool) []string{
    if _,ok := memo[beginWord];ok{
        return nil
    }
    if beginWord == endWord {
        return tmp
    }

    if index >= len(wordList) {
        return nil
    }

    for i:=index;i<len(wordList);i++{
        if(!OneDiff(beginWord,wordList[i])){
            continue
        }
        t := wordList[i]
        wordList[i],wordList[index] = wordList[index],wordList[i]
        newTmp := make([]string,len(tmp))
        copy(newTmp,tmp)
        newTmp = append(newTmp,t)
        result := dfs(t,endWord,index+1,newTmp,wordList,memo)
        if result != nil{
            return result
        }
        wordList[i],wordList[index] = wordList[index],wordList[i]
        memo[t] = true
    }
    return nil
}

func OneDiff(w1,w2 string) bool{
    diff := 0
    for i:=0;i<len(w1);i++{
        if w1[i]  != w2[i]{
            diff++
        }               
    }
    if diff == 1{
        return true
    }
    return false
}

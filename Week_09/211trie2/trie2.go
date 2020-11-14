
package main

import(
    "fmt"
)
func main(){
    word := "apple"
    obj := Constructor();
    obj.AddWord(word);
    param_2 := obj.Search(word);
    fmt.Println("ret:",param_2)
}

type WordDictionary struct {
   ok [27]bool
   have [27]*WordDictionary

}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
    return WordDictionary{}
}

func Num(c byte) int{
    var ret int
    if c == '.'{
        ret = 26
    }
    if c >= 'a' && c <= 'z'{
        ret = int(c - 'a')
    }
    return ret
}
func (this *WordDictionary) AddWord(word string)  {
    for i :=0;i<len(word);i++{
        n := Num(word[i]) 
        if this.have[n] == nil{
            this.have[n] = new(WordDictionary)
        }
        if i+1 == len(word){
            this.ok[n] = true 
        } 
        this = this.have[n]
    }

}

/** Returns if the word is in the trie. */
func (this *WordDictionary) Search(word string) bool {
    return this.Find(this,word,0)
}

func (this *WordDictionary) Find(that *WordDictionary,word string,i int) bool {
    n := Num(word[i])
    if n == '.'{
        for i:=0;i<27;i++{
            if this.have[i] != nil{
                if i+1 == len(word){
                    return true
                }else if this.Find(that,word,i+1){
                    return true
                }
            }
        }
        return false
    }

    if this.have[Num(word[i])] != nil{
        if i+1 == len(word){
            return this.ok[Num(word[i])] 
        }
        return  this.Find(this.have[Num(word[i])],word,i+1)
        
    }
    return false 
}




/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *WordDictionary) StartsWith(prefix string) bool {
    for _,c := range prefix{
        if this.have[c - 'a'] == nil{
            return false 
        }
        this = this.have[c - 'a'] 
    }
    return true 
}



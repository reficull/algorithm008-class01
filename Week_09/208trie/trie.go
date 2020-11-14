package main

import(
    "fmt"
)
func main(){
    word := "apple"
    obj := Constructor();
    obj.Insert(word);
    param_2 := obj.Search(word);
    param_3 := obj.StartsWith(prefix);
}

type Trie struct {
   isEnd bool
   next [26]*Trie

}


/** Initialize your data structure here. */
func Constructor() Trie {
    t := Trie{}
    return t
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    if this.Search(word){return}
    for _,c := range word{
        if this.next[c - 'a'] == nil{
            this.next[c - 'a'] = new(Trie)
        }
        this = this.next[c - 'a']
    }
    this.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    for _,c := range word{
        if this.next[c - 'a'] == nil{
            return false
        }
       this = this.next[c - 'a']
        
    
    }
    if this.isEnd == false{
        return false
    }

    return true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    for _,c := range prefix{
        if this.next[c - 'a'] == nil{
            return false 
        }
        this = this.next[c - 'a'] 
    }
    return true 
}


/**
 * Your Trie object will be instantiated and called as such:
  * obj := Constructor();
   * obj.Insert(word);
    * param_2 := obj.Search(word);
     * param_3 := obj.StartsWith(prefix);
      */

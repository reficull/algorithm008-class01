#include <iostream>
#include <string.h>

class Solution {
    int hash[10][10];
    int cnt;
    void doAdd(int r,int c, int n, int val){
        if(r<0 || r>=n) return;
        if(c<0 || c>=n) return;
        hash[r][c]+=val;
    }
    void add(int r,int c,int n,int val){
        int i;
        for(i=0;i<n;++i){
            doAdd(r,i,n,val);// 行增加标签
            doAdd(i,c,n,val);// 列增加标签
        }
        for(i=0;i<n;++i){
            doAdd(r+i,c+i,n,val);
            doAdd(r+i,c-i,n,val);
            doAdd(r-i,c+i,n,val);
            doAdd(r-i,c-i,n,val);
            
        }
    }
    void dfs(int depth,int maxDepth){
        if(depth == maxDepth){
            cnt++;
            return;
        }
        //i 是列数
        for(int i = 0;i<maxDepth;++i){
            if(hash[depth][i] == 0){
                add(depth,i,maxDepth,1);
                dfs(depth + 1,maxDepth);
                add(depth,i,maxDepth,-1);
            }
        }
    }
public:
    int totalNQueens(int n) {
        cnt = 0;
        memset(hash,0,sizeof(hash));
        dfs(0,n);
        return cnt;
    }
};

int main(){
	Solution *s = new Solution();
    int cnt = s->totalNQueens(4);
    std::cout << "4:" << cnt << std::endl;
}

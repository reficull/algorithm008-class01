#include <iostream>
#include <vector>
#include <string.h>
using namespace std;

const size_t N = 10;

class Solution {
    
    vector<vector<string>> matrix;
    int hash[N][N];
    int cnt;
    void doAdd(int r,int c, int n, int val){
        if(r<0 || r>=n) return;
        if(c<0 || c>=n) return;
        
        //if(hash[r][c]<1)
        hash[r][c]+=val;
    }
    void add(int r,int c,int n,int val){
        int i;
        for(i=0;i<n;++i){
            doAdd(r,i,n,val);// 行增加标签
            doAdd(i,c,n,val);// 列增加标签
        }
        for(i=0;i<n;++i){
            doAdd(r+i,c+i,n,val); // diaganal "\"
            doAdd(r+i,c-i,n,val); // diaganal "/"
            doAdd(r-i,c+i,n,val); // diaganal "/"
            doAdd(r-i,c-i,n,val); // diaganal "\"
            
        }
    }
    void dfs(int depth,int maxDepth){
        if(depth == maxDepth){
            printHash(maxDepth);
            //makeRet(maxDepth);
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
/*
    void makeRet(int n){
        if(matrix.size() == 0)
            matrix = new vector<vector<string>>();
        string s = new String();
        vector<string> board = new vector<string>();
        for(int c=0;c<n;++c){
            for(int r=0;r<n;++r){
                if(hash[r][c]>=6)
                    s.append("Q");
                else
                    s.append(".");
            }
            board.append(s);
        }
        matrix.append(board);
    }
*/
    void printHash(int n){
        for(int c=0;c<n;++c){
            for(int r=0;r<n;++r){
                if(hash[r][c]>=6)
                    cout << "Q";
                else
                    cout << ".";
                //cout << hash[r][c] ;
            }
            cout << endl;
        }
        cout << "================" << endl;
        cout << endl;
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
    int n = 1;
    cin >> n;
    int cnt = s->totalNQueens(n);
    cout << "N:" << n << ", count:" << cnt << endl;
}

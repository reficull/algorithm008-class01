package unionSearch

type struct CS{
    P []int
}
func (CS *) Init(p []int){
    l:=len(p)
    CS.P = make([]int,l)
    for i:=0;i<l;i++{
        CS.P[i]= p[i]
    }
}

func (CS *) union(p []int,i,j){
    p1 := CS.Parent(i)
    p2 := CS.Parent(j)
    p[p1] = p2
}

func (CS *) Parent(p []int,i int){
    root := i
    for  CS.P[root] != root{
        root = CS.P[root]
    }
    //路径压缩
    for CS.P[i] != i{
        x:=i
        i = CS.P[i]
        CS.P[x] = root
    }
    return root
}

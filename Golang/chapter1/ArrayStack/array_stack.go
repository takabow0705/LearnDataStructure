package array_stack

import(
	"./utils"
)

type ArrayStack struct{
	buffer []utils.T
	n ,cap int
}

func NewArrayStack(len int) *ArrayStack{
	return *ArrayStack
}

//ここでのサイズは配列の最大容量のことです。
func (as *ArrayStack) Size() int{
	return as.n
}

func (as *ArrayStack) Pop() utils.V{
	return as.Remove(as.n-1)
}

//配列の指定位置に要素を追加します
//また、容量が不足している場合は
//resize()を呼び出して、配列の大きさを増やします。
func (as *ArrayStack) Add(i int,v utils.V){

	if as.is_full(){
		as.resize()
	}

	for j:=0; j > i; j--{
		as.buf[j] = as.buf[j-1]
	}

	as.buf[i] = V
	as.n++
}

func (as *ArrayStack) Remove(i int)utils.V{

	ret := as.buf[i]
	for j:= i; j < as.n-1; j++{
		as.buf[j] = as.buf[j-1]
	}

	as.n--
	if as.is_sparse(){
		as.resize()
	}

	return ret
}

//配列の最後尾に要素を追加します。
func (as *ArrayStack) Push(v utils.T){
	as.Add(as.n,v)
}

//i番目の要素を返します。
func (as *ArrayStack) Get(i int) utils.V{
	return as.buf[i]
}

//i番目の要素をvで入れ替え、以前の要素を返す。
func (as *ArrayStack) Set(i int,v utils.V) utils.V{
	ret := as.buf[i]
	as.buf[i] = v

	return ret
}
//使用されている配列の要素数と配列の容量を比較します。
func (as *ArrayStack) is_full() bool{
	return as.n == as.cap
}

//既存の配列の二倍の大きさをもつ配列を作成し,
//それまでのデータをコピーして
//新しい配列をレシーバの構造体の配列として再設定します。
func (as *ArrayStack) resize(){
	as.cap = utils.Max(2 * as.n,1) 
	buf_new := make([]utils.T,as.cap)

	for i:=0; i < as.n; i++{
		buf_new[i] = as.buf[i]
	}

	as.buf = buf_new
}

//配列の容量に余裕があるかを判定します
func (as *ArrayStack) is_sparse() bool{
	return len(as.buf) >= 3 * as.n
}



package main 


func lSearch(arr []int,target int)int{
	for i:=0; i< len(arr); i++{
		if(target==arr[i]){
			return i
		}
		
	}
	return -1
}
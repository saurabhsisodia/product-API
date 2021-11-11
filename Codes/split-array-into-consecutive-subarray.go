func isPossible(nums []int) bool {
	// remaining elements frequency
	mp := make(map[int]int)

	// number of sequence which can store current element as next element of their sequence
	next := make(map[int]int)

	for _,v:=range nums{
		mp[v]++
	}
	for i:=0;i<len(nums);i++{
		curr := nums[i]
		if mp[curr]==0{   // if current element has already used
			continue
		}
		if next[curr]>0{// if current element can be appended to existing sequence
			next[curr]--// old sequence decreases by one,bcoz cuurent element will be appended to that
			mp[curr]--// current element just used in previous sequence
			next[curr+1]++//now, this sequence take values from current+1
		}else if mp[curr+1]>0 && mp[curr+2]>0{// start new sequence with curr element, but next 2 consecutive 
													// should exists to make length >=3
			mp[curr]--
			mp[curr+1]--
			mp[curr+2]--
			next[curr+3]++// this sequence will take values from curr+3
		}else{
			return false
		}
	}
	return true
}
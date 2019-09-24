package main

import "fmt"

//func nextGreaterElement(nums1 []int, nums2 []int) []int {
//	//利用一个栈来查找父集合中每个元素的Next Greater Element，
//	// 找到了就存放到HashMap中，最后遍历子集合，
//	// 如果HashMap中没有说明不存在，赋值-1。
//	/*
//	 //辅助栈，存放待查找结果的元素，查找到的立即出栈
//        Stack<Integer> stack = new Stack<>();
//        //key存放元素，value存放找到的第一个大于它的值
//        Map<Integer, Integer> map = new HashMap<>();
//        //当栈顶元素大于当前元素时，入栈；当栈顶元素小于当前元素时，说明栈顶元素找到了第一个大于的值，出栈，然后继续出栈直到栈顶元素大于当前元素，将当前元素入栈。
//        for (int i = 0; i < nums.length; i++) {
//            while(!stack.isEmpty() && stack.peek() < nums[i]){
//                map.put(stack.pop(), nums[i]);
//            }
//            stack.push(nums[i]);
//        }
//        //ans数组存放结果
//        int[] ans = new int[findNums.length];
//　　　　 //遍历findNums，在map中查找结果，不存在说明没有大于它的第一个元素，赋值为-1
//        for (int i = 0; i < findNums.length; i++) {
//            ans[i] = map.getOrDefault(findNums[i], -1);
//        }
//        return ans;
//	 */
//	//辅助栈，存放待查找结果的元素，查找到的立即出栈
//	stack := NewStack()
//	//key存放元素，value存放找到的第一个大于它的值
//	tempMap := make(map[int]int)
//	//当栈顶元素大于当前元素时，入栈；当栈顶元素小于当前元素时，说明栈顶元素找到了第一个大于的值，出栈，然后继续出栈直到栈顶元素大于当前元素，将当前元素入栈。
//	for i := 0; i < len(nums2); i++ {
//		while(!stack.isEmpty() && stack.peek() < nums[i]){
//			tempMap[stack.Pop()] = nums2[i]
//		}
//		stack.Push(nums2[i])
//	}
//
//
//
//
//
//}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0, len(nums2))
	findMap := make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums2[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			findMap[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	rst := make([]int, len(nums1))
	for k, v := range nums1 {
		if val, ok := findMap[v]; ok {
			rst[k] = val
		} else {
			rst[k] = -1
		}
	}
	return rst
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

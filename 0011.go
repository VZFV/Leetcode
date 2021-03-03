package main

func maxArea(height []int) int {
	start, end, max := 0, len(height)-1, 0
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		temp := high * width
		if temp > max{
			max = temp
		}
	}
	return max
}
func main() {

}

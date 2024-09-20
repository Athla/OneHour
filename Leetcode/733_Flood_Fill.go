package leetcode

// How tf do that shi work
// check if current spot is different from the desiresd
// Call the dfs (aka paint) in the point. Pass the Image, Coords (x, y), src color (to avoid repaint), desired color
// Check if {
// x > 0; x < img_size; same with y; and if the current point is == src color
//	if this condition is met, repaint and call the paint in the correct adjacent parts

//			X +- 1
//			Y +- 1
//	}
// By doing that, it correctly fills the img.

// Time Complexy O(m * n), where m and n represent the amount of rows and columns
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	src := image[sr][sc]
	if src != color {
		dfs(image, sr, sc, src, color)
	}

	return image

}

func dfs(i [][]int, x, y, src, new int) {
	if x < 0 || x >= len(i) || y < 0 || y >= len(i[0]) || i[x][y] != src {
		return
	}

	i[x][y] = new
	dfs(i, x-1, y, src, new)
	dfs(i, x+1, y, src, new)
	dfs(i, x, y+1, src, new)
	dfs(i, x, y-1, src, new)
}

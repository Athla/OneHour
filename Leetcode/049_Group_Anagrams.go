package leetcode

type Word [26]byte

func groupAnagrams(strs []string) [][]string {
	strWord := func(s string) (w Word) {
		for i := range s {
			w[s[i]-'a']++
		}
		return w
	}
	group := make(map[Word][]string)

	for _, v := range strs {
		w := strWord(v)
		group[w] = append(group[w], v)
	}

	res := make([][]string, 0, len(group))
	for _, v := range group {
		res = append(res, v)
	}

	return res
}

package gopractice

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i, pI := range points {
		distCache := map[int]int{}
		for j, pJ := range points {
			if i == j {
				continue
			}
			distIJ := (pI[0]-pJ[0])*(pI[0]-pJ[0]) + (pI[1]-pJ[1])*(pI[1]-pJ[1])
			if v, ok := distCache[distIJ]; ok {
				distCache[distIJ] = v + 1
			} else {
				distCache[distIJ] = 1
			}
		}
		for _, v := range distCache {
			// 假设已有4个point与i距离一样，当发现第五个point与i距离一样时，这个point会增加4×2个数量的boomerang
			//（与i和其余四个point），所以，总和便是1×2+2×2+3×2+4×2=5×4。
			res += v * (v - 1)
		}
	}
	return res
}

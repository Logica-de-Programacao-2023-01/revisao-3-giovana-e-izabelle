package q3

//Você está subindo uma escada. Leva n passos para chegar ao topo.
//
//A cada vez, você pode subir 1 ou 2 degraus. De quantas maneiras distintas você pode subir até o topo?

func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	do := make([]int, n+1)
	do[1] = 1
	do[2] = 2

	for i := 3; i <= n; i++ {
		do[i] = do[i-1] + do[i-2]
	}
	return do[n]

}

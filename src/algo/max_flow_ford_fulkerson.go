package algo

import "math"

/*
	https://github.com/indy256/codelibrary/blob/master/java/graphs/flows/MaxFlowFordFulkerson.java
*/
func (cm *CapacityMatrix) MaxFlowFordFulkerson(s, t uint32) float64 {
	flow := float64(0)
	for {
		df := cm.fordFulkersonFindPath(make([]bool, cm.Size()), s, t, math.MaxFloat64)
		if df == 0 {
			return flow
		}
		flow += df
	}
}

func (cm *CapacityMatrix) fordFulkersonFindPath(vis []bool, u, t uint32, f float64) float64 {
	if u == t {
		return f
	}

	capM := *cm

	vis[u] = true
	for v, vi := range vis {
		if !vi && capM[u][v] > 0 {
			df := cm.fordFulkersonFindPath(vis, uint32(v), t, minFF(f, capM[u][v]))
			if df > 0 {
				capM[u][v] -= df
				capM[v][u] += df
				return df
			}
		}
	}

	return 0
}

func minFF(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

package lottery

import (
	"fmt"
	"hash/fnv"
	"sort"
)

// CalculateResult 根据双色球开奖号码计算结果值
func CalculateResult(redBalls []int) int {
	hashValue := calculateHashValue(redBalls)
	return int(hashValue)
}

// SelectClosestValues 根据结果值选择离其最近的指定个数的数轴上的值
func SelectClosestValues(result int, values []int, numValues int) []int {
	// 对值列表按照与结果值的绝对值距离进行排序
	sort.Slice(values, func(i, j int) bool {
		return abs(values[i]-result) < abs(values[j]-result)
	})

	// 选择指定个数的值
	selectedValues := make([]int, 0, numValues)
	for _, value := range values {
		selectedValues = append(selectedValues, value)
		if len(selectedValues) == numValues {
			break
		}
	}

	return selectedValues
}

// calculateHashValue 计算哈希值
func calculateHashValue(redBalls []int) uint32 {
	hash := fnv.New32a()
	for _, ball := range redBalls {
		hash.Write([]byte(fmt.Sprintf("%d", ball)))
	}
	return hash.Sum32()
}

// abs 返回一个整数的绝对值
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

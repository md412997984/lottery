package lottery

import (
	"fmt"
	"testing"
)

func TestSelectIDs(t *testing.T) {
	redBalls := []int{1, 9, 13, 14, 25, 6} //双色球
	// 计算得到的结果值
	userPool := []int{1000, 1200, 1300, 1400, 1600, 1800, 20000, 99999} // 用户池中的ID列表
	numIDs := 3                                                         // 需要选择的ID的个数
	minID := userPool[0]                                                // 用户池中的最小ID
	maxID := userPool[len(userPool)-1]                                  // 用户池中的最大ID
	// 将哈希值映射到最小ID和最大ID之间的值
	result := minID + int(CalculateResult(redBalls)%(int(maxID-minID+1)))
	fmt.Println(result)

	selectedIDs := SelectClosestValues(result, userPool, numIDs) 

	if len(selectedIDs) != numIDs {
		t.Errorf("Expected %d IDs, got %d", numIDs, len(selectedIDs))
	}

	for _, id := range selectedIDs {
		t.Errorf("got: %d", id)
	}
}

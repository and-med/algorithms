package main

import (
	"fmt"
	"math"
)

type MaxHeap struct {
	pos      int
	heap     []int
	elements map[int]map[int]struct{}
}

func NewMaxHeap(n int) *MaxHeap {
	heap := make([]int, n)
	elements := make(map[int]map[int]struct{}, n)
	return &MaxHeap{0, heap, elements}
}

func (h *MaxHeap) up(i int) {
	if i == 0 {
		return
	}

	parent := h.parent(i)

	if h.heap[i] > h.heap[parent] {
		h.swap(i, parent)
		h.up(parent)
	}
}

func (h *MaxHeap) down(i int) {
	left, right := h.left(i), h.right(i)

	if left < h.pos && right < h.pos {
		leftBiggerRight := h.heap[left] > h.heap[right]
		rightBiggerLeft := h.heap[right] > h.heap[left]
		leftBiggerMiddle := h.heap[left] > h.heap[i]
		middleBiggerLeft := h.heap[i] > h.heap[left]
		rightBiggerMiddle := h.heap[right] > h.heap[i]
		middleBiggerRight := h.heap[i] > h.heap[right]

		if leftBiggerMiddle && leftBiggerRight {
			// left biggest
			h.swap(i, left)
			h.down(left)
		} else if rightBiggerMiddle && rightBiggerLeft {
			// right biggest
			h.swap(i, right)
			h.down(right)
		} else if !leftBiggerRight && !rightBiggerLeft && leftBiggerMiddle {
			// right == left && left > middle, left goes up
			h.swap(i, left)
			h.down(left)
		} else if !leftBiggerMiddle && !middleBiggerLeft && rightBiggerMiddle {
			// left == middle and right > middle, right goes up
			h.swap(i, right)
			h.down(right)
		} else if !rightBiggerMiddle && middleBiggerRight && leftBiggerMiddle {
			// right == middle and left > middle, left goes up
			h.swap(i, left)
			h.down(left)
		}
	} else if left < h.pos && h.heap[left] > h.heap[i] {
		// left bigger middle
		h.swap(i, left)
		h.down(left)
	} else if right < h.pos && h.heap[right] > h.heap[i] {
		// right bigger middle
		h.swap(i, right)
		h.down(right)
	}
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) left(i int) int {
	return i*2 + 1
}

func (h *MaxHeap) right(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) swap(i, j int) {
	left, right := h.heap[i], h.heap[j]

	delete(h.elements[left], i)
	delete(h.elements[right], j)
	h.elements[left][j], h.elements[right][i] = struct{}{}, struct{}{}

	h.heap[i], h.heap[j] = right, left
}

func (h *MaxHeap) Insert(el int) {
	if h.pos >= len(h.heap) {
		panic("no space in the heap")
	}

	h.heap[h.pos] = el
	if _, ok := h.elements[el]; !ok {
		h.elements[el] = make(map[int]struct{})
	}
	h.elements[el][h.pos] = struct{}{}

	h.up(h.pos)
	h.pos++
}

func (h *MaxHeap) Remove(el int) {
	if h.pos == 0 {
		panic("heap is empty")
	}

	idx := h.anyIdx(el)
	h.pos--

	if idx != h.pos {
		h.swap(idx, h.pos)
		h.up(idx)
		h.down(idx)
	}

	delete(h.elements[el], h.pos)
}

func (h *MaxHeap) anyIdx(el int) int {
	for k := range h.elements[el] {
		return k
	}

	panic(fmt.Sprintf("element %d is not present in the heap", el))
}

func (h *MaxHeap) Max() int {
	if h.pos == 0 {
		panic("heap is empty")
	}

	return h.heap[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	resp := []int{}
	heap := NewMaxHeap(k)

	for i := 0; i < k; i++ {
		heap.Insert(nums[i])
	}

	resp = append(resp, heap.Max())

	for i := k; i < len(nums); i++ {
		heap.Remove(nums[i-k])
		heap.Insert(nums[i])

		resp = append(resp, heap.Max())
	}

	return resp
}

func actualMax(nums []int, start int, k int) int {
	m := math.MinInt

	for i := start; i < min(start+k, len(nums)); i++ {
		m = max(m, nums[i])
	}

	return m
}

func main() {
	nums := []int{-1593, 3371, 6405, 5234, 6697, 3518, -8220, 7740, 5823, 1001, 404, -785, -5954, -8006, -9402, -7573, -2322, -7657, -8302, -8287, -5663, 4517, -5903, -3727, -5396, -162, 2529, 6097, 5013, 2699, 5077, 1576, -9073, 5535, 9331, 7453, -4334, -4686, -1317, -6065, 754, -5163, 3015, 8178, -4164, -67, -4889, -7023, -4539, 3828, -5804, -467, 9724, -3626, 6487, -2846, 6523, 9651, -8091, 8957, -1942, -8926, 7131, -4201, 105, -2979, 6955, -570, -4604, 1793, 5922, -8046, -7356, -3063, 9640, 1912, 6811, -7467, 9953, 9123, 3170, -1552, 3241, 1796, 3243, -4055, -1210, -5562, -6402, -277, 9614, -8966, -8499, 3595, -8070, -8297, 3081, 3159, -2901, -4642, -7989, -5279, 7631, -4407, 4476, -757, 6029, 197, 4179, -7365, -4298, -102, -3548, 3177, -6364, 1231, -3934, 3093, -7392, 7560, 7831, -1788, 1880, -1616, 7022, -5740, 7549, -2758, 8934, -2350, 5504, -3679, -9999, 4918, 7628, 5476, 7832, 1269, 8210, 6857, 4086, -3696, -8767, -7476, -642, 6431, -6828, -6346, -7793, -4153, -7970, -3765, 2527, -4405, 2412, 8924, -1296, 1285, 2439, -9358, -6053, -5889, -7343, -7631, 633, -8189, -3366, -7278, 8130, -8639, -6427, -5156, 3559, -9393, -1194, -9088, -9235, 6072, -8774, 2976, -4903, 3086, -6751, 2655, -5180, 6093, -3129, 2444, 9795, 7689, -6890, -90, 9064, 4435, 1174, -8797, 5480, -5636, 2204, 4995, -6964, -7668, -2393, 9809, 2506, 9131, 6347, 1566, -9269, -9052, 6368, -9852, 1213, -1483, 5672, 3448, -574, -8406, -2435, 6733, -3211, -9129, 9782, 9164, 6141, 7800, 3729, -6609, 9474, 1517, 5754, 8620, 6566, 7879, 2083, -9404, -5046, 3768, 6161, 6248, 9268, -5874, 6758, -7470, 8611, -5722, -4391, -9667, -1823, 8943, -6621, -4438, 9694, 8007, -5749, -4161, 7838, -796, -7379, 8487, -4382, 3660, -5602, 2443, -384, 3676, -1785, -3098, 8275, 4924, -148, 8497, 364, -399, 2570, 2767, -1025, -5420, -2062, -5732, 3081, 3676, 6792, -9713, 9670, 3416, 1542, 8571, 3798, 270, -3318, 4746, 6320, -5553, -504, 9195, -15, -5104, -6992, -8579, -5062, 7442, 1895, 1355, 3866, 3706, 1038, -8064, -6764, 9495, 9069, -4636, 5557, 4749, 6992, -2723, -1327, -7825, -7192, -9884, 5998, -3497, 9484, 2617, -2431, -2895, -8214, -2906, 2643, -3852, 8188, 6260, 6850, -3406, 5791, 6027, 4701, 8864, -3845, -5252, -2062, -6362, -8018, 6158, -8812, 5472, -549, 833, 7118, 4612, 2513, -7864, -6850, 7707, 6186, 7287, -8494, -163, -5243, 7876, -5293, -4170, -805, 96, 1452, -8907, 6941, -5109, -4416, 7947, -8871, 5695, -1456, 9721, -2536, 8685, -9987, -416, 558, 6659, 3221, -8800, -5955, -5520, -8484, 5032, -8816, -9327, -7234, -2772, 3597, 4832, 1980, -7449, -5818, 4230, -3736, 4357, 5838, -6529, -1765, 5277, 8698, 7195, -3784, -1265, 3219, 7049, -9323, -6916, 8355, -4419, 7074, 6648, 7242, 6864, 8303, -7108, 8948, 9072, 8890, 6778, -1936, -7149, 1986, -6030, 7729, 2156, -4990, -3562, -7109, 1686, 381, -6502, -857, 284, -3099, -6614, -1935, -6209, -4111, -1882, 2380, 7460, 5121, -4799, -1274, -2332, 7166, 3951, -2691, -2471, -3743, 1163, -2241, 3197, 3037, -5990, 3532, -1986, 8411, -7337, 2040, -4613, 1526, 5146, -9292, -8753, -7271, 5228, -3405, -1842, -4373, -9095, -832, -4516, 7156, 9964, 9393, -3142, -4997, 4329, 7847, 8190, -2081, -3759, -732, 6218, 1654, -8730, 9091, 1553, -52, 6551, -7790, 6728, -1154, 2703, -2924, -7446, 5611, -3046, 472, -1904, 7191, 2585, -12, -2131, 854, -1464, 7175, 7627, 8227, -5903, -7308, -9053, 2132, 2895, 3096, -6293, -1292, 3899, -3624, -1276, -8243, -1968, 1954, -7292, 4035, 9677, -3763, 2826, 8368, -5775, -4520, 9410, 7334, -1686, 1450, 6707, 3674, 5099, 5181, -1135, 1815, -1081, -2434, -9297, -5397, -27, 9859, 6004, -5856, 6158, 4272, 6873, 5821, 3863, -6287, 6803, 2943, -2278, 1025, 8834, -4957, -5226, 2343, -2029, -8973, 1086, 451, -1538, -235, -6924, -3165, 3981, 5314, 3613, 1093, 6799, 8250, -7243, -5078, 1842, -8341, 4406, -6562, -3511, 9095, -6835, 8318, 2762, 6496, -8991, -1082, 6203, -4170, 8297, 8568, 4433, -9642}

	fmt.Println(maxSlidingWindow(nums, 113))
}

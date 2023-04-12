package heap

// 堆
// 大根堆、小根堆
// 父节点：(i-1)/2 , 子节点（左）：i*2+1 , 子节点（右）：i*2+2

// MaxHeap 定义堆
type MaxHeap struct{
	heap []int
	// heapSize heap 的元素个数
	heapSize int
}

// NewHeap 初始化
func NewHeap()*MaxHeap{
	return &MaxHeap{
		heap:     make([]int, 0),
		heapSize: 0,
	}
}

// IsEmpty 堆为空
func (h *MaxHeap) IsEmpty() bool{
	return h.heapSize ==0
}

// Push 插入元素
func (h *MaxHeap) Push(v int) {
	
	if h.heapSize == len(h.heap) { // 堆满了
		h.heap = append(h.heap, v)
	}else{ // 堆没满
		// heapSize < len(h.heap)
		h.heap[h.heapSize] = v // 直接赋值
	}
	h.HeapInsert(h.heapSize)
	h.heapSize++
	
}

// Pop 弹出堆顶元素
func (h *MaxHeap) Pop() int{
	ans := h.heap[0]
	h.heapSize-- // 精髓
	if h.heapSize >0 {
		h.Swap(0,h.heapSize) // 此时heapSize已经减1是最后一个元素
		// 从0位置往下做调整
		h.Heapify(0,h.heapSize)
	}	
	return ans
}

// HeapInsert从i位置往上调整，如果大于父节点就交换
func (h *MaxHeap)HeapInsert(i int) {
	for h.heap[i] > h.heap[(i-1)/2]{ // 父节点小于子节点
		h.Swap(i,(i-1)/2) // 交换
		i=(i-1)/2 // i指向父节点
	}
}

func (h *MaxHeap) Swap(a,b int)(int,int){
	h.heap[a],h.heap[b]=h.heap[b],h.heap[a]
	return h.heap[a],h.heap[b]
}

// 从i位置往下调整
func (h *MaxHeap) Heapify(i,heapSize int){
	left := i*2 +1 // 左孩子节点
	for left < h.heapSize { // 存在左孩子，左右中先定左孩子更大
		// largest 设两个孩子的更大值的下标为左孩子	
		largest := left 
		if left+1 < heapSize && h.heap[left+1]>h.heap[left]{ // 存在右孩子且大于左孩子
			largest = left +1 
		}
		// 此时largest 确实为两个孩子更大值的下标
		if h.heap[largest] <= h.heap[i]{  // 孩子节点小于父节点
				largest = i
				// liargest为 父与子节点之间更大值的下标
		}
		if largest == i { // i位置已经最大
			break
		}
		/// 父亲往下走
		h.Swap(largest,i) //交换Max(子节点) 和 父节点
		i = largest // 把largest（原来Max(子节点）的下标） 赋值给i（从i位置开始往下调整）
		left = i*2 +1
	}
}
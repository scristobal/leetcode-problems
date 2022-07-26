/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *       
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */

func printLinkedListInReverse(head ImmutableListNode) {

    // compute the size of the list
    size := 0
    for node := head; node != nil; node = node.getNext() {
        size++
    }
    

    // get equidistant nodes
    step := int(math.Sqrt(float64(size))) + 1
    
    heads := make([]ImmutableListNode, step)
    
    k := 0
    for j, node := 0, head; node != nil; j, node = j+1, node.getNext() {
        if j % step == 0 { 
            heads[k] = node
            k++
        }
    }
    
    // print last chunck, might be smaller
    if k>0 {
        sizeLast := 0
        for node := heads[k-1]; node != nil; node = node.getNext() {
            sizeLast++
        }
        printNReverse(heads[k-1], sizeLast) 
    }
    
    // print the rest in chunks
    for i:=k-2; i>=0; i-- {
        printNReverse(heads[i], step)
    }
    
}

func printNReverse(head ImmutableListNode, n int) {
    for i:=0; i<n ; i++ {
        node := head
        for j:=0; j<(n-i-1) && node.getNext() != nil ; j++ {
            node = node.getNext()
        }
        node.printValue()
    }
}
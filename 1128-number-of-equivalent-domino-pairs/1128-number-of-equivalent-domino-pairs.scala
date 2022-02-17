import scala.annotation.tailrec

object Solution {

  def dominoEquivalent(domino: Array[Int], otherDomino: Array[Int]): Boolean =
    (domino(0) == otherDomino(0) && domino(1) == otherDomino(1)) ||
      (domino(0) == otherDomino(1) && domino(1) == otherDomino(0))

  @tailrec
  def numEquivDominoPairsTail(dominoes: Array[Array[Int]], acc: Int): Int = {

    if (dominoes.length == 1) return acc

    val cur = dominoes.tail.count(
      dominoEquivalent(_, dominoes.head)
    ) + acc
      
    numEquivDominoPairsTail(dominoes.tail, cur)

  }
    
  def numEquivDominoPairs(dominoes: Array[Array[Int]]): Int = {
    numEquivDominoPairsTail(dominoes, 0)
  }
    
}
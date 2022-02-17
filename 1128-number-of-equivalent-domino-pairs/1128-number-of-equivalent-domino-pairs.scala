object Solution {

  def dominoEquivalent(domino: Array[Int], otherDomino: Array[Int]): Boolean =
    (domino(0) == otherDomino(0) && domino(1) == otherDomino(1)) ||
      (domino(0) == otherDomino(1) && domino(1) == otherDomino(0))

  def howManyTimesIn(domino: Array[Int], dominoes: Array[Array[Int]]): Int =
    dominoes.count(dominoEquivalent(_, domino))

  def numEquivDominoPairs(dominoes: Array[Array[Int]]): Int = {

    if (dominoes.length == 1) return 0

    howManyTimesIn(dominoes.head, dominoes.tail) + numEquivDominoPairs(
      dominoes.tail
    )

  }
    
}
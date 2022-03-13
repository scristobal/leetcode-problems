
object Solution {
 def hash(d: Array[Int]): Int = {
    if (d(0) > d(1)) {
      return (d(0) * 10) + d(1)
    }
    (d(1) * 10) + d(0)
  }

  def numEquivDominoPairs(dominoes: Array[Array[Int]]): Int = {
    val cache = collection.mutable.Map.empty[Int, Int]

    dominoes.foreach(d => {
      val h = hash(d)
    
      cache.update(h, cache.getOrElse(h, 0) + 1)

    })

    var res = 0
      
    cache.foreach{case (_, v) => {
      if (v > 1) {
        res = res + v * (v - 1) / 2
      }
    }}

    res

  }

    
}



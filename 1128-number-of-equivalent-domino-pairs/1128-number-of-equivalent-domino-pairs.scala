
object Solution {
 def hash(d: Array[Int]): Int = {
    if (d(0) > d(1)) {
      return (d(0) * 10) + d(1)
    }
    (d(1) * 10) + d(0)
  }

  def numEquivDominoPairs(dominoes: Array[Array[Int]]): Int = {
    val cache = collection.mutable.Map.empty[Int, Int].withDefaultValue(0)

    dominoes.foreach(d => {
      val h = hash(d)
    
      cache.update(h, cache(h) + 1)

    })
      
    cache.values.map(c =>  c * (c - 1) / 2).sum

  }

    
}



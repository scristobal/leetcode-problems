<h2><a href="https://leetcode.com/problems/custom-interval/">2805. Custom Interval</a></h2><h3>Medium</h3><hr><div><p><strong>Function&nbsp;</strong><code>customInterval</code></p>

<p>Given a function <code>fn</code>, a number <code>delay</code> and a number <code>period</code>, return&nbsp;a number&nbsp;<code>id</code>. <code>customInterval</code>&nbsp;is a function that should execute the provided function <code>fn</code> at intervals based on a linear pattern defined by the formula <code>delay&nbsp;+ period&nbsp;* count</code>.&nbsp;The <code>count</code> in the formula&nbsp;represents the number of times the interval has been&nbsp;executed starting from an initial value of 0.</p>

<p><strong>Function </strong><code>customClearInterval</code>&nbsp;</p>

<p>Given the&nbsp;<code>id</code>. <code>id</code>&nbsp;is the&nbsp;returned value from&nbsp;the function&nbsp;<code>customInterval</code>. <code>customClearInterval</code>&nbsp;should stop executing&nbsp;provided function <code>fn</code> at intervals.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre><strong>Input:</strong> delay = 50, period = 20, stopTime = 225
<strong>Output:</strong> [50,120,210]
<strong>Explanation:</strong> 
const t = performance.now()&nbsp;&nbsp;
const result = []
&nbsp; &nbsp; &nbsp; &nbsp;&nbsp;
const fn = () =&gt; {
    result.push(Math.floor(performance.now() - t))
}
const id = customInterval(fn, delay, period)
        
setTimeout(() =&gt; {
    customClearInterval(id)
}, 225)

50 + 20 * 0 = 50 // 50ms - 1st function call
50 + 20&nbsp;* 1 = 70 // 50ms + 70ms = 120ms - 2nd function call
50 + 20 * 2 = 90 // 50ms + 70ms + 90ms = 210ms - 3rd function call
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre><strong>Input:</strong> delay = 20, period = 20, stopTime = 150
<strong>Output:</strong> [20,60,120]
<strong>Explanation:</strong> 
20 + 20 * 0 = 20 // 20ms - 1st function call
20 + 20&nbsp;* 1 = 40 // 20ms + 40ms = 60ms - 2nd function call
20 + 20 * 2 = 60 // 20ms + 40ms + 60ms = 120ms - 3rd function call
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre><strong>Input:</strong> delay = 100, period = 200, stopTime = 500
<strong>Output:</strong> [100,400]
<strong>Explanation:</strong> 
100 + 200 * 0 = 100 // 100ms - 1st function call
100 + 200&nbsp;* 1 = 300 // 100ms + 300ms = 400ms - 2nd function call
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>20 &lt;= delay, period &lt;= 250</code></li>
	<li><code>20 &lt;= stopTime &lt;= 1000</code></li>
</ul>
</div>
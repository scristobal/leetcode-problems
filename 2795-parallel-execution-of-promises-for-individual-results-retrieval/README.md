<h2><a href="https://leetcode.com/problems/parallel-execution-of-promises-for-individual-results-retrieval/">2795. Parallel Execution of Promises for Individual Results Retrieval</a></h2><h3>Medium</h3><hr><div><p>Given an array&nbsp;<code>functions</code>, return a promise <code>promise</code>. <code>functions</code>&nbsp;is an array of functions that return promises <code>fnPromise.</code>&nbsp;Each <code>fnPromise</code>&nbsp;can be resolved or rejected.&nbsp;&nbsp;</p>

<p>If&nbsp;<code>fnPromise</code> is resolved:</p>

<p>&nbsp; &nbsp; <code>obj = { status: "fulfilled", value: <em>resolved value</em>}</code></p>

<p>If&nbsp;<code>fnPromise</code> is rejected:</p>

<p>&nbsp; &nbsp;&nbsp;<code>obj = { status: "rejected", reason: <em>reason of rejection (catched error message)</em>}</code></p>

<p>The <code>promise</code>&nbsp;should resolve with an array of these objects <code>obj</code>.&nbsp;Each <code>obj</code> in the array should correspond&nbsp;to the promises in the original array function, <strong>maintaining the same order</strong>.</p>

<p>Try to implement it without using the built-in method&nbsp;<code>Promise.allSettled()</code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre><strong>Input:</strong> functions = [
    () =&gt; new Promise(resolve =&gt; setTimeout(() =&gt; resolve(15), 100))
]
<strong>Output: </strong>{"t":100,"values":[{"status":"fulfilled","value":15}]}
<strong>Explanation:</strong> 
const time = performance.now()
const promise = promiseAllSettled(functions);
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
promise.then(res =&gt; {
    const out = {t: Math.floor(performance.now() - time), values: res}
    console.log(out) // {"t":100,"values":[{"status":"fulfilled","value":15}]}
})

The returned promise resolves within 100 milliseconds. Since promise from the array functions is fulfilled, the resolved value of the returned promise is set to [{"status":"fulfilled","value":15}].
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre><strong>Input:</strong> functions = [
    () =&gt; new Promise(resolve =&gt; setTimeout(() =&gt; resolve(20), 100)), 
    () =&gt; new Promise(resolve =&gt; setTimeout(() =&gt; resolve(15), 100))
]
<strong>Output: 
</strong>{
    "t":100,
    "values": [
&nbsp;       {"status":"fulfilled","value":20},
&nbsp;       {"status":"fulfilled","value":15}
    ]
}
<strong>Explanation:</strong> The returned promise resolves within 100 milliseconds, because the resolution time is determined by the promise that takes the longest time to fulfill. Since promises from the array functions are fulfilled, the resolved value of the returned promise is set to [{"status":"fulfilled","value":20},{"status":"fulfilled","value":15}].
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre><strong>Input:</strong> functions = [
&nbsp;   () =&gt; new Promise(resolve =&gt; setTimeout(() =&gt; resolve(30), 200)), 
&nbsp;   () =&gt; new Promise((resolve, reject) =&gt; setTimeout(() =&gt; reject("Error"), 100))
]
<strong>Output:</strong>
{
    "t":200,
    "values": [
        {"status":"fulfilled","value":30},
        {"status":"rejected","reason":"Error"}
    ]
}
<strong>Explanation:</strong> The returned promise resolves within 200 milliseconds, as its resolution time is determined by the promise that takes the longest time to fulfill. Since one promise from the array function is fulfilled and another is rejected, the resolved value of the returned promise is set to an array containing objects in the following order: [{"status":"fulfilled","value":30}, {"status":"rejected","reason":"Error"}]. Each object in the array corresponds to the promises in the original array function, maintaining the same order.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= functions.length &lt;= 10</code></li>
</ul>
</div>
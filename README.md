# codejam

This is my playground project for 
[Google CodeJam](https://codingcompetitions.withgoogle.com/codejam) / 
[KickStart](https://codingcompetitions.withgoogle.com/kickstart) / 
[CodeForces](https://codeforces.com/contests) tasks.
Individual solutions are under respectively named src subfolders, published under MIT license.

I try to avoid cryptic blobs of instructions and crowded one-liners, keeping solutions understandable and sanely named.
So, the code is sometimes more understandable than sportsmanlike, and I prefer it to be this way.

All recent solutions are in Go with minimal testing automation done in [bash](run).
Several years ago there'd been gradle automation for some Java solutions, still retained but not looked after.

I am using [GoLand](https://www.jetbrains.com/?from=codejam) to format/inspect/tidy up the code and carry out usual git tasks.
There are some handy live templates saving couple of keystrokes during the contest, too.

[<img src="jetbrains-variant-4.svg" alt="Powered by Jetbrains">](https://www.jetbrains.com/?from=codejam)

## Codebook

There're some pre-canned bits and pieces which look a bit like golang codebook.

The [template](src/template/solution.go) alleviates some input/output buffering/performance and 
general verbosity issues of Go.

Another interesting bit is [Competitive Companion Companion](src/z_comp3/main.go), 
  which mints a fresh copy of the template,
  with test-cases scraped directly off the contest platform UI. 
This integrates with [Competitive Companion](https://github.com/jmerle/competitive-companion#explanation) browser extension.

* [google's b+tree](src/algo/btree.go), with several fancy feature commits stripped for brevity (Apache 2.0)
* networks modelled as [adjacency list](src/algo/adjacency_list.go) and [capacity matrix](src/algo/capacity_matrix.go)
* [global min-cut](src/algo/global_min_cut.go) algorithm, ford-fulkerson [max flow](src/algo/max_flow_ford_fulkerson.go)
* [largest common subsequence](src/algo/lcs.go)
* [sequentially enumerating permutations](src/algo/permutations.go)
* [sequentially enumerating all subsets with a bitset](src/algo/subset_flags.go)
* and, of course, our old friend, [union-find](src/algo/union_find.go)

Most of this stuff passed in several submissions over time and got solid green results.
Note that there are other defects which you may still observe in your particular instance.

Also, there're notable pieces which you'd need to rip out of specific solutions:
* [prime factorization](src/hr_woc34_25_gcd_and_sum/solution.go:39) - not completely naive, but nothing _really_ efficient as well
* [wheel prime generation](src/cj16_0q_c_coin_jam/Main.java:15)
* [trie (java)](src/hr_ctci_50_tries_contacts/Solution.java:10) - not really sure it's good for anything, the tests were rather weak


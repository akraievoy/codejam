# codejam

This is my playground project for codejam / hackerrank / codeforces tasks.
Individual solutions are under respectively named src subfolders.
Expect to see more test cases and unit-tests in submissions done recently - I automated some bits to run tests locally.
I now mostly use GoLang, having started from Java.

I try to avoid cryptic blobs of instructions and crowded one-liners, keeping stuff understandable and sanely named.
So, the code is sometimes more understandable than sportsmanlike, and I prefer it to be this way.

There're some pre-canned bits and pieces which look a bit like golang codebook:

* [google's b+tree](src/algo/btree.go), with several fancy feature commits stripped for brevity
* networks modelled as [adjacency list](src/algo/adjacency_list.go) and [capacity matrix](src/algo/capacity_matrix.go)
* [global min-cut](src/algo/global_min_cut.go) algorithm, ford-fulkerson [max flow](src/algo/max_flow_ford_fulkerson.go)
* [largest common subsequence](src/algo/lcs.go)
* [sequentially enumerating permutations](src/algo/permutations.go)
* [sequentially enumerating all subsets with a bitset](src/algo/subset_flags.go)
* and, of course, our old friend, [union-find](src/algo/union_find.go)

Most of this stuff was used in several submissions over time and got solid green results.
Note that there are other defects which you may still observe in your particular instance.

Also there're notable pieces which you'd need to rip out of specific solutions:
* [prime factorization](src/hr_woc34_25_gcd_and_sum/solution.go:39) - not completely naive, but nothing relly deep as well
* [wheel prime generation](src/cj16_0q_c_coin_jam/Main.java:15)
* [trie (java)](src/hr_ctci_50_tries_contacts/Solution.java:10) - not really sure it's good for anything, the tests were rather weak


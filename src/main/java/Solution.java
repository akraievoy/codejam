import javax.annotation.Nonnull;
import java.util.*;

public class Solution {
  private static final int[] PRIMES =
    {
      2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
      31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
      73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
      127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
      179, 181, 191, 193, 197, 199
    };

  public final static class ModRem implements Comparable<ModRem> {
    public final int modulo;
    public final int remainder;

    public ModRem(int modulo, int remainder) {
      this.modulo = modulo;
      this.remainder = remainder;
    }

/*
    public long key() {
      //  we know that modulo is less than 2^16 = 65536, so we don't loose information stomping stuff into long
      return (((long)modulo) << 16) + ((long)remainder);
    }
*/

    @Override public int compareTo(@Nonnull ModRem that) {
      final int byModulo = Integer.compare(this.modulo, that.modulo);
      return byModulo == 0 ? Integer.compare(this.remainder, that.remainder) : byModulo;
    }

    @Override public boolean equals(Object o) {
      if (this == o) return true;
      if (o == null || getClass() != o.getClass()) return false;
      final ModRem modRem = (ModRem) o;
      return modulo == modRem.modulo && remainder == modRem.remainder;
    }

    @Override public int hashCode() {
      int result = modulo;
      result = 31 * result + remainder;
      return result;
    }
  }

  public static void main(String[] args) {
    //  the idea is to cache matching indexes
    //    for modRemKeys of all prime power / modulo pairs
    //    that we actually use in queries
    final Map<ModRem, int[]> modRemToIndexes = new HashMap<>(1000, 0.5f);

    final Scanner in = new Scanner(System.in);
    final int n = in.nextInt();
    final int queries = in.nextInt();

    final int[] a = new int[n];
    for (int i = 0; i < n; i++) {
      a[i] = in.nextInt();
    }
    final int[] lefts = new int[queries];
    final int[] rights = new int[queries];
    final List<List<ModRem>> queryModRems = new ArrayList<>();
    for (int q = 0; q < queries; q++) {
      lefts[q] = in.nextInt();
      rights[q] = in.nextInt();
      int queryModulo = in.nextInt();
      final int queryRemainder = in.nextInt();

      final List<ModRem> queryModRemsLocal = new ArrayList<>();
      int primeTriedIdx = 0;
      while (queryModulo > 1 && primeTriedIdx < PRIMES.length) {
        int primeTried = PRIMES[primeTriedIdx];
        primeTriedIdx++;
        int power = 1;
        while (queryModulo % primeTried == 0) {
          queryModulo /= primeTried;
          power *= primeTried;
        }
        if (power > 1) {
          queryModRemsLocal.add(new ModRem(power, queryRemainder % power));
        }
      }
      //  that's a prime, we tested all divisors below sqrt(40k)
      if (queryModulo > 1) {
        queryModRemsLocal.add(new ModRem(queryModulo, queryRemainder % queryModulo));
      }
      Collections.sort(queryModRemsLocal);
      queryModRems.add(queryModRemsLocal);
    }

    //  a bit of lousy collections code here:
    //    we won't have more than 120k unique mod/rem combinations
    //      even for the most carefully crafted adverse/evil test case
    //        each query has to use products of globally disjoint sets of mod/rems with mods we keep globally co-prime
    //          worst case is 7 mods (2,3,5,7,11,13 = 30030), average would be like 3 for 40k queries
    final Set<ModRem> uniqueModRemsSet = new HashSet<>(1000, 0.5f);
    for (int q = 0; q < queries; q++) {
      uniqueModRemsSet.addAll(queryModRems.get(q));
    }

    final ModRem[] uniqueModRems = uniqueModRemsSet.toArray(new ModRem[uniqueModRemsSet.size()]);
    Arrays.sort(uniqueModRems);

    if (uniqueModRems.length > 0) {
      final BitSet interestingRems = new BitSet();
      int mod = uniqueModRems[0].modulo;
      interestingRems.set(uniqueModRems[0].remainder);
      for (int mrI = 1; mrI < uniqueModRems.length; mrI++) {
        final ModRem modRem = uniqueModRems[mrI];
        if (modRem.modulo == mod) {
          interestingRems.set(modRem.remainder);
        } else {
          //   this won't be called more than 5k times (divisors we factor queries into are powers of primes under 40k)
          populateIndexes(a, mod, interestingRems, modRemToIndexes);
          mod = modRem.modulo;
          interestingRems.clear();
          interestingRems.set(modRem.remainder);
        }
      }
      populateIndexes(a, mod, interestingRems, modRemToIndexes);
    }

    final List<Integer> indexBuf = new ArrayList<>(40000);
    for (int q = 0; q < queries; q++) {
      final int left = lefts[q];
      final int right = rights[q];
      final List<ModRem> modRems = queryModRems.get(q);
      if (modRems.isEmpty()) {
        //  someone tries to query by module = 1?
        //  ohhhh... ohhhh...
        //  at least module is not zero, that would be really PATHETIC
        System.out.println(right - left + 1);
        continue;
      }
      final int[] sparsest = modRemToIndexes.get(modRems.get(modRems.size() - 1));
      if (modRems.size() == 1) {
        final int leftSearch = Arrays.binarySearch(sparsest, left);
        final int leftPos = leftSearch >= 0 ? leftSearch : -leftSearch - 1;
        final int rightSearch = Arrays.binarySearch(sparsest, right + 1);
        final int rightPos = rightSearch >= 0 ? rightSearch : -rightSearch - 1;
        System.out.println(rightPos - leftPos);
      } else {
        indexBuf.clear();
        final int leftSearch = Arrays.binarySearch(sparsest, left);
        final int leftPos = leftSearch >= 0 ? leftSearch : -leftSearch - 1;
        for (int i = leftPos; i < sparsest.length; i++) {
          int idx = sparsest[i];
          if (idx > right) {
            break;
          }
          indexBuf.add(idx);
        }
        for (int modRemI = modRems.size() - 2; !indexBuf.isEmpty() && modRemI >= 0; modRemI--) {
          final int lefter = indexBuf.get(0);
          final int righter = indexBuf.get(indexBuf.size() - 1);
          final int[] modRemIdx = modRemToIndexes.get(modRems.get(modRemI));
        }
        System.out.println(indexBuf.size());
      }
    }
  }

  private static void populateIndexes(int[] a, int mod, BitSet interestingRems, Map<ModRem, int[]> modRemToIndexes) {
    int[] indexesSizes = new int[mod];
    int[][] indexes = new int[mod][];
    for (int rem = 0; rem < mod; rem++) {
      if (interestingRems.get(rem)) {
        indexes[rem] = new int[a.length];
      }
    }
    for (int i = 0; i < a.length; i++) {
      final int aRem = a[i] % mod;
      if (indexes[aRem] != null) {
        indexes[aRem][indexesSizes[aRem]]= i;
        indexesSizes[aRem] += 1;
      }
    }
    for (int rem = 0; rem < mod; rem++) {
      if (indexes[rem] != null) {
        final int[] resIndexes = new int[indexesSizes[rem]];
        System.arraycopy(indexes[rem], 0, resIndexes, 0, resIndexes.length);
        modRemToIndexes.put(new ModRem(mod, rem), resIndexes);
      }
    }
  }
}

import java.io.BufferedWriter;
import java.io.OutputStreamWriter;
import java.io.PrintWriter;
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
  public static final boolean DEBUG = false;

  public final static class ModRem implements Comparable<ModRem> {
    public final int modulo;
    public final int remainder;

    public ModRem(int modulo, int remainder) {
      this.modulo = modulo;
      this.remainder = remainder;
    }

    @Override public int compareTo(ModRem that) {
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
    final Map<ModRem, int[]> modRemToIndexes = new HashMap<>(40000, 0.5f);

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
    final long dedupeStarted = System.currentTimeMillis();
    final Set<ModRem> uniqueModRemsSet = new HashSet<>(40000, 0.5f);
    for (int q = 0; q < queries; q++) {
      uniqueModRemsSet.addAll(queryModRems.get(q));
    }

    final ModRem[] uniqueModRems = uniqueModRemsSet.toArray(new ModRem[uniqueModRemsSet.size()]);
    Arrays.sort(uniqueModRems);
    final long dedupeEnded = System.currentTimeMillis();

    if (DEBUG) {
      System.err.println("uniqueModRems.length = " + uniqueModRems.length + " in " + (dedupeEnded-dedupeStarted) + " msec");
    }

    int divisionScans = 0;
    if (uniqueModRems.length > 0) {
      final int[] interestingRems = new int[40000];
      int mod = uniqueModRems[0].modulo;
      int interestingRemsSize = 1;
      interestingRems[0] = uniqueModRems[0].remainder;
      for (int mrI = 1; mrI < uniqueModRems.length; mrI++) {
        final ModRem modRem = uniqueModRems[mrI];
        if (modRem.modulo == mod) {
          interestingRems[interestingRemsSize] = modRem.remainder;
          interestingRemsSize += 1;
        } else {
          //   this won't be called more than 5k times (divisors we factor queries into are powers of primes under 40k)
          populateIndexes(a, mod, interestingRems, interestingRemsSize, modRemToIndexes);
          divisionScans += 1;
          mod = modRem.modulo;
          interestingRemsSize = 1;
          interestingRems[0] = modRem.remainder;
        }
      }
      populateIndexes(a, mod, interestingRems, interestingRemsSize, modRemToIndexes);
      divisionScans += 1;
    }

    if (DEBUG) {
      System.err.println("divisionScans = " + divisionScans);
    }

    final PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out), 4 * 1024 * 1024));

    int intersectionSize;
    int[] intersection = new int[40000];
    int[] intersectionTemp = new int[40000];
    for (int q = 0; q < queries; q++) {
      final int left = lefts[q];
      final int right = rights[q];
      final List<ModRem> modRems = queryModRems.get(q);

      if (modRems.isEmpty()) {
        //  someone tries to query by module = 1?
        //  ohhhh... ohhhh...
        //  at least module is not zero, that would be really PATHETIC
        out.println(right - left + 1);
        continue;
      }

      final int[] sparsest = modRemToIndexes.get(modRems.get(modRems.size() - 1));
      if (modRems.size() == 1) {
        final int leftSearch = Arrays.binarySearch(sparsest, left);
        final int leftPos = leftSearch >= 0 ? leftSearch : -leftSearch - 1;
        final int rightSearch = Arrays.binarySearch(sparsest, right + 1);
        final int rightPos = rightSearch >= 0 ? rightSearch : -rightSearch - 1;
        out.println(rightPos - leftPos);
      } else {

        intersectionSize = 0;
        final int leftSearch = Arrays.binarySearch(sparsest, left);
        final int leftPos = leftSearch >= 0 ? leftSearch : -leftSearch - 1;
        for (int i = leftPos; i < sparsest.length; i++) {
          int idx = sparsest[i];
          if (idx > right) {
            break;
          }
          intersection[intersectionSize] = idx;
          intersectionSize += 1;
        }

        for (int modRemI = modRems.size() - 2; intersectionSize > 0 && modRemI >= 0; modRemI--) {
          final int lefter = intersection[0];
          final int righter = intersection[intersectionSize - 1];
          final int[] modRemIdx = modRemToIndexes.get(modRems.get(modRemI));

          final int lefterSearch = Arrays.binarySearch(modRemIdx, lefter);
          final int lefterPos = lefterSearch >= 0 ? lefterSearch : -lefterSearch - 1;
          final int righterSearch = Arrays.binarySearch(modRemIdx, righter + 1);
          final int righterPos = righterSearch >= 0 ? righterSearch : -righterSearch - 1;

          int intersectionPos = 0;
          int intersectionTempSize = 0;
          for (int idxPos = lefterPos; intersectionPos < intersectionSize && idxPos < righterPos; ) {
            final int compare = Integer.compare(intersection[intersectionPos], modRemIdx[idxPos]);
            if (compare < 0) {
              intersectionPos++;
            } else if (compare == 0) {
              intersectionTemp[intersectionTempSize] = intersection[intersectionPos];
              intersectionTempSize += 1;
              idxPos++;
              intersectionPos++;
            } else {
              idxPos++;
            }
          }

          final int[] swapTemp = intersection;
          intersection = intersectionTemp;
          intersectionTemp = swapTemp;
          intersectionSize = intersectionTempSize;
        }

        out.println(intersectionSize);
      }
    }

    out.flush();
    out.close();
  }

  private static void populateIndexes(int[] a, int mod, int[] interestingRems, int interestingRemsSize, Map<ModRem, int[]> modRemToIndexes) {
    int[] indexesSizes = new int[mod];
    int[][] indexes = new int[mod][];
    for (int interestingRemIdx = 0; interestingRemIdx < interestingRemsSize; interestingRemIdx++) {
      indexes[interestingRems[interestingRemIdx]] = new int[a.length];
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

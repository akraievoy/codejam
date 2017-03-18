import java.io.BufferedWriter;
import java.io.ByteArrayInputStream;
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
    Random random = null;
    boolean selfTest = false;
    if (args.length > 0) {
      if (args[0].equals("test1")) {
        System.setIn(new ByteArrayInputStream((
          "5 3\n" +
          "250 501 5000 5 4\n" +
          "0 4 5 0\n" +  // --> 3
          "0 4 10 0\n" + // --> 2
          "0 4 3 2"      // --> 2
        ).getBytes()));
      } else if (args[0].equals("test2")) {
        System.setIn(new ByteArrayInputStream((
          "100 6\n" +
          "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99\n" +
          "0 99 33 0\n" + //  --> 4
          "0 99 33 1\n" + //  --> 3
          "0 99 33 2\n" + //  --> 3
          "0 99 7 5\n" +  //  --> 14
          "0 99 98 1\n" + //  --> 2
          "0 99 19 2"     //  --> 6
        ).getBytes()));
      } else if (args[0].equals("test3")) {
        System.setIn(new ByteArrayInputStream((
          "100 12\n" +
          "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99\n" +
          "0 0 100 0\n" + //  --> 1
          "0 1 100 0\n" + //  --> 1
          "1 1 100 0\n" + //  --> 0
          "49 51 100 50\n" + //  --> 1
          "49 50 100 50\n" + //  --> 1
          "50 50 100 50\n" + //  --> 1
          "50 51 100 50\n" + //  --> 1
          "51 51 100 50\n" + //  --> 0
          "49 49 100 50\n" + //  --> 0
          "98 98 100 99\n" +  //  --> 0
          "98 99 100 99\n" + //  --> 1
          "99 99 100 99"     //  --> 1
        ).getBytes()));
      } else if (args[0].equals("test4")) {
        selfTest = true;
        final long seed;
        if (args.length > 1) {
          seed = Long.parseLong(args[1]);
        } else {
          seed = System.currentTimeMillis();
        }
        System.out.println("seed = " + seed);
        random = new Random(seed);
      }
    }

    //  the idea is to cache matching indexes
    //    for modRemKeys of all prime power / modulo pairs
    //    that we actually use in queries
    final Map<ModRem, int[]> modRemToIndexes = new HashMap<>(1000, 0.5f);

    final Scanner in = new Scanner(System.in);
    final int n = selfTest ? 40000 : in.nextInt();
    final int queries = selfTest ? 40000 : in.nextInt();

    final int[] a = new int[n];
    for (int i = 0; i < n; i++) {
      a[i] = selfTest ? i : in.nextInt();
    }
    final int[] lefts = new int[queries];
    final int[] rights = new int[queries];
    final List<List<ModRem>> queryModRems = new ArrayList<>();
    for (int q = 0; q < queries; q++) {
      lefts[q] = selfTest ? random.nextInt(20000) : in.nextInt();
      rights[q] = selfTest ? 20000 + random.nextInt(20000) : in.nextInt();
      int queryModulo = selfTest ? 1 + random.nextInt(40000) : in.nextInt();
      final int queryRemainder = selfTest ? random.nextInt(queryModulo) : in.nextInt();

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

    if (selfTest) {
      System.out.println("uniqueModRems.length = " + uniqueModRems.length);
    }

    int divisionScans = 0;
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
          divisionScans += 1;
          mod = modRem.modulo;
          interestingRems.clear();
          interestingRems.set(modRem.remainder);
        }
      }
      populateIndexes(a, mod, interestingRems, modRemToIndexes);
      divisionScans += 1;
    }

    if (selfTest) {
      System.out.println("divisionScans = " + divisionScans);
    }

    final PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out), 4 * 1024 * 1024));

    List<Integer> intersection = new ArrayList<>(40000);
    List<Integer> intersectionTemp = new ArrayList<>(40000);
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

        intersection.clear();
        final int leftSearch = Arrays.binarySearch(sparsest, left);
        final int leftPos = leftSearch >= 0 ? leftSearch : -leftSearch - 1;
        for (int i = leftPos; i < sparsest.length; i++) {
          int idx = sparsest[i];
          if (idx > right) {
            break;
          }
          intersection.add(idx);
        }

        for (int modRemI = modRems.size() - 2; !intersection.isEmpty() && modRemI >= 0; modRemI--) {
          final int lefter = intersection.get(0);
          final int righter = intersection.get(intersection.size() - 1);
          final int[] modRemIdx = modRemToIndexes.get(modRems.get(modRemI));

          final int lefterSearch = Arrays.binarySearch(modRemIdx, lefter);
          final int lefterPos = lefterSearch >= 0 ? lefterSearch : -lefterSearch - 1;
          final int righterSearch = Arrays.binarySearch(modRemIdx, righter + 1);
          final int righterPos = righterSearch >= 0 ? righterSearch : -righterSearch - 1;

          int intersectionPos = 0;
          intersectionTemp.clear();
          for (int idxPos = lefterPos; intersectionPos < intersection.size() && idxPos < righterPos; ) {
            final int compare = Integer.compare(intersection.get(intersectionPos), modRemIdx[idxPos]);
            if (compare < 0) {
              intersectionPos++;
            } else if (compare == 0) {
              intersectionTemp.add(intersection.get(intersectionPos));
              idxPos++;
              intersectionPos++;
            } else {
              idxPos++;
            }
          }

          final List<Integer> swapTemp = intersection;
          intersection = intersectionTemp;
          intersectionTemp = swapTemp;
        }

        out.println(intersection.size());
      }
    }

    out.flush();
    out.close();
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

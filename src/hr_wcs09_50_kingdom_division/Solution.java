import java.math.BigInteger;
import java.util.Arrays;
import java.util.Scanner;

public class Solution {
  private final static int MOD = 1000000007;
  private final static BigInteger TWO = BigInteger.valueOf(2);

  private static long link(int from, int into) {
    return (((long) from) << 24) + into;
  }

  private static int from(long link) {
    return (int) (link >> 24);
  }

  private static int into(long link) {
    return (int) (link & ((1L << 24) - 1));
  }

  private static BigInteger[] coloringsDFS(
    final int vertexCur,
    final int vertexPrev,
    final long[] links,
    final int[] fingers
  ) {
    BigInteger allPossibleColorings = TWO;
    BigInteger allOpenColorings = TWO;

    for (int f = fingers[vertexCur]; f < fingers[vertexCur + 1]; f++) {
      final int vertexNext = into(links[f]);
      if (vertexNext == vertexPrev) {
        continue;
      }
      final BigInteger[] colorings = coloringsDFS(vertexNext, vertexCur, links, fingers);
      final BigInteger coloringsOpen = colorings[0];
      final BigInteger coloringsClosed = colorings[1];

      allPossibleColorings =
        allPossibleColorings.multiply(
          // all independent closed colorings of current subtree
          coloringsClosed.add(
            // all same-color open colorings of current subtree
            coloringsOpen.shiftRight(1)
          )
        );

      //  closed colorings for different color, thus dividing by two
      allOpenColorings =
        allOpenColorings.multiply(coloringsClosed).shiftRight(1);

    }

    return new BigInteger[] {
      allOpenColorings,
      allPossibleColorings.subtract(allOpenColorings)
    };
  }

  public static void main(String[] args) {
    final Scanner in = new Scanner(System.in);
    final int n = in.nextInt();

    final long[] links = new long[2 * n - 2];
    for(int a1 = 0; a1 < n-1; a1++){
      final int u = in.nextInt() - 1;
      final int v = in.nextInt() - 1;
      links[a1 * 2] = link(u, v);
      links[a1 * 2 + 1] = link(v,u);
    }

    Arrays.sort(links);
    final int[] fingers = new int[n+1];
    fingers[0] = 0;
    fingers[n] = links.length;
    int fromPrev = 0;
    for (int e = 1; e < links.length; e++) {
      final int fromCur = from(links[e]);
      if (from(links[e]) != fromPrev) {
          fingers[fromCur] = e;
        fromPrev = fromCur;
      }
    }

    final BigInteger[] colorings = coloringsDFS(0, -1, links, fingers);

    System.out.println(colorings[1].mod(BigInteger.valueOf(MOD)));
  }
}

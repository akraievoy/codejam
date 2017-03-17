import java.util.Scanner;

public class Main {

  public static void main(String[] args) {
    final Scanner in = new Scanner(System.in);
    final int poles = in.nextInt() /*5000*/;
    final int stacks = in.nextInt() /*2500*/;

    final int[] heights = new int[poles];
    final int[] weights = new int[poles];
    for(int poleI = 0; poleI < poles; poleI++) {
      heights[poleI] = in.nextInt() /*1 + poleI*/;
      weights[poleI] = in.nextInt() /*1*/;
    }

    final long[][] poleDropCosts = new long[poles][poles];
    for (int pI = 0; pI < poles; pI++) {
      for (int pJ = 0; pJ < poles; pJ++) {
        poleDropCosts[pI][pJ] = pJ > pI ? Long.MIN_VALUE : weights[pI] * (heights[pI] - heights[pJ]);
      }
    }

    final long[][] rangeDropCosts = new long[poles][poles];
    for (int pI = poles-1; pI >= 0; pI--) {
      for (int pJ = 0; pJ < poles; pJ++) {
        rangeDropCosts[pI][pJ] =
          pJ > pI ?
            Long.MIN_VALUE :
            poleDropCosts[pI][pJ] +
              (pI == poles-1 ? 0 : rangeDropCosts[pI+1][pJ]);
      }
    }

    final long[][] subCosts = new long[poles][poles];
    for (int pI = 0; pI < poles; pI++) {
      for (int pJ = 0; pJ < poles; pJ++) {
        subCosts[pI][pJ] = Long.MAX_VALUE;
      }
    }

    for (int lowestPole = poles - 1; lowestPole >= 0; lowestPole-- ) {
      //  everything drops to lowest pole
      int subStacks = 1;
      subCosts[lowestPole][lowestPole] = rangeDropCosts[lowestPole][lowestPole];
      //  all poles have own stacks
      subStacks = poles - lowestPole;
      subCosts[lowestPole + subStacks - 1][lowestPole] = 0;
    }

    for (int subStacks = 2; subStacks <= stacks; subStacks++) {
      for (int lowestPole = 0; lowestPole < poles - subStacks; lowestPole++ ) {
        long subCost = Long.MAX_VALUE;
        final long fullRangeCost = rangeDropCosts[lowestPole][lowestPole];
        for (int nextStackPole = lowestPole + 1; nextStackPole < poles - subStacks + 2; nextStackPole++) {
          final long nextCost =
            fullRangeCost
             - rangeDropCosts[nextStackPole][lowestPole]
             + subCosts[nextStackPole + subStacks - 2][nextStackPole];

          if (subCost == Long.MAX_VALUE || subCost > nextCost) {
            subCost = nextCost;
          } else if (nextCost > 2 * subCost) {
            break; // the cost function is convex: if we see significant increase - we bail
          }
        }
        subCosts[lowestPole+subStacks - 1][lowestPole] = subCost;
      }
    }

    System.out.println(subCosts[stacks-1][0]);
  }
}

/*
3 1
20 1
30 1
40 1

6 2
1 1
2 2
3 1
4 1
5 2
6 1

6 3
1 1
2 1
3 1
4 1
5 1
6 1
 */

import java.io.*;
import java.util.Arrays;
import java.util.Optional;
import java.util.Scanner;
import java.util.function.BiFunction;

public class Main implements BiFunction<Scanner, PrintWriter, Optional<String>> {

    @Override
    public Optional<String> apply(Scanner in, PrintWriter out) {
        final int A = in.nextInt();
        final int N = in.nextInt();
        final int[] mots = new int[N];

        for (int i = 0; i < N; i++) {
            mots[i] = in.nextInt();
        }

        if (A == 1) {
            // this mot can not absorb anything, this we have to delete all other mots
            return Optional.of(String.valueOf(N));
        }

        Arrays.sort(mots);

        //  how many extra mots we have to add to absorb mot #i
        final int[] motsAdded = new int[N];
        //  what weight would our mot have after absorption of mot #i
        final long[] weight = new long[N];

        for (int pos = 0; pos < N; pos++) {
            long curWeight;
            if (pos > 0) {
                motsAdded[pos] = motsAdded[pos - 1];
                curWeight = weight[pos - 1];
            } else {
                motsAdded[pos] = 0;
                curWeight = A;
            }
            while (curWeight >= 2 && curWeight <= mots[pos]) {
                curWeight += curWeight - 1;
                motsAdded[pos]++;
            }
            weight[pos] = curWeight + mots[pos];
        }

        if (motsAdded[0] >= N) {
            // even to absorb first guy we have to make N ops, so it's better to nuke'em all
            return Optional.of(String.valueOf(N));
        }

        int minOps = motsAdded[N - 1];
        for (int pos = 0; pos < N - 1; pos++) {
            final int addedTillPosDeletedSince = motsAdded[pos] + N - pos - 1;
            if (addedTillPosDeletedSince < minOps) {
                minOps = addedTillPosDeletedSince;
            }
        }

        return Optional.of(String.valueOf(minOps));
    }

    public static void main(String[] args) {
        final Main main = new Main();

        try (
            final Scanner in = new Scanner(new BufferedReader(new InputStreamReader(System.in)));
            final PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out)));
        ) {
            final int testCount = in.nextInt();

            for (int index = 0; index < testCount; index++) {
                out.print("Case #" + (index+1)+": ");
                Optional<String> optResult = main.apply(in, out);
                if (optResult.isPresent()) {
                    out.println(optResult.get());
                } else {
                    out.println();
                }
            }
        }
    }
}

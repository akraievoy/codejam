import java.io.*;
import java.math.BigInteger;
import java.util.*;
import java.util.function.BiFunction;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class Main implements BiFunction<Scanner, PrintWriter, Optional<String>> {

    private static BigInteger tileIndexForUsed(int[] usedTiles, BigInteger originalTiles) {
        BigInteger tileIndex = BigInteger.ZERO;
        for (int usedTile : usedTiles) {
            tileIndex = tileIndex.multiply(originalTiles).add(BigInteger.valueOf(usedTile));
        }
        return tileIndex.add(BigInteger.ONE);
    }

    public Optional<String> apply(Scanner in, PrintWriter out) {
        final int originalTiles = in.nextInt(); // aka K
        final int transforms = in.nextInt(); // aka C
        final int maxTilesToOpen = in.nextInt(); // aka S

        final BigInteger originalTilesBigInt = BigInteger.valueOf(originalTiles);

        //  we are able to test maximally `transforms` tiles on each open, so
        if (maxTilesToOpen * transforms < originalTiles) {
            return Optional.of("IMPOSSIBLE"); // Tom Cruise to the rescue
        }

        final int tilesToOpen = (originalTiles + transforms - 1) / transforms;

        String result = IntStream.range(0, tilesToOpen)
            .mapToObj(index -> {
                final int[] usedTiles = new int[transforms];
                for (int usedTileIndex = 0; usedTileIndex < usedTiles.length; usedTileIndex++) {
                    usedTiles[usedTileIndex] = (index * transforms + usedTileIndex) % originalTiles;
                }
                return tileIndexForUsed(usedTiles, originalTilesBigInt).toString();
            })
            .collect(Collectors.joining(" "));

        return Optional.of(result);
    }


    public static void main(String[] args) {
        final Main main = new Main();

        try (
            final Scanner in = new Scanner(new BufferedReader(new InputStreamReader(System.in)));
            final PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out)));
        ) {
            final int testCount = in.nextInt();

            for (int index = 0; index < testCount; index++) {
                out.print("Case #" + (index + 1) + ": ");
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

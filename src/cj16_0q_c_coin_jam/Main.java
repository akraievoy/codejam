import java.io.*;
import java.math.BigInteger;
import java.util.*;
import java.util.function.BiFunction;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import java.util.stream.LongStream;

public class Main implements BiFunction<Scanner, PrintWriter, Optional<String>> {
    @SuppressWarnings("FieldCanBeLocal")
    private static long PRIME_LIMIT = (long) 100 * 1000 * 1000;
    private static final int MAX_LOW_BASE = 3;
    private static final int MAX_LEN = 32;

    public static long[] generatePrimes(long limit) {
        return generatePrimes(
            limit,
            210,
            new long[]{
                2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53,
                59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131,
                137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199
            }
        );
    }

    private static long[] generatePrimes(long limit, long upTo, long[] knownPrimes) {
        if (upTo * upTo >= limit) {
            long[] newPrimes = new long[knownPrimes.length];
            int taken = 0;
            for (long tested = upTo + 1; tested <= limit; tested += 2) {
                boolean isPrime = nonTrivialDivisor(tested, knownPrimes) == 0;
                if (isPrime) {
                    if (newPrimes.length == taken) {
                        long[] newPrimesLarger = new long[newPrimes.length * 2];
                        System.arraycopy(newPrimes, 0, newPrimesLarger, 0, newPrimes.length);
                        newPrimes = newPrimesLarger;
                    }
                    newPrimes[taken] = tested;
                    taken++;
                }
            }
            final long[] primesRes = new long[knownPrimes.length + taken];
            System.arraycopy(knownPrimes, 0, primesRes, 0, knownPrimes.length);
            System.arraycopy(newPrimes, 0, primesRes, knownPrimes.length, taken);
            return primesRes;
        } else {
            long[] morePrimes = generatePrimes(upTo * upTo, upTo, knownPrimes);
            return generatePrimes(limit, upTo * upTo, morePrimes);
        }
    }

    private static long nonTrivialDivisor(long tested, long[] knownPrimes) {
        final int testLimit = (int) Math.sqrt(tested);
        for (long curPrime : knownPrimes) {
            if (tested % curPrime == 0) {
                return curPrime;
            }
            if (curPrime > testLimit) {
                return 0;
            }
        }
        return 0;
    }

    private BigInteger nonTrivialDivisor(BigInteger coinValue, BigInteger[] primesUpToLimit) {
        for (BigInteger curPrime : primesUpToLimit) {
            if (coinValue.mod(curPrime).equals(BigInteger.ZERO)) {
                return curPrime;
            }
        }
        return BigInteger.ZERO;
    }

    private static final long[] PRIMES_UP_TO_LIMIT = generatePrimes();
    private static final long MAX_KNOWN_PRIME = PRIMES_UP_TO_LIMIT[PRIMES_UP_TO_LIMIT.length - 1];
    private static final BigInteger MAX_KNOWN_PRIME_BIGINT = BigInteger.valueOf(MAX_KNOWN_PRIME);
    private static final BigInteger MAX_KNOWN_PRIME_SQUARED_BIGINT = MAX_KNOWN_PRIME_BIGINT.pow(2);

    private static long[] generatePrimes() {
        return generatePrimes(PRIME_LIMIT);
    }

    private static final BigInteger[] PRIMES_UP_TO_LIMIT_BIGINT = generatePrimesBigInt();

    private static BigInteger[] generatePrimesBigInt() {
        final BigInteger[] bigintPrimes = new BigInteger[PRIMES_UP_TO_LIMIT.length];
        for (int i = 0; i < bigintPrimes.length; i++) {
            bigintPrimes[i] = BigInteger.valueOf(PRIMES_UP_TO_LIMIT[i]);
        }
        return bigintPrimes;
    }

    private final long[][] lowBasePowers = generateLowPowers();

    private long[][] generateLowPowers() {
        final long[][] lowPowers = new long[MAX_LOW_BASE + 1][MAX_LEN];

        IntStream.range(0, MAX_LOW_BASE + 1).forEach(base ->
            IntStream.range(0, MAX_LEN).forEach(power ->
                lowPowers[base][power] = LongStream.range(0, power).map(idx -> base).reduce((a, b) -> a * b).orElseGet(() -> 1L)
            )
        );

        return lowPowers;
    }

    private final BigInteger[][] highBasePowers = generateHighPowers();

    private BigInteger[][] generateHighPowers() {
        final BigInteger[][] highPowers = new BigInteger[11][MAX_LEN];

        IntStream.range(0, 11).forEach(base ->
            IntStream.range(0, MAX_LEN).forEach(power ->
                highPowers[base][power] = BigInteger.valueOf(base).pow(power)
            )
        );

        return highPowers;
    }

    private static final List<Integer> ALL_POWERS = IntStream.range(2, 11).boxed().collect(Collectors.toList());


    public Optional<String> apply(Scanner in, PrintWriter out) {
        int len = in.nextInt();
        int num = in.nextInt();

        final boolean[] bits = new boolean[len];
        long binValue = lowBasePowers[2][0] + lowBasePowers[2][len - 1];

        final long limit = 2 * lowBasePowers[2][len - 1];
        final List<String> found = new ArrayList<>();
        while (binValue < limit && found.size() < num) {
            final long currentValue = binValue;
            for (int pow = 0; pow < len; pow++) {
                bits[pow] = (currentValue & lowBasePowers[2][pow]) > 0;
            }
            Optional<String> optProvedCoin =
                foldLeft(
                    ALL_POWERS,
                    Optional.of(Long.toBinaryString(currentValue)),
                    (optCoin, nextBase) -> {
                        if (!optCoin.isPresent()) {
                            return Optional.empty(); //  None
                        }
                        String prevCoinState = optCoin.get();
                        BigInteger nonTrivialDivisor;
                        if (nextBase <= MAX_LOW_BASE) {
                            long coinValue;
                            if (nextBase == 2) {
                                coinValue = currentValue;
                            } else {
                                coinValue = 0;
                                for (int pow = 0; pow < len; pow++) {
                                    if (bits[pow]) {
                                        coinValue += lowBasePowers[nextBase][pow];
                                    }
                                }
                            }
                            if (coinValue <= MAX_KNOWN_PRIME) {
                                if (Arrays.binarySearch(PRIMES_UP_TO_LIMIT, coinValue) < 0) {
                                    nonTrivialDivisor = BigInteger.valueOf(nonTrivialDivisor(coinValue, PRIMES_UP_TO_LIMIT));
                                } else {
                                    nonTrivialDivisor = BigInteger.ZERO;
                                }
                            } else {
                                nonTrivialDivisor = BigInteger.valueOf(nonTrivialDivisor(coinValue, PRIMES_UP_TO_LIMIT));
                            }
                        } else {
                            BigInteger coinValue = BigInteger.ZERO;
                            for (int pow = 0; pow < len; pow++) {
                                if (bits[pow]) {
                                    coinValue = coinValue.add(highBasePowers[nextBase][pow]);
                                }
                            }
                            if (coinValue.compareTo(MAX_KNOWN_PRIME_BIGINT) <= 0) {
                                long longValue = coinValue.longValue();
                                if (Arrays.binarySearch(PRIMES_UP_TO_LIMIT, longValue) < 0) {
                                    nonTrivialDivisor = BigInteger.valueOf(nonTrivialDivisor(longValue, PRIMES_UP_TO_LIMIT));
                                } else {
                                    nonTrivialDivisor = BigInteger.ZERO;
                                }
                            } else if (coinValue.compareTo(MAX_KNOWN_PRIME_SQUARED_BIGINT) <= 0) {
                                nonTrivialDivisor = BigInteger.valueOf(nonTrivialDivisor(coinValue.longValue(), PRIMES_UP_TO_LIMIT));
                            } else {
                                nonTrivialDivisor = nonTrivialDivisor(coinValue, PRIMES_UP_TO_LIMIT_BIGINT);
                            }
                        }
                        if (nonTrivialDivisor.equals(BigInteger.ZERO)) {
                            return Optional.empty();
                        } else {
                            return Optional.of(prevCoinState + " " + nonTrivialDivisor);
                        }
                    }
                );
            if (optProvedCoin.isPresent()) {
                found.add(optProvedCoin.get());
            }
            binValue += 2;
        }

        final StringBuilder res = new StringBuilder();
        for (String coin : found) {
            res.append("\n");
            res.append(coin);
        }
        return Optional.of(res.toString());
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

    //  java 8 + the same :
    //    http://stackoverflow.com/questions/15787042/strange-reduce-method-group-in-jdk8-bulk-collection-operations-library
    //  or
    //    http://www.functionaljava.org/features.html
    //  or
    //    http://blog.jooq.org/2014/09/10/when-the-java-8-streams-api-is-not-enough/
    //  or, even, use code of those guys who do God's work
    //    https://github.com/goldmansachs/gs-collections
    //    https://github.com/goldmansachs/gs-collections-kata
    public static <E, R> R foldLeft(Iterable<E> elems, R init, BiFunction<R, E, R> foldFun) {
        R acc = init;
        for (E elem : elems) {
            acc = foldFun.apply(acc, elem);
        }
        return acc;
    }
}

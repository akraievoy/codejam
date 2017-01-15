import static org.testng.Assert.*;

import org.testng.annotations.Test;

import java.io.PrintWriter;
import java.io.StringReader;
import java.io.StringWriter;
import java.util.Arrays;
import java.util.Optional;
import java.util.Scanner;

public class MainTest {
    @Test
    public void testGeneratePrimes() {
        long[] primes = Main.generatePrimes((long) 1000 * 1000);
        assertTrue(Arrays.binarySearch(primes, 247) < 0);
        assertEquals(primes.length, 78498);
        assertEquals(primes[primes.length - 1], 999983);
    }

    @Test
    public void testSillyCase0() throws Exception {
        assertEquals(
            run("6 3"),
            "\n" +
                "100001 3 2 5 2 7 2 3 2 11\n" +
                "100011 5 13 3 31 43 3 73 7 3\n" +
                "100111 3 2 5 2 7 2 3 2 11"
        );
    }


    private static String run(String... lines) {
        final StringWriter inWriter = new StringWriter();
        final PrintWriter inPrintWriter = new PrintWriter(inWriter);
        for (String line : lines) {
            inPrintWriter.println(line);
        }
        inPrintWriter.close();

        final StringWriter stringWriter = new StringWriter();
        final PrintWriter printWriter = new PrintWriter(stringWriter);
        final Optional<String> maybeSingleLine = new Main().apply(
            new Scanner(new StringReader(inWriter.toString())),
            printWriter
        );
        maybeSingleLine.ifPresent(printWriter::print);
        printWriter.close();

        return stringWriter.toString();
    }

}
import org.testng.annotations.Test;

import java.io.PrintWriter;
import java.io.StringReader;
import java.io.StringWriter;
import java.util.Optional;
import java.util.Scanner;

import static org.testng.Assert.*;

public class MainTest {

    @Test
    public void testSillyCase0() throws Exception {
        assertEquals(run("2 2", "2 1"), "0");
    }

    @Test
    public void testSillyCase1() throws Exception {
        assertEquals(run("2 4", "2 1 1 6"), "1");
    }

    @Test
    public void testSillyCase2() throws Exception {
        assertEquals(run("10 4", "25 20 9 100"), "2");
    }

    @Test
    public void testSillyCase3() throws Exception {
        assertEquals(run("1 4", "1 1 1 1"), "4");
    }

    private static String run(String... lines) {
        final String lineSep = System.getProperty("line.separator");
        final StringBuilder buf = new StringBuilder();
        for (String line : lines) {
            buf.append(line);
            buf.append(lineSep);
        }

        final StringWriter stringWriter = new StringWriter();
        final PrintWriter printWriter = new PrintWriter(stringWriter);
        final Optional<String> maybeSingleLine = new Main().apply(
            new Scanner(new StringReader(buf.toString())),
            printWriter
        );
        maybeSingleLine.ifPresent(printWriter::print);
        printWriter.flush();
        printWriter.close();

        return stringWriter.toString();
    }
}

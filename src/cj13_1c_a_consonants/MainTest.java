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
        assertEquals(run("quartz 3"), "4");
    }

    @Test
    public void testSillyCase1() throws Exception {
        assertEquals(run("straight 3"), "11");
    }

    @Test
    public void testSillyCase2() throws Exception {
        assertEquals(run("gcj 2"), "3");
    }

    @Test
    public void testSillyCase3() throws Exception {
        assertEquals(run("tsetse 2"), "11");
    }

    @Test
    public void testSillyCase4() throws Exception {
        assertEquals(run("abbbabbba 3"), "20");
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

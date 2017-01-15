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
        assertEquals(run("?6? ?0?"), "069 100");
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
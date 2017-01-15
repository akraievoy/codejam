package org.akraievoy.codejam;

import com.google.common.base.Stopwatch;
import com.google.common.base.Throwables;
import org.akraievoy.style.tuples.Tuple2;

import java.io.*;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;
import java.util.concurrent.*;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class Main implements Callable<String> {
  private final String input;

  public Main(String input) {
    this.input = input;
  }

  @Override public String call() throws Exception {
    return ">>> please solve for input " + input + " <<<";
  }

  public static void main(String[] args) throws IOException, InterruptedException, ExecutionException, TimeoutException {
    final ExecutorService runners =
      Executors.newFixedThreadPool(Runtime.getRuntime().availableProcessors());
    final BufferedReader filesChanged = new BufferedReader(new InputStreamReader(System.in));

    System.out.println("[?] Reading list of input file names from System.in");

    String inputFile;
    while ((inputFile = filesChanged.readLine()) != null) {
      if (!inputFile.endsWith(".in")) {
        System.out.println("[!] Skipping file '" + inputFile + "', as its name does not look as valid codejam input");
        continue;
      } else {
        System.out.println("[#] Started processing file '" + inputFile + "'");
      }

      final String outFileName = inputFile + "." + System.currentTimeMillis() + ".out.txt";

      try (
        final BufferedReader in = openReader(inputFile);
        final PrintWriter solutionOut = new PrintWriter(new FileWriter(outFileName))
      ) {
        final Stopwatch sinceStart = Stopwatch.createStarted();
        final int testCount = Integer.parseInt(in.readLine().trim());

        final List<Tuple2<Integer, CompletableFuture<String>>> indexedFutures = IntStream
          .range(1, testCount + 1)
          .mapToObj(testIndex -> {
            try {
              final String input = in.readLine().trim();
              final CompletableFuture<String> resultFuture =
                CompletableFuture.supplyAsync(
                  () -> {
                    try {
                      final Stopwatch stopwatch = Stopwatch.createStarted();
                      final String result = new Main(input).call();
                      System.out.println(" ~  Time for #" + testIndex + ": " + stopwatch.toString());
                      return result;
                    } catch (Exception e) {
                      throw Throwables.propagate(e);
                    }
                  },
                  runners
                );
              return new Tuple2<>(testIndex, resultFuture);
            } catch (IOException e) {
              throw Throwables.propagate(e);
            }
          })
          .collect(Collectors.toList());

        System.out.println("[~] Finished parsing and submitting in " + sinceStart.toString());

        final Stopwatch sinceParsed = Stopwatch.createStarted();

        final List<CompletableFuture<String>> allFutures =
          indexedFutures.stream()
            .map(tuple -> tuple._2)
            .collect(Collectors.toList());

        final CompletableFuture<Void> futureOfAll =
          CompletableFuture.allOf(allFutures.toArray(new CompletableFuture<?>[allFutures.size()]));

        futureOfAll.get(8, TimeUnit.MINUTES);

        System.out.println("[~] Completed waiting for individual tasks in " + sinceParsed.toString() + " since parse, " + sinceStart.toString() + " since start");
        final Stopwatch sinceComplete = Stopwatch.createStarted();

        indexedFutures.forEach( tuple -> {
          final int index = tuple._1;
          final String result;
          try {
            result = tuple._2.get();
          } catch (InterruptedException | ExecutionException e) {
            throw Throwables.propagate(e);
          }
          solutionOut.println("Case #"+index+": " + result);
        });

        System.out.println("[#] Written solutions in " + sinceComplete + " since complete, " + sinceStart.toString() + " since start");
      } catch (Throwable t) {
        System.err.println("[!] failed with " + t.getMessage());
        t.printStackTrace();
      } finally {
        System.out.println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-");
      }
    }

    runners.shutdown();
    runners.awaitTermination(12, TimeUnit.MINUTES);
  }

  static BufferedReader openReader(String inputFile) throws IOException {
    if (Files.exists(Paths.get(inputFile))) {
      return new BufferedReader(new FileReader(Paths.get(inputFile).toFile()));
    } else {
      throw new IllegalStateException("file " + inputFile + " does not exist");
    }
  }

}

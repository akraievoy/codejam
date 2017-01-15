package org.akraievoy.style;

import java.util.function.BiFunction;

public class Collections {
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

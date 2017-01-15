package org.akraievoy.style.tuples;

import javax.annotation.Nonnull;

public final class Tuple2<A1,A2> {
  @Nonnull public final A1 _1;
  @Nonnull public final A2 _2;

  public Tuple2(@Nonnull A1 _1, @Nonnull A2 _2) {
    this._1 = _1;
    this._2 = _2;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;

    final Tuple2 tuple2 = (Tuple2) o;
    return _1.equals(tuple2._1) && _2.equals(tuple2._2);
  }

  @Override
  public int hashCode() {
    int result = _1.hashCode();
    result = 31 * result + _2.hashCode();
    return result;
  }

  @Override public String toString() {
    return "(" + _1 + "," + _2 + ')';
  }
}

package org.akraievoy.style.tuples;

public class Tuple4<E0, E1, E2, E3> {
  public final E0 _0;
  public final E1 _1;
  public final E2 _2;
  public final E3 _3;

  public Tuple4(E0 _0, E1 _1, E2 _2, E3 _3) {
    this._0 = _0;
    this._1 = _1;
    this._2 = _2;
    this._3 = _3;
  }

  @Override public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;

    Tuple4<?, ?, ?, ?> tuple4 = (Tuple4<?, ?, ?, ?>) o;

    return _0.equals(tuple4._0) && _1.equals(tuple4._1) && _2.equals(tuple4._2) && _3.equals(tuple4._3);
  }

  @Override public int hashCode() {
    int result = _0.hashCode();
    result = 31 * result + _1.hashCode();
    result = 31 * result + _2.hashCode();
    result = 31 * result + _3.hashCode();
    return result;
  }

  @Override public String toString() {
    return "(" + _0 + "," + _1 + "," + _2 + "," + _3 + ')';
  }
}

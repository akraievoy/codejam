br_ren() {
  git checkout "$1" && \
  git checkout -b "$2" && \
  git push origin "$2" && \
  git branch -D "$1" && \
  git push origin ":$1" &&
  echo "renamed branch $1 to $2"
}

with import <nixpkgs> {};
mkShell {
  nativeBuildInputs = [
    go
    just
    cobra-cli
  ];
}

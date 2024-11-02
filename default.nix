{ pkgs ? import <nixpkgs> {} }:

pkgs.buildGoModule rec {
    pname = "memodir";
    version = "0.0.1";

    # only change this if go.mod changes! (i think)
    vendorHash = "sha256-pc0iLRslEwnKbRYrQJBMD1w0flr7ZweJ8iBDvSk/w2M=";

    src = ./.;

    meta = with pkgs.lib; {
      description = "A possibly blazingly fast CLI tool to remember and query selected directories";
      homepage = "https://github.com/sawcce/memodir";
      license = licenses.mit;
      # maintainers = with maintainers; [ # me's not on nixpkgs];
      mainProgram = "memodir";
    };
}
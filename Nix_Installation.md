# Nix installation Guide!

> [!NOTE]
> the maintainer of the nix version (and possibly a nixpkg), [lucoso](https://github.com/luwucoso), is not very good with nix! if you know any better ways to implement any of this, please make a pr or open an issue!

you have two ways to install this, cloning the repo and using the default.nix or taking the default.nix and replacing the src with a fetchFromGitHub

## 1: clone the repo
```
git clone https://github.com/sawcce/memodir.git
```

unless you want to use fetchFromGitHub, in that case just copy the default.nix to a folder named `memodir` and change the `src = ./.` to 
```nix
src = fetchFromGitHub {
	owner = "sawcce";
	repo = "memodir";
	rev = "main"; # probably should pick an actual revision instead of main
	hash = pkgs.lib.fakeHash
}
```
the build should error out saying that the hash doesn't match but that's expected! just replace the `pkgs.lib.fakeHash` to the actual hash the command says, inbetween quotation marks


## 2: make an overlay

now to use memodir as a normal program with `pkgs.memodir` you need to make an overlay

you should check out the wiki page for overlays, [over here](https://wiki.nixos.org/wiki/Overlays), but here's a simplified version

make another nix file, preferably outside the memodir directory, and put this in it:
```nix
# overlay.nix
self: super:
{
  memodir = self.callPackage ./memodir { }; # "./memodir" should be the directory that contains the memodir's default.nix 
}
```

## 3: using the overlay

now if you're using a shell.nix put this in:
```nix
import <nixpkgs> { overlays = [ (import ./overlay.nix) ]; }
# ...
```
assuming that the overlay.nix is in the same directory as the shell.nix

now you can use it just like any other package in the shell.nix!

if you want to use it in flakes you can check [this section of the wiki](https://wiki.nixos.org/wiki/Overlays#In_a_Nix_flake)
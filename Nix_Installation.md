# Nix Installation Guide! 
> [!NOTE] 
> The maintainer of the nix version (and possibly a nixpkg), [Lucoso](https://github.com/luwucoso), is not very good with nix! If you know any better ways to implement any of this, please make a PR or open an issue! 

You have two ways to install this: cloning the repo and using the default.nix or taking the default.nix and replacing the src with a fetchFromGitHub 

## 1: Clone the repo 
``` git clone https://github.com/sawcce/memodir.git ``` 
Unless you want to use fetchFromGitHub, in that case just copy the default.nix to a folder named `memodir` and change the `src = ./.` to 
```nix 
src = fetchFromGitHub { 
	owner = "sawcce"; 
	repo = "memodir"; 
	rev = "main"; # probably should pick an actual revision instead of main 
	hash = pkgs.lib.fakeHash 
	} 
``` 
The build should error out saying that the hash doesn't match, but that's expected! Just replace the `pkgs.lib.fakeHash` to the actual hash the command says, in between quotation marks. 

## 2: Make an overlay 
now to use memodir as a normal program with 'pkgs.memodir' You need to make an overlay; 
you should check out the wiki page for overlays, [over here](https://wiki.nixos.org/wiki/Overlays), but here's a simplified version: 

Make another nix file, preferably outside the memodir directory, and put this in it: 
```nix 
# overlay.nix 
self: super: { 
	memodir = self.callPackage ./memodir { }; # "./memodir" should be the directory that contains the memodir's default.nix 
} 
``` 

## 3: Using the overlay

If you want to use it in flakes, you can check [this section of the wiki](https://wiki.nixos.org/wiki/Overlays#In_a_Nix_flake)

Now if you're using a shell.nix put this in: 
```nix 
import { overlays = [ (import ./overlay.nix) ]; } 
# ... 
``` 
assuming that the overlay.nix is in the same directory as the shell.nix
now you can use it just like any other package in the shell.nix!  
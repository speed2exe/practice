# AppFlowy shell.nix (rust-lib)

{ pkgs ? import <nixpkgs> {} }:

# let
#    unstable = import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/nixos-unstable.tar.gz") {};
# in

pkgs.mkShell {

  packages = with pkgs; [
    go
  ];

  shellHook = ''
  '';
}

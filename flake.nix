{
  description = "go dev env";

  inputs = {
    # Nix Packages collection & NixOS
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    supportedSystems = [
      "x86_64-linux"
      "aarch64-linux"
      "x86_64-darwin"
      "aarch64-darwin"
    ];
    forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
  in {
    devShells = forAllSystems (
      system: let
        pkgs = import nixpkgs {
          inherit system;
        };
      in {
        default = pkgs.mkShell {
          packages = with pkgs; [
            # Go Programming language
            go

            # Official language server for the Go language
            gopls

            # Additional tools for Go development
            gotools

            # Collection of tools and libraries for working with Go code, including linters and static analysis
            go-tools

            # Fast linters Runner for Go
            golangci-lint

            # Debugger for the Go programming language
            delve
          ];

          shellHook = ''
            echo "go environment loaded!"
            echo "  compile linux:   GOOS=linux GOARCH=amd64 go build -ldflags=\"-s -w\" -o app main.go"
            echo "  compile windows: GOOS=windows GOARCH=amd64 go build -ldflags=\"-s -w\" -o app.exe main.go"
            echo "  compile macOS:   GOOS=darwin GOARCH=arm64 go build -ldflags=\"-s -w\" -o app-mac main.go"
            echo ""
            echo "useful commands:"
            echo "  tidy up deps: go mod tidy"
            echo "  format code: go fmt ./..."
            echo "  run linters: golangci-lint run"
          '';
        };
      }
    );
  };
}

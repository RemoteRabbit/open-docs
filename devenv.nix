{ pkgs, ... }:

{
  languages.go = {
    enable = true;
    version = "1.26.4";
  };

  packages = with pkgs; [
    git
    opentofu
    prek
  ];

  scripts = {
    build.exec = "go build -o open-doc .";
    format.exec = ''
      find . -type f -name '*.go' -not -path './.devenv/*' -exec gofmt -w {} +
    '';
    verify.exec = ''
      set -eu

      unformatted_files="$(find . -type f -name '*.go' -not -path './.devenv/*' -exec gofmt -l {} +)"
      if [ -n "$unformatted_files" ]; then
        echo "The following Go files need gofmt:" >&2
        echo "$unformatted_files" >&2
        exit 1
      fi

      go build -o open-doc .
      go vet ./...
      go test ./...

      smoke_output="$(mktemp)"
      trap 'rm -f "$smoke_output"' EXIT
      ./open-doc -o "$smoke_output" ./examples/vpc
      test -s "$smoke_output"
    '';
  };

  enterShell = ''
    if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
      prek install >/dev/null
    fi

    echo "open-doc development environment"
    echo "  $(go version)"
    echo "  $(tofu version | head -n 1)"
    echo "  $(prek --version)"
    echo "Run 'verify' before submitting changes."
  '';

  enterTest = ''
    verify
  '';
}

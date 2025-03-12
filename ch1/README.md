**_go run_**

Use go run when you want to treat a Go program like a script and run the source code immediately.

```zsh
go run hello.go
```

**_go build_**

This creates an executable, the name of the binary matches the name of the file or package that you passed in. If you want a different name for your application
or if you want to store it in a different location, use the -o flag.

Use go build to create a binary that is distributed for other people to use.

```zsh
go build -o hello_word hello.go
```

**_getting third party go tools_**

```zsh
go install github.com/rakyll/hey@latest
```

**_makefiles_**

```zsh
make
```

By entering a single command, we make sure the code was formatted correctly, vet the code for nonobvious errors, and compile. We can also run the linter with _make lint_,
vet the code with _make vet_ or just run the formatter with _make fmt_.

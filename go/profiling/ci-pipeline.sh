echo "Grab dependencies, please wait..."
go get -u github.com/kisielk/errcheck
go get -u github.com/fzipp/gocyclo
go get -u golang.org/x/lint/golint
go get -u github.com/securego/gosec/cmd/gosec
echo "Ok, inoking go tools.."
errcheck -blank .
gocyclo -over 5 .
golint
gosec .
echo "Done :)"
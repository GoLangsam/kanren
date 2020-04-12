@Echo On

call doc.cmd
@Echo On

go generate .\...
go fmt  .\...
go vet  .\...
go test	.\...

@Echo.
@Echo finished :-)

@pause

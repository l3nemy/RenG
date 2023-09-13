module github.com/RenG-Visual-Novel-Engine/RenG

go 1.19.2

require internal/RVM v0.0.0

replace internal/RVM => ./internal/RVM

require internal/compiler v0.0.0

replace internal/compiler => ./internal/compiler

require (
	github.com/kaitai-io/kaitai_struct_go_runtime v0.10.0 // indirect
	golang.org/x/text v0.3.7 // indirect
)

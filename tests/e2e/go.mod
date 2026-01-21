module github.com/dachrisch/agent-deck-dev/tests/e2e

go 1.25.5

require (
	github.com/asheshgoplani/agent-deck v0.0.0-00010101000000-000000000000
	github.com/creack/pty v1.1.24
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/sahilm/fuzzy v0.1.1 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/term v0.37.0 // indirect
	golang.org/x/time v0.14.0 // indirect
)

replace github.com/asheshgoplani/agent-deck => ../../agent-deck

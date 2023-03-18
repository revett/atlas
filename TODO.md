# ⏳ Backlog

- [ ] Convert to full TUI charmbracelet/bubbletea app
  - https://github.com/charmbracelet/bubbletea
- [ ] Release v0.1.0 via GoReleaser and Homebrew
  - https://goreleaser.com/customization/homebrew
  - https://www.terminalizer.com/docs#usage
  - https://asciinema.org/docs/usage
  - https://doitlive.readthedocs.io/en/stable
  - https://github.com/christopher-dG/asciiscript
  - https://github.com/PierreMarchand20/asciinema_automation
  - https://github.com/paxtonhare/demo-magic
  - https://gabygoldberg.medium.com/the-building-in-public-how-to-guide-219d417f00c1
- [ ] Refactor `note.NewNote()`
- [ ] Add autocomplete to new note CLI command
- [ ] Command to convert note from base type A to base type B, e.g. `scratch` to `meeting`
- [ ] Command to quickly archive/unarchive a note
- [ ] Check that no other notes with the same ID exist
- [ ] Validate and fix old note ID's that weren't human-readable
- [ ] Write documentation on contributing
- [ ] Add live reload within local development container
  - https://github.com/cosmtrek/air
- [ ] Add command to reset IDs
- [ ] Create `atlas init` command
- [ ] Add validation for hanging templates

# ✅ Completed

> Most recent at the top.

- [x] Make templates optional in validation
- [x] Add install and usage documentation in README
- [x] Use `path` config value when creating new notes etc
- [x] Include configuration file with root path value that allows `atlas` to be run from anywhere
- [x] Improve developer experience by having ephemeral knowledge base locally
- [x] Improve the word list used for note
- [x] Remove `review` as note base type

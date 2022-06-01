<p align="center">
  <img src="./assets/skype-squid-emoji.png" alt="Squid Emoji" width="120px">
</p>

# Sepias

The CLI I use to manage my notes / personal knowledge management
([PKM](https://en.wikipedia.org/wiki/Personal_knowledge_management)).

## Design Principles

- Note hierarchy (`area.language.go.errors`) inspired by
  [Dendron](https://wiki.dendron.so)
- Schema inspired by the [PARA method](https://fortelabs.co/blog/para)
- Simplicity, originally started as a
  [bash script](https://github.com/revett/sepias/commit/ef85ce9d70593c8b3dfb8e1c2aab56300bbee959)

## Workflow

- VS Code with a pinned terminal editor in tab 1 position
- `CMD-P` for navigating to notes within VS Code
- `CMD-SHIFT-F` for searching all notes
- [Markdown Memo](https://github.com/svsool/memo) for creating backlinks in VS
  Code
- [Reflow Markdown](https://github.com/marvhen/reflowmarkdown) for quickly
  formatting text to a preferred line length
- [cSpell](https://github.com/streetsidesoftware/vscode-spell-checker) for spell
  checking within markdown files
- [TODO Highlight](https://github.com/wayou/vscode-todo-highlight) to highlight
  `todo-` annotations within markdown files, as I use this as a placeholder when
  needing to create additional notes

## Schema

All notes fall under a specific schema type:

- **S**ystem - Checklists to follow when carrying out repeated tasks
- **E**ntity - Person, company, location etc
- **P**roject - Notes linked to a goal, with a deadline (taken from the
  [PARA method](https://fortelabs.co/blog/para))
- **I**nterview - Notes from interviewing candidates
- **A**rea - Sphere of activity with a standard to be maintained over time
  (taken from the [PARA method](https://fortelabs.co/blog/para))
- **S**cratch - Random note

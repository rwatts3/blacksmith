---
title: blacksmith generate source
enterprise: false
docker: false
---

# `blacksmith generate source`

This command generates the files for a source. It is the recommended way for a
developer to create a source quickly and efficiently.

**Example:**
```bash
$ blacksmith generate source --name stripe

```

**Related ressources:**
- ETL guides >
  [Sources](/blacksmith/guides/extraction/sources)

## Required flags

- `--name [source]`: Set the name of the source. It shall be a valid name only
  containing lowercase letters (`a-z`), underscores (`_`), and dashes (`-`).

  **Example:**
  ```bash
  $ blacksmith generate source --name stripe

  ```

## Optional flags

- `--path [path]`: Relative path where the source's files will be generated. If
  the directories don't exist, they will automatically be created (if possible).
  If files already exist at the path with the same name of the ones generated by
  the CLI, an error will be prompted. The CLI will never override existing files.

  **Example:**
  ```bash
  $ blacksmith generate source --name stripe --path ./sources/stripe

  ```

- `--migrations`: Enable the source to have migrations. This adds the appropriate
  methods to the source to handle migrations for the `wanderer` adapter.

  **Example:**
  ```bash
  $ blacksmith generate source --name stripe --migrations

  ```

- `--no-comments`: Remove the comments on the files generated.

  **Example:**
  ```bash
  $ blacksmith generate source --name stripe --no-comments

  ```
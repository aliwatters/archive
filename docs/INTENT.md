# What archive is for

**This repository preserves publicly safe, date-labeled snapshots of historical `aliwatters` repositories and their source-revision provenance.**

## Why this document exists

This document settles whether `archive` is a runnable product or an archive of separate projects. It explains how to read its dated directories, manifest, provenance records, and retained archival helper without treating any one snapshot as the repository's primary application.

## What it does

- Stores 19 public snapshots in date-prefixed directories such as `2023-hello-service`, `2022-nextjs-blog`, and `2015-go-course-01`.
- Records each snapshot in `manifests/public-snapshots.tsv` with its source URL, captured commit, last-active time, visibility, sensitivity classification, and capture date.
- Keeps an `ARCHIVE.md` beside each public snapshot; for example, `2023-hello-service/ARCHIVE.md` records the exact source revision `136b97ba86b1cf4872853c044df421755214ff9d`.
- Retains `deprecate-repo.sh`, a legacy archival helper that clones a named `aliwatters` repository, relocates its contents under a repository-named directory with `git mv`, and merges it into `archive`.

## What it is not for

- It is not one integrated application or service: there is no root package manifest or application entry point, while independent entry points such as `2023-hello-service/src/main.ts` and `2022-learn-rust/chapter-01/hello_cargo/src/main.rs` remain inside separate snapshots.
- It is not a shared build or test suite: snapshot-specific commands such as `2023-hello-service/package.json`'s `test` and `test:e2e` belong to that snapshot, and the retained CI workflow is scoped to `2021-hello-github-actions/.github/workflows/main.yml`.
- It is not a deployable stack: Dockerfiles and Compose files are contained within individual snapshots, including `2022-docker-php-minimal/docker-compose.yml` and `2023-hello-service/docker-compose.yml`, rather than defining a root deployment.
- It does not finalize the disposition of the source repositories: every row in `manifests/public-snapshots.tsv` has `final_action` set to `pending-confirmation`.

## How to tell it is working

- Each row in `manifests/public-snapshots.tsv` identifies an existing date-prefixed snapshot directory and supplies a source URL and captured commit.
- The 19 public snapshot directories each contain an `ARCHIVE.md` with provenance fields for the source, captured commit, last-active time, visibility, sensitivity bucket, capture time, and final source-repository action.
- The captured directory contents retain their original project structures—for example, the Nest service has `src/`, `test/`, and `package.json`, while the Rust snapshot retains chapter-level `Cargo.toml` files.
- `git log` includes `8035b7c chore: add dated public archive snapshots` and the older imported histories, providing a repository-level record of archival additions.

## Where it fits

The manifest connects this repository to the original public GitHub source URLs and identifies the source revision preserved for each snapshot. Its public boundary is explicit: all manifest records are classified `public` and `public-safe`; the README states that sensitive or uncertain repositories are kept in a private archive instead.

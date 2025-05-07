## hobot

Hetzner Robot API CLI

### Synopsis
A CLI to interact with the Hetzner Robot API - Copyright 2025 Fabian Topfstedt

### Usage 

All of the usage docs live in [docs/index.md](docs/hobot.md) ↗

### Output formats

Hobot supports theses output formats when interacting with Hetzner’s Robot API:

    table
    A CSV-compliant, comma-separated list with a header line (omit headers with
    --no-headers).
    Field order is not guaranteed! Use the next option to explicitly lock in
    the fields (and their order) when building your tooling.

    table=Foo,Bar,Baz
    Same as above, but you specify exactly which fields to output (and in what
    order).

    json
    Pretty-printed JSON output.

    yaml
    YAML-formatted output.

In the spirit of UNIX pipelines, use the right tool for each format:

    For CSV/table data:
    – [cut](https://www.gnu.org/software/coreutils/manual/html_node/The-cut-command.html)
    – [awk](https://www.gnu.org/software/gawk/manual/gawk.html)
    – [csvtk](https://github.com/shenwei356/csvtk)

    For JSON data:
    – [jq](https://github.com/jqlang/jq)

    For YAML data:
    – [yq](https://github.com/mikefarah/yq)

Pipe Hobot’s output into these utilities when building your scripts!

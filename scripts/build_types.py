"""Simple script to run the auto-gen types from the schema"""

import argparse
import json
import os
import re
import shlex
import subprocess

from collections import defaultdict
from dataclasses import dataclass
from typing import Literal

DIR = os.path.dirname(__file__)
ROOT_DIR = os.path.dirname(DIR)
SCHEMA_DIR = os.path.join(ROOT_DIR, "schemas")
GO_DIR = os.path.join(ROOT_DIR, "backend", "models")
TS_DIR = os.path.join(ROOT_DIR, "frontend", "src", "types")
GO_CONSTR_PATH = os.path.join(os.path.dirname(__file__), "generate_constructors.go")

parser = argparse.ArgumentParser()
parser.add_argument("--root-schema", "-r")
parser.add_argument("--debug", "-d", action="store_true")
parser.add_argument("--verbose", "-v", action="store_true")
args = parser.parse_args()

SubReturn = Literal["out", "err", "both"]
""" String literal type representing the output choices for cmdx """


# pylint: disable=C0103
@dataclass
class SubReturns:
    """Enum class for SubReturn values"""

    OUT: SubReturn = "out"
    ERR: SubReturn = "err"
    BOTH: SubReturn = "both"


# pylint: enable=C0103


def run_command(
    cmd: list[str] | str, rtrn: SubReturn = "out", print_out: bool = True
) -> str | tuple[str, str]:
    """Executes a command and returns the output or error

    Args:
        cmd (list[str] | str): - A list of strings that make up the command or a string
            that will be split by spaces
        rtrn (SubReturn, optional): What outputs to return. If both, it will return a
            tuple of (stdout, stderr)Defaults to 'out'.

    Returns:
        str | tuple[str, str]: The output of the command or a tuple of (stdout, stderr)
    """

    if isinstance(cmd, str):
        cmd = shlex.split(cmd)

    try:
        process = subprocess.run(
            cmd,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            encoding="utf-8",
            check=True,
        )
    except subprocess.CalledProcessError as e:
        # Print and handle the errors here if needed
        process = e

    stdout = process.stdout
    stderr = process.stderr
    if print_out:
        if stdout:
            print(stdout)
        if stderr:
            print("\nERROR:\n", stderr)

    if rtrn == "out":
        return process.stdout
    if rtrn == "err":
        return process.stderr

    return process.stdout, process.stderr


def read_json(path: str) -> dict:
    with open(path, "r", encoding="UTF-8") as fs:
        return json.load(fs)


def flatten_json(data: object, parent_key: str = "") -> dict[str, object]:
    """Recursively flattens a nested JSON-like structure (dicts and lists) into dot-separated keys.

    Args:
        data: The JSON data (typically parsed into dicts, lists, and primitives).
        parent_key: The base key used for recursion (used internally).

    Returns:
        A flattened dictionary where keys are dot-separated paths to primitive values.
    """
    items: dict[str, object] = {}

    if isinstance(data, dict):
        for key, value in data.items():
            full_key = f"{parent_key}.{key}" if parent_key else key
            items.update(flatten_json(value, full_key))
    elif isinstance(data, list):
        for index, value in enumerate(data):
            full_key = f"{parent_key}.{index}" if parent_key else str(index)
            items.update(flatten_json(value, full_key))
    else:
        items[parent_key] = data

    return items


def to_pascal_case(text: str) -> str:
    """Convert a string to PascalCase."""
    return re.sub(r"(_|-|^)([a-z])", lambda m: m.group(2).upper(), text).replace(
        " ", ""
    )


def write_log(stdout: str, stderr: str) -> None:
    with open(os.path.join(DIR, "stdout"), "w", encoding="UTF-8") as fs:
        fs.write(stdout)
    with open(os.path.join(DIR, "stderr"), "w", encoding="UTF-8") as fs:
        fs.write(stderr)


def compile_go(schema_name: str, out_name: str = "gen_types.go") -> None:
    schema_path = os.path.join(SCHEMA_DIR, f"{schema_name}.schema.json")

    print(f"Generating types for {schema_path}..")

    debug = " --debug=print-schema-resolving" if args.debug else ""
    go_path = os.path.join(GO_DIR, out_name)

    cmd = f"quicktype --package models{debug} --top-level {to_pascal_case(schema_name)}  -s schema {schema_path} -o {go_path}"
    if args.verbose:
        print(cmd)
    if args.debug:
        out, err = run_command(cmd, SubReturns.BOTH, False)
        write_log(out, err)
    else:
        run_command(cmd)


def compile_enums() -> None:
    print("Compiling detected Enums for TS")

    schemas = [
        read_json(os.path.join(SCHEMA_DIR, filename))
        for filename in os.listdir(SCHEMA_DIR)
        if filename.endswith(".json")
    ]

    enum_lists = defaultdict(lambda: defaultdict(list))
    # need title, prop key, and enums
    for schema in schemas:
        flat = flatten_json(schema)
        title = schema.get("title")
        # print(json.dumps(flat, indent=4))
        for key, value in flat.items():
            path = key.split(".")
            if len(path) < 4:
                continue

            if path[-2] != "enum" or path[-4] != "properties":
                continue

            enum_lists[path[-3]][title].append(value)

    enum_output = ""

    for prop_name, enums in enum_lists.items():
        for s_name, enum_values in enums.items():

            enum_name = f"{s_name}{to_pascal_case(prop_name)}"
            enum_output += f"export enum {enum_name} {{\n"
            for enum in enum_values:
                enum_output += f"    {to_pascal_case(enum)} = '{enum}',\n"
            enum_output += "}\n\n"

    ts_path = os.path.join(TS_DIR, "auto-enums.ts")
    with open(ts_path, "w", encoding="UTF-8") as fs:
        fs.write(enum_output)


def compile_schemas() -> None:

    for go_file in os.listdir(GO_DIR):
        go_path = os.path.join(GO_DIR, go_file)
        os.remove(go_path)

    """ for schema_path in schemas: """
    root_schema = args.root_schema or "app-state"
    compile_go(root_schema)

    # run_command(f"quicktype -s schema {schema_path} -o {ts_path} ")
    print("Types generated successfully")

    print("Generating constructors")
    run_command(f"go run {GO_CONSTR_PATH} {GO_DIR}")

    compile_enums()


def main():
    compile_schemas()


if __name__ == "__main__":
    main()

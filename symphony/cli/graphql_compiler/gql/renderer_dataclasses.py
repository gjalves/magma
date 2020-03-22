#!/usr/bin/env python3
from typing import Dict

from graphql import GraphQLSchema

from .query_parser import (
    ParsedField,
    ParsedObject,
    ParsedOperation,
    ParsedQuery,
    ParsedVariableDefinition,
)
from .utils_codegen import (
    CodeChunk,
    get_enum_filename,
    get_fragment_filename,
    get_input_filename,
)


class DataclassesRenderer:
    def __init__(self, schema: GraphQLSchema) -> None:
        self.schema = schema

    def render(self, parsed_query: ParsedQuery) -> str:
        buffer = CodeChunk()
        buffer.write("#!/usr/bin/env python3")
        buffer.write("# @" + "generated AUTOGENERATED file. Do not Change!")
        buffer.write("")
        buffer.write("from dataclasses import dataclass")
        buffer.write("from datetime import datetime")
        buffer.write("from gql.gql.datetime_utils import DATETIME_FIELD")
        buffer.write("from gql.gql.graphql_client import GraphqlClient")
        buffer.write("from functools import partial")
        buffer.write("from numbers import Number")
        buffer.write("from typing import Any, Callable, List, Mapping, Optional")
        buffer.write("")
        buffer.write("from dataclasses_json import DataClassJsonMixin")
        buffer.write("")
        for fragment_name in sorted(set(parsed_query.used_fragments)):
            buffer.write(
                f"from .{get_fragment_filename(fragment_name)} import {fragment_name}, QUERY as {fragment_name}Query"
            )
        enum_names = set()
        for enum in parsed_query.enums:
            enum_names.add(enum.name)
        if enum_names:
            buffer.write("from gql.gql.enum_utils import enum_field")
            for enum_name in sorted(enum_names):
                buffer.write(f"from .{get_enum_filename(enum_name)} import {enum_name}")
            buffer.write("")
        input_object_names = set()
        for input_object in parsed_query.input_objects:
            input_object_names.add(input_object.name)
        if input_object_names:
            for input_object_name in sorted(input_object_names):
                buffer.write(
                    f"from .{get_input_filename(input_object_name)} "
                    f"import {input_object_name}"
                )
            buffer.write("")

        sorted_objects = sorted(
            parsed_query.objects,
            key=lambda obj: 1 if isinstance(obj, ParsedOperation) else 0,
        )
        for obj in sorted_objects:
            buffer.write("")
            if isinstance(obj, ParsedObject):
                self.__render_object(parsed_query, buffer, obj)
            elif isinstance(obj, ParsedOperation):
                self.__render_operation(parsed_query, buffer, obj)

        if parsed_query.fragment_objects:
            if parsed_query.used_fragments:
                queries = [
                    f"{fragment_name}Query"
                    for fragment_name in sorted(set(parsed_query.used_fragments))
                ]
                buffer.write(f'QUERY: str = {" + ".join(queries)} + """')
            else:
                buffer.write('QUERY: str = """')
            buffer.write(parsed_query.query)
            buffer.write('"""')
            buffer.write("")

        for obj in parsed_query.fragment_objects:
            self.__render_fragment(parsed_query, buffer, obj)

        return str(buffer)

    def render_enums(self, parsed_query: ParsedQuery) -> Dict[str, str]:
        result = {}

        for enum in parsed_query.enums + parsed_query.internal_enums:
            buffer = CodeChunk()
            buffer.write("#!/usr/bin/env python3")
            buffer.write("# @" + "generated AUTOGENERATED file. Do not Change!")
            buffer.write("")
            buffer.write("from enum import Enum")
            buffer.write("")
            with buffer.write_block(f"class {enum.name}(Enum):"):
                for value_name, value in enum.values.items():
                    if isinstance(value, str):
                        value = f'"{value}"'

                    buffer.write(f"{value_name} = {value}")
                buffer.write('MISSING_ENUM = ""')
                buffer.write("")
                buffer.write("@classmethod")
                with buffer.write_block(
                    f'def _missing_(cls, value: str) -> "{enum.name}":'
                ):
                    buffer.write("return cls.MISSING_ENUM")
            buffer.write("")
            result[enum.name] = str(buffer)

        return result

    def render_input_objects(self, parsed_query: ParsedQuery) -> Dict[str, str]:
        result = {}

        for input_object in parsed_query.input_objects + parsed_query.internal_inputs:
            buffer = CodeChunk()
            buffer.write("#!/usr/bin/env python3")
            buffer.write("# @" + "generated AUTOGENERATED file. Do not Change!")
            buffer.write("")
            buffer.write("from dataclasses import dataclass")
            buffer.write("from datetime import datetime")
            buffer.write("from functools import partial")
            buffer.write("from gql.gql.datetime_utils import DATETIME_FIELD")
            buffer.write("from numbers import Number")
            buffer.write("from typing import Any, Callable, List, Mapping, Optional")
            buffer.write("")
            buffer.write("from dataclasses_json import DataClassJsonMixin")
            buffer.write("")
            enum_names = set()
            for enum in input_object.input_enums:
                enum_names.add(enum.name)
            if enum_names:
                buffer.write("from gql.gql.enum_utils import enum_field")
                for enum_name in sorted(enum_names):
                    buffer.write(
                        f"from .{get_enum_filename(enum_name)} import {enum_name}"
                    )
                buffer.write("")
            input_object_names = set()
            for input_dep in input_object.inputs:
                input_object_names.add(input_dep.name)
            for input_object_name in sorted(input_object_names):
                buffer.write(
                    f"from .{get_input_filename(input_object_name)} "
                    f"import {input_object_name}"
                )

            self.__render_object(parsed_query, buffer, input_object)
            buffer.write("")
            result[input_object.name] = str(buffer)

        return result

    def __render_object(
        self, parsed_query: ParsedQuery, buffer: CodeChunk, obj: ParsedObject
    ) -> None:
        class_parents = (
            "(DataClassJsonMixin)" if not obj.parents else f'({", ".join(obj.parents)})'
        )

        buffer.write("@dataclass")
        with buffer.write_block(f"class {obj.name}{class_parents}:"):
            # render child objects
            children_names = set()
            for child_object in obj.children:
                if child_object.name not in children_names:
                    self.__render_object(parsed_query, buffer, child_object)
                children_names.add(child_object.name)

            # render fields
            sorted_fields = sorted(obj.fields, key=lambda f: 1 if f.nullable else 0)
            for field in sorted_fields:
                self.__render_field(parsed_query, buffer, field)

            # pass if not children or fields
            if not (obj.children or obj.fields):
                buffer.write("pass")

        buffer.write("")

    def __render_fragment(
        self, parsed_query: ParsedQuery, buffer: CodeChunk, obj: ParsedObject
    ) -> None:
        class_parents = (
            "(DataClassJsonMixin)" if not obj.parents else f'({", ".join(obj.parents)})'
        )

        buffer.write("@dataclass")
        with buffer.write_block(f"class {obj.name}{class_parents}:"):

            # render child objects
            children_names = set()
            for child_object in obj.children:
                if child_object.name not in children_names:
                    self.__render_object(parsed_query, buffer, child_object)
                children_names.add(child_object.name)

            # render fields
            sorted_fields = sorted(obj.fields, key=lambda f: 1 if f.nullable else 0)
            for field in sorted_fields:
                self.__render_field(parsed_query, buffer, field)

        buffer.write("")

    def __render_operation(
        self, parsed_query: ParsedQuery, buffer: CodeChunk, parsed_op: ParsedOperation
    ) -> None:
        buffer.write("@dataclass")
        with buffer.write_block(f"class {parsed_op.name}(DataClassJsonMixin):"):
            # Render children
            for child_object in parsed_op.children:
                self.__render_object(parsed_query, buffer, child_object)

            # operation fields
            buffer.write(f"data: {parsed_op.name}Data")
            buffer.write("")

            # Execution functions
            if parsed_op.variables:
                vars_args = ", " + ", ".join(
                    [
                        self.__render_variable_definition(var)
                        for var in parsed_op.variables
                    ]
                )
                variables_dict = (
                    "{"
                    + ", ".join(
                        f'"{var.name}": {var.name}' for var in parsed_op.variables
                    )
                    + "}"
                )
            else:
                vars_args = ""
                variables_dict = "{}"

            if len(parsed_query.used_fragments):
                queries = [
                    f"{fragment_name}Query"
                    for fragment_name in sorted(set(parsed_query.used_fragments))
                ]
                buffer.write(f'__QUERY__: str = {" + ".join(queries)} + """')
            else:
                buffer.write('__QUERY__: str = """')
            buffer.write(parsed_query.query)
            buffer.write('"""')
            buffer.write("")

            buffer.write("@classmethod")
            buffer.write("# fmt: off")
            with buffer.write_block(
                f"def execute(cls, client: GraphqlClient{vars_args})"
                f" -> {parsed_op.name}Data:"
            ):
                buffer.write("# fmt: off")
                buffer.write(f"variables = {variables_dict}")
                buffer.write(
                    "response_text = client.call(cls.__QUERY__, " "variables=variables)"
                )
                buffer.write("return cls.from_json(response_text).data")

            buffer.write("")

    @staticmethod
    def __render_variable_definition(var: ParsedVariableDefinition):
        var_type = var.type

        if var_type == "DateTime":
            var_type = "datetime"
        elif var_type == "Cursor":
            var_type = "str"

        if var.is_list:
            return f"{var.name}: List[{var_type}] = []"

        if not var.nullable:
            return f"{var.name}: {var_type}"

        return f'{var.name}: Optional[{var_type}] = {var.default_value or "None"}'

    @staticmethod
    def __render_field(
        parsed_query: ParsedQuery, buffer: CodeChunk, field: ParsedField
    ) -> None:
        enum_names = [e.name for e in parsed_query.enums + parsed_query.internal_enums]
        is_enum = field.type in enum_names
        suffix = ""
        field_type = field.type

        if is_enum:
            suffix = f" = enum_field({field.type})"

        if field.type == "DateTime":
            suffix = " = DATETIME_FIELD"
            field_type = "datetime"

        if field.nullable:
            suffix = f" = {field.default_value}"
            buffer.write(f"{field.name}: Optional[{field_type}]{suffix}")
        else:
            buffer.write(f"{field.name}: {field_type}{suffix}")

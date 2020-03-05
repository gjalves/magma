#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!
# pyre-strict

from dataclasses import dataclass
from datetime import datetime
from functools import partial
from gql.gql.datetime_utils import DATETIME_FIELD
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin

from gql.gql.enum_utils import enum_field
from .survey_status_enum import SurveyStatus

from .survey_question_response_input import SurveyQuestionResponse
@dataclass
class SurveyCreateData(DataClassJsonMixin):
    name: str
    completionTimestamp: int
    locationID: str
    surveyResponses: List[SurveyQuestionResponse]
    ownerName: Optional[str] = None
    creationTimestamp: Optional[int] = None
    status: Optional[SurveyStatus] = None


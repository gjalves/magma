#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from enum import Enum

class ImageEntity(Enum):
    LOCATION = "LOCATION"
    WORK_ORDER = "WORK_ORDER"
    SITE_SURVEY = "SITE_SURVEY"
    EQUIPMENT = "EQUIPMENT"
    MISSING_ENUM = ""

    @classmethod
    def _missing_(cls, value: str) -> "ImageEntity":
        return cls.MISSING_ENUM

import pytest
from src.alchemy import YakSokController


def test1():
    ysctrl = YakSokController.factory()
    print(ysctrl.metadata)
    print(ysctrl.table['yaksoks'])
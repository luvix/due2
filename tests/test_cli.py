import pytest


def test_main():
    pytest.main(['-h'])
    pytest.main(['--help'])
    pytest.main(['--list'])
    pytest.main(['--list-all'])

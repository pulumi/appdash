#!/usr/bin/env python

from setuptools import setup

setup(
    name = 'Appdash',
    version = '1.0',
    description = 'Appdash Python Integration',
    author = 'Sourcegraph',
    author_email = 'hi@sourcegraph.com',
    url = 'https://github.com/pulumi/appdash',
    packages = ['appdash'],
    install_requires = ['basictracer'],
)
